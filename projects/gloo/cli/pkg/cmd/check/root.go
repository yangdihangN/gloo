package check

import (
	"fmt"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/spf13/cobra"
)

func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Checks for any configuration errors.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return check(opts)
		},
	}
	pflags := cmd.PersistentFlags()
	flagutils.AddMetadataFlags(pflags, &opts.Metadata)
	flagutils.AddFileFlag(cmd.LocalFlags(), &opts.Top.File)
	flagutils.AddCheckFlags(pflags, &opts.Check)
	cliutils.ApplyOptions(cmd, optionsFunc)
	cmd.AddCommand(
		upstreamCmd(opts),
	)
	return cmd
}

func upstreamCmd(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "upstream",
		Aliases: []string{"up"},
		Short:   "Check upstreams for any configuration errors.",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := checkUpstreams(opts, args)
			if err != nil {
				return err
			}
			fmt.Println(resp.String())
			return nil
		},
	}
	return cmd
}

func check(opts *options.Options) error {
	if opts.Top.File != "" {
		return checkFile(opts.Top.File)
	}
	return checkAllResources(opts)
}

func checkAllResources(opts *options.Options) error {
	// TODO - future PR
	// TODO: check gloo-related pods for errors
	// TODO: parse error logs for gloo
	// TODO: parse error logs for discovery
	// TODO: parse error logs for gateway
	// TODO: parse error logs for proxy
	// TODO: check virtual services for errors
	// TODO: check upstreams for errors
	// TODO: check proxy CRDs for errors
	return nil
}

func checkResource(ref core.Metadata) error {
	// TODO check a specific resource
	return nil
}

func checkFile(file string) error {
	// TODO[PRE-MERGE]
	return nil
}
