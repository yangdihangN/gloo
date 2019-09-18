package kube

import (
	"context"
	"encoding/base64"

	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type TlsSecret struct {
	SecretName, SecretNamespace string
	PrivateKeyKey, CaCertKey    string
	PrivateKey, CaCert          []byte
}

func CreateTlsSecret(ctx context.Context, kube kubernetes.Interface, secretCfg TlsSecret) error {
	secret := makeTlsSecret(secretCfg)

	contextutils.LoggerFrom(ctx).Infow("creating TLS secret", zap.String("secret", secret.Name))

	_, err := kube.CoreV1().Secrets(secretCfg.SecretNamespace).Create(secret)
	return err
}

func makeTlsSecret(args TlsSecret) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      args.SecretName,
			Namespace: args.SecretNamespace,
		},
		Type: v1.SecretTypeTLS,
		Data: map[string][]byte{
			args.PrivateKeyKey: []byte(base64.StdEncoding.EncodeToString(args.PrivateKey)),
			args.CaCertKey:     []byte(base64.StdEncoding.EncodeToString(args.CaCert)),
		},
	}
}
