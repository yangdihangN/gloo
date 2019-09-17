package check

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

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
	// Setup
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

	// Save the logs in a tar file
	fs := afero.NewOsFs()
	dir, err := afero.TempDir(fs, "", "")
	if err != nil {
		return err
	}
	defer fs.RemoveAll(dir)
	storageClient := debugutils.NewFileStorageClient(fs)
	if err = logCollector.SaveLogs(storageClient, dir, logRequests); err != nil {
		return err
	}
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

	// Print out the error logs
	responses, err := StreamLogs(logRequests)
	if err != nil {
		return err
	}
	for _, response := range responses {
		response := response
		scanner := bufio.NewScanner(response.Response)
		errorLogs := ""
		for scanner.Scan() {
			if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower("error")) {
				errorLogs += scanner.Text() + "\n"
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		if errorLogs != "" {
			fmt.Printf("Container name: %s\n", response.LogMeta.ContainerName)
			fmt.Printf("ResourceID: %s\n", response.ResourceId())
			fmt.Printf("%s\n", errorLogs)
		}
		response.Response.Close()
	}

	return nil
}

func StreamLogs(requests []*debugutils.LogsRequest) ([]*debugutils.LogsResponse, error) {
	result := make([]*debugutils.LogsResponse, 0, len(requests))
	eg := errgroup.Group{}
	lock := sync.Mutex{}
	for _, request := range requests {
		// necessary to shadow this variable so that it is unique within the goroutine
		restRequest := request
		eg.Go(func() error {
			reader, err := restRequest.Request.Stream()
			if err != nil {
				return err
			}
			lock.Lock()
			defer lock.Unlock()
			result = append(result, &debugutils.LogsResponse{
				LogMeta: debugutils.LogMeta{
					PodMeta:       restRequest.PodMeta,
					ContainerName: restRequest.ContainerName,
				},
				Response: reader,
			})
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return result, nil
}
