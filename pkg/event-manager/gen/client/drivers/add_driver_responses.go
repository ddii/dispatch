///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// AddDriverReader is a Reader for the AddDriver structure.
type AddDriverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDriverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddDriverCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddDriverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddDriverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddDriverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddDriverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDriverCreated creates a AddDriverCreated with default headers values
func NewAddDriverCreated() *AddDriverCreated {
	return &AddDriverCreated{}
}

/*AddDriverCreated handles this case with default header values.

Driver created
*/
type AddDriverCreated struct {
	Payload *models.Driver
}

func (o *AddDriverCreated) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] addDriverCreated  %+v", 201, o.Payload)
}

func (o *AddDriverCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Driver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverBadRequest creates a AddDriverBadRequest with default headers values
func NewAddDriverBadRequest() *AddDriverBadRequest {
	return &AddDriverBadRequest{}
}

/*AddDriverBadRequest handles this case with default header values.

Invalid input
*/
type AddDriverBadRequest struct {
	Payload *models.Error
}

func (o *AddDriverBadRequest) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] addDriverBadRequest  %+v", 400, o.Payload)
}

func (o *AddDriverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverUnauthorized creates a AddDriverUnauthorized with default headers values
func NewAddDriverUnauthorized() *AddDriverUnauthorized {
	return &AddDriverUnauthorized{}
}

/*AddDriverUnauthorized handles this case with default header values.

Unauthorized Request
*/
type AddDriverUnauthorized struct {
	Payload *models.Error
}

func (o *AddDriverUnauthorized) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] addDriverUnauthorized  %+v", 401, o.Payload)
}

func (o *AddDriverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverInternalServerError creates a AddDriverInternalServerError with default headers values
func NewAddDriverInternalServerError() *AddDriverInternalServerError {
	return &AddDriverInternalServerError{}
}

/*AddDriverInternalServerError handles this case with default header values.

Internal server error
*/
type AddDriverInternalServerError struct {
	Payload *models.Error
}

func (o *AddDriverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] addDriverInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDriverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDriverDefault creates a AddDriverDefault with default headers values
func NewAddDriverDefault(code int) *AddDriverDefault {
	return &AddDriverDefault{
		_statusCode: code,
	}
}

/*AddDriverDefault handles this case with default header values.

Unknown error
*/
type AddDriverDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add driver default response
func (o *AddDriverDefault) Code() int {
	return o._statusCode
}

func (o *AddDriverDefault) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] addDriver default  %+v", o._statusCode, o.Payload)
}

func (o *AddDriverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}