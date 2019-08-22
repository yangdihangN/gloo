package main

import (
	"github.com/solo-io/go-utils/log"
	"os"
)

const (
	i
)

func main() {
	var version, repoPrefixOverride, globalPullPolicy string
	if len(os.Args) < 2 {
		panic("Must provide version as argument")
	} else {
		version = os.Args[1]

		if len(os.Args) >= 3 {
			repoPrefixOverride = os.Args[2]
		}
		if len(os.Args) >= 4 {
			globalPullPolicy = os.Args[3]
		}
	}

	log.Printf("Generating Levant Variables.")
	if err := generateLevantVariables(version, repoPrefixOverride, globalPullPolicy); err != nil {
		log.Fatalf("generating variables.yaml failed!: %v", err)
	}
}


type Variables struct {
	Datacenter    string `json:"datacenter"`
	Region        string `json:"region"`
	Config        Config `json:"config"`
	Consul        Consul `json:"consul"`
	Vault         Vault  `json:"vault"`
	DockerNetwork string `json:"dockerNetwork"`
	Gloo          Gloo   `json:"gloo"`
	Discovery     Task   `json:"discovery"`
	Gateway       Task   `json:"gateway"`
	GatewayProxy  Task   `json:"gatewayProxy"`
}

type Config struct {
	Namespace   string `json:"namespace"`
	RefreshRate string `json:"refreshRate"`
}

type Consul struct {
	Address string `json:"address"`
}

type Vault struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}

type Gloo struct {
	XdsPort int `json:"xdsPort"`
	Task
}

type Task struct {
	Image          Image `json:"image"`
	CPULimit       int   `json:"cpuLimit"`
	MemLimit       int   `json:"memLimit"`
	BandwidthLimit int   `json:"bandwidthLimit"`
	Replicas       int   `json:"replicas"`
}

type Image struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}
