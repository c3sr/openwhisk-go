package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/models"
)

// UpdateRuleReader is a Reader for the UpdateRule structure.
type UpdateRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 413:
		result := NewUpdateRuleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRuleOK creates a UpdateRuleOK with default headers values
func NewUpdateRuleOK() *UpdateRuleOK {
	return &UpdateRuleOK{}
}

/*UpdateRuleOK handles this case with default header values.

Updated Item
*/
type UpdateRuleOK struct {
	Payload *models.ItemID
}

func (o *UpdateRuleOK) Error() string {
	return fmt.Sprintf("[PUT /namespaces/{namespace}/rules/{ruleName}][%d] updateRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ItemID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleBadRequest creates a UpdateRuleBadRequest with default headers values
func NewUpdateRuleBadRequest() *UpdateRuleBadRequest {
	return &UpdateRuleBadRequest{}
}

/*UpdateRuleBadRequest handles this case with default header values.

Bad request
*/
type UpdateRuleBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *UpdateRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /namespaces/{namespace}/rules/{ruleName}][%d] updateRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleUnauthorized creates a UpdateRuleUnauthorized with default headers values
func NewUpdateRuleUnauthorized() *UpdateRuleUnauthorized {
	return &UpdateRuleUnauthorized{}
}

/*UpdateRuleUnauthorized handles this case with default header values.

Unauthorized request
*/
type UpdateRuleUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *UpdateRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /namespaces/{namespace}/rules/{ruleName}][%d] updateRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleConflict creates a UpdateRuleConflict with default headers values
func NewUpdateRuleConflict() *UpdateRuleConflict {
	return &UpdateRuleConflict{}
}

/*UpdateRuleConflict handles this case with default header values.

Conflicting item already exists
*/
type UpdateRuleConflict struct {
	Payload *models.ErrorMessage
}

func (o *UpdateRuleConflict) Error() string {
	return fmt.Sprintf("[PUT /namespaces/{namespace}/rules/{ruleName}][%d] updateRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleRequestEntityTooLarge creates a UpdateRuleRequestEntityTooLarge with default headers values
func NewUpdateRuleRequestEntityTooLarge() *UpdateRuleRequestEntityTooLarge {
	return &UpdateRuleRequestEntityTooLarge{}
}

/*UpdateRuleRequestEntityTooLarge handles this case with default header values.

Request entity too large
*/
type UpdateRuleRequestEntityTooLarge struct {
	Payload *models.ErrorMessage
}

func (o *UpdateRuleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /namespaces/{namespace}/rules/{ruleName}][%d] updateRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateRuleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleInternalServerError creates a UpdateRuleInternalServerError with default headers values
func NewUpdateRuleInternalServerError() *UpdateRuleInternalServerError {
	return &UpdateRuleInternalServerError{}
}

/*UpdateRuleInternalServerError handles this case with default header values.

Server error
*/
type UpdateRuleInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *UpdateRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /namespaces/{namespace}/rules/{ruleName}][%d] updateRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}