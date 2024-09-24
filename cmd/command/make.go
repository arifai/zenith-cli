package command

import (
	"fmt"
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/arifai/zenith-cli/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"text/template"
)

type MakeCommand struct {
	FilePath    string
	FileType    string
	FeatureName string
}

var makeCommand = &cobra.Command{
	Use:   "make [feature_name]",
	Short: "Generate boilerplate code for a new feature.",
	Long: "Generate boilerplate code for a new feature, including router, handler, model, repository, service, and migration files. " +
		"This command helps you quickly scaffold the essential components of a new feature in your application.",
	Args:     cobra.ExactArgs(1),
	Example:  "zen make account",
	Run:      runMake,
	PostRunE: postRunE,
}

func runMake(_ *cobra.Command, args []string) {
	m := &MakeCommand{FeatureName: args[0]}

	if !utils.CheckGoModFileExists() {
		printer.Red("🚫 Go module not found. Please ensure that this is a Go project and a go.mod file exists in the root directory.")
		os.Exit(1)
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
			printer.Yellow("⚠️ Skipping %s, template file template/%s.tmpl does not exist.", m.FileType, m.FileType)
			continue
		}

		m.genFile()
	}
}

func (m *MakeCommand) genFile() {
	snakeCaseFeatureName := utils.ToSnakeCase(m.FeatureName)
	filePath := filepath.Join("internal", m.FilePath, snakeCaseFeatureName+".go")

	if _, err := os.Stat(filePath); err == nil {
		printer.Yellow("⚠️ File %s already exists.", filePath)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		printer.Red("🚫 Failed to create directory %s: %v.", filepath.Dir(filePath), err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		printer.Red("🚫 Failed to create file %s: %v.", filePath, err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	if err := m.parseTemplate(file); err != nil {
		printer.Red("🚫 Failed to execute template: %v", err)
		os.Exit(1)
	} else {
		printer.Green("✨ Successfully generated file %s.", filePath)
	}
}

func (m *MakeCommand) parseTemplate(file *os.File) error {
	templateFile := fmt.Sprintf("template/%s.tmpl", m.FileType)

	moduleName, err := utils.GetModuleName(".")
	if err != nil {
		printer.Red("🚫 Error getting module name: %v", err)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"ToSnakeCase": utils.ToSnakeCase,
		"ToCamelCase": func(str string) string {
			return utils.ConvertCase(str, false)
		},
		"ToPascalCase": func(str string) string {
			return utils.ConvertCase(str, true)
		},
	}

	tmpl, err := template.New(filepath.Base(templateFile)).Funcs(funcMap).ParseFiles(templateFile)
	if err != nil {
		printer.Red("🚫 Failed to parse template file %s: %v", templateFile, err)
		os.Exit(1)
	}

	data := map[string]interface{}{"FeatureName": m.FeatureName, "ModuleName": moduleName}
	if err := tmpl.Execute(file, data); err != nil {
		printer.Red("🚫 Failed to execute template for file %s: %v", templateFile, err)
		os.Exit(1)
	}

	return nil
}

func (m *MakeCommand) templateExists() bool {
	templateFile := fmt.Sprintf("template/%s.tmpl", m.FileType)
	if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func postRunE(cmd *cobra.Command, args []string) error {
	if err := utils.GoFmt(); err != nil {
		return err
	}

	return nil
}
