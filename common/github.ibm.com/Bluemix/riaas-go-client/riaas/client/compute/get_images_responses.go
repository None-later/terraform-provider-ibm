// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetImagesReader is a Reader for the GetImages structure.
type GetImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImagesOK creates a GetImagesOK with default headers values
func NewGetImagesOK() *GetImagesOK {
	return &GetImagesOK{}
}

/*GetImagesOK handles this case with default header values.

dummy
*/
type GetImagesOK struct {
	Payload *GetImagesOKBody
}

func (o *GetImagesOK) Error() string {
	return fmt.Sprintf("[GET /images][%d] getImagesOK  %+v", 200, o.Payload)
}

func (o *GetImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetImagesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesInternalServerError creates a GetImagesInternalServerError with default headers values
func NewGetImagesInternalServerError() *GetImagesInternalServerError {
	return &GetImagesInternalServerError{}
}

/*GetImagesInternalServerError handles this case with default header values.

error
*/
type GetImagesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images][%d] getImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetImagesOKBody ImageCollection
swagger:model GetImagesOKBody
*/
type GetImagesOKBody struct {

	// first
	First *GetImagesOKBodyFirst `json:"first,omitempty"`

	// Collection of images
	// Required: true
	Images []*models.Image `json:"images"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// next
	Next *GetImagesOKBodyNext `json:"next,omitempty"`

	// The total number of resources across all pages
	// Minimum: 0
	TotalCount *int64 `json:"total_count,omitempty"`
}

// Validate validates this get images o k body
func (o *GetImagesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetImagesOKBody) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(o.First) { // not required
		return nil
	}

	if o.First != nil {
		if err := o.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getImagesOK" + "." + "first")
			}
			return err
		}
	}

	return nil
}

func (o *GetImagesOKBody) validateImages(formats strfmt.Registry) error {

	if err := validate.Required("getImagesOK"+"."+"images", "body", o.Images); err != nil {
		return err
	}

	for i := 0; i < len(o.Images); i++ {
		if swag.IsZero(o.Images[i]) { // not required
			continue
		}

		if o.Images[i] != nil {
			if err := o.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getImagesOK" + "." + "images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetImagesOKBody) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(o.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("getImagesOK"+"."+"limit", "body", int64(o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("getImagesOK"+"."+"limit", "body", int64(o.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (o *GetImagesOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if o.Next != nil {
		if err := o.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getImagesOK" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (o *GetImagesOKBody) validateTotalCount(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("getImagesOK"+"."+"total_count", "body", int64(*o.TotalCount), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetImagesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetImagesOKBody) UnmarshalBinary(b []byte) error {
	var res GetImagesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetImagesOKBodyFirst A reference to the first page of resources
swagger:model GetImagesOKBodyFirst
*/
type GetImagesOKBodyFirst struct {

	// The URL for the first page of resources
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`
}

// Validate validates this get images o k body first
func (o *GetImagesOKBodyFirst) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetImagesOKBodyFirst) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("getImagesOK"+"."+"first"+"."+"href", "body", o.Href); err != nil {
		return err
	}

	if err := validate.Pattern("getImagesOK"+"."+"first"+"."+"href", "body", string(*o.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetImagesOKBodyFirst) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetImagesOKBodyFirst) UnmarshalBinary(b []byte) error {
	var res GetImagesOKBodyFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetImagesOKBodyNext A reference to the next page of resources; this reference is included for all pages except the last page
swagger:model GetImagesOKBodyNext
*/
type GetImagesOKBodyNext struct {

	// The URL for the next page of resources
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`
}

// Validate validates this get images o k body next
func (o *GetImagesOKBodyNext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetImagesOKBodyNext) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("getImagesOK"+"."+"next"+"."+"href", "body", o.Href); err != nil {
		return err
	}

	if err := validate.Pattern("getImagesOK"+"."+"next"+"."+"href", "body", string(*o.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetImagesOKBodyNext) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetImagesOKBodyNext) UnmarshalBinary(b []byte) error {
	var res GetImagesOKBodyNext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
