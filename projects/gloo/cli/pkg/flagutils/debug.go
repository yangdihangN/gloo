package flagutils

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/spf13/pflag"
)

func AddDebugFlags(set *pflag.FlagSet, debug *options.Debug) {
	set.StringVarP(&debug.HelmChartOverride, "file", "f", "", "Install Gloo from this Helm chart archive file rather than from a release")
	set.StringVarP(&debug.Namespace, "namespace", "n", defaults.GlooSystem, "namespace to install gloo into")
	set.StringVarP(&debug.ArchiveLocation, "output", "o", "", "output directory to save Gloo debug archive dump")
}