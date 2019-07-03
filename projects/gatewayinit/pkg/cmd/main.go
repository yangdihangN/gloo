package cmd

import (
	"context"

	"github.com/solo-io/gloo/projects/gatewayinit/pkg/conversion/convertgateway"
	"github.com/solo-io/gloo/projects/gatewayinit/pkg/setup"
	"github.com/solo-io/go-utils/contextutils"
)

func main() {
	ctx := contextutils.WithLogger(context.Background(), "gateway-init")
	clientSet := setup.MustClientSet(ctx)
	gatewayLadder := convertgateway.NewLadder(
		ctx,
		"gloo-system",
		clientSet.V1Gateway,
		clientSet.V2alpha1Gateway,
		convertgateway.NewV2alpha1Converter(),
	)
	gatewayLadder.Climb()
}
