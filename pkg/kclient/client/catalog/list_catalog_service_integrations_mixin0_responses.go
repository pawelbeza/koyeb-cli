// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ListCatalogServiceIntegrationsMixin0Reader is a Reader for the ListCatalogServiceIntegrationsMixin0 structure.
type ListCatalogServiceIntegrationsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCatalogServiceIntegrationsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCatalogServiceIntegrationsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCatalogServiceIntegrationsMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCatalogServiceIntegrationsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCatalogServiceIntegrationsMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCatalogServiceIntegrationsMixin0OK creates a ListCatalogServiceIntegrationsMixin0OK with default headers values
func NewListCatalogServiceIntegrationsMixin0OK() *ListCatalogServiceIntegrationsMixin0OK {
	return &ListCatalogServiceIntegrationsMixin0OK{}
}

/*ListCatalogServiceIntegrationsMixin0OK handles this case with default header values.

A successful response.
*/
type ListCatalogServiceIntegrationsMixin0OK struct {
	Payload *models.StorageListCatalogIntegrationsReply
}

func (o *ListCatalogServiceIntegrationsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsMixin0OK  %+v", 200, o.Payload)
}

func (o *ListCatalogServiceIntegrationsMixin0OK) GetPayload() *models.StorageListCatalogIntegrationsReply {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageListCatalogIntegrationsReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServiceIntegrationsMixin0BadRequest creates a ListCatalogServiceIntegrationsMixin0BadRequest with default headers values
func NewListCatalogServiceIntegrationsMixin0BadRequest() *ListCatalogServiceIntegrationsMixin0BadRequest {
	return &ListCatalogServiceIntegrationsMixin0BadRequest{}
}

/*ListCatalogServiceIntegrationsMixin0BadRequest handles this case with default header values.

Validation error
*/
type ListCatalogServiceIntegrationsMixin0BadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ListCatalogServiceIntegrationsMixin0BadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *ListCatalogServiceIntegrationsMixin0BadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServiceIntegrationsMixin0Forbidden creates a ListCatalogServiceIntegrationsMixin0Forbidden with default headers values
func NewListCatalogServiceIntegrationsMixin0Forbidden() *ListCatalogServiceIntegrationsMixin0Forbidden {
	return &ListCatalogServiceIntegrationsMixin0Forbidden{}
}

/*ListCatalogServiceIntegrationsMixin0Forbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ListCatalogServiceIntegrationsMixin0Forbidden struct {
	Payload *models.CommonError
}

func (o *ListCatalogServiceIntegrationsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *ListCatalogServiceIntegrationsMixin0Forbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogServiceIntegrationsMixin0NotFound creates a ListCatalogServiceIntegrationsMixin0NotFound with default headers values
func NewListCatalogServiceIntegrationsMixin0NotFound() *ListCatalogServiceIntegrationsMixin0NotFound {
	return &ListCatalogServiceIntegrationsMixin0NotFound{}
}

/*ListCatalogServiceIntegrationsMixin0NotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ListCatalogServiceIntegrationsMixin0NotFound struct {
	Payload *models.CommonError
}

func (o *ListCatalogServiceIntegrationsMixin0NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/catalog/services/{id}/integrations][%d] listCatalogServiceIntegrationsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *ListCatalogServiceIntegrationsMixin0NotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListCatalogServiceIntegrationsMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
