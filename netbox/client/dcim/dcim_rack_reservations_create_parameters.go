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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewDcimRackReservationsCreateParams creates a new DcimRackReservationsCreateParams object
// with the default values initialized.
func NewDcimRackReservationsCreateParams() *DcimRackReservationsCreateParams {
	var ()
	return &DcimRackReservationsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsCreateParamsWithTimeout creates a new DcimRackReservationsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackReservationsCreateParamsWithTimeout(timeout time.Duration) *DcimRackReservationsCreateParams {
	var ()
	return &DcimRackReservationsCreateParams{

		timeout: timeout,
	}
}

// NewDcimRackReservationsCreateParamsWithContext creates a new DcimRackReservationsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackReservationsCreateParamsWithContext(ctx context.Context) *DcimRackReservationsCreateParams {
	var ()
	return &DcimRackReservationsCreateParams{

		Context: ctx,
	}
}

// NewDcimRackReservationsCreateParamsWithHTTPClient creates a new DcimRackReservationsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackReservationsCreateParamsWithHTTPClient(client *http.Client) *DcimRackReservationsCreateParams {
	var ()
	return &DcimRackReservationsCreateParams{
		HTTPClient: client,
	}
}

/*DcimRackReservationsCreateParams contains all the parameters to send to the API endpoint
for the dcim rack reservations create operation typically these are written to a http.Request
*/
type DcimRackReservationsCreateParams struct {

	/*Data*/
	Data *models.WritableRackReservation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) WithTimeout(timeout time.Duration) *DcimRackReservationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) WithContext(ctx context.Context) *DcimRackReservationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) WithHTTPClient(client *http.Client) *DcimRackReservationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) WithData(data *models.WritableRackReservation) *DcimRackReservationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack reservations create params
func (o *DcimRackReservationsCreateParams) SetData(data *models.WritableRackReservation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
