// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountBillingInfoReply account billing info reply
// swagger:model accountBillingInfoReply
type AccountBillingInfoReply struct {

	// billing info
	BillingInfo *AccountBillingInfo `json:"billing_info,omitempty"`
}

// Validate validates this account billing info reply
func (m *AccountBillingInfoReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountBillingInfoReply) validateBillingInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingInfo) { // not required
		return nil
	}

	if m.BillingInfo != nil {
		if err := m.BillingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountBillingInfoReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountBillingInfoReply) UnmarshalBinary(b []byte) error {
	var res AccountBillingInfoReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
