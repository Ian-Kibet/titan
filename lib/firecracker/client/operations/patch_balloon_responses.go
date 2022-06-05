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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

<<<<<<< HEAD
	models "titan/lib/firecracker/client/models"
=======
	models "github.com/Ian-Kibet/firecracker-go-sdk/client/models"
>>>>>>> b8aa219df3977843c18fb0ce7af8af072b1bf0b8
)

// PatchBalloonReader is a Reader for the PatchBalloon structure.
type PatchBalloonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBalloonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchBalloonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchBalloonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchBalloonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchBalloonNoContent creates a PatchBalloonNoContent with default headers values
func NewPatchBalloonNoContent() *PatchBalloonNoContent {
	return &PatchBalloonNoContent{}
}

/*PatchBalloonNoContent handles this case with default header values.

Balloon device updated
*/
type PatchBalloonNoContent struct {
}

func (o *PatchBalloonNoContent) Error() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloonNoContent ", 204)
}

func (o *PatchBalloonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchBalloonBadRequest creates a PatchBalloonBadRequest with default headers values
func NewPatchBalloonBadRequest() *PatchBalloonBadRequest {
	return &PatchBalloonBadRequest{}
}

/*PatchBalloonBadRequest handles this case with default header values.

Balloon device cannot be updated due to bad input
*/
type PatchBalloonBadRequest struct {
	Payload *models.Error
}

func (o *PatchBalloonBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloonBadRequest  %+v", 400, o.Payload)
}

func (o *PatchBalloonBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBalloonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBalloonDefault creates a PatchBalloonDefault with default headers values
func NewPatchBalloonDefault(code int) *PatchBalloonDefault {
	return &PatchBalloonDefault{
		_statusCode: code,
	}
}

/*PatchBalloonDefault handles this case with default header values.

Internal server error
*/
type PatchBalloonDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch balloon default response
func (o *PatchBalloonDefault) Code() int {
	return o._statusCode
}

func (o *PatchBalloonDefault) Error() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloon default  %+v", o._statusCode, o.Payload)
}

func (o *PatchBalloonDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBalloonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}