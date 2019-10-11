package plugins

import (
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/envoyproxy/go-control-plane/pkg/conversion"
	"github.com/gogo/protobuf/proto"
)

func NewStagedFilter(name string, stage FilterStage) StagedHttpFilter {
	s, _ := NewStagedFilterWithConfig(name, nil, stage)
	return s
}

func NewStagedFilterWithConfig(name string, config proto.Message, stage FilterStage) (StagedHttpFilter, error) {

	s := StagedHttpFilter{
		HttpFilter: &envoyhttp.HttpFilter{
			Name: name,
		},
		Stage: stage,
	}

	if config != nil {
		marshalledConf, err := conversion.MessageToStruct(config)
		if err != nil {
			// this should NEVER HAPPEN!
			return StagedHttpFilter{}, err
		}

		s.HttpFilter.ConfigType = &envoyhttp.HttpFilter_Config{
			Config: marshalledConf,
		}
	}

	return s, nil
}
