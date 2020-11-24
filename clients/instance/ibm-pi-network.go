package instance

import (
	"fmt"
	"github.com/bkhadars/power-go-client/errors"
	"github.com/bkhadars/power-go-client/ibmpisession"
	"github.com/bkhadars/power-go-client/power/client/p_cloud_networks"
	"github.com/bkhadars/power-go-client/power/models"
	"log"
	"time"
)

const getNetwork = "Get Networks for [ Network id (%s) ] and CloudInstance id - (%s) ] "
const getAllPortPrint = "Get Networks for [ Network id (%s) ] and CloudInstance id - (%s) ] "
const getPortPrint = "Get Port for [ Port id - (%s) and  Network id -  (%s)  and CloudInstance id - (%s) ] "

type IBMPINetworkClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerImageClient ...
func NewIBMPINetworkClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPINetworkClient {
	return &IBMPINetworkClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

func (f *IBMPINetworkClient) Get(id, powerinstanceid string, timeout time.Duration) (*models.Network, error) {
	log.Printf(getNetwork, id, powerinstanceid)
	params := p_cloud_networks.NewPcloudNetworksGetParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *IBMPINetworkClient) Create(name string, networktype string, cidr string, dnsservers []string, gateway string, startip string, endip string, powerinstanceid string, timeout time.Duration) (*models.Network, *models.Network, error) {

	var body = models.NetworkCreate{}

	body.Name = name
	body.Type = &networktype

	if networktype == "vlan" {
		var ipbody = []*models.IPAddressRange{
			{&endip, &startip},
		}
		body.IPAddressRanges = ipbody
		body.Gateway = gateway
		body.Cidr = cidr
	}
	body.DNSServers = dnsservers

	log.Printf("Printing the body %+v", body)
	params := p_cloud_networks.NewPcloudNetworksPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	_, resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	//log.Printf("The error is %d ",resp.Payload.VlanID)
	if err != nil {
		return nil, nil, errors.ToError(err)
	}

	if resp != nil {
		log.Printf("Failed to create the network ")
	}

	return resp.Payload, nil, nil
}

func (f *IBMPINetworkClient) GetPublic(cloud_instance_id string, timeout time.Duration) (*models.Networks, error) {

	filterQuery := "type=\"pub-vlan\""
	params := p_cloud_networks.NewPcloudNetworksGetallParamsWithTimeout(timeout).WithCloudInstanceID(cloud_instance_id).WithFilter(&filterQuery)

	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGetall(params, ibmpisession.NewAuth(f.session, cloud_instance_id))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *IBMPINetworkClient) Delete(id string, powerinstanceid string, timeout time.Duration) error {
	params := p_cloud_networks.NewPcloudNetworksDeleteParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(id)
	_, err := f.session.Power.PCloudNetworks.PcloudNetworksDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// New Function for Ports

//Get all
func (f *IBMPINetworkClient) GetAllPort(id string, powerinstanceid string, timeout time.Duration) (*models.NetworkPorts, error) {

	log.Printf(getAllPortPrint, id, powerinstanceid)
	params := p_cloud_networks.NewPcloudNetworksPortsGetallParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	log.Printf("Printing the response %s", len(resp.Payload.Ports))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the GetNetworkPorts Operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

// Get Port

func (f *IBMPINetworkClient) GetPort(id string, powerinstanceid string, network_port_id string, timeout time.Duration) (*models.NetworkPort, error) {
	log.Printf(getPortPrint, network_port_id, id, powerinstanceid)
	params := p_cloud_networks.NewPcloudNetworksPortsGetParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(id).WithPortID(network_port_id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the GetNetworkPort Operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

//Create

func (f *IBMPINetworkClient) CreatePort(id string, powerinstanceid string, networkportdef *p_cloud_networks.PcloudNetworksPortsPostParams, timeout time.Duration) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsPostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(id).WithBody(networkportdef.Body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to create the network port")
		return nil, fmt.Errorf("Failed to create the network port for cloudinstance id [%s]", powerinstanceid)
	}
	return resp.Payload, nil
}

// Delete

func (f *IBMPINetworkClient) DeletePort(networkid string, powerinstanceid string, portid string, timeout time.Duration) (*models.Object, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsDeleteParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(networkid).WithPortID(portid)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to create the network port")
		return nil, errors.ToError(err)
	}
	return &resp.Payload, nil
}

//Attach Port to the PVM Instance

func (f *IBMPINetworkClient) AttachPort(powerinstanceid, network_id, port_id, description, pvminstanceid string, timeout time.Duration) (*models.NetworkPort, error) {

	var body = models.NetworkPortUpdate{}
	body.Description = &description
	body.PvmInstanceID = &pvminstanceid
	log.Printf("Attaching the port with id [%s] that belongs to cloud instance [%s] and networkid [%s] to the instance with id [%s]", port_id, powerinstanceid, network_id, pvminstanceid)
	params := p_cloud_networks.NewPcloudNetworksPortsPutParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(network_id).WithPortID(port_id).WithBody(&body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		log.Printf("Failed to attach the port to the pvminstance")
		return nil, fmt.Errorf("Failed to attach the port [%s] to the pvminstance [%s]", port_id, pvminstanceid)
	}

	return resp.Payload, nil
}

// Detach Port from the PVM Instance

func (f *IBMPINetworkClient) DetachPort(powerinstanceid, network_id, port_id string, timeout time.Duration) (*models.NetworkPort, error) {
	log.Printf("Calling the detach port ")

	emptyPVM := ""
	body := &models.NetworkPortUpdate{
		PvmInstanceID: &emptyPVM,
	}

	params := p_cloud_networks.NewPcloudNetworksPortsPutParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithNetworkID(network_id).WithPortID(port_id).WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		log.Printf("Failed to detach the port to the pvminstance")
		return nil, fmt.Errorf("Failed to detach the port [%s]", port_id)
	}

	return resp.Payload, nil
}
