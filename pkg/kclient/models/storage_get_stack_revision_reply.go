// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageGetStackRevisionReply storage get stack revision reply
// swagger:model storageGetStackRevisionReply
type StorageGetStackRevisionReply struct {

	// revision
	Revision *StorageStackRevision `json:"revision,omitempty"`
}

// Validate validates this storage get stack revision reply
func (m *StorageGetStackRevisionReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageGetStackRevisionReply) validateRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageGetStackRevisionReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageGetStackRevisionReply) UnmarshalBinary(b []byte) error {
	var res StorageGetStackRevisionReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}