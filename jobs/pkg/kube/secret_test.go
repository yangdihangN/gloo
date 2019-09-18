package kube_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	. "github.com/solo-io/gloo/jobs/pkg/kube"
)

var _ = Describe("Secret", func() {
	It("creates a tls secret from the provided certs", func() {
		kube := fake.NewSimpleClientset()
		secretCfg := TlsSecret{
			SecretName:      "mysecret",
			SecretNamespace: "mynamespace",
			PrivateKeyKey:   "key.pem",
			CaCertKey:       "ca.pem",
			PrivateKey:      []byte{1, 2, 3},
			CaCert:          []byte{1, 2, 3},
		}

		err := CreateTlsSecret(context.TODO(), kube, secretCfg)
		Expect(err).NotTo(HaveOccurred())

		secret, err := kube.CoreV1().Secrets(secretCfg.SecretNamespace).Get(secretCfg.SecretName, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(secret).To(Equal(&v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "mysecret",
				Namespace: "mynamespace",
			},
			Data: map[string][]byte{"key.pem": []byte("AQID"), "ca.pem": []byte("AQID")},
			Type: "kubernetes.io/tls",
		}))
	})
})
