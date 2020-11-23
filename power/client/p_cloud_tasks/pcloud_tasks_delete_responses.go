// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bkhadars/power-go-client/power/models"
)

// PcloudTasksDeleteReader is a Reader for the PcloudTasksDelete structure.
type PcloudTasksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTasksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudTasksDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudTasksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewPcloudTasksDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudTasksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudTasksDeleteOK creates a PcloudTasksDeleteOK with default headers values
func NewPcloudTasksDeleteOK() *PcloudTasksDeleteOK {
	return &PcloudTasksDeleteOK{}
}

/*PcloudTasksDeleteOK handles this case with default header values.

OK
*/
type PcloudTasksDeleteOK struct {
	Payload models.Object
}

func (o *PcloudTasksDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tasks/{task_id}][%d] pcloudTasksDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudTasksDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksDeleteBadRequest creates a PcloudTasksDeleteBadRequest with default headers values
func NewPcloudTasksDeleteBadRequest() *PcloudTasksDeleteBadRequest {
	return &PcloudTasksDeleteBadRequest{}
}

/*PcloudTasksDeleteBadRequest handles this case with default header values.

Bad Request
*/
type PcloudTasksDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTasksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tasks/{task_id}][%d] pcloudTasksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTasksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksDeleteGone creates a PcloudTasksDeleteGone with default headers values
func NewPcloudTasksDeleteGone() *PcloudTasksDeleteGone {
	return &PcloudTasksDeleteGone{}
}

/*PcloudTasksDeleteGone handles this case with default header values.

Gone
*/
type PcloudTasksDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudTasksDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tasks/{task_id}][%d] pcloudTasksDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudTasksDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksDeleteInternalServerError creates a PcloudTasksDeleteInternalServerError with default headers values
func NewPcloudTasksDeleteInternalServerError() *PcloudTasksDeleteInternalServerError {
	return &PcloudTasksDeleteInternalServerError{}
}

/*PcloudTasksDeleteInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudTasksDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTasksDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tasks/{task_id}][%d] pcloudTasksDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTasksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
