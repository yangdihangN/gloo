package main

import (
	"fmt"
	"log"
	"time"

	"github.com/solo-io/go-utils/kubeutils"

	"github.com/gogo/protobuf/types"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func main() {
	str := kubeutils.SanitizeNameV2("gloo-_nomad-task-c6fae9f9-5ed8-1040-1026-b547433195a5-gloo-gloo")

	log.Print(str)

	_, err := writeSettings("/data", "gloo-system")
	if err != nil {
		panic(err)
	}
}

func writeSettings(dataDir, writeNamespace string) (*gloov1.Settings, error) {
	settings := &gloov1.Settings{
		ConfigSource: &gloov1.Settings_ConsulKvSource{
			ConsulKvSource: &gloov1.Settings_ConsulKv{},
		},
		SecretSource: &gloov1.Settings_DirectorySecretSource{
			DirectorySecretSource: &gloov1.Settings_Directory{
				Directory: dataDir,
			},
		},
		//SecretSource: &gloov1.Settings_VaultSecretSource{
		//	VaultSecretSource: &gloov1.Settings_VaultSecrets{
		//		Address: "http://host.docker.internal:8200",
		//		Token:   "root",
		//	},
		//},
		ArtifactSource: &gloov1.Settings_DirectoryArtifactSource{
			DirectoryArtifactSource: &gloov1.Settings_Directory{
				Directory: dataDir,
			},
		},
		Consul: &gloov1.Settings_ConsulConfiguration{
			Address:          "host.docker.internal:8500",
			ServiceDiscovery: &gloov1.Settings_ConsulConfiguration_ServiceDiscoveryOptions{},
		},
		BindAddr:           "0.0.0.0:9977",
		RefreshRate:        types.DurationProto(time.Minute),
		DiscoveryNamespace: writeNamespace,
		Metadata:           core.Metadata{Namespace: writeNamespace, Name: "default"},
	}
	yam, err := protoutils.MarshalYAML(settings)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(yam))
	return settings, nil
}
