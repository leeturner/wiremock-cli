package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewGetScenariosCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "scenarios",
		Short:   "Get all scenarios",
		Long:    `Get all scenarios`,
		Example: "wm get scenarios",
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			scenarios, err := wm.GetScenarios()
			if err != nil {
				return err
			}
			cmd.Println(scenarios)
			return nil
		},
	}
}
