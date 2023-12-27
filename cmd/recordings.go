package cmd

import "github.com/spf13/cobra"

func NewRecordingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recordings",
		Short: "Stub mapping record and snapshot functions",
		Long:  "Stub mapping record and snapshot functions",
	}

	cmd.AddCommand(NewRecordingsStatusCommand())

	return cmd
}
