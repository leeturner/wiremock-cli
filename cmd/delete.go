package cmd

import "github.com/spf13/cobra"

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes resources from the Wiremock server",
		Long:  "Deletes resources from the Wiremock server",
	}

	cmd.AddCommand(NewDeleteMappingsCommand())

	return cmd
}
