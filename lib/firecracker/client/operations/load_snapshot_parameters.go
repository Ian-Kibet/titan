// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

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

<<<<<<< HEAD
	models "titan/lib/firecracker/client/models"
=======
	models "github.com/Ian-Kibet/firecracker-go-sdk/client/models"
>>>>>>> b8aa219df3977843c18fb0ce7af8af072b1bf0b8
)

// NewLoadSnapshotParams creates a new LoadSnapshotParams object
// with the default values initialized.
func NewLoadSnapshotParams() *LoadSnapshotParams {
	var ()
	return &LoadSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadSnapshotParamsWithTimeout creates a new LoadSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadSnapshotParamsWithTimeout(timeout time.Duration) *LoadSnapshotParams {
	var ()
	return &LoadSnapshotParams{

		timeout: timeout,
	}
}

// NewLoadSnapshotParamsWithContext creates a new LoadSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadSnapshotParamsWithContext(ctx context.Context) *LoadSnapshotParams {
	var ()
	return &LoadSnapshotParams{

		Context: ctx,
	}
}

// NewLoadSnapshotParamsWithHTTPClient creates a new LoadSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadSnapshotParamsWithHTTPClient(client *http.Client) *LoadSnapshotParams {
	var ()
	return &LoadSnapshotParams{
		HTTPClient: client,
	}
}

/*LoadSnapshotParams contains all the parameters to send to the API endpoint
for the load snapshot operation typically these are written to a http.Request
*/
type LoadSnapshotParams struct {

	/*Body
	  The configuration used for loading a snaphot.

	*/
	Body *models.SnapshotLoadParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load snapshot params
func (o *LoadSnapshotParams) WithTimeout(timeout time.Duration) *LoadSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load snapshot params
func (o *LoadSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load snapshot params
func (o *LoadSnapshotParams) WithContext(ctx context.Context) *LoadSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load snapshot params
func (o *LoadSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load snapshot params
func (o *LoadSnapshotParams) WithHTTPClient(client *http.Client) *LoadSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load snapshot params
func (o *LoadSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the load snapshot params
func (o *LoadSnapshotParams) WithBody(body *models.SnapshotLoadParams) *LoadSnapshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the load snapshot params
func (o *LoadSnapshotParams) SetBody(body *models.SnapshotLoadParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LoadSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}