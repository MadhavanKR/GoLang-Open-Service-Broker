package osb_services

import (
	"errors"
	"github.com/MadhavanKR/go-osb/pkg/datastore"
	"github.com/golang/glog"
	osb "sigs.k8s.io/go-open-service-broker-client/v2"
	"github.com/MadhavanKR/osb-broker-lib/pkg/broker"
	"time"
)

func (brokerContext *BrokerContext) Bind(request *osb.BindRequest, c *broker.RequestContext) (*broker.BindResponse, error) {
	_, instanceFetchErr := brokerContext.DataStore.GetServiceInstance(request.InstanceID)
	if instanceFetchErr != nil {
		glog.Errorf("error creating bind: %v\n", instanceFetchErr)
		return nil, instanceFetchErr
	}

	if bindData, _ := brokerContext.DataStore.GetServiceBinding(request.BindingID); bindData != nil {
		glog.Errorf("binding %s already exists\n", request.BindingID)
		return &broker.BindResponse{
			BindResponse: osb.BindResponse{
				Async:           false,
				Credentials:     bindData.Credentials,
				SyslogDrainURL:  nil,
				RouteServiceURL: nil,
				VolumeMounts:    nil,
				OperationKey:    nil,
				Endpoints:       nil,
			},
			Exists:       true,
		}, nil
	}

	serviceBindingData := datastore.ServiceBindingData{
		BaseData:        datastore.BaseData{
			ID:                    request.BindingID,
			ServiceId:             request.ServiceID,
			PlanId:                request.BindingID,
			Parameters:            request.Parameters,
			Context:               request.Context,
			Name:                  request.BindingID,
			Status:                1,
			Description:           "created bind",
			LastModifiedTimestamp: time.Now().String(),
		},
		InstanceId:      request.InstanceID,
		ApplicationName: *request.AppGUID,
		Credentials:     map[string]interface{}{
			"username": "sample-user",
			"password": "sample-password",
			"url": "sample-url",
		},
	}
	insertBindErr := brokerContext.DataStore.InsertServiceBinding(serviceBindingData)
	if insertBindErr != nil {
		glog.Errorf("failed to persisit binding data for %s: %v\n", request.BindingID, insertBindErr)
		return nil, insertBindErr
	}
	glog.Infof("successfully created bind for %s\n", request.BindingID)
	return &broker.BindResponse{
		BindResponse: osb.BindResponse{
			Async:           false,
			Credentials:     serviceBindingData.Credentials,
			SyslogDrainURL:  nil,
			RouteServiceURL: nil,
			VolumeMounts:    nil,
			OperationKey:    nil,
			Endpoints:       nil,
		},
		Exists:       false,
	}, nil
}

func (brokerContext *BrokerContext) GetBinding(request *osb.GetBindingRequest, c *broker.RequestContext) (*broker.GetBindingResponse, error) {
	if bindData, _ := brokerContext.DataStore.GetServiceBinding(request.BindingID); bindData != nil {
		glog.Infof("binding %s found \n", request.BindingID)
		return &broker.GetBindingResponse{
			osb.GetBindingResponse{
				Credentials:     bindData.Credentials,
				SyslogDrainURL:  nil,
				RouteServiceURL: nil,
				VolumeMounts:    nil,
				Parameters:      bindData.Parameters,
				Endpoints:       nil,
			},
		}, nil
	}
	glog.Infof("no data found for binding %s\n", request.BindingID)
	return nil, errors.New("no data found for binding " + request.BindingID)
}

func (brokerContext *BrokerContext) BindingLastOperation(request *osb.BindingLastOperationRequest, c *broker.RequestContext) (*broker.LastOperationResponse, error) {
	if bindData, _ := brokerContext.DataStore.GetServiceBinding(request.BindingID); bindData != nil {
		var state osb.LastOperationState
		switch bindData.Status {
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
			Description: &bindData.Description,
			PollDelay:   nil,
		}}, nil
	}
	glog.Errorf("no data found for binding %s\n", request.BindingID)
	return nil, errors.New("no data found for binding " + request.BindingID)
}

func (brokerContext *BrokerContext) Unbind(request *osb.UnbindRequest, c *broker.RequestContext) (*broker.UnbindResponse, error) {
	deleteBindErr := brokerContext.DataStore.DeleteServiceBinding(request.BindingID)
	if deleteBindErr != nil {
		glog.Errorf("error while deleteing bind %s: %v\n", request.BindingID, deleteBindErr)
		return nil, deleteBindErr
	}
	return &broker.UnbindResponse{osb.UnbindResponse{
		Async:        false,
		OperationKey: nil,
	}}, nil
}
