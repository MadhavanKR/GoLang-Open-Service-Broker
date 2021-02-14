package main

import (
	"context"
	"flag"
	"github.com/MadhavanKR/go-osb/pkg/osb_services"
	"github.com/MadhavanKR/go-osb/pkg/server"
	"github.com/golang/glog"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	brokerContext := &osb_services.BrokerContext{
		BrokerName:    "sample-broker",
		InstanceCount: 0,
		BindingCount:  0,
	}
	brokerServer, brokerServerCreateErr := server.GetHttpServer(brokerContext)
	if brokerServerCreateErr != nil {
		log.Printf("error while initializing http server for broker: %v\n", brokerServer)
		os.Exit(1)
	}
	log.Printf("starting broker server\n")
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	go cancelOnInterrupt(ctx, cancelFunc)
	var sample string
	flag.String(sample, "sample", "")
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infoln("am i visible")
	err := brokerServer.Run(ctx, ":2527")
	glog.Error(err)
}

func cancelOnInterrupt(ctx context.Context, f context.CancelFunc) {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-term:
			glog.Infof("Received SIGTERM, exiting gracefully...")
			f()
			os.Exit(0)
		case <-ctx.Done():
			os.Exit(0)
		}
	}
}
