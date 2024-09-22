package command

import (
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/spf13/cobra"
	"os"
)

var makeCommand = &cobra.Command{
	Use:     "make [module_name]",
	Aliases: []string{"m"},
	Short:   "Generate a new module.",
	Long:    "Generate a new module with specified name.",
	Args:    cobra.ExactArgs(1),
	Example: "zen make account",
	Run:     run,
}

func run(_ *cobra.Command, args []string) {
	moduleName := args[0]
	printer.Yellow("🚀Generating module %s", moduleName)

	file, err := os.Create(moduleName)
	if err != nil {
		printer.Red("🔥Failed to generate new module:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			printer.Red("🔥Failed to close module:", err)
			os.Exit(1)
		}
	}(file)

	printer.Green("✨Successfully generate new module %s", moduleName)
	os.Exit(0)
}
