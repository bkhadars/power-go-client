// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// PcloudCloudinstancesVolumesGetReader is a Reader for the PcloudCloudinstancesVolumesGet structure.
type PcloudCloudinstancesVolumesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudinstancesVolumesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudinstancesVolumesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudCloudinstancesVolumesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudinstancesVolumesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesGetOK creates a PcloudCloudinstancesVolumesGetOK with default headers values
func NewPcloudCloudinstancesVolumesGetOK() *PcloudCloudinstancesVolumesGetOK {
	return &PcloudCloudinstancesVolumesGetOK{}
}

/*PcloudCloudinstancesVolumesGetOK handles this case with default header values.

OK
*/
type PcloudCloudinstancesVolumesGetOK struct {
	Payload *models.Volume
}

func (o *PcloudCloudinstancesVolumesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetBadRequest creates a PcloudCloudinstancesVolumesGetBadRequest with default headers values
func NewPcloudCloudinstancesVolumesGetBadRequest() *PcloudCloudinstancesVolumesGetBadRequest {
	return &PcloudCloudinstancesVolumesGetBadRequest{}
}

/*PcloudCloudinstancesVolumesGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetNotFound creates a PcloudCloudinstancesVolumesGetNotFound with default headers values
func NewPcloudCloudinstancesVolumesGetNotFound() *PcloudCloudinstancesVolumesGetNotFound {
	return &PcloudCloudinstancesVolumesGetNotFound{}
}

/*PcloudCloudinstancesVolumesGetNotFound handles this case with default header values.

Not Found
*/
type PcloudCloudinstancesVolumesGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetInternalServerError creates a PcloudCloudinstancesVolumesGetInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesGetInternalServerError() *PcloudCloudinstancesVolumesGetInternalServerError {
	return &PcloudCloudinstancesVolumesGetInternalServerError{}
}

/*PcloudCloudinstancesVolumesGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesVolumesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
