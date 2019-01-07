package kube2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"

	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

var _ = Describe("Kube2e", func() {

	BeforeEach(func() {

	})

	It("works", func() {

		cfg, err := kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())

		kube, err := kubernetes.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		kubeIngressClient := kube.ExtensionsV1beta1().Ingresses(namespace)
		backend := &v1beta1.IngressBackend{
			ServiceName: "foo",
			ServicePort: intstr.IntOrString{
				IntVal: 8080,
			},
		}
		kubeIng, err := kubeIngressClient.Create(&v1beta1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "rusty",
				Namespace: namespace,
			},
			Spec: v1beta1.IngressSpec{
				Backend: backend,
				TLS: []v1beta1.IngressTLS{
					{
						Hosts:      []string{"some.host"},
						SecretName: "doesntexistanyway",
					},
				},
				Rules: []v1beta1.IngressRule{
					{
						Host: "some.host",
						IngressRuleValue: v1beta1.IngressRuleValue{
							HTTP: &v1beta1.HTTPIngressRuleValue{
								Paths: []v1beta1.HTTPIngressPath{
									{
										Backend: *backend,
									},
								},
							},
						},
					},
				},
			},
		})
		log.Printf("%v", kubeIng)
	})
})
