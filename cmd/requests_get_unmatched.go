package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewGetUnmatchedRequestsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unmatched",
		Short:   " Get all unmatched requests",
		Long:    ` Get all unmatched requests in journal`,
		Example: `wm requests get unmatched`,
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			requests, err := wm.GetUnmatchedRequests()
			if err != nil {
				return err
			}
			if requests != "" {
				cmd.Println(requests)
			}
			return nil
		},
	}

	return cmd
}
