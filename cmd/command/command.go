package command

import (
	"github.com/spf13/cobra"
	"os"
)

var command = &cobra.Command{
	Use:     "zen",
	Short:   "Zenith CLI for creating and managing projects.",
	Version: "1.0.0",
	Long: "Zenith CLI is a command-line tool for creating and managing projects based on the Zenith boilerplate. " +
		"It streamlines the setup of project modules, allowing you to quickly generate boilerplate code.",
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
	command.AddCommand(makeCommand, createCommand)
}
