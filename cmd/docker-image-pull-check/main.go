package main

import (
	// 	"context"
	"os"
	"strings"
	// 	"os/signal"
	// 	"path/filepath"
	// 	"syscall"
	// 	"time"
	// 	kh "github.com/Comcast/kuberhealthy/v2/pkg/checks/external/checkclient"
	// 	"github.com/Comcast/kuberhealthy/v2/pkg/kubeClient"
	// 	log "github.com/sirupsen/logrus"
	// 	"k8s.io/client-go/kubernetes"
	//  "github.com/docker/docker"
)

var (
	images []string
	imagesEnv string
)

func init() {
	parseImages()
}

func main() {
	go listenForInterrupts

	pullImages()
}

// listenForInterrupts watches the signal and done channels for termination.
func listenForInterrupts(ctx context.Context) {

	// Relay incoming OS interrupt signals to the signalChan.
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	sig := <-signalChan // This is a blocking operation -- the routine will stop here until there is something sent down the channel.
	log.Infoln("Received an interrupt signal from the signal channel.")
	log.Debugln("Signal received was:", sig.String())

	log.Debugln("Cancelling context.")
	ctxCancel() // Causes all functions within the check to return without error and abort. NOT an error
	// condition; this is a response to an external shutdown signal.

	// Clean up pods here.
	log.Infoln("Shutting down.")

	select {
	case sig = <-signalChan:
		// If there is an interrupt signal, interrupt the run.
		log.Warnln("Received a second interrupt signal from the signal channel.")
		log.Debugln("Signal received was:", sig.String())
	}

	os.Exit(0)
}