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
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConsoleServerPort writable console server port
// swagger:model WritableConsoleServerPort
type WritableConsoleServerPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	//         Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Connection status
	// Enum: [false true]
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Type string `json:"type,omitempty"`
}

// Validate validates this writable console server port
func (m *WritableConsoleServerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *WritableConsoleServerPort) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

var writableConsoleServerPortTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsoleServerPortTypeConnectionStatusPropEnum = append(writableConsoleServerPortTypeConnectionStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableConsoleServerPort) validateConnectionStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writableConsoleServerPortTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsoleServerPort) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsoleServerPort) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

var writableConsoleServerPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsoleServerPortTypeTypePropEnum = append(writableConsoleServerPortTypeTypePropEnum, v)
	}
}

const (

	// WritableConsoleServerPortTypeDe9 captures enum value "de-9"
	WritableConsoleServerPortTypeDe9 string = "de-9"

	// WritableConsoleServerPortTypeDb25 captures enum value "db-25"
	WritableConsoleServerPortTypeDb25 string = "db-25"

	// WritableConsoleServerPortTypeRj11 captures enum value "rj-11"
	WritableConsoleServerPortTypeRj11 string = "rj-11"

	// WritableConsoleServerPortTypeRj12 captures enum value "rj-12"
	WritableConsoleServerPortTypeRj12 string = "rj-12"

	// WritableConsoleServerPortTypeRj45 captures enum value "rj-45"
	WritableConsoleServerPortTypeRj45 string = "rj-45"

	// WritableConsoleServerPortTypeUsba captures enum value "usb-a"
	WritableConsoleServerPortTypeUsba string = "usb-a"

	// WritableConsoleServerPortTypeUsbb captures enum value "usb-b"
	WritableConsoleServerPortTypeUsbb string = "usb-b"

	// WritableConsoleServerPortTypeUsbc captures enum value "usb-c"
	WritableConsoleServerPortTypeUsbc string = "usb-c"

	// WritableConsoleServerPortTypeUsbMinia captures enum value "usb-mini-a"
	WritableConsoleServerPortTypeUsbMinia string = "usb-mini-a"

	// WritableConsoleServerPortTypeUsbMinib captures enum value "usb-mini-b"
	WritableConsoleServerPortTypeUsbMinib string = "usb-mini-b"

	// WritableConsoleServerPortTypeUsbMicroa captures enum value "usb-micro-a"
	WritableConsoleServerPortTypeUsbMicroa string = "usb-micro-a"

	// WritableConsoleServerPortTypeUsbMicrob captures enum value "usb-micro-b"
	WritableConsoleServerPortTypeUsbMicrob string = "usb-micro-b"

	// WritableConsoleServerPortTypeOther captures enum value "other"
	WritableConsoleServerPortTypeOther string = "other"
)

// prop value enum
func (m *WritableConsoleServerPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableConsoleServerPortTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsoleServerPort) validateType(formats strfmt.Registry) error {

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
func (m *WritableConsoleServerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConsoleServerPort) UnmarshalBinary(b []byte) error {
	var res WritableConsoleServerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
