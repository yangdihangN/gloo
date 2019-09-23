package skutils

import (
	"encoding/json"

	"github.com/solo-io/go-utils/errors"
	v1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
)

// convert raw Kube JSON to a Solo-Kit resource
func UnmarshalResource(kubeJson []byte, resource resources.Resource) error {
	var resourceCrd v1.Resource
	if err := json.Unmarshal(kubeJson, &resourceCrd); err != nil {
		return errors.Wrapf(err, "unmarshalling from raw json")
	}
	resource.SetMetadata(kubeutils.FromKubeMeta(resourceCrd.ObjectMeta))
	if withStatus, ok := resource.(resources.InputResource); ok {
		resources.UpdateStatus(withStatus, func(status *core.Status) {
			*status = resourceCrd.Status
		})
	}

	if resourceCrd.Spec != nil {
		if err := protoutils.UnmarshalMap(*resourceCrd.Spec, resource); err != nil {
			return errors.Wrapf(err, "parsing resource from crd spec %v in namespace %v into %T", resourceCrd.Name, resourceCrd.Namespace, resource)
		}
	}

	return nil
}
