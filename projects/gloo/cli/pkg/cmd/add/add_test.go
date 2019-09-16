package add_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/add"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"

	"github.com/hashicorp/consul/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
	"github.com/solo-io/gloo/test/services"
)

var _ = Describe("Add", func() {
	if os.Getenv("RUN_CONSUL_TESTS") != "1" {
		log.Print("This test downloads and runs consul and is disabled by default. To enable, set RUN_CONSUL_TESTS=1 in your env.")
		return
	}

	var (
		consulFactory  *services.ConsulFactory
		consulInstance *services.ConsulInstance
		client         *api.Client
	)

	BeforeSuite(func() {
		var err error
		consulFactory, err = services.NewConsulFactory()
		Expect(err).NotTo(HaveOccurred())
		client, err = api.NewClient(api.DefaultConfig())
		Expect(err).NotTo(HaveOccurred())

	})

	AfterSuite(func() {
		_ = consulFactory.Clean()
	})

	BeforeEach(func() {
		helpers.UseDefaultClients()
		var err error
		// Start Consul
		consulInstance, err = consulFactory.NewConsulInstance()
		Expect(err).NotTo(HaveOccurred())
		err = consulInstance.Run()
		Expect(err).NotTo(HaveOccurred())
		// wait for consul to start
		Eventually(func() error {
			_, err := client.KV().Put(&api.KVPair{Key: "test"}, nil)
			return err
		}).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		if consulInstance != nil {
			err := consulInstance.Clean()
			Expect(err).NotTo(HaveOccurred())
		}
		helpers.UseDefaultClients()
	})

	Context("consul storage backend", func() {
		It("works", func() {
			err := testutils.Glooctl("add route " +
				"--path-prefix / " +
				"--dest-name petstore " +
				"--prefix-rewrite /api/pets " +
				"--use-consul")
			Expect(err).NotTo(HaveOccurred())
			kv, _, err := client.KV().Get("gloo/gateway.solo.io/v1/VirtualService/gloo-system/default", nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(kv).NotTo(BeNil())
		})
	})
})

func TestRootCmd(t *testing.T) {
	type args struct {
		opts        *options.Options
		optionsFunc []cliutils.OptionsFunc
	}
	tests := []struct {
		name string
		args args
		want *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add.RootCmd(tt.args.opts, tt.args.optionsFunc...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoute(t *testing.T) {
	type args struct {
		opts        *options.Options
		optionsFunc []cliutils.OptionsFunc
	}
	tests := []struct {
		name string
		args args
		want *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add.Route(tt.args.opts, tt.args.optionsFunc...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_actionFromInput(t *testing.T) {
	type args struct {
		input options.InputRoute
	}
	tests := []struct {
		name    string
		args    args
		want    *gatewayv1.Route_RouteAction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add.actionFromInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("actionFromInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actionFromInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addRoute(t *testing.T) {
	type args struct {
		opts *options.Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := add.addRoute(tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("addRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_destSpecFromInput(t *testing.T) {
	type args struct {
		input options.DestinationSpec
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.DestinationSpec
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add.destSpecFromInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("destSpecFromInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("destSpecFromInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matcherFromInput(t *testing.T) {
	type args struct {
		input options.RouteMatchers
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.Matcher
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add.matcherFromInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("matcherFromInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matcherFromInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pluginsFromInput(t *testing.T) {
	type args struct {
		input options.RoutePlugins
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.RoutePlugins
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add.pluginsFromInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("pluginsFromInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pluginsFromInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
