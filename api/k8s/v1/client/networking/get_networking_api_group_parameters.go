// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking

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

// NewGetNetworkingAPIGroupParams creates a new GetNetworkingAPIGroupParams object
// with the default values initialized.
func NewGetNetworkingAPIGroupParams() *GetNetworkingAPIGroupParams {

	return &GetNetworkingAPIGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkingAPIGroupParamsWithTimeout creates a new GetNetworkingAPIGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkingAPIGroupParamsWithTimeout(timeout time.Duration) *GetNetworkingAPIGroupParams {

	return &GetNetworkingAPIGroupParams{

		timeout: timeout,
	}
}

// NewGetNetworkingAPIGroupParamsWithContext creates a new GetNetworkingAPIGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkingAPIGroupParamsWithContext(ctx context.Context) *GetNetworkingAPIGroupParams {

	return &GetNetworkingAPIGroupParams{

		Context: ctx,
	}
}

// NewGetNetworkingAPIGroupParamsWithHTTPClient creates a new GetNetworkingAPIGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkingAPIGroupParamsWithHTTPClient(client *http.Client) *GetNetworkingAPIGroupParams {

	return &GetNetworkingAPIGroupParams{
		HTTPClient: client,
	}
}

/*GetNetworkingAPIGroupParams contains all the parameters to send to the API endpoint
for the get networking API group operation typically these are written to a http.Request
*/
type GetNetworkingAPIGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networking API group params
func (o *GetNetworkingAPIGroupParams) WithTimeout(timeout time.Duration) *GetNetworkingAPIGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networking API group params
func (o *GetNetworkingAPIGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networking API group params
func (o *GetNetworkingAPIGroupParams) WithContext(ctx context.Context) *GetNetworkingAPIGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networking API group params
func (o *GetNetworkingAPIGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networking API group params
func (o *GetNetworkingAPIGroupParams) WithHTTPClient(client *http.Client) *GetNetworkingAPIGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networking API group params
func (o *GetNetworkingAPIGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkingAPIGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
