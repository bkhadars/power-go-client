package instance

import (
	"fmt"
	"github.com/bkhadars/power-go-client/errors"
	"github.com/bkhadars/power-go-client/ibmpisession"
	"github.com/bkhadars/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/bkhadars/power-go-client/power/client/p_cloud_snapshots"
	"github.com/bkhadars/power-go-client/power/models"
	"log"
	"time"
)

type IBMPISnapshotClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewSnapShotClient ...
func NewIBMPISnapshotClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPISnapshotClient {
	return &IBMPISnapshotClient{
		sess, powerinstanceid,
	}
}

//Get information about a single snapshot only
func (f *IBMPISnapshotClient) Get(id, powerinstanceid string, timeout time.Duration) (*models.Snapshot, error) {

	log.Printf("Calling the Snapshotget Method..")
	log.Printf("The input snapshot name is %s and  to the cloudinstance id %s", id, powerinstanceid)

	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsGetParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithSnapshotID(id)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Delete ...
func (f *IBMPISnapshotClient) Delete(id string, powerinstanceid string, timeout time.Duration) error {
	//var cloudinstanceid = f.session.PowerServiceInstance
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsDeleteParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithSnapshotID(id)
	_, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Update..
func (f *IBMPISnapshotClient) Update(id, powerinstanceid string, snapshotdef *models.SnapshotUpdate, timeout time.Duration) (models.Object, error) {

	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsPutParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithSnapshotID(id).WithBody(snapshotdef)

	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// All snapshots part of an instance

func (f *IBMPISnapshotClient) GetAll(id, powerinstanceid string, timeout time.Duration) (*models.Snapshots, error) {

	log.Printf("Calling the Power Snapshots GetAll Method")
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsGetallParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

// Restore a Snapshot

func (f *IBMPISnapshotClient) Create(pvminstanceid, powerinstanceid, snapshotid, restorefailAction string, timeout time.Duration) (*models.Snapshot, error) {
	log.Printf("Calling the Power Snapshots Restore Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsRestorePostParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid).WithCloudInstanceID(pvminstanceid).WithSnapshotID(snapshotid).WithRestoreFailAction(&restorefailAction)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsRestorePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, fmt.Errorf("Failed to create the restore")
	}
	return resp.Payload, nil
}
