package check

import (
	"path/filepath"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/go-utils/debugutils"
	"github.com/solo-io/go-utils/tarutils"
	"github.com/spf13/afero"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Filename = "/tmp/gloo-system-logs.tgz"
)

func DebugResources(opts *options.Options) error {
	if opts.Top.File == "" {
		opts.Top.File = Filename
	}

	pods, err := helpers.MustKubeClient().CoreV1().Pods(opts.Metadata.Namespace).List(metav1.ListOptions{
		LabelSelector: "gloo",
	})
	if err != nil {
		return err
	}
	resources, err := debugutils.ConvertPodsToUnstructured(pods)
	if err != nil {
		return err
	}
	logCollector, err := debugutils.DefaultLogCollector()
	if err != nil {
		return err
	}
	logRequests, err := logCollector.GetLogRequests(resources)
	if err != nil {
		return err
	}

	fs := afero.NewOsFs()
	dir, err := afero.TempDir(fs, "", "")
	if err != nil {
		return err
	}
	defer fs.RemoveAll(dir)

	// Request the logs and save them
	storageClient := debugutils.NewFileStorageClient(fs)
	if err = logCollector.SaveLogs(storageClient, dir, logRequests); err != nil {
		return err
	}

	// Tar the logs
	tarball, err := afero.TempFile(fs, "", "")
	defer fs.Remove(tarball.Name())
	if err != nil {
		return err
	}
	if err := tarutils.Tar(dir, fs, tarball); err != nil {
		return err
	}

	if err := storageClient.Save(filepath.Dir(opts.Top.File), &debugutils.StorageObject{
		Name:     filepath.Base(opts.Top.File),
		Resource: tarball,
	}); err != nil {
		return err
	}

	return nil
}
