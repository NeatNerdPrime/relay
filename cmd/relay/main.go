package main

import (
	"os"

	"github.com/puppetlabs/relay/pkg/config"
	"github.com/puppetlabs/relay/pkg/format/error"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:           "relay",
		Short:         "Relay by Puppet",
		Args:          cobra.MinimumNArgs(1),
		SilenceErrors: true,
		Long: `Relay connects your tools, APIs, and infrastructure 
to automate common tasks through simple event driven workflows.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// This turns off usage info in json output mode
			cfg, cfgerr := config.GetConfig(cmd.Flags())

			if cfgerr == nil && cfg.Out == config.OutputTypeJSON {
				cmd.SilenceUsage = true
			}
		},
	}

	cmd.PersistentFlags().BoolP("debug", "d", false, "print debugging information")
	cmd.PersistentFlags().StringP("out", "o", "text", "output type: (text|json)")
	// Config flag is hidden for now
	cmd.PersistentFlags().StringP("config", "c", "", "path to config file (default is $HOME.config/relay)")
	cmd.PersistentFlags().MarkHidden("config")

	cmd.AddCommand(NewAuthCommand())

	if err := cmd.Execute(); err != nil {
		error.FormatError(err, cmd)

		os.Exit(1)
	}
}