package controller

import "github.com/pkg/errors"

var (
	FailedToListGatewayResources = func(err error, version, namespace string) error {
		return errors.Wrapf(err, "Failed to list %v gateway resources in %v", version, namespace)
	}

	FailedToDeleteGateway = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to delete %v gateway %v.%v", version, namespace)
	}

	FailedToWriteGateway = func(err error, version, namespace, name string) error {
		return errors.Wrapf(err, "Failed to write %v gateway %v.%v", version, namespace)
	}
)
