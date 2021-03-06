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

// PcloudV2VolumescloneGetReader is a Reader for the PcloudV2VolumescloneGet structure.
type PcloudV2VolumescloneGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudV2VolumescloneGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudV2VolumescloneGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudV2VolumescloneGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudV2VolumescloneGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudV2VolumescloneGetOK creates a PcloudV2VolumescloneGetOK with default headers values
func NewPcloudV2VolumescloneGetOK() *PcloudV2VolumescloneGetOK {
	return &PcloudV2VolumescloneGetOK{}
}

/*PcloudV2VolumescloneGetOK handles this case with default header values.

OK
*/
type PcloudV2VolumescloneGetOK struct {
	Payload *models.VolumesCloneDetail
}

func (o *PcloudV2VolumescloneGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesCloneDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetBadRequest creates a PcloudV2VolumescloneGetBadRequest with default headers values
func NewPcloudV2VolumescloneGetBadRequest() *PcloudV2VolumescloneGetBadRequest {
	return &PcloudV2VolumescloneGetBadRequest{}
}

/*PcloudV2VolumescloneGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudV2VolumescloneGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumescloneGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetNotFound creates a PcloudV2VolumescloneGetNotFound with default headers values
func NewPcloudV2VolumescloneGetNotFound() *PcloudV2VolumescloneGetNotFound {
	return &PcloudV2VolumescloneGetNotFound{}
}

/*PcloudV2VolumescloneGetNotFound handles this case with default header values.

Not Found
*/
type PcloudV2VolumescloneGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetInternalServerError creates a PcloudV2VolumescloneGetInternalServerError with default headers values
func NewPcloudV2VolumescloneGetInternalServerError() *PcloudV2VolumescloneGetInternalServerError {
	return &PcloudV2VolumescloneGetInternalServerError{}
}

/*PcloudV2VolumescloneGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}][%d] pcloudV2VolumescloneGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
