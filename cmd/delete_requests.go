package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

var deleteRequestId = ""

func NewDeleteRequestsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "requests",
		Short: "Delete all requests in journal",
		Long:  `Delete all requests in journal - if an id is specified, only that request is deleted`,
		Example: `wm delete requests
wm delete requests --id 0baca68a-0112-4f26-8529-ac12d8eb3720
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			requests, err := wm.DeleteRequests(deleteRequestId)
			if err != nil {
				return err
			}
			if requests != "" {
				cmd.Println(requests)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&deleteRequestId, "id", "i", "", "Specify the id of the specific request you want to delete")

	return cmd
}
