// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package circuits

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

	models "github.com/netbox-community/go-netbox/models"
)

// NewCircuitsProvidersPartialUpdateParams creates a new CircuitsProvidersPartialUpdateParams object
// with the default values initialized.
func NewCircuitsProvidersPartialUpdateParams() *CircuitsProvidersPartialUpdateParams {
	var ()
	return &CircuitsProvidersPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersPartialUpdateParamsWithTimeout creates a new CircuitsProvidersPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsProvidersPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersPartialUpdateParams {
	var ()
	return &CircuitsProvidersPartialUpdateParams{

		timeout: timeout,
	}
}

// NewCircuitsProvidersPartialUpdateParamsWithContext creates a new CircuitsProvidersPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsProvidersPartialUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersPartialUpdateParams {
	var ()
	return &CircuitsProvidersPartialUpdateParams{

		Context: ctx,
	}
}

// NewCircuitsProvidersPartialUpdateParamsWithHTTPClient creates a new CircuitsProvidersPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsProvidersPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersPartialUpdateParams {
	var ()
	return &CircuitsProvidersPartialUpdateParams{
		HTTPClient: client,
	}
}

/*CircuitsProvidersPartialUpdateParams contains all the parameters to send to the API endpoint
for the circuits providers partial update operation typically these are written to a http.Request
*/
type CircuitsProvidersPartialUpdateParams struct {

	/*Data*/
	Data *models.Provider
	/*ID
	  A unique integer value identifying this provider.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithData(data *models.Provider) *CircuitsProvidersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WithID adds the id to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) WithID(id int64) *CircuitsProvidersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers partial update params
func (o *CircuitsProvidersPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
