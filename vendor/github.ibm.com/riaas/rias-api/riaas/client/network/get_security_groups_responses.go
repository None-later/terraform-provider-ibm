// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetSecurityGroupsReader is a Reader for the GetSecurityGroups structure.
type GetSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSecurityGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetSecurityGroupsOK creates a GetSecurityGroupsOK with default headers values
func NewGetSecurityGroupsOK() *GetSecurityGroupsOK {
	return &GetSecurityGroupsOK{}
}

/*GetSecurityGroupsOK handles this case with default header values.

dummy
*/
type GetSecurityGroupsOK struct {
	Payload *models.GetSecurityGroupsOKBody
}

func (o *GetSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /security_groups][%d] getSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSecurityGroupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsDefault creates a GetSecurityGroupsDefault with default headers values
func NewGetSecurityGroupsDefault(code int) *GetSecurityGroupsDefault {
	return &GetSecurityGroupsDefault{
		_statusCode: code,
	}
}

/*GetSecurityGroupsDefault handles this case with default header values.

unexpectederror
*/
type GetSecurityGroupsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get security groups default response
func (o *GetSecurityGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GetSecurityGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /security_groups][%d] GetSecurityGroups default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecurityGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}