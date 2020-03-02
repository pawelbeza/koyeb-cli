// Code generated by go-swagger; DO NOT EDIT.

package managed_stores

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// ListManagedStoresReader is a Reader for the ListManagedStores structure.
type ListManagedStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListManagedStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListManagedStoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListManagedStoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListManagedStoresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListManagedStoresNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListManagedStoresOK creates a ListManagedStoresOK with default headers values
func NewListManagedStoresOK() *ListManagedStoresOK {
	return &ListManagedStoresOK{}
}

/*ListManagedStoresOK handles this case with default header values.

A successful response.
*/
type ListManagedStoresOK struct {
	Payload *models.StorageListManagedStoresReply
}

func (o *ListManagedStoresOK) Error() string {
	return fmt.Sprintf("[GET /v1/managed_stores][%d] listManagedStoresOK  %+v", 200, o.Payload)
}

func (o *ListManagedStoresOK) GetPayload() *models.StorageListManagedStoresReply {
	return o.Payload
}

func (o *ListManagedStoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageListManagedStoresReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManagedStoresBadRequest creates a ListManagedStoresBadRequest with default headers values
func NewListManagedStoresBadRequest() *ListManagedStoresBadRequest {
	return &ListManagedStoresBadRequest{}
}

/*ListManagedStoresBadRequest handles this case with default header values.

Validation error
*/
type ListManagedStoresBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *ListManagedStoresBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/managed_stores][%d] listManagedStoresBadRequest  %+v", 400, o.Payload)
}

func (o *ListManagedStoresBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *ListManagedStoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManagedStoresForbidden creates a ListManagedStoresForbidden with default headers values
func NewListManagedStoresForbidden() *ListManagedStoresForbidden {
	return &ListManagedStoresForbidden{}
}

/*ListManagedStoresForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type ListManagedStoresForbidden struct {
	Payload *models.CommonError
}

func (o *ListManagedStoresForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/managed_stores][%d] listManagedStoresForbidden  %+v", 403, o.Payload)
}

func (o *ListManagedStoresForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListManagedStoresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManagedStoresNotFound creates a ListManagedStoresNotFound with default headers values
func NewListManagedStoresNotFound() *ListManagedStoresNotFound {
	return &ListManagedStoresNotFound{}
}

/*ListManagedStoresNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type ListManagedStoresNotFound struct {
	Payload *models.CommonError
}

func (o *ListManagedStoresNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/managed_stores][%d] listManagedStoresNotFound  %+v", 404, o.Payload)
}

func (o *ListManagedStoresNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *ListManagedStoresNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
