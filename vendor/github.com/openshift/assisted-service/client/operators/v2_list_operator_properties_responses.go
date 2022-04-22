// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListOperatorPropertiesReader is a Reader for the V2ListOperatorProperties structure.
type V2ListOperatorPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListOperatorPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListOperatorPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListOperatorPropertiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListOperatorPropertiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2ListOperatorPropertiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListOperatorPropertiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListOperatorPropertiesOK creates a V2ListOperatorPropertiesOK with default headers values
func NewV2ListOperatorPropertiesOK() *V2ListOperatorPropertiesOK {
	return &V2ListOperatorPropertiesOK{}
}

/* V2ListOperatorPropertiesOK describes a response with status code 200, with default header values.

Success.
*/
type V2ListOperatorPropertiesOK struct {
	Payload models.OperatorProperties
}

func (o *V2ListOperatorPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators/{operator_name}][%d] v2ListOperatorPropertiesOK  %+v", 200, o.Payload)
}
func (o *V2ListOperatorPropertiesOK) GetPayload() models.OperatorProperties {
	return o.Payload
}

func (o *V2ListOperatorPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOperatorPropertiesUnauthorized creates a V2ListOperatorPropertiesUnauthorized with default headers values
func NewV2ListOperatorPropertiesUnauthorized() *V2ListOperatorPropertiesUnauthorized {
	return &V2ListOperatorPropertiesUnauthorized{}
}

/* V2ListOperatorPropertiesUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2ListOperatorPropertiesUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2ListOperatorPropertiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators/{operator_name}][%d] v2ListOperatorPropertiesUnauthorized  %+v", 401, o.Payload)
}
func (o *V2ListOperatorPropertiesUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListOperatorPropertiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOperatorPropertiesForbidden creates a V2ListOperatorPropertiesForbidden with default headers values
func NewV2ListOperatorPropertiesForbidden() *V2ListOperatorPropertiesForbidden {
	return &V2ListOperatorPropertiesForbidden{}
}

/* V2ListOperatorPropertiesForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2ListOperatorPropertiesForbidden struct {
	Payload *models.InfraError
}

func (o *V2ListOperatorPropertiesForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators/{operator_name}][%d] v2ListOperatorPropertiesForbidden  %+v", 403, o.Payload)
}
func (o *V2ListOperatorPropertiesForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListOperatorPropertiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOperatorPropertiesNotFound creates a V2ListOperatorPropertiesNotFound with default headers values
func NewV2ListOperatorPropertiesNotFound() *V2ListOperatorPropertiesNotFound {
	return &V2ListOperatorPropertiesNotFound{}
}

/* V2ListOperatorPropertiesNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2ListOperatorPropertiesNotFound struct {
	Payload *models.Error
}

func (o *V2ListOperatorPropertiesNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators/{operator_name}][%d] v2ListOperatorPropertiesNotFound  %+v", 404, o.Payload)
}
func (o *V2ListOperatorPropertiesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListOperatorPropertiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOperatorPropertiesInternalServerError creates a V2ListOperatorPropertiesInternalServerError with default headers values
func NewV2ListOperatorPropertiesInternalServerError() *V2ListOperatorPropertiesInternalServerError {
	return &V2ListOperatorPropertiesInternalServerError{}
}

/* V2ListOperatorPropertiesInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ListOperatorPropertiesInternalServerError struct {
	Payload *models.Error
}

func (o *V2ListOperatorPropertiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/supported-operators/{operator_name}][%d] v2ListOperatorPropertiesInternalServerError  %+v", 500, o.Payload)
}
func (o *V2ListOperatorPropertiesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListOperatorPropertiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}