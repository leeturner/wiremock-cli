package cmd

import "github.com/spf13/cobra"

func NewMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mappings",
		Short: "Handles all mappings commands for the Wiremock server",
		Long:  "Handles all mappings commands for the Wiremock server",
	}

	cmd.AddCommand(
		NewGetMappingsCommand(),
		NewDeleteMappingsCommand(),
		NewImportMappingsCommand(),
		NewSaveMappingsCommand(),
		NewResetMappingsCommand(),
	)

	return cmd
}
