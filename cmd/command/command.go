package command

import (
	"github.com/spf13/cobra"
	"os"
)

var command = &cobra.Command{
	Use:     "zen",
	Short:   "Zenith CLI",
	Long:    "Zenith is a CLI tool for creating project module for Zenith template.",
	Version: "1.0.0",
}

func Execute() {
	if err := command.Execute(); err != nil {
		err := command.Help()
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	command.AddCommand(makeCommand)
}
