package osb_services

import (
	"errors"
	"fmt"
	"github.com/MadhavanKR/go-osb/pkg/datastore"
	"github.com/MadhavanKR/osb-broker-lib/pkg/broker"
	"github.com/golang/glog"
	osb "sigs.k8s.io/go-open-service-broker-client/v2"
	"time"
)

func (brokerContext *BrokerContext) ValidateBrokerAPIVersion(version string) error {
	return nil
}

func (brokerContext *BrokerContext) Provision(request *osb.ProvisionRequest, c *broker.RequestContext) (*broker.ProvisionResponse, error) {
	serviceInstanceData := datastore.ServiceInstanceData{
		BaseData: datastore.BaseData{
			ID:                    request.InstanceID,
			PlanId:                request.PlanID,
			ServiceId:             request.ServiceID,
			Parameters:            request.Parameters,
			Context:               request.Context,
			Name:                  request.InstanceID,
			Status:                1,
			Description:           "instance successfully provisioned",
			LastModifiedTimestamp: time.Now().String(),
		},
		DashboardURL: "https://sample-go-osb.com",
	}
	instanceData, _ := brokerContext.DataStore.GetServiceInstance(request.InstanceID)
	if instanceData != nil {
		errorMessage := fmt.Sprintf("instance id %s already exists", request.InstanceID)
		glog.Errorln(errorMessage)
		return &broker.ProvisionResponse{
			ProvisionResponse: osb.ProvisionResponse{
				Async:        false,
				DashboardURL: &serviceInstanceData.DashboardURL,
				OperationKey: nil,
			},
			Exists: true,
		}, nil
	}
	saveInstanceErr := brokerContext.DataStore.InsertServiceInstance(serviceInstanceData)
	if saveInstanceErr != nil {
		glog.Errorf("error while persisting instance details for %s: %v\n", request.InstanceID, saveInstanceErr)
		return nil, saveInstanceErr
	}
	glog.Infof("successfuly provisioned instance %s\n", request.InstanceID)
	return &broker.ProvisionResponse{
		ProvisionResponse: osb.ProvisionResponse{
			Async:        false,
			DashboardURL: &serviceInstanceData.DashboardURL,
			OperationKey: nil,
		},
		Exists: false,
	}, nil
}

func (brokerContext *BrokerContext) Deprovision(request *osb.DeprovisionRequest, c *broker.RequestContext) (*broker.DeprovisionResponse, error) {
	deleteInstanceErr := brokerContext.DataStore.DeleteServiceInstance(request.InstanceID)
	if deleteInstanceErr != nil {
		errorMessage := fmt.Sprintf("error while deleting instance %d: %v", request.InstanceID, deleteInstanceErr)
		glog.Errorf(errorMessage)
		return nil, deleteInstanceErr
	}
	glog.Infof("successfully deleted instance %s\n", request.InstanceID)
	return &broker.DeprovisionResponse{osb.DeprovisionResponse{
		Async:        false,
		OperationKey: nil,
	}}, nil
}

func (brokerContext *BrokerContext) LastOperation(request *osb.LastOperationRequest, c *broker.RequestContext) (*broker.LastOperationResponse, error) {
	instanceData, instanceFetchErr := brokerContext.DataStore.GetServiceInstance(request.InstanceID)
	if instanceData != nil {
		var state osb.LastOperationState
		switch instanceData.Status {
		case 0:
			state = osb.StateInProgress
			break
		case 1:
			state = osb.StateSucceeded
			break
		default:
			state = osb.StateFailed
		}
		return &broker.LastOperationResponse{osb.LastOperationResponse{
			State:       state,
			Description: &instanceData.Description,
			PollDelay:   nil,
		}}, nil
	}
	glog.Errorf("error while fetching instance details for %d: %v\n", request.InstanceID, instanceFetchErr)
	return nil, instanceFetchErr
}

func (brokerContext *BrokerContext) Update(request *osb.UpdateInstanceRequest, c *broker.RequestContext) (*broker.UpdateInstanceResponse, error) {
	return nil, errors.New("update not supported")
}
