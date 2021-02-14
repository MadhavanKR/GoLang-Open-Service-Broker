package osb_services

import (
	"github.com/MadhavanKR/osb-broker-lib/pkg/broker"
	"log"
	osb "sigs.k8s.io/go-open-service-broker-client/v2"
)

func (brokerContext *BrokerContext) GetCatalog(c *broker.RequestContext) (*broker.CatalogResponse, error) {
	log.Println("fetching catalog..")
	catalog := osb.CatalogResponse{Services: getServices()}
	return &broker.CatalogResponse{catalog}, nil
}

func getServices() []osb.Service {
	trueBool := true
	falseBool := false
	serviceList := make([]osb.Service, 0)
	service := osb.Service{
		ID:                   "sample-service-id",
		Name:                 "sample-service-name",
		Description:          "this is a sample server",
		Tags:                 nil,
		Requires:             nil,
		Bindable:             trueBool,
		InstancesRetrievable: true,
		BindingsRetrievable:  true,
		PlanUpdatable:        &falseBool,
		Plans:                getPlans("sample-service-id"),
		DashboardClient:      nil,
		Metadata:             nil,
	}
	serviceList = append(serviceList, service)
	return serviceList
}

func getPlans(serviceId string) []osb.Plan {
	trueBool := true
	planList := make([]osb.Plan, 0)
	switch serviceId {
	case "sample-service-id":
		samplePlan1 := osb.Plan{
			ID:          "sample-plan-id",
			Name:        "sample-plan-name",
			Description: "this is a sample plan",
			Free:        &trueBool,
			Bindable:    &trueBool,
			Metadata:    nil,
			Schemas:     nil,
		}
		planList = append(planList, samplePlan1)
		break
	default:
		log.Printf("%s is an unknown service, no plans available\n", serviceId)
	}
	return planList
}