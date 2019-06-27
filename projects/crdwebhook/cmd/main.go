package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/solo-io/go-utils/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/conversion"

)

func main() {

	sigStop := make(chan os.Signal)
	signal.Notify(sigStop, syscall.SIGTERM)

	stop := make(chan struct{})
	go func() {
		select {
		case <- sigStop:
			stop <- struct{}{}
		}
	}()

	server := webhook.Server{}
	converter := &conversion.Webhook{}
	server.Register("/crdconvert", converter)
	if err := server.Start(stop); err != nil {
		log.Fatalf("%v", err)
	}
}