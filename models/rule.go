package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Rule rule
// swagger:model Rule
type Rule struct {

	// Name of the action
	// Required: true
	// Min Length: 1
	Action *string `json:"action"`

	// Name of the item
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// Namespace of the item
	// Required: true
	// Min Length: 1
	Namespace *string `json:"namespace"`

	// Whether to publish the item or not
	// Required: true
	Publish *bool `json:"publish"`

	// Status of a rule
	// Required: true
	Status *string `json:"status"`

	// Name of the trigger
	// Required: true
	// Min Length: 1
	Trigger *string `json:"trigger"`

	// Semantic version of the item
	// Required: true
	// Min Length: 1
	Version *string `json:"version"`
}

// Validate validates this rule
func (m *Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublish(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
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

func (m *Rule) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	if err := validate.MinLength("action", "body", string(*m.Action), 1); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	if err := validate.MinLength("namespace", "body", string(*m.Namespace), 1); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validatePublish(formats strfmt.Registry) error {

	if err := validate.Required("publish", "body", m.Publish); err != nil {
		return err
	}

	return nil
}

var ruleTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","activating","deactivating"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleTypeStatusPropEnum = append(ruleTypeStatusPropEnum, v)
	}
}

const (
	// RuleStatusActive captures enum value "active"
	RuleStatusActive string = "active"
	// RuleStatusInactive captures enum value "inactive"
	RuleStatusInactive string = "inactive"
	// RuleStatusActivating captures enum value "activating"
	RuleStatusActivating string = "activating"
	// RuleStatusDeactivating captures enum value "deactivating"
	RuleStatusDeactivating string = "deactivating"
)

// prop value enum
func (m *Rule) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ruleTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Rule) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateTrigger(formats strfmt.Registry) error {

	if err := validate.Required("trigger", "body", m.Trigger); err != nil {
		return err
	}

	if err := validate.MinLength("trigger", "body", string(*m.Trigger), 1); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}
