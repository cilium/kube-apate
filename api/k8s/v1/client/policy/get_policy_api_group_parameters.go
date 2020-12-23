// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

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

// NewGetPolicyAPIGroupParams creates a new GetPolicyAPIGroupParams object
// with the default values initialized.
func NewGetPolicyAPIGroupParams() *GetPolicyAPIGroupParams {

	return &GetPolicyAPIGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyAPIGroupParamsWithTimeout creates a new GetPolicyAPIGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyAPIGroupParamsWithTimeout(timeout time.Duration) *GetPolicyAPIGroupParams {

	return &GetPolicyAPIGroupParams{

		timeout: timeout,
	}
}

// NewGetPolicyAPIGroupParamsWithContext creates a new GetPolicyAPIGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyAPIGroupParamsWithContext(ctx context.Context) *GetPolicyAPIGroupParams {

	return &GetPolicyAPIGroupParams{

		Context: ctx,
	}
}

// NewGetPolicyAPIGroupParamsWithHTTPClient creates a new GetPolicyAPIGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyAPIGroupParamsWithHTTPClient(client *http.Client) *GetPolicyAPIGroupParams {

	return &GetPolicyAPIGroupParams{
		HTTPClient: client,
	}
}

/*GetPolicyAPIGroupParams contains all the parameters to send to the API endpoint
for the get policy API group operation typically these are written to a http.Request
*/
type GetPolicyAPIGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy API group params
func (o *GetPolicyAPIGroupParams) WithTimeout(timeout time.Duration) *GetPolicyAPIGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy API group params
func (o *GetPolicyAPIGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy API group params
func (o *GetPolicyAPIGroupParams) WithContext(ctx context.Context) *GetPolicyAPIGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy API group params
func (o *GetPolicyAPIGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy API group params
func (o *GetPolicyAPIGroupParams) WithHTTPClient(client *http.Client) *GetPolicyAPIGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy API group params
func (o *GetPolicyAPIGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyAPIGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}