package upstreamconn

import (
	"math"
	"time"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/solo-io/gloo/pkg/utils/gogoutils"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

type Plugin struct{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *Plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {

	cfg := in.GetUpstreamSpec().GetConnectionConfig()
	if cfg == nil {
		return nil
	}

	if cfg.MaxRequestsPerConnection > 0 {
		out.MaxRequestsPerConnection = &wrappers.UInt32Value{
			Value: cfg.MaxRequestsPerConnection,
		}
	}

	if cfg.ConnectTimeout != nil {
		out.ConnectTimeout = gogoutils.DurationStdToProto(cfg.ConnectTimeout)
	}

	if cfg.TcpKeepalive != nil {
		out.UpstreamConnectionOptions = &envoyapi.UpstreamConnectionOptions{
			TcpKeepalive: convertTcpKeepAlive(cfg.TcpKeepalive),
		}
	}

	return nil
}

func convertTcpKeepAlive(tcp *v1.ConnectionConfig_TcpKeepAlive) *envoycore.TcpKeepalive {
	var probes *wrappers.UInt32Value
	if tcp.KeepaliveProbes > 0 {
		probes = &wrappers.UInt32Value{
			Value: tcp.KeepaliveProbes,
		}
	}
	return &envoycore.TcpKeepalive{
		KeepaliveInterval: roundToSecond(tcp.KeepaliveInterval),
		KeepaliveTime:     roundToSecond(tcp.KeepaliveTime),
		KeepaliveProbes:   probes,
	}
}

func roundToSecond(d *time.Duration) *wrappers.UInt32Value {
	if d == nil {
		return nil
	}

	// round up
	seconds := math.Round(d.Seconds() + 0.4999)
	return &wrappers.UInt32Value{
		Value: uint32(seconds),
	}

}
