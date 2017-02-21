package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PackagePut A restricted Package view that elides properties that are auto-assigned or derived from the URI (i.e., the namespace and name).
// swagger:model PackagePut
type PackagePut struct {

	// annotations on the item
	Annotations []*KeyValue `json:"annotations"`

	// binding
	Binding *PackageBinding `json:"binding,omitempty"`

	// parameter for the package
	Parameters []*KeyValue `json:"parameters"`

	// Whether to publish the item or not
	Publish bool `json:"publish,omitempty"`

	// Semantic version of the item
	// Min Length: 1
	Version string `json:"version,omitempty"`
}

// Validate validates this package put
func (m *PackagePut) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBinding(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackagePut) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	for i := 0; i < len(m.Annotations); i++ {

		if swag.IsZero(m.Annotations[i]) { // not required
			continue
		}

		if m.Annotations[i] != nil {

			if err := m.Annotations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PackagePut) validateBinding(formats strfmt.Registry) error {

	if swag.IsZero(m.Binding) { // not required
		return nil
	}

	if m.Binding != nil {

		if err := m.Binding.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PackagePut) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {

		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {

			if err := m.Parameters[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PackagePut) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", string(m.Version), 1); err != nil {
		return err
	}

	return nil
}
