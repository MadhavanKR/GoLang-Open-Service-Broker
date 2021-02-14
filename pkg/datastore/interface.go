package datastore

type IDatastore interface {
	GetServiceInstance(instanceId string) (*ServiceInstanceData, error)
	UpdateServiceInstance(data ServiceInstanceData) error
	InsertServiceInstance(data ServiceInstanceData) error
	DeleteServiceInstance(instanceId string) error

	GetServiceBinding(bindingId string) (*ServiceBindingData, error)
	UpdateServiceBinding(data ServiceBindingData) error
	InsertServiceBinding(data ServiceBindingData) error
	DeleteServiceBinding(bindingId string) error
}

