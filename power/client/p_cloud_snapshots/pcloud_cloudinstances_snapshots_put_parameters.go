// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// NewPcloudCloudinstancesSnapshotsPutParams creates a new PcloudCloudinstancesSnapshotsPutParams object
// with the default values initialized.
func NewPcloudCloudinstancesSnapshotsPutParams() *PcloudCloudinstancesSnapshotsPutParams {
	var ()
	return &PcloudCloudinstancesSnapshotsPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesSnapshotsPutParamsWithTimeout creates a new PcloudCloudinstancesSnapshotsPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesSnapshotsPutParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesSnapshotsPutParams {
	var ()
	return &PcloudCloudinstancesSnapshotsPutParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesSnapshotsPutParamsWithContext creates a new PcloudCloudinstancesSnapshotsPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesSnapshotsPutParamsWithContext(ctx context.Context) *PcloudCloudinstancesSnapshotsPutParams {
	var ()
	return &PcloudCloudinstancesSnapshotsPutParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesSnapshotsPutParamsWithHTTPClient creates a new PcloudCloudinstancesSnapshotsPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesSnapshotsPutParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesSnapshotsPutParams {
	var ()
	return &PcloudCloudinstancesSnapshotsPutParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesSnapshotsPutParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances snapshots put operation typically these are written to a http.Request
*/
type PcloudCloudinstancesSnapshotsPutParams struct {

	/*Body
	  Parameters for the update of a  PVM instance snapshot

	*/
	Body *models.SnapshotUpdate
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*SnapshotID
	  PVM Instance snapshot id

	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesSnapshotsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) WithContext(ctx context.Context) *PcloudCloudinstancesSnapshotsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesSnapshotsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) WithBody(body *models.SnapshotUpdate) *PcloudCloudinstancesSnapshotsPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) SetBody(body *models.SnapshotUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesSnapshotsPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithSnapshotID adds the snapshotID to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) WithSnapshotID(snapshotID string) *PcloudCloudinstancesSnapshotsPutParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the pcloud cloudinstances snapshots put params
func (o *PcloudCloudinstancesSnapshotsPutParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesSnapshotsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param snapshot_id
	if err := r.SetPathParam("snapshot_id", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
