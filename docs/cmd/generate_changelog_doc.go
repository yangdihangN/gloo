package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"

	"github.com/solo-io/go-utils/errors"

	"github.com/solo-io/go-utils/changelogutils"
)

func main() {
	ctx := context.Background()
	app := rootApp(ctx)
	if err := app.Execute(); err != nil {
		fmt.Printf("unable to run: %v\n", err)
		os.Exit(1)
	}
}

type options struct {
	ctx              context.Context
	HugoDataSoloOpts HugoDataSoloOpts
}
type HugoDataSoloOpts struct {
	product string
	version string
	// if set, will override the version when rendering the
	callLatest bool
	noScope    bool
}

func rootApp(ctx context.Context) *cobra.Command {
	opts := &options{
		ctx: ctx,
	}
	app := &cobra.Command{
		Use: "docs-util",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	app.AddCommand(writeVersionScopeDataForHugo(opts))
	app.AddCommand(changelogMdCmd(opts))
	return app
}
func changelogMdCmd(opts *options) *cobra.Command {
	app := &cobra.Command{
		Use:   "gen-changelog-md",
		Short: "generate a markdown file from changelogs",
		RunE: func(cmd *cobra.Command, args []string) error {

			return generateChangelogMd(opts, args)
			return nil
		},
	}
	app.PersistentFlags().StringVar(&opts.HugoDataSoloOpts.version, "version", "", "version of docs and code")
	app.PersistentFlags().StringVar(&opts.HugoDataSoloOpts.product, "product", "gloo", "product to which the docs refer (defaults to gloo)")
	app.PersistentFlags().BoolVar(&opts.HugoDataSoloOpts.noScope, "no-scope", false, "if set, will not nest the served docs by product or version")
	app.PersistentFlags().BoolVar(&opts.HugoDataSoloOpts.callLatest, "call-latest", false, "if set, will use the string 'latest' in the scope, rather than the particular release version")
	return app
}

type HugoDataSoloYaml struct {
	DocsVersion string `yaml:"DocsVersion"`
}

const hugoDataSoloFilename = "data/Solo.yaml"

func writeVersionScopeDataForHugo(opts *options) *cobra.Command {
	app := &cobra.Command{
		Use:   "gen-version-scope-data",
		Short: "generate a data file for Hugo that indicates the docs version",
		RunE: func(cmd *cobra.Command, args []string) error {
			data := &HugoDataSoloYaml{}
			var err error
			data.DocsVersion, err = getDocsVersionFromOpts(opts.HugoDataSoloOpts)
			if err != nil {
				return err
			}
			marshalled, err := yaml.Marshal(data)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(hugoDataSoloFilename, marshalled, 0x644)
		},
	}
	return app
}

func getDocsVersionFromOpts(hugoOpts HugoDataSoloOpts) (string, error) {
	if hugoOpts.noScope {
		return "", nil
	}
	if hugoOpts.version == "" || hugoOpts.product == "" {
		return "", errors.New("must provide a version and product for scoped docs generation")
	}
	return fmt.Sprintf("/%v/%v", hugoOpts.product, hugoOpts.version), nil
}

const (
	glooDocGen  = "gloo"
	glooEDocGen = "glooe"
)

var (
	InvalidInputError = func(arg string) error {
		return errors.Errorf("invalid input, must provide exactly one argument, either '%v' or '%v', (provided %v)",
			glooDocGen,
			glooEDocGen,
			arg)
	}
)

func generateChangelogMd(opts *options, args []string) error {
	if len(args) != 2 {
		return InvalidInputError(fmt.Sprintf("%v", len(args)-1))
	}
	target := args[1]

	var repoRootPath, repo, changelogDirPath string
	switch target {
	case glooDocGen:
		repoRootPath = ".."
		repo = "gloo"
		changelogDirPath = "../changelog"
	case glooEDocGen:
		repoRootPath = "../../solo-projects"
		repo = "solo-projects"
		changelogDirPath = "../../solo-projects/changelog"
	default:
		return InvalidInputError(target)
	}

	// consider writing to stdout to enhance makefile/io readability `go run cmd/main.go > changelogSummary.md`
	owner := "solo-io"
	w := os.Stdout
	err := changelogutils.GenerateChangelogFromLocalDirectory(opts.ctx, repoRootPath, owner, repo, changelogDirPath, w)
	return err
}
