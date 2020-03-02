// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// RefreshTokenReader is a Reader for the RefreshToken structure.
type RefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRefreshTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRefreshTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRefreshTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRefreshTokenOK creates a RefreshTokenOK with default headers values
func NewRefreshTokenOK() *RefreshTokenOK {
	return &RefreshTokenOK{}
}

/*RefreshTokenOK handles this case with default header values.

A successful response.
*/
type RefreshTokenOK struct {
	Payload *models.AccountLoginReply
}

func (o *RefreshTokenOK) Error() string {
	return fmt.Sprintf("[PUT /v1/account/refresh][%d] refreshTokenOK  %+v", 200, o.Payload)
}

func (o *RefreshTokenOK) GetPayload() *models.AccountLoginReply {
	return o.Payload
}

func (o *RefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountLoginReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshTokenBadRequest creates a RefreshTokenBadRequest with default headers values
func NewRefreshTokenBadRequest() *RefreshTokenBadRequest {
	return &RefreshTokenBadRequest{}
}

/*RefreshTokenBadRequest handles this case with default header values.

Validation error
*/
type RefreshTokenBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *RefreshTokenBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/account/refresh][%d] refreshTokenBadRequest  %+v", 400, o.Payload)
}

func (o *RefreshTokenBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *RefreshTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshTokenForbidden creates a RefreshTokenForbidden with default headers values
func NewRefreshTokenForbidden() *RefreshTokenForbidden {
	return &RefreshTokenForbidden{}
}

/*RefreshTokenForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type RefreshTokenForbidden struct {
	Payload *models.CommonError
}

func (o *RefreshTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/account/refresh][%d] refreshTokenForbidden  %+v", 403, o.Payload)
}

func (o *RefreshTokenForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *RefreshTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshTokenNotFound creates a RefreshTokenNotFound with default headers values
func NewRefreshTokenNotFound() *RefreshTokenNotFound {
	return &RefreshTokenNotFound{}
}

/*RefreshTokenNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type RefreshTokenNotFound struct {
	Payload *models.CommonError
}

func (o *RefreshTokenNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/account/refresh][%d] refreshTokenNotFound  %+v", 404, o.Payload)
}

func (o *RefreshTokenNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *RefreshTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
