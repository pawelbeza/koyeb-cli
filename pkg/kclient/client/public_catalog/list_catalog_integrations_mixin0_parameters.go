// Code generated by go-swagger; DO NOT EDIT.

package public_catalog

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

// NewListCatalogIntegrationsMixin0Params creates a new ListCatalogIntegrationsMixin0Params object
// with the default values initialized.
func NewListCatalogIntegrationsMixin0Params() *ListCatalogIntegrationsMixin0Params {
	var ()
	return &ListCatalogIntegrationsMixin0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCatalogIntegrationsMixin0ParamsWithTimeout creates a new ListCatalogIntegrationsMixin0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCatalogIntegrationsMixin0ParamsWithTimeout(timeout time.Duration) *ListCatalogIntegrationsMixin0Params {
	var ()
	return &ListCatalogIntegrationsMixin0Params{

		timeout: timeout,
	}
}

// NewListCatalogIntegrationsMixin0ParamsWithContext creates a new ListCatalogIntegrationsMixin0Params object
// with the default values initialized, and the ability to set a context for a request
func NewListCatalogIntegrationsMixin0ParamsWithContext(ctx context.Context) *ListCatalogIntegrationsMixin0Params {
	var ()
	return &ListCatalogIntegrationsMixin0Params{

		Context: ctx,
	}
}

// NewListCatalogIntegrationsMixin0ParamsWithHTTPClient creates a new ListCatalogIntegrationsMixin0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCatalogIntegrationsMixin0ParamsWithHTTPClient(client *http.Client) *ListCatalogIntegrationsMixin0Params {
	var ()
	return &ListCatalogIntegrationsMixin0Params{
		HTTPClient: client,
	}
}

/*ListCatalogIntegrationsMixin0Params contains all the parameters to send to the API endpoint
for the list catalog integrations mixin0 operation typically these are written to a http.Request
*/
type ListCatalogIntegrationsMixin0Params struct {

	/*Limit*/
	Limit *string
	/*Name*/
	Name *string
	/*Offset*/
	Offset *string
	/*Tags*/
	Tags *string
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithTimeout(timeout time.Duration) *ListCatalogIntegrationsMixin0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithContext(ctx context.Context) *ListCatalogIntegrationsMixin0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithHTTPClient(client *http.Client) *ListCatalogIntegrationsMixin0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithLimit(limit *string) *ListCatalogIntegrationsMixin0Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithName(name *string) *ListCatalogIntegrationsMixin0Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithOffset(offset *string) *ListCatalogIntegrationsMixin0Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetOffset(offset *string) {
	o.Offset = offset
}

// WithTags adds the tags to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithTags(tags *string) *ListCatalogIntegrationsMixin0Params {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetTags(tags *string) {
	o.Tags = tags
}

// WithType adds the typeVar to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) WithType(typeVar *string) *ListCatalogIntegrationsMixin0Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list catalog integrations mixin0 params
func (o *ListCatalogIntegrationsMixin0Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListCatalogIntegrationsMixin0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Tags != nil {

		// query param tags
		var qrTags string
		if o.Tags != nil {
			qrTags = *o.Tags
		}
		qTags := qrTags
		if qTags != "" {
			if err := r.SetQueryParam("tags", qTags); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
