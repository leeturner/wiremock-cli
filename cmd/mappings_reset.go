package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewResetMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reset",
		Short:   "Reset stub mappings",
		Long:    `Restores stub mappings to the defaults defined back in the backing store`,
		Example: `wm mappings reset`,
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			body, err := wm.ResetMappings()
			if err != nil {
				return err
			}
			if body != "" {
				cmd.Println(body)
			}
			return nil
		},
	}

	return cmd
}
