// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// PcloudPvminstancesPostReader is a Reader for the PcloudPvminstancesPost structure.
type PcloudPvminstancesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPcloudPvminstancesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewPcloudPvminstancesPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPcloudPvminstancesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudPvminstancesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesPostOK creates a PcloudPvminstancesPostOK with default headers values
func NewPcloudPvminstancesPostOK() *PcloudPvminstancesPostOK {
	return &PcloudPvminstancesPostOK{}
}

/*PcloudPvminstancesPostOK handles this case with default header values.

OK
*/
type PcloudPvminstancesPostOK struct {
	Payload models.PVMInstanceList
}

func (o *PcloudPvminstancesPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostCreated creates a PcloudPvminstancesPostCreated with default headers values
func NewPcloudPvminstancesPostCreated() *PcloudPvminstancesPostCreated {
	return &PcloudPvminstancesPostCreated{}
}

/*PcloudPvminstancesPostCreated handles this case with default header values.

Created
*/
type PcloudPvminstancesPostCreated struct {
	Payload models.PVMInstanceList
}

func (o *PcloudPvminstancesPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostCreated  %+v", 201, o.Payload)
}

func (o *PcloudPvminstancesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostAccepted creates a PcloudPvminstancesPostAccepted with default headers values
func NewPcloudPvminstancesPostAccepted() *PcloudPvminstancesPostAccepted {
	return &PcloudPvminstancesPostAccepted{}
}

/*PcloudPvminstancesPostAccepted handles this case with default header values.

Accepted
*/
type PcloudPvminstancesPostAccepted struct {
	Payload models.PVMInstanceList
}

func (o *PcloudPvminstancesPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudPvminstancesPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostBadRequest creates a PcloudPvminstancesPostBadRequest with default headers values
func NewPcloudPvminstancesPostBadRequest() *PcloudPvminstancesPostBadRequest {
	return &PcloudPvminstancesPostBadRequest{}
}

/*PcloudPvminstancesPostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostConflict creates a PcloudPvminstancesPostConflict with default headers values
func NewPcloudPvminstancesPostConflict() *PcloudPvminstancesPostConflict {
	return &PcloudPvminstancesPostConflict{}
}

/*PcloudPvminstancesPostConflict handles this case with default header values.

Conflict
*/
type PcloudPvminstancesPostConflict struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostUnprocessableEntity creates a PcloudPvminstancesPostUnprocessableEntity with default headers values
func NewPcloudPvminstancesPostUnprocessableEntity() *PcloudPvminstancesPostUnprocessableEntity {
	return &PcloudPvminstancesPostUnprocessableEntity{}
}

/*PcloudPvminstancesPostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesPostInternalServerError creates a PcloudPvminstancesPostInternalServerError with default headers values
func NewPcloudPvminstancesPostInternalServerError() *PcloudPvminstancesPostInternalServerError {
	return &PcloudPvminstancesPostInternalServerError{}
}

/*PcloudPvminstancesPostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
