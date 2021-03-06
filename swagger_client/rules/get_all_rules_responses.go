package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/swagger_models"
)

// GetAllRulesReader is a Reader for the GetAllRules structure.
type GetAllRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAllRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllRulesOK creates a GetAllRulesOK with default headers values
func NewGetAllRulesOK() *GetAllRulesOK {
	return &GetAllRulesOK{}
}

/*GetAllRulesOK handles this case with default header values.

Rules response
*/
type GetAllRulesOK struct {
	Payload []*swagger_models.EntityBrief
}

func (o *GetAllRulesOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/rules][%d] getAllRulesOK  %+v", 200, o.Payload)
}

func (o *GetAllRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllRulesUnauthorized creates a GetAllRulesUnauthorized with default headers values
func NewGetAllRulesUnauthorized() *GetAllRulesUnauthorized {
	return &GetAllRulesUnauthorized{}
}

/*GetAllRulesUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetAllRulesUnauthorized struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetAllRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/rules][%d] getAllRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllRulesInternalServerError creates a GetAllRulesInternalServerError with default headers values
func NewGetAllRulesInternalServerError() *GetAllRulesInternalServerError {
	return &GetAllRulesInternalServerError{}
}

/*GetAllRulesInternalServerError handles this case with default header values.

Server error
*/
type GetAllRulesInternalServerError struct {
	Payload *swagger_models.ErrorMessage
}

func (o *GetAllRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/rules][%d] getAllRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
