package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ActionLimits Limits on a specific action
// swagger:model ActionLimits
type ActionLimits struct {

	// memory
	Memory *int32 `json:"memory,omitempty"`

	// timeout
	Timeout *int32 `json:"timeout,omitempty"`
}

// Validate validates this action limits
func (m *ActionLimits) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
