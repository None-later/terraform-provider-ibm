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

// ShareProfile share profile
// swagger:model ShareProfile
type ShareProfile struct {

	// The CRN for this profile
	Crn string `json:"crn,omitempty"`

	// family
	Family string `json:"family,omitempty"`

	// The platform generation of the share profile
	// Enum: [gt gc]
	Generation *string `json:"generation,omitempty"`

	// The URL for this share profile
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The user-defined name for share profile.
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this share profile
func (m *ShareProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var shareProfileTypeGenerationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["gt","gc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shareProfileTypeGenerationPropEnum = append(shareProfileTypeGenerationPropEnum, v)
	}
}

const (

	// ShareProfileGenerationGt captures enum value "gt"
	ShareProfileGenerationGt string = "gt"

	// ShareProfileGenerationGc captures enum value "gc"
	ShareProfileGenerationGc string = "gc"
)

// prop value enum
func (m *ShareProfile) validateGenerationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, shareProfileTypeGenerationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ShareProfile) validateGeneration(formats strfmt.Registry) error {

	if swag.IsZero(m.Generation) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenerationEnum("generation", "body", *m.Generation); err != nil {
		return err
	}

	return nil
}

func (m *ShareProfile) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *ShareProfile) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShareProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShareProfile) UnmarshalBinary(b []byte) error {
	var res ShareProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
