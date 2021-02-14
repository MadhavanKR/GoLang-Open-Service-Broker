package osb_services

import "github.com/MadhavanKR/go-osb/pkg/datastore"

type BrokerContext struct {
	BrokerName    string
	InstanceCount int
	BindingCount  int
	DataStore datastore.IDatastore
}
