package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

var (
	deleteMappingId = ""
)

func NewDeleteMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete stub mappings",
		Long:  `Delete stub mappings - if an id is specified, only that mapping is deleted`,
		Example: `wm mappings delete
wm mappings delete --id 0baca68a-0112-4f26-8529-ac12d8eb3720
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			_, err := wm.DeleteMappings(deleteMappingId)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&deleteMappingId, "id", "i", "", "Specify the id of the specific mapping you want to delete")

	return cmd
}
