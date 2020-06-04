// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewListCatalogStoreParams creates a new ListCatalogStoreParams object
// with the default values initialized.
func NewListCatalogStoreParams() *ListCatalogStoreParams {

	return &ListCatalogStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCatalogStoreParamsWithTimeout creates a new ListCatalogStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCatalogStoreParamsWithTimeout(timeout time.Duration) *ListCatalogStoreParams {

	return &ListCatalogStoreParams{

		timeout: timeout,
	}
}

// NewListCatalogStoreParamsWithContext creates a new ListCatalogStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCatalogStoreParamsWithContext(ctx context.Context) *ListCatalogStoreParams {

	return &ListCatalogStoreParams{

		Context: ctx,
	}
}

// NewListCatalogStoreParamsWithHTTPClient creates a new ListCatalogStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCatalogStoreParamsWithHTTPClient(client *http.Client) *ListCatalogStoreParams {

	return &ListCatalogStoreParams{
		HTTPClient: client,
	}
}

/*ListCatalogStoreParams contains all the parameters to send to the API endpoint
for the list catalog store operation typically these are written to a http.Request
*/
type ListCatalogStoreParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list catalog store params
func (o *ListCatalogStoreParams) WithTimeout(timeout time.Duration) *ListCatalogStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list catalog store params
func (o *ListCatalogStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list catalog store params
func (o *ListCatalogStoreParams) WithContext(ctx context.Context) *ListCatalogStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list catalog store params
func (o *ListCatalogStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list catalog store params
func (o *ListCatalogStoreParams) WithHTTPClient(client *http.Client) *ListCatalogStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list catalog store params
func (o *ListCatalogStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCatalogStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
