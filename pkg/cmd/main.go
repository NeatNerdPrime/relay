package cmd

import (
	"os"

	"github.com/puppetlabs/relay/pkg/client"
	"github.com/puppetlabs/relay/pkg/config"
	"github.com/puppetlabs/relay/pkg/dialog"
	"github.com/puppetlabs/relay/pkg/format"
	"github.com/spf13/cobra"
)

// Config is the configuration that our commands should use. We can assume that
// it's been configured accordingly by the time that a command executres.
var Config *config.Config

// Client is the client that we should use based on the configuration. If the
// configuration can't be loaded then we can't assume that the client is
// loaded.
var Client *client.Client

// Dialog is the UI to use derrived from the current configuration.
var Dialog dialog.Dialog

func Execute() {
	cmd := &cobra.Command{
		Use:           "relay",
		Short:         "Relay by Puppet",
		Args:          cobra.MinimumNArgs(1),
		SilenceErrors: true,
		Long: `Relay connects your tools, APIs, and infrastructure 
to automate common tasks through simple event driven workflows.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// This turns off usage info in json output mode
			cfg, err := config.FromFlags(cmd.Flags())

			if err != nil {
				// What kind of error could this be? We will abort accordingly.
				return err
			} else if err == nil && cfg.Out == config.OutputTypeJSON {
				cmd.SilenceUsage = true
			}

			// We have a config that we can assume is good to use.
			Config = cfg
			Client = client.NewClient(Config)
			Dialog = dialog.FromConfig(Config)

			return nil
		},
	}

	cmd.PersistentFlags().BoolP("debug", "d", false, "print debugging information")
	cmd.PersistentFlags().BoolP("yes", "y", false, "skip confirmation prompts")
	cmd.PersistentFlags().StringP("out", "o", "text", "output type: (text|json)")

	// allow the user to override the default configuration location if they
	// can find the flag. likely figured out from reading this comment, actually...
	cmd.PersistentFlags().StringP("config", "c", "", "path to config file (default is $HOME.config/relay)")
	cmd.PersistentFlags().MarkHidden("config")

	cmd.AddCommand(newAuthCommand())
	cmd.AddCommand(newWorkflowCommand())

	if err := cmd.Execute(); err != nil {
		Dialog.Error(format.Error(err, cmd))
		os.Exit(1)
	}
}
