// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetCodeVersionParams creates a new GetCodeVersionParams object
// with the default values initialized.
func NewGetCodeVersionParams() *GetCodeVersionParams {

	return &GetCodeVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCodeVersionParamsWithTimeout creates a new GetCodeVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCodeVersionParamsWithTimeout(timeout time.Duration) *GetCodeVersionParams {

	return &GetCodeVersionParams{

		timeout: timeout,
	}
}

// NewGetCodeVersionParamsWithContext creates a new GetCodeVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCodeVersionParamsWithContext(ctx context.Context) *GetCodeVersionParams {

	return &GetCodeVersionParams{

		Context: ctx,
	}
}

// NewGetCodeVersionParamsWithHTTPClient creates a new GetCodeVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCodeVersionParamsWithHTTPClient(client *http.Client) *GetCodeVersionParams {

	return &GetCodeVersionParams{
		HTTPClient: client,
	}
}

/*GetCodeVersionParams contains all the parameters to send to the API endpoint
for the get code version operation typically these are written to a http.Request
*/
type GetCodeVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get code version params
func (o *GetCodeVersionParams) WithTimeout(timeout time.Duration) *GetCodeVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get code version params
func (o *GetCodeVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get code version params
func (o *GetCodeVersionParams) WithContext(ctx context.Context) *GetCodeVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get code version params
func (o *GetCodeVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get code version params
func (o *GetCodeVersionParams) WithHTTPClient(client *http.Client) *GetCodeVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get code version params
func (o *GetCodeVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCodeVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
