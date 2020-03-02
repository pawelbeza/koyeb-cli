// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ListCatalogServicesMixin0Reader is a Reader for the ListCatalogServicesMixin0 structure.
type ListCatalogServicesMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCatalogServicesMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCatalogServicesMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCatalogServicesMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCatalogServicesMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCatalogServicesMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCatalogServicesMixin0OK creates a ListCatalogServicesMixin0OK with default headers values
func NewListCatalogServicesMixin0OK() *ListCatalogServicesMixin0OK {
	return &ListCatalogServicesMixin0OK{}
}

/*ListCatalogServicesMixin0OK handles this case with default header values.

A successful response.
*/
type ListCatalogServicesMixin0OK struct {
	Payload *models.StorageListCatalogServicesReply
}

func (o *ListCatalogServicesMixin0OK) Error() string {
	return fmt.Sprintf("[GET /v1/public/catalog/services][%d] listCatalogServicesMixin0OK  %+v", 200, o.Payload)
}

func (o *ListCatalogServicesMixin0OK) GetPayload() *models.StorageListCatalogServicesReply {
	return o.Payload
}

func (o *ListCatalogServicesMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageListCatalogServicesReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServicesMixin0BadRequest creates a ListCatalogServicesMixin0BadRequest with default headers values
func NewListCatalogServicesMixin0BadRequest() *ListCatalogServicesMixin0BadRequest {
	return &ListCatalogServicesMixin0BadRequest{}
}

/*ListCatalogServicesMixin0BadRequest handles this case with default header values.

Validation error
*/
type ListCatalogServicesMixin0BadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ListCatalogServicesMixin0BadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/public/catalog/services][%d] listCatalogServicesMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *ListCatalogServicesMixin0BadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ListCatalogServicesMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServicesMixin0Forbidden creates a ListCatalogServicesMixin0Forbidden with default headers values
func NewListCatalogServicesMixin0Forbidden() *ListCatalogServicesMixin0Forbidden {
	return &ListCatalogServicesMixin0Forbidden{}
}

/*ListCatalogServicesMixin0Forbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ListCatalogServicesMixin0Forbidden struct {
	Payload *models.CommonError
}

func (o *ListCatalogServicesMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /v1/public/catalog/services][%d] listCatalogServicesMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *ListCatalogServicesMixin0Forbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListCatalogServicesMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServicesMixin0NotFound creates a ListCatalogServicesMixin0NotFound with default headers values
func NewListCatalogServicesMixin0NotFound() *ListCatalogServicesMixin0NotFound {
	return &ListCatalogServicesMixin0NotFound{}
}

/*ListCatalogServicesMixin0NotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ListCatalogServicesMixin0NotFound struct {
	Payload *models.CommonError
}

func (o *ListCatalogServicesMixin0NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/public/catalog/services][%d] listCatalogServicesMixin0NotFound  %+v", 404, o.Payload)
}

func (o *ListCatalogServicesMixin0NotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListCatalogServicesMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
