package command

import (
	"fmt"
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/arifai/zenith-cli/pkg/utils"
	"github.com/arifai/zenith-cli/tmpl"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

type Make struct {
	FilePath    string
	FileType    string
	FeatureName string
}

var (
	MakeCommand = &cobra.Command{
		Use:   "make [feature_name]",
		Short: "Generate boilerplate code for a new feature.",
		Long: "Generate boilerplate code for a new feature, including router, handler, model, repository, service, and migration files. " +
			"This command helps you quickly scaffold the essential components of a new feature in your application.",
		Args:     cobra.ExactArgs(1),
		Example:  "zen make account",
		RunE:     runMake,
		PostRunE: PostRunE,
	}
)

func init() {
	MakeCommand.Flags().StringP("org", "o", "", "specify your organization name")
}

func runMake(_ *cobra.Command, args []string) error {
	m := &Make{FeatureName: args[0]}

	if !utils.CheckGoModFileExists() {
		return fmt.Errorf("go module not found. Please ensure that this is a Go project and a go.mod file exists in the root directory")
	}

	paths := map[string]string{
		"router":     "api/router",
		"handler":    "handler",
		"model":      "model",
		"repository": "repository",
		"service":    "service",
		"migration":  "model/migration",
		"types_req":  "types/request",
	}

	for fileType, filePath := range paths {
		m.FileType = fileType
		m.FilePath = filePath

		if !m.templateExists() {
			printer.Yellow("⚠️ Skipping %s, tmpl file tmpl/%s.tmpl does not exist.", m.FileType, m.FileType)
			continue
		}

		err := m.genFile()
		if err != nil {
			printer.Red("🚫 Error: %v\n", err)
			return err
		}
	}
	return nil
}

func (m *Make) genFile() error {
	snakeCaseFeatureName := utils.ToSnakeCase(m.FeatureName)
	filePath := filepath.Join("internal", m.FilePath, snakeCaseFeatureName+".go")

	if _, err := os.Stat(filePath); err == nil {
		printer.Yellow("📦 File %s already exists.", filePath)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", filepath.Dir(filePath), err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	if err := m.parseTemplate(file); err != nil {
		return fmt.Errorf("failed to execute tmpl: %v", err)
	} else {
		printer.Green("✨ Successfully generated file %s.", filePath)
	}
	return nil
}

func (m *Make) parseTemplate(file *os.File) error {
	tmplFile := fmt.Sprintf("%s.tmpl", m.FileType)

	funcMap := template.FuncMap{
		"ToSnakeCase": utils.ToSnakeCase,
		"ToCamelCase": func(str string) string {
			return utils.ConvertCase(str, false)
		},
		"ToPascalCase": func(str string) string {
			return utils.ConvertCase(str, true)
		},
	}

	t, err := template.New(tmplFile).Funcs(funcMap).ParseFS(tmpl.TemplateFile, tmplFile)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	moduleName, err := utils.GetModuleName(".")
	if err != nil {
		return fmt.Errorf("error getting module name: %v", err)
	}

	data := map[string]interface{}{
		"FeatureName": m.FeatureName,
		"ModuleName":  moduleName,
	}

	if err := t.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

func (m *Make) templateExists() bool {
	templateFile := fmt.Sprintf("%s.tmpl", m.FileType)
	_, err := fs.Stat(tmpl.TemplateFile, templateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
