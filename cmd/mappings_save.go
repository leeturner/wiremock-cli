package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewSaveMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "save",
		Short:   "Persist stub mappings",
		Long:    `Save all persistent stub mappings to the backing store`,
		Example: `wm mappings save`,
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			body, err := wm.SaveMappings()
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
