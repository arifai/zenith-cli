package command

import (
	"github.com/arifai/zenith-cli/pkg/printer"
	"github.com/arifai/zenith-cli/pkg/utils"
	"github.com/spf13/cobra"
)

func PostRunE(_ *cobra.Command, _ []string) error {
	if err := utils.GoFmt(); err != nil {
		printer.Red("🚫 Error: %v\n", err)
		return err
	}

	return nil
}
