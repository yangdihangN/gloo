package cmd

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway/pkg/conversion"
	"github.com/solo-io/gloo/projects/gatewayinit/pkg/setup"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

func main() {
	ctx := contextutils.WithLogger(context.Background(), "gateway-init")
	clientSet := setup.MustClientSet(ctx)
	gatewayLadder := conversion.NewLadder(
		ctx,
		"gloo-system",
		clientSet.V1Gateway,
		clientSet.V2alpha1Gateway,
		conversion.NewV2alpha1Converter(),
	)
	if err := gatewayLadder.Climb(); err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("Failed to upgrade existing gateway resources.", zap.Error(err))
	}
}
