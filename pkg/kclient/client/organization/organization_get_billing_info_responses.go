// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/kclient/models"
)

// OrganizationGetBillingInfoReader is a Reader for the OrganizationGetBillingInfo structure.
type OrganizationGetBillingInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationGetBillingInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationGetBillingInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationGetBillingInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationGetBillingInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationGetBillingInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrganizationGetBillingInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationGetBillingInfoOK creates a OrganizationGetBillingInfoOK with default headers values
func NewOrganizationGetBillingInfoOK() *OrganizationGetBillingInfoOK {
	return &OrganizationGetBillingInfoOK{}
}

/*OrganizationGetBillingInfoOK handles this case with default header values.

A successful response.
*/
type OrganizationGetBillingInfoOK struct {
	Payload *models.AccountBillingInfoReply
}

func (o *OrganizationGetBillingInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/account/billing][%d] organizationGetBillingInfoOK  %+v", 200, o.Payload)
}

func (o *OrganizationGetBillingInfoOK) GetPayload() *models.AccountBillingInfoReply {
	return o.Payload
}

func (o *OrganizationGetBillingInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountBillingInfoReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetBillingInfoBadRequest creates a OrganizationGetBillingInfoBadRequest with default headers values
func NewOrganizationGetBillingInfoBadRequest() *OrganizationGetBillingInfoBadRequest {
	return &OrganizationGetBillingInfoBadRequest{}
}

/*OrganizationGetBillingInfoBadRequest handles this case with default header values.

Validation error
*/
type OrganizationGetBillingInfoBadRequest struct {
	Payload *models.CommonErrorWithFields
}

func (o *OrganizationGetBillingInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/account/billing][%d] organizationGetBillingInfoBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationGetBillingInfoBadRequest) GetPayload() *models.CommonErrorWithFields {
	return o.Payload
}

func (o *OrganizationGetBillingInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonErrorWithFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetBillingInfoForbidden creates a OrganizationGetBillingInfoForbidden with default headers values
func NewOrganizationGetBillingInfoForbidden() *OrganizationGetBillingInfoForbidden {
	return &OrganizationGetBillingInfoForbidden{}
}

/*OrganizationGetBillingInfoForbidden handles this case with default header values.

Returned when the user does not have permission to access the resource.
*/
type OrganizationGetBillingInfoForbidden struct {
	Payload *models.CommonError
}

func (o *OrganizationGetBillingInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/account/billing][%d] organizationGetBillingInfoForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationGetBillingInfoForbidden) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationGetBillingInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetBillingInfoNotFound creates a OrganizationGetBillingInfoNotFound with default headers values
func NewOrganizationGetBillingInfoNotFound() *OrganizationGetBillingInfoNotFound {
	return &OrganizationGetBillingInfoNotFound{}
}

/*OrganizationGetBillingInfoNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type OrganizationGetBillingInfoNotFound struct {
	Payload *models.CommonError
}

func (o *OrganizationGetBillingInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/account/billing][%d] organizationGetBillingInfoNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationGetBillingInfoNotFound) GetPayload() *models.CommonError {
	return o.Payload
}

func (o *OrganizationGetBillingInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationGetBillingInfoDefault creates a OrganizationGetBillingInfoDefault with default headers values
func NewOrganizationGetBillingInfoDefault(code int) *OrganizationGetBillingInfoDefault {
	return &OrganizationGetBillingInfoDefault{
		_statusCode: code,
	}
}

/*OrganizationGetBillingInfoDefault handles this case with default header values.

An unexpected error response
*/
type OrganizationGetBillingInfoDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the organization get billing info default response
func (o *OrganizationGetBillingInfoDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationGetBillingInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1/account/billing][%d] organization_GetBillingInfo default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationGetBillingInfoDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *OrganizationGetBillingInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}