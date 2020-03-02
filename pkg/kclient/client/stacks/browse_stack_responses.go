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

// BrowseStackReader is a Reader for the BrowseStack structure.
type BrowseStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BrowseStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBrowseStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBrowseStackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBrowseStackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBrowseStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBrowseStackOK creates a BrowseStackOK with default headers values
func NewBrowseStackOK() *BrowseStackOK {
	return &BrowseStackOK{}
}

/*BrowseStackOK handles this case with default header values.

A successful response.
*/
type BrowseStackOK struct {
	Payload *models.StorageBrowseStackReply
}

func (o *BrowseStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/browse/{path}][%d] browseStackOK  %+v", 200, o.Payload)
}

func (o *BrowseStackOK) GetPayload() *models.StorageBrowseStackReply {
	return o.Payload
}

func (o *BrowseStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageBrowseStackReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBrowseStackBadRequest creates a BrowseStackBadRequest with default headers values
func NewBrowseStackBadRequest() *BrowseStackBadRequest {
	return &BrowseStackBadRequest{}
}

/*BrowseStackBadRequest handles this case with default header values.

Validation error
*/
type BrowseStackBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *BrowseStackBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/browse/{path}][%d] browseStackBadRequest  %+v", 400, o.Payload)
}

func (o *BrowseStackBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *BrowseStackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBrowseStackForbidden creates a BrowseStackForbidden with default headers values
func NewBrowseStackForbidden() *BrowseStackForbidden {
	return &BrowseStackForbidden{}
}

/*BrowseStackForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type BrowseStackForbidden struct {
	Payload *models.CommonError
}

func (o *BrowseStackForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/browse/{path}][%d] browseStackForbidden  %+v", 403, o.Payload)
}

func (o *BrowseStackForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *BrowseStackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBrowseStackNotFound creates a BrowseStackNotFound with default headers values
func NewBrowseStackNotFound() *BrowseStackNotFound {
	return &BrowseStackNotFound{}
}

/*BrowseStackNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type BrowseStackNotFound struct {
	Payload *models.CommonError
}

func (o *BrowseStackNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/browse/{path}][%d] browseStackNotFound  %+v", 404, o.Payload)
}

func (o *BrowseStackNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *BrowseStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
