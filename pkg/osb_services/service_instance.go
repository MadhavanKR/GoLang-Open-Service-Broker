package osb_services

import (
	osb "sigs.k8s.io/go-open-service-broker-client/v2"
	"github.com/MadhavanKR/osb-broker-lib/pkg/broker"
)

func (brokerContext *BrokerContext) ValidateBrokerAPIVersion(version string) error {
	return nil
}

func (brokerContext *BrokerContext) Provision(request *osb.ProvisionRequest, c *broker.RequestContext) (*broker.ProvisionResponse, error) {
	return nil, nil
}

func (brokerContext *BrokerContext) Deprovision(request *osb.DeprovisionRequest, c *broker.RequestContext) (*broker.DeprovisionResponse, error) {
	return nil, nil
}

func (brokerContext *BrokerContext) LastOperation(request *osb.LastOperationRequest, c *broker.RequestContext) (*broker.LastOperationResponse, error) {
	return nil, nil
}

func (BrokerContext *BrokerContext) Update(request *osb.UpdateInstanceRequest, c *broker.RequestContext) (*broker.UpdateInstanceResponse, error) {
	return nil, nil
}
