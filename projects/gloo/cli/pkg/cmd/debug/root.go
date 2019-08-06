package debug

import (
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

const (
	gateway = "gateway"
	ingress = "ingress"
	knative = "knative"
)
var (
	validArgs = []string{gateway, ingress, knative}
)

func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:     constants.DEBUG_COMMAND.Use,
		Aliases: constants.DEBUG_COMMAND.Aliases,
		Short:   constants.DEBUG_COMMAND.Short,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cliutil.NoSubcommandError
		},
	}
	pflags := cmd.PersistentFlags()
	flagutils.AddMetadataFlags(pflags, &opts.Metadata)

	cmd.AddCommand(Archive(opts))
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}
