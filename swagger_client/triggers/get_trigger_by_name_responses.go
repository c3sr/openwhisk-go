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

// GetTriggerByNameReader is a Reader for the GetTriggerByName structure.
type GetTriggerByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTriggerByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTriggerByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTriggerByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTriggerByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTriggerByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTriggerByNameOK creates a GetTriggerByNameOK with default headers values
func NewGetTriggerByNameOK() *GetTriggerByNameOK {
	return &GetTriggerByNameOK{}
}

/*GetTriggerByNameOK handles this case with default header values.

Returned trigger
*/
type GetTriggerByNameOK struct {
	Payload *swagger_models.Trigger
}

func (o *GetTriggerByNameOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers/{triggerName}][%d] getTriggerByNameOK  %+v", 200, o.Payload)
}

func (o *GetTriggerByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerByNameUnauthorized creates a GetTriggerByNameUnauthorized with default headers values
func NewGetTriggerByNameUnauthorized() *GetTriggerByNameUnauthorized {
	return &GetTriggerByNameUnauthorized{}
}

/*GetTriggerByNameUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetTriggerByNameUnauthorized struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetTriggerByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers/{triggerName}][%d] getTriggerByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTriggerByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerByNameNotFound creates a GetTriggerByNameNotFound with default headers values
func NewGetTriggerByNameNotFound() *GetTriggerByNameNotFound {
	return &GetTriggerByNameNotFound{}
}

/*GetTriggerByNameNotFound handles this case with default header values.

Item not found
*/
type GetTriggerByNameNotFound struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetTriggerByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers/{triggerName}][%d] getTriggerByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetTriggerByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerByNameInternalServerError creates a GetTriggerByNameInternalServerError with default headers values
func NewGetTriggerByNameInternalServerError() *GetTriggerByNameInternalServerError {
	return &GetTriggerByNameInternalServerError{}
}

/*GetTriggerByNameInternalServerError handles this case with default header values.

Server error
*/
type GetTriggerByNameInternalServerError struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetTriggerByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers/{triggerName}][%d] getTriggerByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTriggerByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
