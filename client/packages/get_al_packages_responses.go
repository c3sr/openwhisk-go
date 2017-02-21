package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/c3sr/openwhisk-go/models"
)

// GetAlPackagesReader is a Reader for the GetAlPackages structure.
type GetAlPackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlPackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlPackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAlPackagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAlPackagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlPackagesOK creates a GetAlPackagesOK with default headers values
func NewGetAlPackagesOK() *GetAlPackagesOK {
	return &GetAlPackagesOK{}
}

/*GetAlPackagesOK handles this case with default header values.

Packages response
*/
type GetAlPackagesOK struct {
	Payload []*models.EntityBrief
}

func (o *GetAlPackagesOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/packages][%d] getAlPackagesOK  %+v", 200, o.Payload)
}

func (o *GetAlPackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlPackagesUnauthorized creates a GetAlPackagesUnauthorized with default headers values
func NewGetAlPackagesUnauthorized() *GetAlPackagesUnauthorized {
	return &GetAlPackagesUnauthorized{}
}

/*GetAlPackagesUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetAlPackagesUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetAlPackagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/packages][%d] getAlPackagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlPackagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlPackagesInternalServerError creates a GetAlPackagesInternalServerError with default headers values
func NewGetAlPackagesInternalServerError() *GetAlPackagesInternalServerError {
	return &GetAlPackagesInternalServerError{}
}

/*GetAlPackagesInternalServerError handles this case with default header values.

Server error
*/
type GetAlPackagesInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetAlPackagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/packages][%d] getAlPackagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlPackagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
