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

	models "agent/lib/firecracker/client/models"
)

// NewPutMmdsParams creates a new PutMmdsParams object
// with the default values initialized.
func NewPutMmdsParams() *PutMmdsParams {
	var ()
	return &PutMmdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMmdsParamsWithTimeout creates a new PutMmdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMmdsParamsWithTimeout(timeout time.Duration) *PutMmdsParams {
	var ()
	return &PutMmdsParams{

		timeout: timeout,
	}
}

// NewPutMmdsParamsWithContext creates a new PutMmdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMmdsParamsWithContext(ctx context.Context) *PutMmdsParams {
	var ()
	return &PutMmdsParams{

		Context: ctx,
	}
}

// NewPutMmdsParamsWithHTTPClient creates a new PutMmdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMmdsParamsWithHTTPClient(client *http.Client) *PutMmdsParams {
	var ()
	return &PutMmdsParams{
		HTTPClient: client,
	}
}

/*PutMmdsParams contains all the parameters to send to the API endpoint
for the put mmds operation typically these are written to a http.Request
*/
type PutMmdsParams struct {

	/*Body
	  The MMDS data store as JSON.

	*/
	Body models.MmdsContentsObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put mmds params
func (o *PutMmdsParams) WithTimeout(timeout time.Duration) *PutMmdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put mmds params
func (o *PutMmdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put mmds params
func (o *PutMmdsParams) WithContext(ctx context.Context) *PutMmdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put mmds params
func (o *PutMmdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put mmds params
func (o *PutMmdsParams) WithHTTPClient(client *http.Client) *PutMmdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put mmds params
func (o *PutMmdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put mmds params
func (o *PutMmdsParams) WithBody(body models.MmdsContentsObject) *PutMmdsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put mmds params
func (o *PutMmdsParams) SetBody(body models.MmdsContentsObject) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutMmdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
