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

// BadRequestError bad request error
//
// swagger:model BadRequestError
type BadRequestError struct {

	// 400
	// Required: true
	Code *int64 `json:"code"`

	// 输入参数错误
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this bad request error
func (m *BadRequestError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BadRequestError) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *BadRequestError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bad request error based on context it is used
func (m *BadRequestError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BadRequestError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequestError) UnmarshalBinary(b []byte) error {
	var res BadRequestError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
