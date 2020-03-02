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

// DeleteS3CredentialReader is a Reader for the DeleteS3Credential structure.
type DeleteS3CredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteS3CredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteS3CredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteS3CredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteS3CredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteS3CredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteS3CredentialOK creates a DeleteS3CredentialOK with default headers values
func NewDeleteS3CredentialOK() *DeleteS3CredentialOK {
	return &DeleteS3CredentialOK{}
}

/*DeleteS3CredentialOK handles this case with default header values.

A successful response.
*/
type DeleteS3CredentialOK struct {
	Payload models.CommonEmpty
}

func (o *DeleteS3CredentialOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/s3_credentials/{id}][%d] deleteS3CredentialOK  %+v", 200, o.Payload)
}

func (o *DeleteS3CredentialOK) GetPayload() models.CommonEmpty {
	return o.Payload
}

func (o *DeleteS3CredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteS3CredentialBadRequest creates a DeleteS3CredentialBadRequest with default headers values
func NewDeleteS3CredentialBadRequest() *DeleteS3CredentialBadRequest {
	return &DeleteS3CredentialBadRequest{}
}

/*DeleteS3CredentialBadRequest handles this case with default header values.

Validation error
*/
type DeleteS3CredentialBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *DeleteS3CredentialBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/s3_credentials/{id}][%d] deleteS3CredentialBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteS3CredentialBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *DeleteS3CredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteS3CredentialForbidden creates a DeleteS3CredentialForbidden with default headers values
func NewDeleteS3CredentialForbidden() *DeleteS3CredentialForbidden {
	return &DeleteS3CredentialForbidden{}
}

/*DeleteS3CredentialForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type DeleteS3CredentialForbidden struct {
	Payload *models.CommonError
}

func (o *DeleteS3CredentialForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/s3_credentials/{id}][%d] deleteS3CredentialForbidden  %+v", 403, o.Payload)
}

func (o *DeleteS3CredentialForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *DeleteS3CredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteS3CredentialNotFound creates a DeleteS3CredentialNotFound with default headers values
func NewDeleteS3CredentialNotFound() *DeleteS3CredentialNotFound {
	return &DeleteS3CredentialNotFound{}
}

/*DeleteS3CredentialNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type DeleteS3CredentialNotFound struct {
	Payload *models.CommonError
}

func (o *DeleteS3CredentialNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/s3_credentials/{id}][%d] deleteS3CredentialNotFound  %+v", 404, o.Payload)
}

func (o *DeleteS3CredentialNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *DeleteS3CredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
