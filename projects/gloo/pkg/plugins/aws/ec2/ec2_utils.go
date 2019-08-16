package ec2

import (
	"context"
	"fmt"
	"strings"

	"github.com/solo-io/go-utils/errors"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	glooec2 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws/ec2"
	aws2 "github.com/solo-io/gloo/projects/gloo/pkg/utils/aws"
)

func getEc2SessionForCredentials(regionConfig *aws.Config, secretRef core.ResourceRef, secrets v1.SecretList) (*session.Session, error) {
	return aws2.GetAwsSession(
		secretRef,
		secrets,
		regionConfig,
	)
}

func GetEc2Client(cred *CredentialSpec, secrets v1.SecretList) (*ec2.EC2, error) {
	var sess *session.Session
	var err error
	regionConfig := &aws.Config{Region: aws.String(cred.Region())}
	secretRef := cred.SecretRef()
	if secretRef == nil {
		sess, err = session.NewSession(regionConfig)
		if err != nil {
			return nil, CreateSessionFromEnvError(err)
		}
	} else {
		sess, err = getEc2SessionForCredentials(regionConfig, *secretRef, secrets)
		if err != nil {
			return nil, CreateSessionFromSecretError(err)
		}
	}
	if cred.Arn() != "" {
		cred := stscreds.NewCredentials(sess, cred.Arn())
		config := &aws.Config{Credentials: cred}
		return ec2.New(sess, config), nil
	}
	return ec2.New(sess), nil
}

func GetInstancesFromDescription(desc *ec2.DescribeInstancesOutput) []*ec2.Instance {
	var instances []*ec2.Instance
	for _, reservation := range desc.Reservations {
		for _, instance := range reservation.Instances {
			if validInstance(instance) {
				instances = append(instances, instance)
			}
		}
	}
	return instances
}

// this filter function defines what gloo considers a valid EC2 instance
func validInstance(instance *ec2.Instance) bool {
	if instance.PublicIpAddress != nil {
		return true
	}
	if instance.PrivateIpAddress != nil {
		return true
	}
	return false
}

// generate an ec2 filter spec for a given upstream.
// not currently used since we are batching API calls by credentials, without filters
func convertFiltersFromSpec(upstreamSpec *glooec2.UpstreamSpec) []*ec2.Filter {
	var filters []*ec2.Filter
	for _, filterSpec := range upstreamSpec.Filters {
		var currentFilter *ec2.Filter
		switch x := filterSpec.Spec.(type) {
		case *glooec2.TagFilter_Key:
			currentFilter = tagFiltersKey(x.Key)
		case *glooec2.TagFilter_KvPair_:
			currentFilter = tagFiltersKeyValue(x.KvPair.Key, x.KvPair.Value)
		}
		filters = append(filters, currentFilter)
	}
	return filters
}

// EC2 Describe Instance filters expect a particular key format:
//   https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html
//   tag:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the
//   tag value as the filter value. For example, to find all credentialMap that have a tag with the key Owner and the value
//   TeamA, specify tag:Owner for the filter name and TeamA for the filter value.
func tagFilterName(tagName string) *string {
	str := fmt.Sprintf("tag:%v", tagName)
	return &str
}

func tagFilterValue(tagValue string) []*string {
	if tagValue == "" {
		return nil
	}
	return []*string{&tagValue}
}

// Helper for getting a filter that selects all instances that have a given tag and tag-value pair
func tagFiltersKeyValue(tagName, tagValue string) *ec2.Filter {
	return &ec2.Filter{
		Name:   tagFilterName(tagName),
		Values: tagFilterValue(tagValue),
	}
}

/*
NOTE on EC2
How to find all instances that have a given tag-key, regardless of the tag value:
  https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html
  tag-key - The key of a tag assigned to the resource. Use this filter to find all credentialMap that have a tag with a
  specific key, regardless of the tag value.
*/
// generate a filter that selects all elements that contain a given tag
func tagFiltersKey(tagName string) *ec2.Filter {
	return &ec2.Filter{
		Name:   aws.String("tag-key"),
		Values: []*string{aws.String(tagName)},
	}
}

var (
	CreateSessionFromEnvError = func(err error) error {
		return errors.Wrapf(err, "unable to create a session with credentials taken from env")
	}

	CreateSessionFromSecretError = func(err error) error {
		return errors.Wrapf(err, "unable to create a session with credentials taken from secret ref")
	}
)

func InstancesForUpstream(ctx context.Context, upstream *v1.Upstream, secrets v1.SecretList) ([]*ec2.Instance, []*ec2.Instance, error) {
	ec2Spec := upstream.UpstreamSpec.GetAwsEc2()
	if ec2Spec == nil {
		return nil, nil, nil
	}
	cred := NewCredentialSpecFromEc2UpstreamSpec(ec2Spec)
	client, err := GetEc2Client(cred, secrets)
	if err != nil {
		return nil, nil, GetClientError(err)
	}
	e := NewEc2InstanceLister()
	all, err := e.ListWithClient(context.Background(), client)
	if err != nil {
		return nil, nil, err
	}
	credGroups, err := getCredGroupsFromUpstreams(v1.UpstreamList{upstream})
	if err != nil {
		return nil, nil, err
	}
	if err := getInstancesForCredentialGroups(ctx, e, secrets, credGroups); err != nil {
		return nil, nil, err
	}
	credGroup := credGroups[cred.GetKey()]
	matchedInstances := filterInstancesForUpstream(ctx, upstream, credGroup)
	matchedMap := make(map[string]bool)
	for _, inst := range matchedInstances {
		matchedMap[aws.StringValue(inst.InstanceId)] = true
	}
	var unmatchedInstances []*ec2.Instance
	for _, inst := range all {
		if _, ok := matchedMap[aws.StringValue(inst.InstanceId)]; !ok {
			unmatchedInstances = append(unmatchedInstances, inst)
		}
	}

	return matchedInstances, unmatchedInstances, nil
}

func SummarizeInstances(matched, unmatched []*ec2.Instance) string {
	summary := fmt.Sprintf("filters matched %v of %v instances for credentials:\n", len(matched), len(unmatched))
	var rows []string

	printInstances := func(title string, list []*ec2.Instance) {
		if len(list) == 0 {
			return
		}
		rows = append(rows, title)
		for _, inst := range list {
			nameContent := ""
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					nameContent = fmt.Sprintf(" (%v)", *tag.Value)
				}
			}
			rows = append(rows, fmt.Sprintf("%v%v", aws.StringValue(inst.InstanceId), nameContent))
		}
	}
	printInstances("instances matching credentials and filters:", matched)
	printInstances("instances matching credentials only:", unmatched)
	summary += strings.Join(rows, "\n")
	return summary
}
