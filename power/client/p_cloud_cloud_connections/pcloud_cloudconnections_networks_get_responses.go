// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// PcloudCloudconnectionsNetworksGetReader is a Reader for the PcloudCloudconnectionsNetworksGet structure.
type PcloudCloudconnectionsNetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsNetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudconnectionsNetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudconnectionsNetworksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudCloudconnectionsNetworksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudconnectionsNetworksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudconnectionsNetworksGetOK creates a PcloudCloudconnectionsNetworksGetOK with default headers values
func NewPcloudCloudconnectionsNetworksGetOK() *PcloudCloudconnectionsNetworksGetOK {
	return &PcloudCloudconnectionsNetworksGetOK{}
}

/*PcloudCloudconnectionsNetworksGetOK handles this case with default header values.

OK
*/
type PcloudCloudconnectionsNetworksGetOK struct {
	Payload *models.Network
}

func (o *PcloudCloudconnectionsNetworksGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksGetBadRequest creates a PcloudCloudconnectionsNetworksGetBadRequest with default headers values
func NewPcloudCloudconnectionsNetworksGetBadRequest() *PcloudCloudconnectionsNetworksGetBadRequest {
	return &PcloudCloudconnectionsNetworksGetBadRequest{}
}

/*PcloudCloudconnectionsNetworksGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudconnectionsNetworksGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksGetNotFound creates a PcloudCloudconnectionsNetworksGetNotFound with default headers values
func NewPcloudCloudconnectionsNetworksGetNotFound() *PcloudCloudconnectionsNetworksGetNotFound {
	return &PcloudCloudconnectionsNetworksGetNotFound{}
}

/*PcloudCloudconnectionsNetworksGetNotFound handles this case with default header values.

Not Found
*/
type PcloudCloudconnectionsNetworksGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksGetInternalServerError creates a PcloudCloudconnectionsNetworksGetInternalServerError with default headers values
func NewPcloudCloudconnectionsNetworksGetInternalServerError() *PcloudCloudconnectionsNetworksGetInternalServerError {
	return &PcloudCloudconnectionsNetworksGetInternalServerError{}
}

/*PcloudCloudconnectionsNetworksGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsNetworksGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
