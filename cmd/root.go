package cmd

import (
	"github.com/spf13/cobra"
)

const Version = "0.1"

var Host string
var Port string
var AdminPrefix string

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wm",
		Short: "wm is a simple command line tool to access the Wiremock admin API",
		Long: "wm is a simple command line tool to access the Wiremock admin API \n" +
			"\n" +
			"\u001B[34m██     ██ ██ ██████  ███████ \u001B[33m███    ███  ██████   ██████ ██   ██ \n" +
			"\u001B[34m██     ██ ██ ██   ██ ██      \u001B[33m████  ████ ██    ██ ██      ██  ██  \n" +
			"\u001B[34m██  █  ██ ██ ██████  █████   \u001B[33m██ ████ ██ ██    ██ ██      █████   \n" +
			"\u001B[34m██ ███ ██ ██ ██   ██ ██      \u001B[33m██  ██  ██ ██    ██ ██      ██  ██  \n" +
			"\u001B[34m ███ ███  ██ ██   ██ ███████ \u001B[33m██      ██  ██████   ██████ ██   ██ \n" +
			"\u001B[0m\n" +
			"----------------------------------------------------------------\n" +
			"|               Cloud: https://wiremock.io/cloud               |\n" +
			"|                                                              |\n" +
			"|               Slack: https://slack.wiremock.org              |\n" +
			"----------------------------------------------------------------\n",
		Version: Version,
	}

	cmd.SetOut(cmd.OutOrStdout())
	cmd.SetErr(cmd.OutOrStderr())
	cmd.SetIn(cmd.InOrStdin())

	cmd.PersistentFlags().StringVarP(&Host, "host", "H", "http://localhost", "Wiremock host")
	cmd.PersistentFlags().StringVarP(&Port, "port", "p", "8080", "Wiremock port")
	cmd.PersistentFlags().StringVarP(&AdminPrefix, "admin-prefix", "a", "__admin", "Wiremock admin url prefix")

	cmd.AddCommand(NewResetCommand(), NewShutdownCommand(), NewGetCommand())
	return cmd
}
