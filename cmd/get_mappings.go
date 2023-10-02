package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

var (
	limit  = 10
	offset = 0
)

func NewGetMappingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mappings",
		Short: "Get stub mappings",
		Long:  `Get stub mappings - if an id is specified then only that mapping is returned`,
		Example: `wm get mappings
wm get mappings --limit 5
wm get mappings --limit 5 --offset 10
wm get mappings 0baca68a-0112-4f26-8529-ac12d8eb3720
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			var mappings string
			var err error
			// we have an id, so we are only looking for a single mapping
			if len(args) == 1 {
				mappings, err = wm.GetMapping(args[0])
			} else {
				// we don't have an id, so we are looking for all mappings
				mappings, err = wm.GetMappings(limit, offset)
			}
			if err != nil {
				return err
			}
			if mappings != "" {
				cmd.Println(mappings)
			}
			return nil
		},
	}

	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit the number of mappings returned (only used if no id is specified)")
	cmd.Flags().IntVarP(&offset, "offset", "o", 0, "Offset the returned mappings by this number (only used if no id is specified)")

	return cmd
}
