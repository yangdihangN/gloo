package version

import (
	"fmt"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/solo-projects/pkg/cliutil"
	"github.com/spf13/cobra"
	"os/exec"
	"sort"
	"strings"
)

func VersionCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.VERSION_COMMAND.Use,
		Short: constants.VERSION_COMMAND.Short,
		Aliases: constants.VERSION_COMMAND.Aliases,
		RunE: func(cmd *cobra.Command, args []string) error {
				return runCommand(cmd, opts)
		},
	}
	cliutils.ApplyOptions(cmd, optionsFunc)
	flagutils.AddNamespaceFlag(cmd.Flags(), &opts.Metadata.Namespace)
	return cmd
}

func getClientVersionString() string {
	return fmt.Sprintf("glooctl community edition %s", version.Version)
}

func getPodVersionsString(namespace string) (string, error) {
	kubectl := exec.Command("kubectl", "get", "pods", "-n", namespace, "-l", "gloo", "-o", "jsonpath={..containers..image}")
	kubectl.Stdout = cliutil.Logger
	kubectl.Stderr = cliutil.Logger
	out, err := kubectl.Output()
	if err != nil {
		return "", err
	}
	images := strings.Split(string(out), " ")
	sort.Strings(images)
	return strings.Join(images, "\n"), nil
}

func runCommand(cmd *cobra.Command, opts *options.Options) error {
	_, err := fmt.Fprintln(cmd.OutOrStdout(), getClientVersionString())
	if err != nil {
		return err
	}
	if opts.Metadata.Namespace == "" {
		return fmt.Errorf("No namespace set, cannot look for server images.")
	}
	fmt.Fprintln(cmd.OutOrStdout(), "gloo server images for namespace", opts.Metadata.Namespace)
	images, err := getPodVersionsString(opts.Metadata.Namespace)
	if err != nil {
		return err
	}
	fmt.Fprintln(cmd.OutOrStdout(), images)
	return nil
}
