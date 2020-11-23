// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

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

// NewPcloudCloudinstancesImagesPostParams creates a new PcloudCloudinstancesImagesPostParams object
// with the default values initialized.
func NewPcloudCloudinstancesImagesPostParams() *PcloudCloudinstancesImagesPostParams {
	var ()
	return &PcloudCloudinstancesImagesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesImagesPostParamsWithTimeout creates a new PcloudCloudinstancesImagesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesImagesPostParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesImagesPostParams {
	var ()
	return &PcloudCloudinstancesImagesPostParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesImagesPostParamsWithContext creates a new PcloudCloudinstancesImagesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesImagesPostParamsWithContext(ctx context.Context) *PcloudCloudinstancesImagesPostParams {
	var ()
	return &PcloudCloudinstancesImagesPostParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesImagesPostParamsWithHTTPClient creates a new PcloudCloudinstancesImagesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesImagesPostParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesImagesPostParams {
	var ()
	return &PcloudCloudinstancesImagesPostParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesImagesPostParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances images post operation typically these are written to a http.Request
*/
type PcloudCloudinstancesImagesPostParams struct {

	/*Body
	  Parameters for the creation of a new image from available images

	*/
	Body *models.CreateImage
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesImagesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) WithContext(ctx context.Context) *PcloudCloudinstancesImagesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesImagesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) WithBody(body *models.CreateImage) *PcloudCloudinstancesImagesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) SetBody(body *models.CreateImage) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesImagesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances images post params
func (o *PcloudCloudinstancesImagesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesImagesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
