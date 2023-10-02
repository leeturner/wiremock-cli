package cmd

import (
	"github.com/leeturner/wiremock-cli/cmd/wiremock"
	"github.com/spf13/cobra"
)

//var resetCmd = NewResetCommand()

func NewResetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset mappings & journal",
		Long:  `Reset mappings to the default state and reset the request journal`,
		RunE: func(cmd *cobra.Command, args []string) error {
			wm := wiremock.Init(Host, Port)
			wm.WithAdminPrefix(AdminPrefix)
			_, err := wm.Reset()
			if err != nil {
				return err
			}
			cmd.Println("Wiremock server reset")
			return nil
		},
	}
}
