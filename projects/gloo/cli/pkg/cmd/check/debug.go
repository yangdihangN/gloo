package check

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/solo-io/go-utils/tarutils"
	"github.com/spf13/afero"

	"golang.org/x/sync/errgroup"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/go-utils/debugutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Filename = "/tmp/gloo-system-logs.tgz"
)

func DebugResources(opts *options.Options) error {
	responses, err := setup(opts)
	if err != nil {
		return err
	}

	fs := afero.NewOsFs()
	dir, err := afero.TempDir(fs, "", "")
	if err != nil {
		return err
	}
	defer fs.RemoveAll(dir)
	storageClient := debugutils.NewFileStorageClient(fs)
	if opts.Top.File == "" {
		opts.Top.File = Filename
	}

	for _, response := range responses {
		response := response
		scanner := bufio.NewScanner(response.Response)
		logs := ""
		for scanner.Scan() {
			line := scanner.Text()
			if opts.Top.ErrorsOnly {
				in := []byte(line)
				var raw map[string]interface{}
				if err = json.Unmarshal(in, &raw); err != nil {
					break
				}
				if raw["level"] == "error" {
					logs += line + "\n"
				}
			} else {
				logs += line + "\n"
			}
		}
		if logs != "" {
			if opts.Top.Zip {
				err = storageClient.Save(dir, &debugutils.StorageObject{
					Resource: strings.NewReader(logs),
					Name:     response.ResourceId(),
				})
				if err != nil {
					return err
				}
			} else {
				fmt.Printf("\n\n\nContainer name: %s\n", response.LogMeta.ContainerName)
				fmt.Printf("ResourceID: %s\n\n", response.ResourceId())
				fmt.Print(logs)
			}
		}

		response.Response.Close()
	}

	if opts.Top.Zip {
		tarball, err := fs.Create(opts.Top.File)
		if err != nil {
			return err
		}
		if err := tarutils.Tar(dir, fs, tarball); err != nil {
			return err
		}
	}

	return nil
}

func setup(opts *options.Options) ([]*debugutils.LogsResponse, error) {
	pods, err := helpers.MustKubeClient().CoreV1().Pods(opts.Metadata.Namespace).List(metav1.ListOptions{
		LabelSelector: "gloo",
	})
	if err != nil {
		return nil, err
	}
	resources, err := debugutils.ConvertPodsToUnstructured(pods)
	if err != nil {
		return nil, err
	}
	logCollector, err := debugutils.DefaultLogCollector()
	if err != nil {
		return nil, err
	}
	logRequests, err := logCollector.GetLogRequests(resources)

	return logCollector.LogRequestBuilder.StreamLogs(logRequests)
}

//func saveZip(fs afero.Fs)

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
