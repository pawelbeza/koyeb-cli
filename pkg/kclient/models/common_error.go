// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonError common error
// swagger:model commonError
type CommonError struct {

	// code
	Code string `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this common error
func (m *CommonError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonError) UnmarshalBinary(b []byte) error {
	var res CommonError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
