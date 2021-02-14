package datastore

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
)

func NewInMemoryDataStore() IDatastore {
	return &InMemoryDataStore{
		InstanceStore: make(map[string]ServiceInstanceData),
		BindingStore:  make(map[string]ServiceBindingData),
	}
}

func (inMemDS *InMemoryDataStore) GetServiceInstance(instanceId string) (*ServiceInstanceData, error) {
	if instanceData, ok := inMemDS.InstanceStore[instanceId]; ok {
		glog.Infof("instance data found for %s\n", instanceId)
		return &instanceData, nil
	} else {
		glog.Errorf("instance data not found for %s\n", instanceId)
		return nil, errors.New(fmt.Sprintf("%s not found\n", instanceId))
	}
}

func (inMemDS *InMemoryDataStore) InsertServiceInstance(instanceData ServiceInstanceData) error {
	if _, ok := inMemDS.InstanceStore[instanceData.ID]; ok {
		glog.Errorf("instance %s already exists\n", instanceData.ID)
		return errors.New(fmt.Sprintf("instance %s already exists\n", instanceData.ID))
	}
	inMemDS.InstanceStore[instanceData.ID] = instanceData
	glog.Infof("successfully stored instance data for %s\n", instanceData.ID)
	return nil
}

func (inMemDS *InMemoryDataStore) UpdateServiceInstance(instanceData ServiceInstanceData) error {
	if _, ok := inMemDS.InstanceStore[instanceData.ID]; !ok {
		glog.Infof("instance data not found for %s\n", instanceData.ID)
		return errors.New(fmt.Sprintf("instance %s not found\n", instanceData.ID))
	}
	inMemDS.InstanceStore[instanceData.ID] = instanceData
	glog.Infof("successfully updated instance data for %s\n", instanceData.ID)
	return nil
}

func (inMemDS *InMemoryDataStore) DeleteServiceInstance(instanceId string) error {
	if _, ok := inMemDS.InstanceStore[instanceId]; !ok {
		glog.Infof("instance data not found for %s\n", instanceId)
		return nil
	}
	//check if instance has any bindings
	bindingCount := 0
	for _, bindingData := range inMemDS.BindingStore {
		if bindingData.InstanceId == instanceId {
			bindingCount++
		}
	}
	if bindingCount > 0 {
		errorMessage := fmt.Sprintf("%s has %d bindings, please delete all bindings before deleting instance", instanceId, bindingCount)
		glog.Errorln(errorMessage)
		return errors.New(errorMessage)
	}
	delete(inMemDS.InstanceStore, instanceId)
	glog.Infof("successfully delete instance data for %s\n", instanceId)
	return nil
}

func (inMemDS *InMemoryDataStore) GetServiceBinding(bindingId string) (*ServiceBindingData, error) {
	if bindingData, ok := inMemDS.BindingStore[bindingId]; ok {
		glog.Infof("binding data found for %s\n", bindingId)
		return &bindingData, nil
	} else {
		glog.Errorf("binding data not found for %s\n", bindingId)
		return nil, errors.New(fmt.Sprintf("%s not found\n", bindingId))
	}
}

func (inMemDS *InMemoryDataStore) UpdateServiceBinding(data ServiceBindingData) error {
	if _, ok := inMemDS.BindingStore[data.ID]; ok {
		glog.Infof("binding data updated for %s\n", data.ID)
		inMemDS.BindingStore[data.ID] = data
		return nil
	}
	glog.Infof("binding data not found for %s\n", data.ID)
	return errors.New(fmt.Sprintf("binding %s not found\n", data.ID))
}

func (inMemDS *InMemoryDataStore) InsertServiceBinding(data ServiceBindingData) error {
	if _, ok := inMemDS.BindingStore[data.ID]; ok {
		glog.Errorf("binding %s already exists\n", data.ID)
		return errors.New(fmt.Sprintf("binding %s already exists\n", data.ID))
	}
	inMemDS.BindingStore[data.ID] = data
	glog.Infof("successfully stored binding data for %s\n", data.ID)
	return nil
}

func (inMemDS *InMemoryDataStore) DeleteServiceBinding(bindingId string) error {
	if _, ok := inMemDS.BindingStore[bindingId]; !ok {
		glog.Infof("binding data not found for %s\n", bindingId)
		return nil
	}
	delete(inMemDS.BindingStore, bindingId)
	glog.Infof("successfully delete binding data for %s\n", bindingId)
	return nil
}