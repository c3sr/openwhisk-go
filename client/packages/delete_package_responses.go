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

// DeletePackageReader is a Reader for the DeletePackage structure.
type DeletePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeletePackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeletePackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeletePackageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePackageOK creates a DeletePackageOK with default headers values
func NewDeletePackageOK() *DeletePackageOK {
	return &DeletePackageOK{}
}

/*DeletePackageOK handles this case with default header values.

Deleted Item
*/
type DeletePackageOK struct {
}

func (o *DeletePackageOK) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/packages/{packageName}][%d] deletePackageOK ", 200)
}

func (o *DeletePackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePackageUnauthorized creates a DeletePackageUnauthorized with default headers values
func NewDeletePackageUnauthorized() *DeletePackageUnauthorized {
	return &DeletePackageUnauthorized{}
}

/*DeletePackageUnauthorized handles this case with default header values.

Unauthorized request
*/
type DeletePackageUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *DeletePackageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/packages/{packageName}][%d] deletePackageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackageNotFound creates a DeletePackageNotFound with default headers values
func NewDeletePackageNotFound() *DeletePackageNotFound {
	return &DeletePackageNotFound{}
}

/*DeletePackageNotFound handles this case with default header values.

Item not found
*/
type DeletePackageNotFound struct {
	Payload *models.ErrorMessage
}

func (o *DeletePackageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/packages/{packageName}][%d] deletePackageNotFound  %+v", 404, o.Payload)
}

func (o *DeletePackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackageInternalServerError creates a DeletePackageInternalServerError with default headers values
func NewDeletePackageInternalServerError() *DeletePackageInternalServerError {
	return &DeletePackageInternalServerError{}
}

/*DeletePackageInternalServerError handles this case with default header values.

Server error
*/
type DeletePackageInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *DeletePackageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /namespaces/{namespace}/packages/{packageName}][%d] deletePackageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePackageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
