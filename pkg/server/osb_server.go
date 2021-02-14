package server

import (
	"github.com/MadhavanKR/osb-broker-lib/pkg/broker"
	"github.com/MadhavanKR/osb-broker-lib/pkg/metrics"
	"github.com/MadhavanKR/osb-broker-lib/pkg/rest"
	"github.com/MadhavanKR/osb-broker-lib/pkg/server"
	prom "github.com/prometheus/client_golang/prometheus"
	"log"
)

func registerHandlers(brokerContext broker.Interface, collector *metrics.OSBMetricsCollector) (*rest.APISurface, error){
	apiSurface, apiSurfaceCreateErr := rest.NewAPISurface(brokerContext, collector)
	if apiSurfaceCreateErr != nil {
		log.Printf("error while creating API surface: %v", apiSurfaceCreateErr)
		return nil, apiSurfaceCreateErr
	}
	log.Printf("created apiSurface successfully\n")
	return apiSurface, nil
}

func GetHttpServer(brokerContext broker.Interface) (*server.Server, error){
	reg := prom.NewRegistry()
	osbMetrics := metrics.New()
	reg.MustRegister(osbMetrics)
	apiSurface, apiSurfaceCreateErr := registerHandlers(brokerContext, osbMetrics)
	if apiSurfaceCreateErr != nil {
		return nil, apiSurfaceCreateErr
	}

	brokerServer := server.New(apiSurface, reg)
	return brokerServer, nil
}
