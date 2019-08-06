package add

import (
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)


func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "Adds configuration to a top-level Gloo resource.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cliutil.NoSubcommandError
		},
	}
	pflags := cmd.PersistentFlags()
	flagutils.AddOutputFlag(pflags, &opts.Top.Output)
	flagutils.AddMetadataFlags(pflags, &opts.Metadata)
	flagutils.AddDryRunFlag(pflags, &opts.Add.DryRun)
	cmd.AddCommand(Route(opts))
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}
