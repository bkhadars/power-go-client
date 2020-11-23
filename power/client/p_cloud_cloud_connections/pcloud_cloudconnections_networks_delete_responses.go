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

// PcloudCloudconnectionsNetworksDeleteReader is a Reader for the PcloudCloudconnectionsNetworksDelete structure.
type PcloudCloudconnectionsNetworksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsNetworksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudconnectionsNetworksDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudconnectionsNetworksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewPcloudCloudconnectionsNetworksDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudconnectionsNetworksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudconnectionsNetworksDeleteOK creates a PcloudCloudconnectionsNetworksDeleteOK with default headers values
func NewPcloudCloudconnectionsNetworksDeleteOK() *PcloudCloudconnectionsNetworksDeleteOK {
	return &PcloudCloudconnectionsNetworksDeleteOK{}
}

/*PcloudCloudconnectionsNetworksDeleteOK handles this case with default header values.

OK
*/
type PcloudCloudconnectionsNetworksDeleteOK struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteBadRequest creates a PcloudCloudconnectionsNetworksDeleteBadRequest with default headers values
func NewPcloudCloudconnectionsNetworksDeleteBadRequest() *PcloudCloudconnectionsNetworksDeleteBadRequest {
	return &PcloudCloudconnectionsNetworksDeleteBadRequest{}
}

/*PcloudCloudconnectionsNetworksDeleteBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudconnectionsNetworksDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteGone creates a PcloudCloudconnectionsNetworksDeleteGone with default headers values
func NewPcloudCloudconnectionsNetworksDeleteGone() *PcloudCloudconnectionsNetworksDeleteGone {
	return &PcloudCloudconnectionsNetworksDeleteGone{}
}

/*PcloudCloudconnectionsNetworksDeleteGone handles this case with default header values.

Gone
*/
type PcloudCloudconnectionsNetworksDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteInternalServerError creates a PcloudCloudconnectionsNetworksDeleteInternalServerError with default headers values
func NewPcloudCloudconnectionsNetworksDeleteInternalServerError() *PcloudCloudconnectionsNetworksDeleteInternalServerError {
	return &PcloudCloudconnectionsNetworksDeleteInternalServerError{}
}

/*PcloudCloudconnectionsNetworksDeleteInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsNetworksDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
