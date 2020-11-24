package instance

import (
	"github.com/bkhadars/power-go-client/errors"
	"github.com/bkhadars/power-go-client/ibmpisession"
	"github.com/bkhadars/power-go-client/power/client/p_cloud_tenants"
	"github.com/bkhadars/power-go-client/power/models"
	"log"
)

type IBMPITenantClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerImageClient ...
func NewIBMPITenantClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPITenantClient {
	return &IBMPITenantClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

func (f *IBMPITenantClient) Get(tenantid string) (*models.Tenant, error) {

	params := p_cloud_tenants.NewPcloudTenantsGetParams().WithTenantID(f.session.UserAccount).WithTenantID(tenantid)
	resp, err := f.session.Power.PCloudTenants.PcloudTenantsGet(params, ibmpisession.NewAuth(f.session, tenantid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
