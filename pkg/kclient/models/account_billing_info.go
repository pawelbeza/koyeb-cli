// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountBillingInfo account billing info
// swagger:model accountBillingInfo
type AccountBillingInfo struct {

	// address1
	Address1 string `json:"address1,omitempty"`

	// address2
	Address2 string `json:"address2,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// company
	Company bool `json:"company,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// plan
	Plan string `json:"plan,omitempty"`

	// postal code
	PostalCode string `json:"postal_code,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// vat number
	VatNumber string `json:"vat_number,omitempty"`
}

// Validate validates this account billing info
func (m *AccountBillingInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountBillingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountBillingInfo) UnmarshalBinary(b []byte) error {
	var res AccountBillingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}