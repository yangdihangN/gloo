package validation

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v2"
	"github.com/solo-io/gloo/projects/gateway/pkg/translator"
	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

type Validator interface {
	v2.ApiSyncer
	Ready() bool
	ValidateGateway(ctx context.Context, gw *v2.Gateway) error
	ValidateVirtualService(ctx context.Context, vs *v1.VirtualService) error
	ValidateRouteTable(ctx context.Context, rt *v1.RouteTable) error
	ValidateDeleteRouteTable(ctx context.Context, rt core.ResourceRef) error
	ValidateDeleteVirtualService(ctx context.Context, vs core.ResourceRef) error
}

type validator struct {
	l                sync.RWMutex
	latestSnapshot   *v2.ApiSnapshot
	translator       translator.Translator
	validationClient validation.ProxyValidationServiceClient
	writeNamespace   string
}

func NewValidator(translator translator.Translator, validationClient validation.ProxyValidationServiceClient, writeNamespace string) *validator {
	return &validator{translator: translator, validationClient: validationClient, writeNamespace: writeNamespace}
}

func (v *validator) Ready() bool {
	return v.latestSnapshot != nil
}

func (v *validator) Sync(_ context.Context, snap *v2.ApiSnapshot) error {
	snapCopy := snap.Clone()
	v.l.Lock()
	v.latestSnapshot = &snapCopy
	v.l.Unlock()
	return nil
}

func (v *validator) validateSnapshot(ctx context.Context, snap *v2.ApiSnapshot, proxyNames []string) error {
	if !v.Ready() {
		return errors.Errorf("validation is yet not available. Waiting for first snapshot")
	}
	gatewaysByProxy := utils.GatewaysByProxyName(snap.Gateways)

	for _, proxyName := range proxyNames {
		gatewayList := gatewaysByProxy[proxyName]
		proxy, resourceErrs := v.translator.Translate(ctx, proxyName, v.writeNamespace, snap, gatewayList)
		if err := resourceErrs.Validate(); err != nil {
			return errors.Wrapf(err, "could not render proxy")
		}

		if v.validationClient == nil {
			contextutils.LoggerFrom(ctx).Warnf("skipping proxy validation check as the " +
				"Proxy validation client has not been initialized. check to ensure that the gateway and gloo processes " +
				"are configured to communicate.")
			return nil
		}

		// validate the proxy with gloo
		proxyReport, err := v.validationClient.ValidateProxy(ctx, proxy)
		if err != nil {
			contextutils.LoggerFrom(ctx).Errorw("failed to validate Proxy with Gloo validation server.", zap.Error(err))
			return errors.Wrapf(err, "failed to validate Proxy with Gloo validation server")
		}

		if proxyErr := validationutils.GetProxyError(proxyReport); proxyErr != nil {
			return errors.Wrapf(proxyErr, "rendered proxy had errors")
		}
	}
	return nil
}

func (v *validator) ValidateVirtualService(ctx context.Context, vs *v1.VirtualService) error {
	if !v.Ready() {
		return errors.Errorf("Gateway validation is yet not available. Waiting for first snapshot")
	}
	v.l.RLock()
	snap := v.latestSnapshot.Clone()
	v.l.RUnlock()

	vsRef := vs.GetMetadata().Ref()

	// TODO: move this to a function when generics become a thing
	var isUpdate bool
	for i, existingVs := range snap.VirtualServices {
		if vsRef == existingVs.GetMetadata().Ref() {
			// replace the existing virtual service in the snapshot
			snap.VirtualServices[i] = vs
			isUpdate = true
			break
		}
	}
	if !isUpdate {
		snap.VirtualServices = append(snap.VirtualServices, vs)
		snap.VirtualServices.Sort()
	}

	proxiesToConsider := proxiesForVirtualService(snap.Gateways, vs)

	if err := v.validateSnapshot(ctx, &snap, proxiesToConsider); err != nil {
		contextutils.LoggerFrom(ctx).Debugw("Rejected %T %v: %v", vs, vsRef, err)
		return errors.Wrapf(err, "validating %T %v", vs, vsRef)
	}

	contextutils.LoggerFrom(ctx).Debugw("Accepted %T %v", vs, vsRef)

	return nil
}

func (v *validator) ValidateRouteTable(ctx context.Context, rt *v1.RouteTable) error {
	panic("implement me")
}

func (v *validator) ValidateDeleteRouteTable(ctx context.Context, rt core.ResourceRef) error {
	panic("implement me")
}

func (v *validator) ValidateGateway(ctx context.Context, gw *v2.Gateway) error {
	panic("implement me")
}

func (v *validator) ValidateDeleteVirtualService(ctx context.Context, vs core.ResourceRef) error {
	panic("implement me")
}

func proxiesForVirtualService(gwList v2.GatewayList, vs *v1.VirtualService) []string {

	gatewaysByProxy := utils.GatewaysByProxyName(gwList)

	var proxiesToConsider []string

	for proxyName, gatewayList := range gatewaysByProxy {
		if gatewayListContainsVirtualService(gatewayList, vs) {
			// we only care about validating this proxy if it contains this virtual service
			proxiesToConsider = append(proxiesToConsider, proxyName)
		}
	}

	return proxiesToConsider
}

func virtualServicesForRouteTable(rt *v1.RouteTable, allVirtualServices v1.VirtualServiceList, allRoutes v1.RouteTableList) v1.VirtualServiceList {
	// this route table + its parents
	refsContainingRouteTable := []core.ResourceRef{rt.Metadata.Ref()}

	// keep going until the ref list stops expanding
	for countedRefs := 0; countedRefs == len(refsContainingRouteTable); countedRefs = len(refsContainingRouteTable) {
		for _, rt := range allRoutes {
			if routesContainRefs(rt.Routes, refsContainingRouteTable...) {
				refsContainingRouteTable = append(refsContainingRouteTable, rt.Metadata.Ref())
			}
		}
	}
	return nil
}

func routesContainRefs(list []*v1.Route, refs ...core.ResourceRef) bool {
	for _, r := range list {
		delegate := r.GetDelegateAction()
		if delegate == nil {
			return false
		}
		for _, ref := range refs {
			if *delegate == ref {
				return true
			}
		}
	}
	return false
}

func gatewayListContainsVirtualService(gwList v2.GatewayList, vs *v1.VirtualService) bool {
	for _, gw := range gwList {
		if translator.GatewayContainsVirtualService(gw, vs) {
			return true
		}
	}

	return false
}
