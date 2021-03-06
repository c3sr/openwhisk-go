package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ActivationInfo activation info
// swagger:model ActivationInfo
type ActivationInfo struct {

	// Activation id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// result
	Result *ActivationInfoResult `json:"result,omitempty"`

	// Standard error from activation
	Stderr string `json:"stderr,omitempty"`

	// Standard output from activation
	Stdout string `json:"stdout,omitempty"`
}

// Validate validates this activation info
func (m *ActivationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivationInfo) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *ActivationInfo) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {

		if err := m.Result.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// ActivationInfoResult Activation result
// swagger:model ActivationInfoResult
type ActivationInfoResult struct {

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this activation info result
func (m *ActivationInfoResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivationInfoResult) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}
