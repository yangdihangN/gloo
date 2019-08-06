package debug

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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

const (
	archiveFileName = "gloo-debug-archive.tar.gz"
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
	flagutils.AddDebugFlags(pflags, &opts.Debug)
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
			Namespace: opts.Debug.Namespace,
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

	var (
		outputFileName string
	)

	if opts.Debug.ArchiveLocation == "" {
		tmp, err := ioutil.TempFile("", "*.tar.gz")
		if err != nil {
			return err
		}
		outputFileName = tmp.Name()

		fmt.Printf("No output location was specified, so the archive is being stored to a temporary file here: %s", tmp.Name())
	} else {
		file, err := os.OpenFile(filepath.Join(opts.Debug.ArchiveLocation, archiveFileName), os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			return err
		}
		outputFileName = file.Name()

		fmt.Printf("Gloo debug archive has been saved to the following location: %s", file.Name())
	}


	if err := aggregator.StreamFromManifest(sortedManifests, opts.Debug.Namespace, outputFileName); err != nil {
		return err
	}
	return nil
}
