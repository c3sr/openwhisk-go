package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/swagger_models"
)

// DeleteTriggerReader is a Reader for the DeleteTrigger structure.
type DeleteTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTriggerOK creates a DeleteTriggerOK with default headers values
func NewDeleteTriggerOK() *DeleteTriggerOK {
	return &DeleteTriggerOK{}
}

/*DeleteTriggerOK handles this case with default header values.

Deleted Item
*/
type DeleteTriggerOK struct {
}

func (o *DeleteTriggerOK) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/triggers/{triggerName}][%d] deleteTriggerOK ", 200)
}

func (o *DeleteTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTriggerUnauthorized creates a DeleteTriggerUnauthorized with default headers values
func NewDeleteTriggerUnauthorized() *DeleteTriggerUnauthorized {
	return &DeleteTriggerUnauthorized{}
}

/*DeleteTriggerUnauthorized handles this case with default header values.

Unauthorized request
*/
type DeleteTriggerUnauthorized struct {
	Payload *swagger_models.ErrorMessage
}

func (o *DeleteTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/triggers/{triggerName}][%d] deleteTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTriggerNotFound creates a DeleteTriggerNotFound with default headers values
func NewDeleteTriggerNotFound() *DeleteTriggerNotFound {
	return &DeleteTriggerNotFound{}
}

/*DeleteTriggerNotFound handles this case with default header values.

Item not found
*/
type DeleteTriggerNotFound struct {
	Payload *swagger_models.ErrorMessage
}

func (o *DeleteTriggerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/triggers/{triggerName}][%d] deleteTriggerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTriggerInternalServerError creates a DeleteTriggerInternalServerError with default headers values
func NewDeleteTriggerInternalServerError() *DeleteTriggerInternalServerError {
	return &DeleteTriggerInternalServerError{}
}

/*DeleteTriggerInternalServerError handles this case with default header values.

Server error
*/
type DeleteTriggerInternalServerError struct {
	Payload *swagger_models.ErrorMessage
}

func (o *DeleteTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/triggers/{triggerName}][%d] deleteTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
