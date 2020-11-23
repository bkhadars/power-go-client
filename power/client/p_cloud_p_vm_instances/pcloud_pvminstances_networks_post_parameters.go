// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// NewPcloudPvminstancesNetworksPostParams creates a new PcloudPvminstancesNetworksPostParams object
// with the default values initialized.
func NewPcloudPvminstancesNetworksPostParams() *PcloudPvminstancesNetworksPostParams {
	var ()
	return &PcloudPvminstancesNetworksPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesNetworksPostParamsWithTimeout creates a new PcloudPvminstancesNetworksPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesNetworksPostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesNetworksPostParams {
	var ()
	return &PcloudPvminstancesNetworksPostParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesNetworksPostParamsWithContext creates a new PcloudPvminstancesNetworksPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesNetworksPostParamsWithContext(ctx context.Context) *PcloudPvminstancesNetworksPostParams {
	var ()
	return &PcloudPvminstancesNetworksPostParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesNetworksPostParamsWithHTTPClient creates a new PcloudPvminstancesNetworksPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesNetworksPostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesNetworksPostParams {
	var ()
	return &PcloudPvminstancesNetworksPostParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesNetworksPostParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances networks post operation typically these are written to a http.Request
*/
type PcloudPvminstancesNetworksPostParams struct {

	/*Body
	  Add network to PVM Instance parameters

	*/
	Body *models.PVMInstanceAddNetwork
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesNetworksPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) WithContext(ctx context.Context) *PcloudPvminstancesNetworksPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesNetworksPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) WithBody(body *models.PVMInstanceAddNetwork) *PcloudPvminstancesNetworksPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) SetBody(body *models.PVMInstanceAddNetwork) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesNetworksPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesNetworksPostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances networks post params
func (o *PcloudPvminstancesNetworksPostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesNetworksPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
