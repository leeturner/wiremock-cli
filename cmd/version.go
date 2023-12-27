package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get the version of the remote Wiremock server",
		Long:  `Get the version of the remote Wiremock server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			version, err := wm.Version()
			if err != nil {
				return err
			}

			if version != "" {
				cmd.Println(version)
			}
			return nil
		},
	}
}
