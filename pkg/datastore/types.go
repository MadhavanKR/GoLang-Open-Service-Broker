package datastore

type BaseData struct {
	ID string `json:"id"`
	ServiceId string `json:"serviceId"`
	PlanId string `json:"planId"`
	Parameters map[string]interface{} `json:"parameters"`
	Context map[string]interface{} `json:"context"`
	Name string `json:"name"`
	Status int `json:"statusCode"`
	Description string `json:"statusDescription"`
	LastModifiedTimestamp string `json:"lastModifiedTimestamp"`
}

type ServiceInstanceData struct {
	BaseData
	DashboardURL string `json:"dashboardUrl"`
}

type ServiceBindingData struct {
	BaseData
	InstanceId string `json:"instanceId"`
	ApplicationName string `json:"applicationName"`
	Credentials map[string]interface{} `json:"credentials"`
}

type InMemoryDataStore struct {
	InstanceStore map[string]ServiceInstanceData `json:"instanceInfo"`
	BindingStore map[string]ServiceBindingData `json:"instanceBinding"`
}