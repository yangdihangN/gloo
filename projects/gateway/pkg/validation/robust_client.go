package validation

import (
	"context"
	"sync"

	"github.com/avast/retry-go"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientConstructor func() (client validation.ProxyValidationServiceClient, e error)

// robustValidationClient wraps a validation.ProxyValidationServiceClient (grpc Connection)
// if a connection error occurs during an api call, the robustValidationClient
// attempts to reestablish the connection & retry the call before returning the error
type robustValidationClient struct {
	lock                      sync.RWMutex
	validationClient          validation.ProxyValidationServiceClient
	constructValidationClient func() (validation.ProxyValidationServiceClient, error)
}

// the constructor returned here is not threadsafe; call from a lock
func RetryOnUnavailableClientConstructor(ctx context.Context, serverAddress string) ClientConstructor {
	var cancel = func() {}
	return func() (client validation.ProxyValidationServiceClient, e error) {
		// cancel the previous client if it exists
		cancel()
		contextutils.LoggerFrom(ctx).Infow("starting proxy validation client... this may take a moment",
			zap.String("validation_server", serverAddress))
		var clientCtx context.Context
		clientCtx, cancel = context.WithCancel(ctx)

		cc, err := grpc.DialContext(clientCtx, serverAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to initialize grpc connection to validation server.")
		}

		return validation.NewProxyValidationServiceClient(cc), nil
	}
}

func NewRobustValidationClient(constructValidationClient func() (validation.ProxyValidationServiceClient, error)) (*robustValidationClient, error) {
	vc, err := constructValidationClient()
	if err != nil {
		return nil, err
	}
	return &robustValidationClient{
		constructValidationClient: constructValidationClient,
		validationClient:          vc,
	}, nil
}

func (c *robustValidationClient) ValidateProxy(ctx context.Context, proxy *validation.ProxyValidationServiceRequest, opts ...grpc.CallOption) (*validation.ProxyValidationServiceResponse, error) {
	var validationClient validation.ProxyValidationServiceClient
	var proxyReport *validation.ProxyValidationServiceResponse
	var reinstantiateClientErr error
	if err := retry.Do(func() error {
		c.lock.RLock()
		defer c.lock.RUnlock()
		validationClient = c.validationClient
		var err error
		proxyReport, err = validationClient.ValidateProxy(ctx, proxy, opts...)
		return err
	}, retry.RetryIf(func(e error) bool {
		if reinstantiateClientErr != nil {
			return false
		}
		return isUnavailableErr(e)
	}), retry.OnRetry(func(n uint, err error) {
		c.lock.Lock()
		defer c.lock.Unlock()
		// if someone already changed my client, do not replace it
		if validationClient == c.validationClient {
			c.validationClient, reinstantiateClientErr = c.constructValidationClient()
		}
	})); err != nil {
		return nil, err
	}
	return proxyReport, nil
}

func isUnavailableErr(err error) bool {
	if err == nil {
		return true
	}

	switch status.Code(err) {
	case codes.Unavailable, codes.FailedPrecondition, codes.Aborted:
		return true
	}
	return false
}
