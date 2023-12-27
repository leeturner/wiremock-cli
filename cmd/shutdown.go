package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

func NewShutdownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "shutdown",
		Short: "Shutdown the Wiremock server",
		Long:  `Shuts down the remote Wiremock server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			_, err := wm.Shutdown()
			if err != nil {
				return err
			}

			cmd.Println("Wiremock server shutdown")
			return nil
		},
	}
}
