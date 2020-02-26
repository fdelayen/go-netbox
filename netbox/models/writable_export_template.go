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

// WritableExportTemplate writable export template
// swagger:model WritableExportTemplate
type WritableExportTemplate struct {

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// File extension
	//
	// Extension to append to the rendered filename
	// Max Length: 15
	FileExtension string `json:"file_extension,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// MIME type
	//
	// Defaults to <code>text/plain</code>
	// Max Length: 50
	MimeType string `json:"mime_type,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Template code
	//
	// The list of objects being exported is passed as a context variable named <code>queryset</code>.
	// Required: true
	// Min Length: 1
	TemplateCode *string `json:"template_code"`

	// Template language
	// Enum: [django jinja2]
	TemplateLanguage string `json:"template_language,omitempty"`
}

// Validate validates this writable export template
func (m *WritableExportTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileExtension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMimeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableExportTemplate) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *WritableExportTemplate) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableExportTemplate) validateFileExtension(formats strfmt.Registry) error {

	if swag.IsZero(m.FileExtension) { // not required
		return nil
	}

	if err := validate.MaxLength("file_extension", "body", string(m.FileExtension), 15); err != nil {
		return err
	}

	return nil
}

func (m *WritableExportTemplate) validateMimeType(formats strfmt.Registry) error {

	if swag.IsZero(m.MimeType) { // not required
		return nil
	}

	if err := validate.MaxLength("mime_type", "body", string(m.MimeType), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableExportTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableExportTemplate) validateTemplateCode(formats strfmt.Registry) error {

	if err := validate.Required("template_code", "body", m.TemplateCode); err != nil {
		return err
	}

	if err := validate.MinLength("template_code", "body", string(*m.TemplateCode), 1); err != nil {
		return err
	}

	return nil
}

var writableExportTemplateTypeTemplateLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["django","jinja2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableExportTemplateTypeTemplateLanguagePropEnum = append(writableExportTemplateTypeTemplateLanguagePropEnum, v)
	}
}

const (

	// WritableExportTemplateTemplateLanguageDjango captures enum value "django"
	WritableExportTemplateTemplateLanguageDjango string = "django"

	// WritableExportTemplateTemplateLanguageJinja2 captures enum value "jinja2"
	WritableExportTemplateTemplateLanguageJinja2 string = "jinja2"
)

// prop value enum
func (m *WritableExportTemplate) validateTemplateLanguageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableExportTemplateTypeTemplateLanguagePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableExportTemplate) validateTemplateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.TemplateLanguage) { // not required
		return nil
	}

	// value enum
	if err := m.validateTemplateLanguageEnum("template_language", "body", m.TemplateLanguage); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableExportTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableExportTemplate) UnmarshalBinary(b []byte) error {
	var res WritableExportTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
