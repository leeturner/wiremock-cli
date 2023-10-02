package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

var (
	limitRequests = 10
	requestId     = ""
	// TODO: implement since parameter
)

func NewGetRequestsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "requests",
		Short: "Get all requests in journal",
		Long:  `Get all requests in journal - if an id is specified, only that request is returned`,
		Example: `wm get requests
wm get requests --limit 5
wm get requests --id 0baca68a-0112-4f26-8529-ac12d8eb3720
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			requests, err := wm.GetRequests(requestId, limitRequests)
			if err != nil {
				return err
			}
			if requests != "" {
				cmd.Println(requests)
			}
			return nil
		},
	}

	cmd.Flags().IntVarP(&limitRequests, "limit", "l", 10, "Limit the number of requests returned (only used if no id is specified)")
	cmd.Flags().StringVarP(&requestId, "id", "i", "", "Specify the id of the specific request you want to return")

	return cmd
}
