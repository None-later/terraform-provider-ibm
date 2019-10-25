// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListenerTemplateWithPoolByName listener template with pool by name
// swagger:model ListenerTemplateWithPoolByName
type ListenerTemplateWithPoolByName struct {

	// certificate instance
	CertificateInstance *ListenerTemplateWithPoolByNameCertificateInstance `json:"certificate_instance,omitempty"`

	// The connection limit of the listener.
	// Maximum: 15000
	// Minimum: 1
	ConnectionLimit int64 `json:"connection_limit,omitempty"`

	// default pool
	DefaultPool *ListenerTemplateWithPoolByNameDefaultPool `json:"default_pool,omitempty"`

	// The listener port number.
	// Maximum: 65535
	// Minimum: 1
	Port int64 `json:"port,omitempty"`

	// The listener protocol.
	// Enum: [http https tcp]
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this listener template with pool by name
func (m *ListenerTemplateWithPoolByName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListenerTemplateWithPoolByName) validateCertificateInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateInstance) { // not required
		return nil
	}

	if m.CertificateInstance != nil {
		if err := m.CertificateInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate_instance")
			}
			return err
		}
	}

	return nil
}

func (m *ListenerTemplateWithPoolByName) validateConnectionLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("connection_limit", "body", int64(m.ConnectionLimit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("connection_limit", "body", int64(m.ConnectionLimit), 15000, false); err != nil {
		return err
	}

	return nil
}

func (m *ListenerTemplateWithPoolByName) validateDefaultPool(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultPool) { // not required
		return nil
	}

	if m.DefaultPool != nil {
		if err := m.DefaultPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_pool")
			}
			return err
		}
	}

	return nil
}

func (m *ListenerTemplateWithPoolByName) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var listenerTemplateWithPoolByNameTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","https","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerTemplateWithPoolByNameTypeProtocolPropEnum = append(listenerTemplateWithPoolByNameTypeProtocolPropEnum, v)
	}
}

const (

	// ListenerTemplateWithPoolByNameProtocolHTTP captures enum value "http"
	ListenerTemplateWithPoolByNameProtocolHTTP string = "http"

	// ListenerTemplateWithPoolByNameProtocolHTTPS captures enum value "https"
	ListenerTemplateWithPoolByNameProtocolHTTPS string = "https"

	// ListenerTemplateWithPoolByNameProtocolTCP captures enum value "tcp"
	ListenerTemplateWithPoolByNameProtocolTCP string = "tcp"
)

// prop value enum
func (m *ListenerTemplateWithPoolByName) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerTemplateWithPoolByNameTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ListenerTemplateWithPoolByName) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerTemplateWithPoolByName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerTemplateWithPoolByName) UnmarshalBinary(b []byte) error {
	var res ListenerTemplateWithPoolByName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ListenerTemplateWithPoolByNameCertificateInstance listener template with pool by name certificate instance
// swagger:model ListenerTemplateWithPoolByNameCertificateInstance
type ListenerTemplateWithPoolByNameCertificateInstance struct {

	// The ceritificate instance's CRN
	Crn string `json:"crn,omitempty"`
}

// Validate validates this listener template with pool by name certificate instance
func (m *ListenerTemplateWithPoolByNameCertificateInstance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListenerTemplateWithPoolByNameCertificateInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerTemplateWithPoolByNameCertificateInstance) UnmarshalBinary(b []byte) error {
	var res ListenerTemplateWithPoolByNameCertificateInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ListenerTemplateWithPoolByNameDefaultPool listener template with pool by name default pool
// swagger:model ListenerTemplateWithPoolByNameDefaultPool
type ListenerTemplateWithPoolByNameDefaultPool struct {

	// The pool's user-defined name.
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this listener template with pool by name default pool
func (m *ListenerTemplateWithPoolByNameDefaultPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListenerTemplateWithPoolByNameDefaultPool) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("default_pool"+"."+"name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerTemplateWithPoolByNameDefaultPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerTemplateWithPoolByNameDefaultPool) UnmarshalBinary(b []byte) error {
	var res ListenerTemplateWithPoolByNameDefaultPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
