///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api-manager/gen/models"
)

// NewUpdateAPIParams creates a new UpdateAPIParams object
// with the default values initialized.
func NewUpdateAPIParams() *UpdateAPIParams {
	var ()
	return &UpdateAPIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPIParamsWithTimeout creates a new UpdateAPIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAPIParamsWithTimeout(timeout time.Duration) *UpdateAPIParams {
	var ()
	return &UpdateAPIParams{

		timeout: timeout,
	}
}

// NewUpdateAPIParamsWithContext creates a new UpdateAPIParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAPIParamsWithContext(ctx context.Context) *UpdateAPIParams {
	var ()
	return &UpdateAPIParams{

		Context: ctx,
	}
}

// NewUpdateAPIParamsWithHTTPClient creates a new UpdateAPIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAPIParamsWithHTTPClient(client *http.Client) *UpdateAPIParams {
	var ()
	return &UpdateAPIParams{
		HTTPClient: client,
	}
}

/*UpdateAPIParams contains all the parameters to send to the API endpoint
for the update API operation typically these are written to a http.Request
*/
type UpdateAPIParams struct {

	/*API
	  Name of API to work on

	*/
	API string
	/*Body
	  API object

	*/
	Body *models.API

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update API params
func (o *UpdateAPIParams) WithTimeout(timeout time.Duration) *UpdateAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update API params
func (o *UpdateAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update API params
func (o *UpdateAPIParams) WithContext(ctx context.Context) *UpdateAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update API params
func (o *UpdateAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update API params
func (o *UpdateAPIParams) WithHTTPClient(client *http.Client) *UpdateAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update API params
func (o *UpdateAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPI adds the api to the update API params
func (o *UpdateAPIParams) WithAPI(api string) *UpdateAPIParams {
	o.SetAPI(api)
	return o
}

// SetAPI adds the api to the update API params
func (o *UpdateAPIParams) SetAPI(api string) {
	o.API = api
}

// WithBody adds the body to the update API params
func (o *UpdateAPIParams) WithBody(body *models.API) *UpdateAPIParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update API params
func (o *UpdateAPIParams) SetBody(body *models.API) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api
	if err := r.SetPathParam("api", o.API); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models.API)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}