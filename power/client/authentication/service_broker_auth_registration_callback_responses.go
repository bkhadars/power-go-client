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

// ServiceBrokerAuthRegistrationCallbackReader is a Reader for the ServiceBrokerAuthRegistrationCallback structure.
type ServiceBrokerAuthRegistrationCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthRegistrationCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBrokerAuthRegistrationCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewServiceBrokerAuthRegistrationCallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceBrokerAuthRegistrationCallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBrokerAuthRegistrationCallbackOK creates a ServiceBrokerAuthRegistrationCallbackOK with default headers values
func NewServiceBrokerAuthRegistrationCallbackOK() *ServiceBrokerAuthRegistrationCallbackOK {
	return &ServiceBrokerAuthRegistrationCallbackOK{}
}

/*ServiceBrokerAuthRegistrationCallbackOK handles this case with default header values.

OK
*/
type ServiceBrokerAuthRegistrationCallbackOK struct {
	Payload *models.AccessToken
}

func (o *ServiceBrokerAuthRegistrationCallbackOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthRegistrationCallbackUnauthorized creates a ServiceBrokerAuthRegistrationCallbackUnauthorized with default headers values
func NewServiceBrokerAuthRegistrationCallbackUnauthorized() *ServiceBrokerAuthRegistrationCallbackUnauthorized {
	return &ServiceBrokerAuthRegistrationCallbackUnauthorized{}
}

/*ServiceBrokerAuthRegistrationCallbackUnauthorized handles this case with default header values.

Unauthorized
*/
type ServiceBrokerAuthRegistrationCallbackUnauthorized struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthRegistrationCallbackInternalServerError creates a ServiceBrokerAuthRegistrationCallbackInternalServerError with default headers values
func NewServiceBrokerAuthRegistrationCallbackInternalServerError() *ServiceBrokerAuthRegistrationCallbackInternalServerError {
	return &ServiceBrokerAuthRegistrationCallbackInternalServerError{}
}

/*ServiceBrokerAuthRegistrationCallbackInternalServerError handles this case with default header values.

Internal Server Error
*/
type ServiceBrokerAuthRegistrationCallbackInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback-registration][%d] serviceBrokerAuthRegistrationCallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationCallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
