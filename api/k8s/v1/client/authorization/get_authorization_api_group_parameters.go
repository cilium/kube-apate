// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization

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

// NewGetAuthorizationAPIGroupParams creates a new GetAuthorizationAPIGroupParams object
// with the default values initialized.
func NewGetAuthorizationAPIGroupParams() *GetAuthorizationAPIGroupParams {

	return &GetAuthorizationAPIGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationAPIGroupParamsWithTimeout creates a new GetAuthorizationAPIGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationAPIGroupParamsWithTimeout(timeout time.Duration) *GetAuthorizationAPIGroupParams {

	return &GetAuthorizationAPIGroupParams{

		timeout: timeout,
	}
}

// NewGetAuthorizationAPIGroupParamsWithContext creates a new GetAuthorizationAPIGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationAPIGroupParamsWithContext(ctx context.Context) *GetAuthorizationAPIGroupParams {

	return &GetAuthorizationAPIGroupParams{

		Context: ctx,
	}
}

// NewGetAuthorizationAPIGroupParamsWithHTTPClient creates a new GetAuthorizationAPIGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationAPIGroupParamsWithHTTPClient(client *http.Client) *GetAuthorizationAPIGroupParams {

	return &GetAuthorizationAPIGroupParams{
		HTTPClient: client,
	}
}

/*GetAuthorizationAPIGroupParams contains all the parameters to send to the API endpoint
for the get authorization API group operation typically these are written to a http.Request
*/
type GetAuthorizationAPIGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization API group params
func (o *GetAuthorizationAPIGroupParams) WithTimeout(timeout time.Duration) *GetAuthorizationAPIGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization API group params
func (o *GetAuthorizationAPIGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization API group params
func (o *GetAuthorizationAPIGroupParams) WithContext(ctx context.Context) *GetAuthorizationAPIGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization API group params
func (o *GetAuthorizationAPIGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization API group params
func (o *GetAuthorizationAPIGroupParams) WithHTTPClient(client *http.Client) *GetAuthorizationAPIGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization API group params
func (o *GetAuthorizationAPIGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationAPIGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
