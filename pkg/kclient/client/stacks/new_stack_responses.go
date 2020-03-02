// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// NewStackReader is a Reader for the NewStack structure.
type NewStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NewStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNewStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNewStackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNewStackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNewStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNewStackOK creates a NewStackOK with default headers values
func NewNewStackOK() *NewStackOK {
	return &NewStackOK{}
}

/*NewStackOK handles this case with default header values.

A successful response.
*/
type NewStackOK struct {
	Payload *models.StorageStackReply
}

func (o *NewStackOK) Error() string {
	return fmt.Sprintf("[POST /v1/stacks][%d] newStackOK  %+v", 200, o.Payload)
}

func (o *NewStackOK) GetPayload() *models.StorageStackReply {
	return o.Payload
}

func (o *NewStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageStackReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewStackBadRequest creates a NewStackBadRequest with default headers values
func NewNewStackBadRequest() *NewStackBadRequest {
	return &NewStackBadRequest{}
}

/*NewStackBadRequest handles this case with default header values.

Validation error
*/
type NewStackBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *NewStackBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/stacks][%d] newStackBadRequest  %+v", 400, o.Payload)
}

func (o *NewStackBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *NewStackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewStackForbidden creates a NewStackForbidden with default headers values
func NewNewStackForbidden() *NewStackForbidden {
	return &NewStackForbidden{}
}

/*NewStackForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type NewStackForbidden struct {
	Payload *models.CommonError
}

func (o *NewStackForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/stacks][%d] newStackForbidden  %+v", 403, o.Payload)
}

func (o *NewStackForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *NewStackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewStackNotFound creates a NewStackNotFound with default headers values
func NewNewStackNotFound() *NewStackNotFound {
	return &NewStackNotFound{}
}

/*NewStackNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type NewStackNotFound struct {
	Payload *models.CommonError
}

func (o *NewStackNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/stacks][%d] newStackNotFound  %+v", 404, o.Payload)
}

func (o *NewStackNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *NewStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
