// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// ServiceBrokerOpenstacksOpenstackGetReader is a Reader for the ServiceBrokerOpenstacksOpenstackGet structure.
type ServiceBrokerOpenstacksOpenstackGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksOpenstackGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBrokerOpenstacksOpenstackGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceBrokerOpenstacksOpenstackGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewServiceBrokerOpenstacksOpenstackGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceBrokerOpenstacksOpenstackGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksOpenstackGetOK creates a ServiceBrokerOpenstacksOpenstackGetOK with default headers values
func NewServiceBrokerOpenstacksOpenstackGetOK() *ServiceBrokerOpenstacksOpenstackGetOK {
	return &ServiceBrokerOpenstacksOpenstackGetOK{}
}

/*ServiceBrokerOpenstacksOpenstackGetOK handles this case with default header values.

OK
*/
type ServiceBrokerOpenstacksOpenstackGetOK struct {
	Payload *models.OpenStackInfo
}

func (o *ServiceBrokerOpenstacksOpenstackGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStackInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksOpenstackGetBadRequest creates a ServiceBrokerOpenstacksOpenstackGetBadRequest with default headers values
func NewServiceBrokerOpenstacksOpenstackGetBadRequest() *ServiceBrokerOpenstacksOpenstackGetBadRequest {
	return &ServiceBrokerOpenstacksOpenstackGetBadRequest{}
}

/*ServiceBrokerOpenstacksOpenstackGetBadRequest handles this case with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksOpenstackGetBadRequest struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksOpenstackGetNotFound creates a ServiceBrokerOpenstacksOpenstackGetNotFound with default headers values
func NewServiceBrokerOpenstacksOpenstackGetNotFound() *ServiceBrokerOpenstacksOpenstackGetNotFound {
	return &ServiceBrokerOpenstacksOpenstackGetNotFound{}
}

/*ServiceBrokerOpenstacksOpenstackGetNotFound handles this case with default header values.

Not Found
*/
type ServiceBrokerOpenstacksOpenstackGetNotFound struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksOpenstackGetInternalServerError creates a ServiceBrokerOpenstacksOpenstackGetInternalServerError with default headers values
func NewServiceBrokerOpenstacksOpenstackGetInternalServerError() *ServiceBrokerOpenstacksOpenstackGetInternalServerError {
	return &ServiceBrokerOpenstacksOpenstackGetInternalServerError{}
}

/*ServiceBrokerOpenstacksOpenstackGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksOpenstackGetInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
