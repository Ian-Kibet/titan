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

	models "agent/lib/firecracker/client/models"
)

// DescribeBalloonConfigReader is a Reader for the DescribeBalloonConfig structure.
type DescribeBalloonConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeBalloonConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeBalloonConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeBalloonConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDescribeBalloonConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeBalloonConfigOK creates a DescribeBalloonConfigOK with default headers values
func NewDescribeBalloonConfigOK() *DescribeBalloonConfigOK {
	return &DescribeBalloonConfigOK{}
}

/*DescribeBalloonConfigOK handles this case with default header values.

The balloon device configuration
*/
type DescribeBalloonConfigOK struct {
	Payload *models.Balloon
}

func (o *DescribeBalloonConfigOK) Error() string {
	return fmt.Sprintf("[GET /balloon][%d] describeBalloonConfigOK  %+v", 200, o.Payload)
}

func (o *DescribeBalloonConfigOK) GetPayload() *models.Balloon {
	return o.Payload
}

func (o *DescribeBalloonConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Balloon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeBalloonConfigBadRequest creates a DescribeBalloonConfigBadRequest with default headers values
func NewDescribeBalloonConfigBadRequest() *DescribeBalloonConfigBadRequest {
	return &DescribeBalloonConfigBadRequest{}
}

/*DescribeBalloonConfigBadRequest handles this case with default header values.

Balloon device not configured.
*/
type DescribeBalloonConfigBadRequest struct {
	Payload *models.Error
}

func (o *DescribeBalloonConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /balloon][%d] describeBalloonConfigBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeBalloonConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeBalloonConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeBalloonConfigDefault creates a DescribeBalloonConfigDefault with default headers values
func NewDescribeBalloonConfigDefault(code int) *DescribeBalloonConfigDefault {
	return &DescribeBalloonConfigDefault{
		_statusCode: code,
	}
}

/*DescribeBalloonConfigDefault handles this case with default header values.

Internal Server Error
*/
type DescribeBalloonConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the describe balloon config default response
func (o *DescribeBalloonConfigDefault) Code() int {
	return o._statusCode
}

func (o *DescribeBalloonConfigDefault) Error() string {
	return fmt.Sprintf("[GET /balloon][%d] describeBalloonConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeBalloonConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeBalloonConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
