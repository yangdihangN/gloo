package defaults

import (
	v1 "k8s.io/api/core/v1"
	"path/filepath"
)

var (
	GlooProxyValidationServerAddr = "gloo:9988"
	ValidationWebhookBindPort     = 443
	ValidationWebhookTlsCertPath  = filepath.Join("/etc", "gateway", "validation-certs", v1.TLSCertKey)
	ValidationWebhookTlsKeyPath   = filepath.Join("/etc", "gateway", "validation-certs", v1.TLSPrivateKeyKey)
)
