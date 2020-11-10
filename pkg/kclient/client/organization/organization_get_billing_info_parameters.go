// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrganizationGetBillingInfoParams creates a new OrganizationGetBillingInfoParams object
// with the default values initialized.
func NewOrganizationGetBillingInfoParams() *OrganizationGetBillingInfoParams {

	return &OrganizationGetBillingInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationGetBillingInfoParamsWithTimeout creates a new OrganizationGetBillingInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrganizationGetBillingInfoParamsWithTimeout(timeout time.Duration) *OrganizationGetBillingInfoParams {

	return &OrganizationGetBillingInfoParams{

		timeout: timeout,
	}
}

// NewOrganizationGetBillingInfoParamsWithContext creates a new OrganizationGetBillingInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrganizationGetBillingInfoParamsWithContext(ctx context.Context) *OrganizationGetBillingInfoParams {

	return &OrganizationGetBillingInfoParams{

		Context: ctx,
	}
}

// NewOrganizationGetBillingInfoParamsWithHTTPClient creates a new OrganizationGetBillingInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrganizationGetBillingInfoParamsWithHTTPClient(client *http.Client) *OrganizationGetBillingInfoParams {

	return &OrganizationGetBillingInfoParams{
		HTTPClient: client,
	}
}

/*OrganizationGetBillingInfoParams contains all the parameters to send to the API endpoint
for the organization get billing info operation typically these are written to a http.Request
*/
type OrganizationGetBillingInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the organization get billing info params
func (o *OrganizationGetBillingInfoParams) WithTimeout(timeout time.Duration) *OrganizationGetBillingInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization get billing info params
func (o *OrganizationGetBillingInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization get billing info params
func (o *OrganizationGetBillingInfoParams) WithContext(ctx context.Context) *OrganizationGetBillingInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization get billing info params
func (o *OrganizationGetBillingInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization get billing info params
func (o *OrganizationGetBillingInfoParams) WithHTTPClient(client *http.Client) *OrganizationGetBillingInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization get billing info params
func (o *OrganizationGetBillingInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationGetBillingInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}