package controller

import "github.com/pkg/errors"

var (
	FailedToListGatewayResources = func(err error, version, namespace string) error {
		return errors.Wrapf(err, "Failed to list %v gateway resources in %v", version, namespace)
	}
)
