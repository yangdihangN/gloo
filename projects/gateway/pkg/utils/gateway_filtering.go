package utils

import (
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"

	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
)

func GatewaysByProxyName(gateways v2.GatewayList) map[string]v2.GatewayList {
	result := make(map[string]v2.GatewayList)
	for _, gw := range gateways {
		proxyNames := gw.ProxyNames
		if len(proxyNames) == 0 {
			proxyNames = []string{defaults.GatewayProxyName}
		}
		for _, name := range proxyNames {
			result[name] = append(result[name], gw)
		}
	}
	return result
}
