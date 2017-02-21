package activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/models"
)

// GetActivationsReader is a Reader for the GetActivations structure.
type GetActivationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActivationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetActivationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetActivationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetActivationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActivationsOK creates a GetActivationsOK with default headers values
func NewGetActivationsOK() *GetActivationsOK {
	return &GetActivationsOK{}
}

/*GetActivationsOK handles this case with default header values.

Activations response
*/
type GetActivationsOK struct {
	Payload []*models.EntityBrief
}

func (o *GetActivationsOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/activations][%d] getActivationsOK  %+v", 200, o.Payload)
}

func (o *GetActivationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActivationsUnauthorized creates a GetActivationsUnauthorized with default headers values
func NewGetActivationsUnauthorized() *GetActivationsUnauthorized {
	return &GetActivationsUnauthorized{}
}

/*GetActivationsUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetActivationsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetActivationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/activations][%d] getActivationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetActivationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActivationsInternalServerError creates a GetActivationsInternalServerError with default headers values
func NewGetActivationsInternalServerError() *GetActivationsInternalServerError {
	return &GetActivationsInternalServerError{}
}

/*GetActivationsInternalServerError handles this case with default header values.

Server error
*/
type GetActivationsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetActivationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/activations][%d] getActivationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetActivationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}