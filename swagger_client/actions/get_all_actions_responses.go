package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/swagger_models"
)

// GetAllActionsReader is a Reader for the GetAllActions structure.
type GetAllActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAllActionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllActionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllActionsOK creates a GetAllActionsOK with default headers values
func NewGetAllActionsOK() *GetAllActionsOK {
	return &GetAllActionsOK{}
}

/*GetAllActionsOK handles this case with default header values.

Actions response
*/
type GetAllActionsOK struct {
	Payload []*swagger_models.EntityBrief
}

func (o *GetAllActionsOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/actions][%d] getAllActionsOK  %+v", 200, o.Payload)
}

func (o *GetAllActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllActionsUnauthorized creates a GetAllActionsUnauthorized with default headers values
func NewGetAllActionsUnauthorized() *GetAllActionsUnauthorized {
	return &GetAllActionsUnauthorized{}
}

/*GetAllActionsUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetAllActionsUnauthorized struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetAllActionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/actions][%d] getAllActionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllActionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllActionsInternalServerError creates a GetAllActionsInternalServerError with default headers values
func NewGetAllActionsInternalServerError() *GetAllActionsInternalServerError {
	return &GetAllActionsInternalServerError{}
}

/*GetAllActionsInternalServerError handles this case with default header values.

Server error
*/
type GetAllActionsInternalServerError struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetAllActionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/actions][%d] getAllActionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllActionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
