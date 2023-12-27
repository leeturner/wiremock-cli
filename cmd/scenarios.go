package cmd

import "github.com/spf13/cobra"

func NewScenariosCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scenarios",
		Short: "Handles all scenarios commands for the Wiremock server",
		Long:  "Handles all scenarios commands for the Wiremock server",
	}

	cmd.AddCommand(NewGetScenariosCommand())

	return cmd
}
