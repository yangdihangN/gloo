package debug

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	installutil "github.com/solo-io/gloo/pkg/cliutil/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/go-utils/debugutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/installutils/helmchart"
	"github.com/spf13/cobra"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/manifest"
	"k8s.io/helm/pkg/renderutil"
	"k8s.io/helm/pkg/tiller"
)

func Archive(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "archive",
		Aliases: []string{"a"},
		Short:   "Zip all Gloo resources",
		Long: "Create an archive of all Gloo relevant gloo resources and logs" +
			"to send along for debugging purposes",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return archive(opts, args)
		},

		ValidArgs: validArgs,
	}

	pflags := cmd.PersistentFlags()
	flagutils.AddInstallFlags(pflags, &opts.Install)
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}

func archive(opts *options.Options, args []string) error {
	installType := args[0]
	var valuesFileName string
	switch installType {
	case ingress:
		valuesFileName = constants.IngressValuesFileName
	case knative:
		valuesFileName = constants.KnativeValuesFileName
	case gateway:
		valuesFileName = constants.GatewayValuesFileName
	default:
		return errors.Errorf("install type %s provided, should be impossible", installType)
	}

	spec, err := install.GetInstallSpec(opts, valuesFileName)
	if err != nil {
		return err
	}

	if path.Ext(spec.HelmArchiveUri) != ".tgz" && !strings.HasSuffix(spec.HelmArchiveUri, ".tar.gz") {
		return errors.Errorf("unsupported file extension for Helm chart URI: [%s]. Extension must either be .tgz or .tar.gz", spec.HelmArchiveUri)
	}

	chart, err := installutil.GetHelmArchive(spec.HelmArchiveUri)
	if err != nil {
		return errors.Wrapf(err, "retrieving gloo helm chart archive")
	}

	values, err := installutil.GetValuesFromFileIncludingExtra(chart, spec.ValueFileName, spec.ExtraValues, spec.ValueCallbacks...)
	if err != nil {
		return errors.Wrapf(err, "retrieving value file: %s", spec.ValueFileName)
	}

	// These are the .Release.* variables used during rendering
	renderOpts := renderutil.Options{
		ReleaseOptions: chartutil.ReleaseOptions{
			Namespace: opts.Install.Namespace,
			Name:      spec.ProductName,
		},
	}

	renderedTemplates, err := renderutil.Render(chart, values, renderOpts)
	if err != nil {
		return err
	}

	for file, man := range renderedTemplates {
		if helmchart.IsEmptyManifest(man) {
			delete(renderedTemplates, file)
		}
	}
	sortedManifests := helmchart.Manifests(tiller.SortByKind(manifest.SplitManifests(renderedTemplates)))

	aggregator, err := debugutils.DefaultAggregator()
	if err != nil {
		return err
	}

	tmp, err := ioutil.TempFile("", "*.tar.gz")
	if err != nil {
		return err
	}

	fmt.Println(tmp.Name())

	if err := aggregator.StreamFromManifest(sortedManifests, "gloo-system", tmp.Name()); err != nil {
		return err
	}
	return nil
}
