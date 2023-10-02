package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewGetRequestsCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "requests",
		Short:   "Get all requests in journal",
		Long:    `Get all requests in journal`,
		Example: "wm get requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			requests, err := wm.GetRequests()
			if err != nil {
				return err
			}
			cmd.Println(requests)
			return nil
		},
	}
}
