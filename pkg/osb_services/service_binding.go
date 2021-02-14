package osb_services

import (
	osb "sigs.k8s.io/go-open-service-broker-client/v2"
	"github.com/MadhavanKR/osb-broker-lib/pkg/broker"
)

func (brokerContext *BrokerContext) Bind(request *osb.BindRequest, c *broker.RequestContext) (*broker.BindResponse, error) {
	return nil, nil
}

func (brokerContext *BrokerContext) GetBinding(request *osb.GetBindingRequest, c *broker.RequestContext) (*broker.GetBindingResponse, error) {
	return nil, nil
}

func (brokerContext *BrokerContext) BindingLastOperation(request *osb.BindingLastOperationRequest, c *broker.RequestContext) (*broker.LastOperationResponse, error) {
	return nil, nil
}

func (brokerContext *BrokerContext) Unbind(request *osb.UnbindRequest, c *broker.RequestContext) (*broker.UnbindResponse, error) {
	return nil, nil
}
