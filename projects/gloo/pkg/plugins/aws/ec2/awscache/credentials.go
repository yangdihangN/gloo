package awscache

import (
	"fmt"
	"strings"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws/glooec2"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// a credential spec represents an AWS client's view into AWS credentialMap
// we expect multiple upstreams to share the same view (so we batch the queries and apply filters locally)
type credentialSpec struct {
	// secretRef identifies the AWS secret that should be used to authenticate the client
	secretRef core.ResourceRef
	// region is the AWS region where our credentialMap live
	region string
	// roleArns are a list of AWS Roles (specified by their Amazon Resource Number (ARN)) which should be assumed when
	// querying for instances available to the upstream
	roleArns []string
}

// https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-ec2
const arnSegmentDelimiter = ":"

func (cs *credentialSpec) getKey() credentialKey {
	// use a very conservative "hash" strategy to avoid having to depend on aws's arn specification
	joinedArns := strings.Join(cs.roleArns, arnSegmentDelimiter)
	return credentialKey(fmt.Sprintf("%v-%v-%v", cs.secretRef.String(), cs.region, joinedArns))
}

func credentialSpecFromUpstreamSpec(ec2Spec *glooec2.UpstreamSpec) *credentialSpec {
	return &credentialSpec{
		secretRef: ec2Spec.SecretRef,
		region:    ec2Spec.Region,
		roleArns:  ec2Spec.RoleArns,
	}
}

// Since "==" is not defined for slices, slices (in particular, the roleArns slice) cannot be used as keys for go maps.
// Instead, we will use a string form. We give it a name for clarity.
type credentialKey string
