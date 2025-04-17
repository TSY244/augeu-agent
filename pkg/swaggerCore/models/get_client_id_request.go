// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetClientIDRequest get client Id request
//
// swagger:model GetClientIdRequest
type GetClientIDRequest struct {

	// client info
	// Required: true
	ClientInfo *ClientInfo `json:"client_info"`

	// 密钥
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this get client Id request
func (m *GetClientIDRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientIDRequest) validateClientInfo(formats strfmt.Registry) error {

	if err := validate.Required("client_info", "body", m.ClientInfo); err != nil {
		return err
	}

	if m.ClientInfo != nil {
		if err := m.ClientInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client_info")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientIDRequest) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", m.Secret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get client Id request based on the context it is used
func (m *GetClientIDRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientIDRequest) contextValidateClientInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientInfo != nil {

		if err := m.ClientInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientIDRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientIDRequest) UnmarshalBinary(b []byte) error {
	var res GetClientIDRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
