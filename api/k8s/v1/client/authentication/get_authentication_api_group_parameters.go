// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication

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

// NewGetAuthenticationAPIGroupParams creates a new GetAuthenticationAPIGroupParams object
// with the default values initialized.
func NewGetAuthenticationAPIGroupParams() *GetAuthenticationAPIGroupParams {

	return &GetAuthenticationAPIGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthenticationAPIGroupParamsWithTimeout creates a new GetAuthenticationAPIGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthenticationAPIGroupParamsWithTimeout(timeout time.Duration) *GetAuthenticationAPIGroupParams {

	return &GetAuthenticationAPIGroupParams{

		timeout: timeout,
	}
}

// NewGetAuthenticationAPIGroupParamsWithContext creates a new GetAuthenticationAPIGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthenticationAPIGroupParamsWithContext(ctx context.Context) *GetAuthenticationAPIGroupParams {

	return &GetAuthenticationAPIGroupParams{

		Context: ctx,
	}
}

// NewGetAuthenticationAPIGroupParamsWithHTTPClient creates a new GetAuthenticationAPIGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthenticationAPIGroupParamsWithHTTPClient(client *http.Client) *GetAuthenticationAPIGroupParams {

	return &GetAuthenticationAPIGroupParams{
		HTTPClient: client,
	}
}

/*GetAuthenticationAPIGroupParams contains all the parameters to send to the API endpoint
for the get authentication API group operation typically these are written to a http.Request
*/
type GetAuthenticationAPIGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authentication API group params
func (o *GetAuthenticationAPIGroupParams) WithTimeout(timeout time.Duration) *GetAuthenticationAPIGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authentication API group params
func (o *GetAuthenticationAPIGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authentication API group params
func (o *GetAuthenticationAPIGroupParams) WithContext(ctx context.Context) *GetAuthenticationAPIGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authentication API group params
func (o *GetAuthenticationAPIGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authentication API group params
func (o *GetAuthenticationAPIGroupParams) WithHTTPClient(client *http.Client) *GetAuthenticationAPIGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authentication API group params
func (o *GetAuthenticationAPIGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthenticationAPIGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
