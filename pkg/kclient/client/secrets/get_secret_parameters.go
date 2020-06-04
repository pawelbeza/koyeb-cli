// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

// NewGetSecretParams creates a new GetSecretParams object
// with the default values initialized.
func NewGetSecretParams() *GetSecretParams {
	var ()
	return &GetSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecretParamsWithTimeout creates a new GetSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecretParamsWithTimeout(timeout time.Duration) *GetSecretParams {
	var ()
	return &GetSecretParams{

		timeout: timeout,
	}
}

// NewGetSecretParamsWithContext creates a new GetSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecretParamsWithContext(ctx context.Context) *GetSecretParams {
	var ()
	return &GetSecretParams{

		Context: ctx,
	}
}

// NewGetSecretParamsWithHTTPClient creates a new GetSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecretParamsWithHTTPClient(client *http.Client) *GetSecretParams {
	var ()
	return &GetSecretParams{
		HTTPClient: client,
	}
}

/*GetSecretParams contains all the parameters to send to the API endpoint
for the get secret operation typically these are written to a http.Request
*/
type GetSecretParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get secret params
func (o *GetSecretParams) WithTimeout(timeout time.Duration) *GetSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get secret params
func (o *GetSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get secret params
func (o *GetSecretParams) WithContext(ctx context.Context) *GetSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get secret params
func (o *GetSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get secret params
func (o *GetSecretParams) WithHTTPClient(client *http.Client) *GetSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get secret params
func (o *GetSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get secret params
func (o *GetSecretParams) WithID(id string) *GetSecretParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get secret params
func (o *GetSecretParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
