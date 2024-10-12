package command

import (
	"fmt"
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/arifai/zenith-cli/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

var (
	templateUrl = "github.com/arifai/zenith"

	CreateCommand = &cobra.Command{
		Use:      "create [module_name]",
		Short:    "Create a new project.",
		Long:     "Creating a new project by cloning the Zenith boilerplate.",
		Args:     cobra.MinimumNArgs(1),
		Example:  "zen create my_project or zen create my_project --org github.com/username",
		RunE:     runCreate,
		PostRunE: PostRunE,
	}
)

func init() {
	CreateCommand.Flags().StringP("org", "o", "", "specify your organization name.")
}

func runCreate(cmd *cobra.Command, args []string) error {
	orgName, _ := cmd.Flags().GetString("org")

	if err := cloneAndSetup(utils.ToSnakeCase(args[0]), orgName); err != nil {
		printer.Red("🚫 Error: %v\n", err)
		return err
	}
	return nil
}

func cloneAndSetup(moduleName, orgName string) error {
	clonePath := filepath.Join(".", moduleName)
	repoUrl := "https://" + templateUrl

	if _, err := os.Stat(clonePath); !os.IsNotExist(err) {
		return fmt.Errorf("folder %s already exists", moduleName)
	}

	if err := utils.RunCommand("git", "clone", repoUrl, clonePath); err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	} else {
		printer.Green("✨ Successfully created %s.", moduleName)
	}

	// Wait for the cloning process to finish
	time.Sleep(3 * time.Second)

	if err := utils.UpdateGoModAndImports(templateUrl, clonePath, moduleName, orgName); err != nil {
		return fmt.Errorf("failed to update module and imports: %v", err)
	}

	if err := deleteFolder(clonePath, ".git"); err != nil {
		return fmt.Errorf("failed to delete .git folder: %v", err)
	}

	if err := deleteFolder(clonePath, ".idea"); err != nil {
		return fmt.Errorf("failed to delete .idea folder: %v", err)
	}

	return nil
}

func deleteFolder(basePath, folderName string) error {
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		printer.Yellow("📁 %s folder does not exist.", folderPath)
		return nil
	}
	printer.Yellow("🗑️ Deleting %s folder...", folderPath)
	return os.RemoveAll(folderPath)
}
