package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

var (
	limitMappings = 10
	offset        = 0
	mappingId     = ""
)

func NewGetMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get stub mappings",
		Long:  `Get stub mappings - if an id is specified, only that mapping is returned`,
		Example: `wm mappings get
wm mappings get --limit 5
wm mappings get --limit 5 --offset 10
wm mappings get --id 0baca68a-0112-4f26-8529-ac12d8eb3720
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			mappings, err := wm.GetMappings(mappingId, limitMappings, offset)
			if err != nil {
				return err
			}
			if mappings != "" {
				cmd.Println(mappings)
			}
			return nil
		},
	}

	cmd.Flags().IntVarP(&limitMappings, "limit", "l", 10, "Limit the number of mappings returned (only used if no id is specified)")
	cmd.Flags().IntVarP(&offset, "offset", "o", 0, "Offset the returned mappings by this number (only used if no id is specified)")
	cmd.Flags().StringVarP(&mappingId, "id", "i", "", "Specify the id of the specific mapping you want to return")

	return cmd
}
