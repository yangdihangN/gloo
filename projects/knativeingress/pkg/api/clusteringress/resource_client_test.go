package clusteringress_test

import (
	"os"
	"time"

	"github.com/knative/serving/pkg/apis/networking/v1alpha1"
	knativeclientset "github.com/knative/serving/pkg/client/clientset/versioned"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/gloo/projects/knativeingress/pkg/api/clusteringress"
	"github.com/solo-io/gloo/projects/knativeingress/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
)

var _ = Describe("ResourceClient", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace string
		cfg       *rest.Config
	)

	BeforeEach(func() {
		namespace = helpers.RandString(8)
		err := setup.SetupKubeForTest(namespace)
		Expect(err).NotTo(HaveOccurred())
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		setup.TeardownKube(namespace)
	})

	It("can CRUD on v1beta1 ingresses", func() {
		knative, err := knativeclientset.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		baseClient := NewResourceClient(knative, &v1.ClusterIngress{})
		ingressClient := v1.NewClusterIngressClientWithBase(baseClient)
		Expect(err).NotTo(HaveOccurred())
		kubeIngressClient := knative.NetworkingV1alpha1().ClusterIngresses()
		kubeIng, err := kubeIngressClient.Create(&v1alpha1.ClusterIngress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "rusty",
				Namespace: namespace,
			},
			Spec: v1alpha1.IngressSpec{
				Rules: []v1alpha1.ClusterIngressRule{
					{
						Hosts: []string{
							"helloworld-go.default.example.com",
							"helloworld-go.default.svc.cluster.local",
							"helloworld-go.default.svc",
							"helloworld-go.default",
						},
						HTTP: &v1alpha1.HTTPClusterIngressRuleValue{
							Paths: []v1alpha1.HTTPClusterIngressPath{
								{
									AppendHeaders: map[string]string{
										"knative-serving-namespace": "default",
										"knative-serving-revision":  "helloworld-go-00001",
									},
									Retries: &v1alpha1.HTTPRetry{
										Attempts: 3,
										PerTryTimeout: &metav1.Duration{
											Duration: time.Minute,
										},
									},
									Splits: []v1alpha1.ClusterIngressBackendSplit{
										{
											Percent: 100,
											ClusterIngressBackend: v1alpha1.ClusterIngressBackend{
												ServiceName:      "activator-service",
												ServiceNamespace: "knative-serving",
												ServicePort:      intstr.IntOrString{IntVal: 80},
											},
										},
									},
									Timeout: &metav1.Duration{
										Duration: time.Minute,
									},
								},
							},
						},
					},
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
		ingressResource, err := ingressClient.Read(kubeIng.Namespace, kubeIng.Name, clients.ReadOpts{})
		Expect(err).NotTo(HaveOccurred())
		convertedIng, err := ToKube(ingressResource)
		Expect(err).NotTo(HaveOccurred())
		Expect(convertedIng.Spec).To(Equal(kubeIng.Spec))
	})
})
