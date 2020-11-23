// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// PcloudNetworksPutReader is a Reader for the PcloudNetworksPut structure.
type PcloudNetworksPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudNetworksPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudNetworksPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudNetworksPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudNetworksPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudNetworksPutOK creates a PcloudNetworksPutOK with default headers values
func NewPcloudNetworksPutOK() *PcloudNetworksPutOK {
	return &PcloudNetworksPutOK{}
}

/*PcloudNetworksPutOK handles this case with default header values.

OK
*/
type PcloudNetworksPutOK struct {
	Payload *models.Network
}

func (o *PcloudNetworksPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksPutOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPutBadRequest creates a PcloudNetworksPutBadRequest with default headers values
func NewPcloudNetworksPutBadRequest() *PcloudNetworksPutBadRequest {
	return &PcloudNetworksPutBadRequest{}
}

/*PcloudNetworksPutBadRequest handles this case with default header values.

Bad Request
*/
type PcloudNetworksPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudNetworksPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPutUnprocessableEntity creates a PcloudNetworksPutUnprocessableEntity with default headers values
func NewPcloudNetworksPutUnprocessableEntity() *PcloudNetworksPutUnprocessableEntity {
	return &PcloudNetworksPutUnprocessableEntity{}
}

/*PcloudNetworksPutUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudNetworksPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudNetworksPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudNetworksPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPutInternalServerError creates a PcloudNetworksPutInternalServerError with default headers values
func NewPcloudNetworksPutInternalServerError() *PcloudNetworksPutInternalServerError {
	return &PcloudNetworksPutInternalServerError{}
}

/*PcloudNetworksPutInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudNetworksPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudNetworksPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}][%d] pcloudNetworksPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
