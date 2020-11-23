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

// PcloudV2VolumescloneStartPostReader is a Reader for the PcloudV2VolumescloneStartPost structure.
type PcloudV2VolumescloneStartPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneStartPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudV2VolumescloneStartPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPcloudV2VolumescloneStartPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudV2VolumescloneStartPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudV2VolumescloneStartPostOK creates a PcloudV2VolumescloneStartPostOK with default headers values
func NewPcloudV2VolumescloneStartPostOK() *PcloudV2VolumescloneStartPostOK {
	return &PcloudV2VolumescloneStartPostOK{}
}

/*PcloudV2VolumescloneStartPostOK handles this case with default header values.

OK
*/
type PcloudV2VolumescloneStartPostOK struct {
	Payload *models.VolumesClone
}

func (o *PcloudV2VolumescloneStartPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneStartPostNotFound creates a PcloudV2VolumescloneStartPostNotFound with default headers values
func NewPcloudV2VolumescloneStartPostNotFound() *PcloudV2VolumescloneStartPostNotFound {
	return &PcloudV2VolumescloneStartPostNotFound{}
}

/*PcloudV2VolumescloneStartPostNotFound handles this case with default header values.

Not Found
*/
type PcloudV2VolumescloneStartPostNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneStartPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneStartPostInternalServerError creates a PcloudV2VolumescloneStartPostInternalServerError with default headers values
func NewPcloudV2VolumescloneStartPostInternalServerError() *PcloudV2VolumescloneStartPostInternalServerError {
	return &PcloudV2VolumescloneStartPostInternalServerError{}
}

/*PcloudV2VolumescloneStartPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneStartPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneStartPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
