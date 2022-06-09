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

// NewPutGuestDriveByIDParams creates a new PutGuestDriveByIDParams object
// with the default values initialized.
func NewPutGuestDriveByIDParams() *PutGuestDriveByIDParams {
	var ()
	return &PutGuestDriveByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutGuestDriveByIDParamsWithTimeout creates a new PutGuestDriveByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutGuestDriveByIDParamsWithTimeout(timeout time.Duration) *PutGuestDriveByIDParams {
	var ()
	return &PutGuestDriveByIDParams{

		timeout: timeout,
	}
}

// NewPutGuestDriveByIDParamsWithContext creates a new PutGuestDriveByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutGuestDriveByIDParamsWithContext(ctx context.Context) *PutGuestDriveByIDParams {
	var ()
	return &PutGuestDriveByIDParams{

		Context: ctx,
	}
}

// NewPutGuestDriveByIDParamsWithHTTPClient creates a new PutGuestDriveByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutGuestDriveByIDParamsWithHTTPClient(client *http.Client) *PutGuestDriveByIDParams {
	var ()
	return &PutGuestDriveByIDParams{
		HTTPClient: client,
	}
}

/*PutGuestDriveByIDParams contains all the parameters to send to the API endpoint
for the put guest drive by ID operation typically these are written to a http.Request
*/
type PutGuestDriveByIDParams struct {

	/*Body
	  Guest drive properties

	*/
	Body *models.Drive
	/*DriveID
	  The id of the guest drive

	*/
	DriveID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) WithTimeout(timeout time.Duration) *PutGuestDriveByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) WithContext(ctx context.Context) *PutGuestDriveByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) WithHTTPClient(client *http.Client) *PutGuestDriveByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) WithBody(body *models.Drive) *PutGuestDriveByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) SetBody(body *models.Drive) {
	o.Body = body
}

// WithDriveID adds the driveID to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) WithDriveID(driveID string) *PutGuestDriveByIDParams {
	o.SetDriveID(driveID)
	return o
}

// SetDriveID adds the driveId to the put guest drive by ID params
func (o *PutGuestDriveByIDParams) SetDriveID(driveID string) {
	o.DriveID = driveID
}

// WriteToRequest writes these params to a swagger request
func (o *PutGuestDriveByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param drive_id
	if err := r.SetPathParam("drive_id", o.DriveID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
