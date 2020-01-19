// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Password password
// swagger:model Password
type Password strfmt.Password

// Validate validates this password
func (m Password) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 100); err != nil {
		return err
	}

	if err := validate.FormatOf("", "body", "password", strfmt.Password(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}