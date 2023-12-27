package cmd

import "github.com/spf13/cobra"

func NewRequestsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "requests",
		Short: "Handles all request commands for the Wiremock server",
		Long:  "Handles all request commands for the Wiremock server",
	}

	cmd.AddCommand(NewGetRequestsCommand(), NewDeleteRequestsCommand())

	return cmd
}
