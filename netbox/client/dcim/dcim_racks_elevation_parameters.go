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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimRacksElevationParams creates a new DcimRacksElevationParams object
// with the default values initialized.
func NewDcimRacksElevationParams() *DcimRacksElevationParams {
	var (
		expandDevicesDefault = bool(true)
		faceDefault          = string("front")
		includeImagesDefault = bool(true)
		legendWidthDefault   = int64(30)
		renderDefault        = string("json")
		unitHeightDefault    = int64(22)
		unitWidthDefault     = int64(220)
	)
	return &DcimRacksElevationParams{
		ExpandDevices: &expandDevicesDefault,
		Face:          &faceDefault,
		IncludeImages: &includeImagesDefault,
		LegendWidth:   &legendWidthDefault,
		Render:        &renderDefault,
		UnitHeight:    &unitHeightDefault,
		UnitWidth:     &unitWidthDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksElevationParamsWithTimeout creates a new DcimRacksElevationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksElevationParamsWithTimeout(timeout time.Duration) *DcimRacksElevationParams {
	var (
		expandDevicesDefault = bool(true)
		faceDefault          = string("front")
		includeImagesDefault = bool(true)
		legendWidthDefault   = int64(30)
		renderDefault        = string("json")
		unitHeightDefault    = int64(22)
		unitWidthDefault     = int64(220)
	)
	return &DcimRacksElevationParams{
		ExpandDevices: &expandDevicesDefault,
		Face:          &faceDefault,
		IncludeImages: &includeImagesDefault,
		LegendWidth:   &legendWidthDefault,
		Render:        &renderDefault,
		UnitHeight:    &unitHeightDefault,
		UnitWidth:     &unitWidthDefault,

		timeout: timeout,
	}
}

// NewDcimRacksElevationParamsWithContext creates a new DcimRacksElevationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksElevationParamsWithContext(ctx context.Context) *DcimRacksElevationParams {
	var (
		expandDevicesDefault = bool(true)
		faceDefault          = string("front")
		includeImagesDefault = bool(true)
		legendWidthDefault   = int64(30)
		renderDefault        = string("json")
		unitHeightDefault    = int64(22)
		unitWidthDefault     = int64(220)
	)
	return &DcimRacksElevationParams{
		ExpandDevices: &expandDevicesDefault,
		Face:          &faceDefault,
		IncludeImages: &includeImagesDefault,
		LegendWidth:   &legendWidthDefault,
		Render:        &renderDefault,
		UnitHeight:    &unitHeightDefault,
		UnitWidth:     &unitWidthDefault,

		Context: ctx,
	}
}

// NewDcimRacksElevationParamsWithHTTPClient creates a new DcimRacksElevationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksElevationParamsWithHTTPClient(client *http.Client) *DcimRacksElevationParams {
	var (
		expandDevicesDefault = bool(true)
		faceDefault          = string("front")
		includeImagesDefault = bool(true)
		legendWidthDefault   = int64(30)
		renderDefault        = string("json")
		unitHeightDefault    = int64(22)
		unitWidthDefault     = int64(220)
	)
	return &DcimRacksElevationParams{
		ExpandDevices: &expandDevicesDefault,
		Face:          &faceDefault,
		IncludeImages: &includeImagesDefault,
		LegendWidth:   &legendWidthDefault,
		Render:        &renderDefault,
		UnitHeight:    &unitHeightDefault,
		UnitWidth:     &unitWidthDefault,
		HTTPClient:    client,
	}
}

/*DcimRacksElevationParams contains all the parameters to send to the API endpoint
for the dcim racks elevation operation typically these are written to a http.Request
*/
type DcimRacksElevationParams struct {

	/*Exclude*/
	Exclude *int64
	/*ExpandDevices*/
	ExpandDevices *bool
	/*Face*/
	Face *string
	/*ID
	  A unique integer value identifying this rack.

	*/
	ID int64
	/*IncludeImages*/
	IncludeImages *bool
	/*LegendWidth*/
	LegendWidth *int64
	/*Q*/
	Q *string
	/*Render*/
	Render *string
	/*UnitHeight*/
	UnitHeight *int64
	/*UnitWidth*/
	UnitWidth *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithTimeout(timeout time.Duration) *DcimRacksElevationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithContext(ctx context.Context) *DcimRacksElevationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithHTTPClient(client *http.Client) *DcimRacksElevationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithExclude(exclude *int64) *DcimRacksElevationParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetExclude(exclude *int64) {
	o.Exclude = exclude
}

// WithExpandDevices adds the expandDevices to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithExpandDevices(expandDevices *bool) *DcimRacksElevationParams {
	o.SetExpandDevices(expandDevices)
	return o
}

// SetExpandDevices adds the expandDevices to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetExpandDevices(expandDevices *bool) {
	o.ExpandDevices = expandDevices
}

// WithFace adds the face to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithFace(face *string) *DcimRacksElevationParams {
	o.SetFace(face)
	return o
}

// SetFace adds the face to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetFace(face *string) {
	o.Face = face
}

// WithID adds the id to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithID(id int64) *DcimRacksElevationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetID(id int64) {
	o.ID = id
}

// WithIncludeImages adds the includeImages to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithIncludeImages(includeImages *bool) *DcimRacksElevationParams {
	o.SetIncludeImages(includeImages)
	return o
}

// SetIncludeImages adds the includeImages to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetIncludeImages(includeImages *bool) {
	o.IncludeImages = includeImages
}

// WithLegendWidth adds the legendWidth to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithLegendWidth(legendWidth *int64) *DcimRacksElevationParams {
	o.SetLegendWidth(legendWidth)
	return o
}

// SetLegendWidth adds the legendWidth to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetLegendWidth(legendWidth *int64) {
	o.LegendWidth = legendWidth
}

// WithQ adds the q to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithQ(q *string) *DcimRacksElevationParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetQ(q *string) {
	o.Q = q
}

// WithRender adds the render to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithRender(render *string) *DcimRacksElevationParams {
	o.SetRender(render)
	return o
}

// SetRender adds the render to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetRender(render *string) {
	o.Render = render
}

// WithUnitHeight adds the unitHeight to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithUnitHeight(unitHeight *int64) *DcimRacksElevationParams {
	o.SetUnitHeight(unitHeight)
	return o
}

// SetUnitHeight adds the unitHeight to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetUnitHeight(unitHeight *int64) {
	o.UnitHeight = unitHeight
}

// WithUnitWidth adds the unitWidth to the dcim racks elevation params
func (o *DcimRacksElevationParams) WithUnitWidth(unitWidth *int64) *DcimRacksElevationParams {
	o.SetUnitWidth(unitWidth)
	return o
}

// SetUnitWidth adds the unitWidth to the dcim racks elevation params
func (o *DcimRacksElevationParams) SetUnitWidth(unitWidth *int64) {
	o.UnitWidth = unitWidth
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksElevationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exclude != nil {

		// query param exclude
		var qrExclude int64
		if o.Exclude != nil {
			qrExclude = *o.Exclude
		}
		qExclude := swag.FormatInt64(qrExclude)
		if qExclude != "" {
			if err := r.SetQueryParam("exclude", qExclude); err != nil {
				return err
			}
		}

	}

	if o.ExpandDevices != nil {

		// query param expand_devices
		var qrExpandDevices bool
		if o.ExpandDevices != nil {
			qrExpandDevices = *o.ExpandDevices
		}
		qExpandDevices := swag.FormatBool(qrExpandDevices)
		if qExpandDevices != "" {
			if err := r.SetQueryParam("expand_devices", qExpandDevices); err != nil {
				return err
			}
		}

	}

	if o.Face != nil {

		// query param face
		var qrFace string
		if o.Face != nil {
			qrFace = *o.Face
		}
		qFace := qrFace
		if qFace != "" {
			if err := r.SetQueryParam("face", qFace); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.IncludeImages != nil {

		// query param include_images
		var qrIncludeImages bool
		if o.IncludeImages != nil {
			qrIncludeImages = *o.IncludeImages
		}
		qIncludeImages := swag.FormatBool(qrIncludeImages)
		if qIncludeImages != "" {
			if err := r.SetQueryParam("include_images", qIncludeImages); err != nil {
				return err
			}
		}

	}

	if o.LegendWidth != nil {

		// query param legend_width
		var qrLegendWidth int64
		if o.LegendWidth != nil {
			qrLegendWidth = *o.LegendWidth
		}
		qLegendWidth := swag.FormatInt64(qrLegendWidth)
		if qLegendWidth != "" {
			if err := r.SetQueryParam("legend_width", qLegendWidth); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Render != nil {

		// query param render
		var qrRender string
		if o.Render != nil {
			qrRender = *o.Render
		}
		qRender := qrRender
		if qRender != "" {
			if err := r.SetQueryParam("render", qRender); err != nil {
				return err
			}
		}

	}

	if o.UnitHeight != nil {

		// query param unit_height
		var qrUnitHeight int64
		if o.UnitHeight != nil {
			qrUnitHeight = *o.UnitHeight
		}
		qUnitHeight := swag.FormatInt64(qrUnitHeight)
		if qUnitHeight != "" {
			if err := r.SetQueryParam("unit_height", qUnitHeight); err != nil {
				return err
			}
		}

	}

	if o.UnitWidth != nil {

		// query param unit_width
		var qrUnitWidth int64
		if o.UnitWidth != nil {
			qrUnitWidth = *o.UnitWidth
		}
		qUnitWidth := swag.FormatInt64(qrUnitWidth)
		if qUnitWidth != "" {
			if err := r.SetQueryParam("unit_width", qUnitWidth); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
