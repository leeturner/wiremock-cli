package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewImportMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import",
		Short:   "Import stub mappings",
		Long:    `Import given stub mappings to the backing store`,
		Example: `wm mappings import `,
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			body, err := wm.ImportMappings()
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
