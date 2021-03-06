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

// PcloudV2VolumescloneCancelPostReader is a Reader for the PcloudV2VolumescloneCancelPost structure.
type PcloudV2VolumescloneCancelPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneCancelPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPcloudV2VolumescloneCancelPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPcloudV2VolumescloneCancelPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudV2VolumescloneCancelPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudV2VolumescloneCancelPostAccepted creates a PcloudV2VolumescloneCancelPostAccepted with default headers values
func NewPcloudV2VolumescloneCancelPostAccepted() *PcloudV2VolumescloneCancelPostAccepted {
	return &PcloudV2VolumescloneCancelPostAccepted{}
}

/*PcloudV2VolumescloneCancelPostAccepted handles this case with default header values.

Accepted
*/
type PcloudV2VolumescloneCancelPostAccepted struct {
	Payload *models.VolumesClone
}

func (o *PcloudV2VolumescloneCancelPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV2VolumescloneCancelPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostNotFound creates a PcloudV2VolumescloneCancelPostNotFound with default headers values
func NewPcloudV2VolumescloneCancelPostNotFound() *PcloudV2VolumescloneCancelPostNotFound {
	return &PcloudV2VolumescloneCancelPostNotFound{}
}

/*PcloudV2VolumescloneCancelPostNotFound handles this case with default header values.

Not Found
*/
type PcloudV2VolumescloneCancelPostNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneCancelPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneCancelPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostInternalServerError creates a PcloudV2VolumescloneCancelPostInternalServerError with default headers values
func NewPcloudV2VolumescloneCancelPostInternalServerError() *PcloudV2VolumescloneCancelPostInternalServerError {
	return &PcloudV2VolumescloneCancelPostInternalServerError{}
}

/*PcloudV2VolumescloneCancelPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneCancelPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneCancelPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneCancelPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
