package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewDocsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "docs",
		Short:  "Generates documentation for the Wiremock CLI",
		Long:   "Generates documentation for the Wiremock CLI",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			rootCmd := NewRootCmd()
			return doc.GenMarkdownTree(rootCmd, "./docs")
		},
	}

	return cmd
}
