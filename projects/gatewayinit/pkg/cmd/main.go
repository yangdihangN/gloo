package cmd

import (
	"context"

	"github.com/solo-io/gloo/projects/gatewayinit/pkg/conversion"
	"github.com/solo-io/gloo/projects/gatewayinit/pkg/setup"
	"github.com/solo-io/go-utils/contextutils"
)

func main() {
	ctx := contextutils.WithLogger(context.Background(), "gateway-init")
	clientSet := setup.MustClientSet(ctx)
	converterList := conversion.NewConverterList(ctx, clientSet, "gloo-system")
	converterList.Convert()
}
