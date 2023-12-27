package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewRecordingsStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "status",
		Short:   "Get recording status",
		Long:    `Get recording status`,
		Example: `wm recordings status `,
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			status, err := wm.GetRecordingsStatus()
			if err != nil {
				return err
			}
			if status != "" {
				cmd.Println(status)
			}
			return nil
		},
	}

	return cmd
}
