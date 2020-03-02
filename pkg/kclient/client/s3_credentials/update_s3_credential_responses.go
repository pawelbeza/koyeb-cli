// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// UpdateS3CredentialReader is a Reader for the UpdateS3Credential structure.
type UpdateS3CredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateS3CredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateS3CredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateS3CredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateS3CredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateS3CredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateS3CredentialOK creates a UpdateS3CredentialOK with default headers values
func NewUpdateS3CredentialOK() *UpdateS3CredentialOK {
	return &UpdateS3CredentialOK{}
}

/*UpdateS3CredentialOK handles this case with default header values.

A successful response.
*/
type UpdateS3CredentialOK struct {
	Payload *models.AccountS3CredentialReply
}

func (o *UpdateS3CredentialOK) Error() string {
	return fmt.Sprintf("[PUT /v1/s3_credentials/{id}][%d] updateS3CredentialOK  %+v", 200, o.Payload)
}

func (o *UpdateS3CredentialOK) GetPayload() *models.AccountS3CredentialReply {
	return o.Payload
}

func (o *UpdateS3CredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountS3CredentialReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateS3CredentialBadRequest creates a UpdateS3CredentialBadRequest with default headers values
func NewUpdateS3CredentialBadRequest() *UpdateS3CredentialBadRequest {
	return &UpdateS3CredentialBadRequest{}
}

/*UpdateS3CredentialBadRequest handles this case with default header values.

Validation error
*/
type UpdateS3CredentialBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *UpdateS3CredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/s3_credentials/{id}][%d] updateS3CredentialBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateS3CredentialBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *UpdateS3CredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateS3CredentialForbidden creates a UpdateS3CredentialForbidden with default headers values
func NewUpdateS3CredentialForbidden() *UpdateS3CredentialForbidden {
	return &UpdateS3CredentialForbidden{}
}

/*UpdateS3CredentialForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type UpdateS3CredentialForbidden struct {
	Payload *models.CommonError
}

func (o *UpdateS3CredentialForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/s3_credentials/{id}][%d] updateS3CredentialForbidden  %+v", 403, o.Payload)
}

func (o *UpdateS3CredentialForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *UpdateS3CredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateS3CredentialNotFound creates a UpdateS3CredentialNotFound with default headers values
func NewUpdateS3CredentialNotFound() *UpdateS3CredentialNotFound {
	return &UpdateS3CredentialNotFound{}
}

/*UpdateS3CredentialNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type UpdateS3CredentialNotFound struct {
	Payload *models.CommonError
}

func (o *UpdateS3CredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/s3_credentials/{id}][%d] updateS3CredentialNotFound  %+v", 404, o.Payload)
}

func (o *UpdateS3CredentialNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *UpdateS3CredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
