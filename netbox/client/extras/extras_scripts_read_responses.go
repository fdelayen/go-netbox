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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasScriptsReadReader is a Reader for the ExtrasScriptsRead structure.
type ExtrasScriptsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasScriptsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasScriptsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasScriptsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasScriptsReadOK creates a ExtrasScriptsReadOK with default headers values
func NewExtrasScriptsReadOK() *ExtrasScriptsReadOK {
	return &ExtrasScriptsReadOK{}
}

/*
ExtrasScriptsReadOK describes a response with status code 200, with default header values.

ExtrasScriptsReadOK extras scripts read o k
*/
type ExtrasScriptsReadOK struct {
}

// IsSuccess returns true when this extras scripts read o k response has a 2xx status code
func (o *ExtrasScriptsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras scripts read o k response has a 3xx status code
func (o *ExtrasScriptsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras scripts read o k response has a 4xx status code
func (o *ExtrasScriptsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras scripts read o k response has a 5xx status code
func (o *ExtrasScriptsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras scripts read o k response a status code equal to that given
func (o *ExtrasScriptsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasScriptsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/{id}/][%d] extrasScriptsReadOK ", 200)
}

func (o *ExtrasScriptsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/scripts/{id}/][%d] extrasScriptsReadOK ", 200)
}

func (o *ExtrasScriptsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasScriptsReadDefault creates a ExtrasScriptsReadDefault with default headers values
func NewExtrasScriptsReadDefault(code int) *ExtrasScriptsReadDefault {
	return &ExtrasScriptsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasScriptsReadDefault describes a response with status code -1, with default header values.

ExtrasScriptsReadDefault extras scripts read default
*/
type ExtrasScriptsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras scripts read default response
func (o *ExtrasScriptsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras scripts read default response has a 2xx status code
func (o *ExtrasScriptsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras scripts read default response has a 3xx status code
func (o *ExtrasScriptsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras scripts read default response has a 4xx status code
func (o *ExtrasScriptsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras scripts read default response has a 5xx status code
func (o *ExtrasScriptsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras scripts read default response a status code equal to that given
func (o *ExtrasScriptsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasScriptsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/{id}/][%d] extras_scripts_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasScriptsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/scripts/{id}/][%d] extras_scripts_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasScriptsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasScriptsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
