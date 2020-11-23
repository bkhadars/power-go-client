// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// NewServiceBindingBindingParams creates a new ServiceBindingBindingParams object
// with the default values initialized.
func NewServiceBindingBindingParams() *ServiceBindingBindingParams {
	var ()
	return &ServiceBindingBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBindingBindingParamsWithTimeout creates a new ServiceBindingBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBindingBindingParamsWithTimeout(timeout time.Duration) *ServiceBindingBindingParams {
	var ()
	return &ServiceBindingBindingParams{

		timeout: timeout,
	}
}

// NewServiceBindingBindingParamsWithContext creates a new ServiceBindingBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBindingBindingParamsWithContext(ctx context.Context) *ServiceBindingBindingParams {
	var ()
	return &ServiceBindingBindingParams{

		Context: ctx,
	}
}

// NewServiceBindingBindingParamsWithHTTPClient creates a new ServiceBindingBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBindingBindingParamsWithHTTPClient(client *http.Client) *ServiceBindingBindingParams {
	var ()
	return &ServiceBindingBindingParams{
		HTTPClient: client,
	}
}

/*ServiceBindingBindingParams contains all the parameters to send to the API endpoint
for the service binding binding operation typically these are written to a http.Request
*/
type ServiceBindingBindingParams struct {

	/*XBrokerAPIOriginatingIdentity
	  identity of the user that initiated the request from the Platform

	*/
	XBrokerAPIOriginatingIdentity *string
	/*XBrokerAPIVersion
	  version number of the Service Broker API that the Platform will use

	*/
	XBrokerAPIVersion string
	/*AcceptsIncomplete
	  asynchronous operations supported

	*/
	AcceptsIncomplete *bool
	/*BindingID
	  binding id of binding to create

	*/
	BindingID string
	/*Body
	  parameters for the requested service binding

	*/
	Body *models.ServiceBindingRequest
	/*InstanceID
	  instance id of instance to provision

	*/
	InstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service binding binding params
func (o *ServiceBindingBindingParams) WithTimeout(timeout time.Duration) *ServiceBindingBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service binding binding params
func (o *ServiceBindingBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service binding binding params
func (o *ServiceBindingBindingParams) WithContext(ctx context.Context) *ServiceBindingBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service binding binding params
func (o *ServiceBindingBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service binding binding params
func (o *ServiceBindingBindingParams) WithHTTPClient(client *http.Client) *ServiceBindingBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service binding binding params
func (o *ServiceBindingBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXBrokerAPIOriginatingIdentity adds the xBrokerAPIOriginatingIdentity to the service binding binding params
func (o *ServiceBindingBindingParams) WithXBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity *string) *ServiceBindingBindingParams {
	o.SetXBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity)
	return o
}

// SetXBrokerAPIOriginatingIdentity adds the xBrokerApiOriginatingIdentity to the service binding binding params
func (o *ServiceBindingBindingParams) SetXBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity *string) {
	o.XBrokerAPIOriginatingIdentity = xBrokerAPIOriginatingIdentity
}

// WithXBrokerAPIVersion adds the xBrokerAPIVersion to the service binding binding params
func (o *ServiceBindingBindingParams) WithXBrokerAPIVersion(xBrokerAPIVersion string) *ServiceBindingBindingParams {
	o.SetXBrokerAPIVersion(xBrokerAPIVersion)
	return o
}

// SetXBrokerAPIVersion adds the xBrokerApiVersion to the service binding binding params
func (o *ServiceBindingBindingParams) SetXBrokerAPIVersion(xBrokerAPIVersion string) {
	o.XBrokerAPIVersion = xBrokerAPIVersion
}

// WithAcceptsIncomplete adds the acceptsIncomplete to the service binding binding params
func (o *ServiceBindingBindingParams) WithAcceptsIncomplete(acceptsIncomplete *bool) *ServiceBindingBindingParams {
	o.SetAcceptsIncomplete(acceptsIncomplete)
	return o
}

// SetAcceptsIncomplete adds the acceptsIncomplete to the service binding binding params
func (o *ServiceBindingBindingParams) SetAcceptsIncomplete(acceptsIncomplete *bool) {
	o.AcceptsIncomplete = acceptsIncomplete
}

// WithBindingID adds the bindingID to the service binding binding params
func (o *ServiceBindingBindingParams) WithBindingID(bindingID string) *ServiceBindingBindingParams {
	o.SetBindingID(bindingID)
	return o
}

// SetBindingID adds the bindingId to the service binding binding params
func (o *ServiceBindingBindingParams) SetBindingID(bindingID string) {
	o.BindingID = bindingID
}

// WithBody adds the body to the service binding binding params
func (o *ServiceBindingBindingParams) WithBody(body *models.ServiceBindingRequest) *ServiceBindingBindingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service binding binding params
func (o *ServiceBindingBindingParams) SetBody(body *models.ServiceBindingRequest) {
	o.Body = body
}

// WithInstanceID adds the instanceID to the service binding binding params
func (o *ServiceBindingBindingParams) WithInstanceID(instanceID string) *ServiceBindingBindingParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the service binding binding params
func (o *ServiceBindingBindingParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBindingBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XBrokerAPIOriginatingIdentity != nil {

		// header param X-Broker-API-Originating-Identity
		if err := r.SetHeaderParam("X-Broker-API-Originating-Identity", *o.XBrokerAPIOriginatingIdentity); err != nil {
			return err
		}

	}

	// header param X-Broker-API-Version
	if err := r.SetHeaderParam("X-Broker-API-Version", o.XBrokerAPIVersion); err != nil {
		return err
	}

	if o.AcceptsIncomplete != nil {

		// query param accepts_incomplete
		var qrAcceptsIncomplete bool
		if o.AcceptsIncomplete != nil {
			qrAcceptsIncomplete = *o.AcceptsIncomplete
		}
		qAcceptsIncomplete := swag.FormatBool(qrAcceptsIncomplete)
		if qAcceptsIncomplete != "" {
			if err := r.SetQueryParam("accepts_incomplete", qAcceptsIncomplete); err != nil {
				return err
			}
		}

	}

	// path param binding_id
	if err := r.SetPathParam("binding_id", o.BindingID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
