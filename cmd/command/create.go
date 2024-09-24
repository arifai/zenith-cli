package command

import (
	"fmt"
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/arifai/zenith-cli/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	templateUrl = "github.com/arifai/zenith"

	createCommand = &cobra.Command{
		Use:     "create [module_name]",
		Short:   "Create a new project.",
		Long:    "Creating a new project by cloning the Zenith boilerplate.",
		Args:    cobra.MinimumNArgs(1),
		Example: "zen create my_project or zen create my_project --org github.com/username",
		Run:     runCreate,
	}
)

func init() {
	createCommand.Flags().StringP("org", "o", "", "specify your organization name.")
}

func runCreate(cmd *cobra.Command, args []string) {

	orgName, _ := cmd.Flags().GetString("org")

	if err := cloneAndSetup(utils.ToSnakeCase(args[0]), orgName); err != nil {
		printer.Red("🚫 Error: %v\n", err)
		os.Exit(1)
	}
}

func cloneAndSetup(moduleName, orgName string) error {
	clonePath := filepath.Join(".", moduleName)
	repoUrl := "https://" + templateUrl

	if _, err := os.Stat(clonePath); !os.IsNotExist(err) {
		printer.Yellow("⚠️ Folder %s is already exists.", moduleName)
		os.Exit(1)
	}

	if err := exec.Command("git", "clone", repoUrl, clonePath).Run(); err != nil {
		printer.Red("🚫 Failed to clone repository: %v.", err)
		os.Exit(1)
	} else {
		printer.Green("✨ Successfully created %s.", moduleName)
	}

	if err := updateGoModAndImports(clonePath, moduleName, orgName); err != nil {
		printer.Red("🚫 Failed to update module and imports: %v", err)
		os.Exit(1)
	}

	return nil
}

func updateGoModAndImports(clonePath, moduleName, orgName string) error {
	if err := updateGoMod(clonePath, moduleName, orgName); err != nil {
		printer.Red("🚫 Failed to update module name: %v", err)
		os.Exit(1)
	}

	var newModuleName string
	if orgName != "" {
		newModuleName = fmt.Sprintf("%s/%s", orgName, moduleName)
	} else {
		newModuleName = moduleName
	}

	if err := updateImports(clonePath, newModuleName); err != nil {
		printer.Red("🚫 Failed to update import paths: %v", err)
		os.Exit(1)
	}

	return nil
}

func updateGoMod(clonePath, moduleName, orgName string) error {
	goModPath := filepath.Join(clonePath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		printer.Red("🚫 Failed to read go.mod: %v", err)
		os.Exit(1)
	}

	var newModuleName string
	if orgName != "" {
		newModuleName = fmt.Sprintf("%s/%s", orgName, moduleName)
	} else {
		newModuleName = moduleName
	}

	newContent := strings.ReplaceAll(string(content), "module "+templateUrl, "module "+newModuleName)
	if err := os.WriteFile(goModPath, []byte(newContent), 0644); err != nil {
		printer.Red("🚫 Failed to write go.mod: %v", err)
		os.Exit(1)
	}

	return nil
}

func updateImports(clonePath, newModuleName string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command",
			"Get-ChildItem -Path '"+clonePath+"' -Filter '*.go' -Recurse | ForEach-Object { (Get-Content $_.FullName) "+
				"-replace '"+templateUrl+"', '"+newModuleName+"' | Set-Content $_.FullName }")
	} else {
		cmd = exec.Command("find", clonePath, "-name", "*.go", "-exec", "sed", "-i", "",
			"s|"+templateUrl+"|"+newModuleName+"|g", "{}", ";")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
