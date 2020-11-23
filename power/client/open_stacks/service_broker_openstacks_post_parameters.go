// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

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

// NewServiceBrokerOpenstacksPostParams creates a new ServiceBrokerOpenstacksPostParams object
// with the default values initialized.
func NewServiceBrokerOpenstacksPostParams() *ServiceBrokerOpenstacksPostParams {
	var ()
	return &ServiceBrokerOpenstacksPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerOpenstacksPostParamsWithTimeout creates a new ServiceBrokerOpenstacksPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBrokerOpenstacksPostParamsWithTimeout(timeout time.Duration) *ServiceBrokerOpenstacksPostParams {
	var ()
	return &ServiceBrokerOpenstacksPostParams{

		timeout: timeout,
	}
}

// NewServiceBrokerOpenstacksPostParamsWithContext creates a new ServiceBrokerOpenstacksPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBrokerOpenstacksPostParamsWithContext(ctx context.Context) *ServiceBrokerOpenstacksPostParams {
	var ()
	return &ServiceBrokerOpenstacksPostParams{

		Context: ctx,
	}
}

// NewServiceBrokerOpenstacksPostParamsWithHTTPClient creates a new ServiceBrokerOpenstacksPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBrokerOpenstacksPostParamsWithHTTPClient(client *http.Client) *ServiceBrokerOpenstacksPostParams {
	var ()
	return &ServiceBrokerOpenstacksPostParams{
		HTTPClient: client,
	}
}

/*ServiceBrokerOpenstacksPostParams contains all the parameters to send to the API endpoint
for the service broker openstacks post operation typically these are written to a http.Request
*/
type ServiceBrokerOpenstacksPostParams struct {

	/*Body
	  Parameters for the creation of a new Open Stack Instance

	*/
	Body *models.OpenStackCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) WithTimeout(timeout time.Duration) *ServiceBrokerOpenstacksPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) WithContext(ctx context.Context) *ServiceBrokerOpenstacksPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) WithHTTPClient(client *http.Client) *ServiceBrokerOpenstacksPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) WithBody(body *models.OpenStackCreate) *ServiceBrokerOpenstacksPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service broker openstacks post params
func (o *ServiceBrokerOpenstacksPostParams) SetBody(body *models.OpenStackCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerOpenstacksPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
