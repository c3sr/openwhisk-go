package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/models"
)

// GetAllTriggersReader is a Reader for the GetAllTriggers structure.
type GetAllTriggersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTriggersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTriggersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAllTriggersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllTriggersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllTriggersOK creates a GetAllTriggersOK with default headers values
func NewGetAllTriggersOK() *GetAllTriggersOK {
	return &GetAllTriggersOK{}
}

/*GetAllTriggersOK handles this case with default header values.

Triggers response
*/
type GetAllTriggersOK struct {
	Payload []*models.EntityBrief
}

func (o *GetAllTriggersOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers][%d] getAllTriggersOK  %+v", 200, o.Payload)
}

func (o *GetAllTriggersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTriggersUnauthorized creates a GetAllTriggersUnauthorized with default headers values
func NewGetAllTriggersUnauthorized() *GetAllTriggersUnauthorized {
	return &GetAllTriggersUnauthorized{}
}

/*GetAllTriggersUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetAllTriggersUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetAllTriggersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers][%d] getAllTriggersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllTriggersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTriggersInternalServerError creates a GetAllTriggersInternalServerError with default headers values
func NewGetAllTriggersInternalServerError() *GetAllTriggersInternalServerError {
	return &GetAllTriggersInternalServerError{}
}

/*GetAllTriggersInternalServerError handles this case with default header values.

Server error
*/
type GetAllTriggersInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetAllTriggersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/triggers][%d] getAllTriggersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllTriggersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
