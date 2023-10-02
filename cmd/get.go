package cmd

import "github.com/spf13/cobra"

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets resources from the Wiremock server",
		Long:  "Gets resources from the Wiremock server",
	}

	cmd.AddCommand(NewGetMappingsCommand(), NewGetScenariosCommand(), NewGetRequestsCommand())

	return cmd
}
