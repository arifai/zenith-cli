package command

import (
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/arifai/zenith-cli/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
)

var (
	templateUrl = "github.com/arifai/zenith"

	CreateCommand = &cobra.Command{
		Use:     "create [module_name]",
		Short:   "Create a new project.",
		Long:    "Creating a new project by cloning the Zenith boilerplate.",
		Args:    cobra.MinimumNArgs(1),
		Example: "zen create my_project or zen create my_project --org github.com/username",
		Run:     runCreate,
	}
)

func init() {
	CreateCommand.Flags().StringP("org", "o", "", "specify your organization name.")
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

	if err := utils.RunCommand("git", "clone", repoUrl, clonePath); err != nil {
		printer.Red("🚫 Failed to clone repository: %v.", err)
		os.Exit(1)
	} else {
		printer.Green("✨ Successfully created %s.", moduleName)
	}

	if err := utils.UpdateGoModAndImports(templateUrl, clonePath, moduleName, orgName); err != nil {
		printer.Red("🚫 Failed to update module and imports: %v.", err)
		os.Exit(1)
	}

	if err := deleteFolder(clonePath, ".git"); err != nil {
		printer.Red("🚫 Failed to delete .git folder: %v.", err)
		os.Exit(1)
	}

	if err := deleteFolder(clonePath, ".idea"); err != nil {
		printer.Red("🚫 Failed to delete .idea folder: %v.", err)
		os.Exit(1)
	}

	return nil
}

func deleteFolder(basePath, folderName string) error {
	folderPath := filepath.Join(basePath, folderName)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return nil
	}

	if runtime.GOOS == "windows" {
		err := utils.RunCommand("cmd", "/C", "rmdir", "/S", "/Q", folderPath)
		if err != nil {
			return err
		}
	} else {
		err := utils.RunCommand("rm", "-rf", folderPath)
		if err != nil {
			return err
		}
	}
	return nil
}
