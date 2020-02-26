// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCable writable cable
// swagger:model WritableCable
type WritableCable struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	// Max Length: 100
	Label string `json:"label,omitempty"`

	// Length
	// Maximum: 32767
	// Minimum: 0
	Length *int64 `json:"length,omitempty"`

	// Length unit
	// Enum: [m cm ft in]
	LengthUnit string `json:"length_unit,omitempty"`

	// Status
	// Enum: [connected planned decommissioning]
	Status string `json:"status,omitempty"`

	// Termination a
	// Read Only: true
	Terminationa map[string]string `json:"termination_a,omitempty"`

	// Termination a id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationaID *int64 `json:"termination_a_id"`

	// Termination a type
	// Required: true
	TerminationaType *string `json:"termination_a_type"`

	// Termination b
	// Read Only: true
	Terminationb map[string]string `json:"termination_b,omitempty"`

	// Termination b id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationbID *int64 `json:"termination_b_id"`

	// Termination b type
	// Required: true
	TerminationbType *string `json:"termination_b_type"`

	// Type
	// Enum: [cat3 cat5 cat5e cat6 cat6a cat7 dac-active dac-passive coaxial mmf mmf-om1 mmf-om2 mmf-om3 mmf-om4 smf smf-os1 smf-os2 aoc power]
	Type string `json:"type,omitempty"`
}

// Validate validates this writable cable
func (m *WritableCable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCable) validateColor(formats strfmt.Registry) error {

	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", string(m.Color), 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", string(m.Color), `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateLength(formats strfmt.Registry) error {

	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if err := validate.MinimumInt("length", "body", int64(*m.Length), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "body", int64(*m.Length), 32767, false); err != nil {
		return err
	}

	return nil
}

var writableCableTypeLengthUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["m","cm","ft","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeLengthUnitPropEnum = append(writableCableTypeLengthUnitPropEnum, v)
	}
}

const (

	// WritableCableLengthUnitM captures enum value "m"
	WritableCableLengthUnitM string = "m"

	// WritableCableLengthUnitCm captures enum value "cm"
	WritableCableLengthUnitCm string = "cm"

	// WritableCableLengthUnitFt captures enum value "ft"
	WritableCableLengthUnitFt string = "ft"

	// WritableCableLengthUnitIn captures enum value "in"
	WritableCableLengthUnitIn string = "in"
)

// prop value enum
func (m *WritableCable) validateLengthUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableCableTypeLengthUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateLengthUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.LengthUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateLengthUnitEnum("length_unit", "body", m.LengthUnit); err != nil {
		return err
	}

	return nil
}

var writableCableTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeStatusPropEnum = append(writableCableTypeStatusPropEnum, v)
	}
}

const (

	// WritableCableStatusConnected captures enum value "connected"
	WritableCableStatusConnected string = "connected"

	// WritableCableStatusPlanned captures enum value "planned"
	WritableCableStatusPlanned string = "planned"

	// WritableCableStatusDecommissioning captures enum value "decommissioning"
	WritableCableStatusDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WritableCable) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableCableTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationaID(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_id", "body", m.TerminationaID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_a_id", "body", int64(*m.TerminationaID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_a_id", "body", int64(*m.TerminationaID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationaType(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_type", "body", m.TerminationaType); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationbID(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_id", "body", m.TerminationbID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_b_id", "body", int64(*m.TerminationbID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_b_id", "body", int64(*m.TerminationbID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCable) validateTerminationbType(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_type", "body", m.TerminationbType); err != nil {
		return err
	}

	return nil
}

var writableCableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cat3","cat5","cat5e","cat6","cat6a","cat7","dac-active","dac-passive","coaxial","mmf","mmf-om1","mmf-om2","mmf-om3","mmf-om4","smf","smf-os1","smf-os2","aoc","power"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCableTypeTypePropEnum = append(writableCableTypeTypePropEnum, v)
	}
}

const (

	// WritableCableTypeCat3 captures enum value "cat3"
	WritableCableTypeCat3 string = "cat3"

	// WritableCableTypeCat5 captures enum value "cat5"
	WritableCableTypeCat5 string = "cat5"

	// WritableCableTypeCat5e captures enum value "cat5e"
	WritableCableTypeCat5e string = "cat5e"

	// WritableCableTypeCat6 captures enum value "cat6"
	WritableCableTypeCat6 string = "cat6"

	// WritableCableTypeCat6a captures enum value "cat6a"
	WritableCableTypeCat6a string = "cat6a"

	// WritableCableTypeCat7 captures enum value "cat7"
	WritableCableTypeCat7 string = "cat7"

	// WritableCableTypeDacActive captures enum value "dac-active"
	WritableCableTypeDacActive string = "dac-active"

	// WritableCableTypeDacPassive captures enum value "dac-passive"
	WritableCableTypeDacPassive string = "dac-passive"

	// WritableCableTypeCoaxial captures enum value "coaxial"
	WritableCableTypeCoaxial string = "coaxial"

	// WritableCableTypeMmf captures enum value "mmf"
	WritableCableTypeMmf string = "mmf"

	// WritableCableTypeMmfOm1 captures enum value "mmf-om1"
	WritableCableTypeMmfOm1 string = "mmf-om1"

	// WritableCableTypeMmfOm2 captures enum value "mmf-om2"
	WritableCableTypeMmfOm2 string = "mmf-om2"

	// WritableCableTypeMmfOm3 captures enum value "mmf-om3"
	WritableCableTypeMmfOm3 string = "mmf-om3"

	// WritableCableTypeMmfOm4 captures enum value "mmf-om4"
	WritableCableTypeMmfOm4 string = "mmf-om4"

	// WritableCableTypeSmf captures enum value "smf"
	WritableCableTypeSmf string = "smf"

	// WritableCableTypeSmfOs1 captures enum value "smf-os1"
	WritableCableTypeSmfOs1 string = "smf-os1"

	// WritableCableTypeSmfOs2 captures enum value "smf-os2"
	WritableCableTypeSmfOs2 string = "smf-os2"

	// WritableCableTypeAoc captures enum value "aoc"
	WritableCableTypeAoc string = "aoc"

	// WritableCableTypePower captures enum value "power"
	WritableCableTypePower string = "power"
)

// prop value enum
func (m *WritableCable) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableCableTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCable) UnmarshalBinary(b []byte) error {
	var res WritableCable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
