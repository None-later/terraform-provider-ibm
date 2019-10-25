// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchSharesIDReader is a Reader for the PatchSharesID structure.
type PatchSharesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSharesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSharesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchSharesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSharesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchSharesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSharesIDOK creates a PatchSharesIDOK with default headers values
func NewPatchSharesIDOK() *PatchSharesIDOK {
	return &PatchSharesIDOK{}
}

/*PatchSharesIDOK handles this case with default header values.

dummy
*/
type PatchSharesIDOK struct {
	Payload *models.Share
}

func (o *PatchSharesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /shares/{id}][%d] patchSharesIdOK  %+v", 200, o.Payload)
}

func (o *PatchSharesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Share)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSharesIDBadRequest creates a PatchSharesIDBadRequest with default headers values
func NewPatchSharesIDBadRequest() *PatchSharesIDBadRequest {
	return &PatchSharesIDBadRequest{}
}

/*PatchSharesIDBadRequest handles this case with default header values.

error
*/
type PatchSharesIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchSharesIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /shares/{id}][%d] patchSharesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSharesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSharesIDNotFound creates a PatchSharesIDNotFound with default headers values
func NewPatchSharesIDNotFound() *PatchSharesIDNotFound {
	return &PatchSharesIDNotFound{}
}

/*PatchSharesIDNotFound handles this case with default header values.

error
*/
type PatchSharesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchSharesIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /shares/{id}][%d] patchSharesIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchSharesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSharesIDInternalServerError creates a PatchSharesIDInternalServerError with default headers values
func NewPatchSharesIDInternalServerError() *PatchSharesIDInternalServerError {
	return &PatchSharesIDInternalServerError{}
}

/*PatchSharesIDInternalServerError handles this case with default header values.

error
*/
type PatchSharesIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PatchSharesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /shares/{id}][%d] patchSharesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSharesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchSharesIDBody SharePatch
swagger:model PatchSharesIDBody
*/
type PatchSharesIDBody struct {

	// The user-defined name for this volume
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this patch shares ID body
func (o *PatchSharesIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchSharesIDBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchSharesIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchSharesIDBody) UnmarshalBinary(b []byte) error {
	var res PatchSharesIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
