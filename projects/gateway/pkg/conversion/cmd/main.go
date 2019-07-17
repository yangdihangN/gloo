package main

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway/pkg/conversion"
	"github.com/solo-io/gloo/projects/gateway/pkg/conversion/setup"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

func main() {
	ctx := contextutils.WithLogger(context.Background(), "gateway-conversion")
	// Set v1 to served for the duration of the job.
	crd := setup.MustSetV1ToServed(ctx)
	defer setup.MustSetV1ToNotServed(ctx, crd)

	clientSet := setup.MustClientSet(ctx)
	gatewayLadder := conversion.NewLadder(
		ctx,
		"gloo-system",
		clientSet.V1Gateway,
		clientSet.V2alpha1Gateway,
		conversion.NewGatewayConverter(),
	)

	if err := gatewayLadder.Climb(); err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Failed to upgrade existing gateway resources.", zap.Error(err))
	}
}
