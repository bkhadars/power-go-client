// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// ServiceBrokerAuthLogoutReader is a Reader for the ServiceBrokerAuthLogout structure.
type ServiceBrokerAuthLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBrokerAuthLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewServiceBrokerAuthLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBrokerAuthLogoutOK creates a ServiceBrokerAuthLogoutOK with default headers values
func NewServiceBrokerAuthLogoutOK() *ServiceBrokerAuthLogoutOK {
	return &ServiceBrokerAuthLogoutOK{}
}

/*ServiceBrokerAuthLogoutOK handles this case with default header values.

OK
*/
type ServiceBrokerAuthLogoutOK struct {
	Payload models.Object
}

func (o *ServiceBrokerAuthLogoutOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/logout][%d] serviceBrokerAuthLogoutOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthLogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthLogoutInternalServerError creates a ServiceBrokerAuthLogoutInternalServerError with default headers values
func NewServiceBrokerAuthLogoutInternalServerError() *ServiceBrokerAuthLogoutInternalServerError {
	return &ServiceBrokerAuthLogoutInternalServerError{}
}

/*ServiceBrokerAuthLogoutInternalServerError handles this case with default header values.

Internal Server Error
*/
type ServiceBrokerAuthLogoutInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthLogoutInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/logout][%d] serviceBrokerAuthLogoutInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthLogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
