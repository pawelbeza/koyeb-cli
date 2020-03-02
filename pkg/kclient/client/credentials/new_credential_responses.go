// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// NewCredentialReader is a Reader for the NewCredential structure.
type NewCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NewCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNewCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNewCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNewCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNewCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNewCredentialOK creates a NewCredentialOK with default headers values
func NewNewCredentialOK() *NewCredentialOK {
	return &NewCredentialOK{}
}

/*NewCredentialOK handles this case with default header values.

A successful response.
*/
type NewCredentialOK struct {
	Payload *models.AccountCredentialReply
}

func (o *NewCredentialOK) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] newCredentialOK  %+v", 200, o.Payload)
}

func (o *NewCredentialOK) GetPayload() *models.AccountCredentialReply {
	return o.Payload
}

func (o *NewCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountCredentialReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewCredentialBadRequest creates a NewCredentialBadRequest with default headers values
func NewNewCredentialBadRequest() *NewCredentialBadRequest {
	return &NewCredentialBadRequest{}
}

/*NewCredentialBadRequest handles this case with default header values.

Validation error
*/
type NewCredentialBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *NewCredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] newCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *NewCredentialBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *NewCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewCredentialForbidden creates a NewCredentialForbidden with default headers values
func NewNewCredentialForbidden() *NewCredentialForbidden {
	return &NewCredentialForbidden{}
}

/*NewCredentialForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type NewCredentialForbidden struct {
	Payload *models.CommonError
}

func (o *NewCredentialForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] newCredentialForbidden  %+v", 403, o.Payload)
}

func (o *NewCredentialForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *NewCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNewCredentialNotFound creates a NewCredentialNotFound with default headers values
func NewNewCredentialNotFound() *NewCredentialNotFound {
	return &NewCredentialNotFound{}
}

/*NewCredentialNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type NewCredentialNotFound struct {
	Payload *models.CommonError
}

func (o *NewCredentialNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/credentials][%d] newCredentialNotFound  %+v", 404, o.Payload)
}

func (o *NewCredentialNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *NewCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
