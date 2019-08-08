package flagutils

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/spf13/pflag"
)

func AddCheckFlags(set *pflag.FlagSet, checkOpts *options.Check) {
	set.BoolVarP(&checkOpts.AllNamespaces, "all-namespaces", "A", false, "check for resources in all namespaces")
	set.StringSliceVar(&checkOpts.NamespaceList, "namespaces", []string{}, "check for resources in namespace list")
}
