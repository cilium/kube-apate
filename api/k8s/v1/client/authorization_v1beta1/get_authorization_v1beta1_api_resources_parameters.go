// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1beta1

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

// NewGetAuthorizationV1beta1APIResourcesParams creates a new GetAuthorizationV1beta1APIResourcesParams object
// with the default values initialized.
func NewGetAuthorizationV1beta1APIResourcesParams() *GetAuthorizationV1beta1APIResourcesParams {

	return &GetAuthorizationV1beta1APIResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationV1beta1APIResourcesParamsWithTimeout creates a new GetAuthorizationV1beta1APIResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationV1beta1APIResourcesParamsWithTimeout(timeout time.Duration) *GetAuthorizationV1beta1APIResourcesParams {

	return &GetAuthorizationV1beta1APIResourcesParams{

		timeout: timeout,
	}
}

// NewGetAuthorizationV1beta1APIResourcesParamsWithContext creates a new GetAuthorizationV1beta1APIResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationV1beta1APIResourcesParamsWithContext(ctx context.Context) *GetAuthorizationV1beta1APIResourcesParams {

	return &GetAuthorizationV1beta1APIResourcesParams{

		Context: ctx,
	}
}

// NewGetAuthorizationV1beta1APIResourcesParamsWithHTTPClient creates a new GetAuthorizationV1beta1APIResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationV1beta1APIResourcesParamsWithHTTPClient(client *http.Client) *GetAuthorizationV1beta1APIResourcesParams {

	return &GetAuthorizationV1beta1APIResourcesParams{
		HTTPClient: client,
	}
}

/*GetAuthorizationV1beta1APIResourcesParams contains all the parameters to send to the API endpoint
for the get authorization v1beta1 API resources operation typically these are written to a http.Request
*/
type GetAuthorizationV1beta1APIResourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization v1beta1 API resources params
func (o *GetAuthorizationV1beta1APIResourcesParams) WithTimeout(timeout time.Duration) *GetAuthorizationV1beta1APIResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization v1beta1 API resources params
func (o *GetAuthorizationV1beta1APIResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization v1beta1 API resources params
func (o *GetAuthorizationV1beta1APIResourcesParams) WithContext(ctx context.Context) *GetAuthorizationV1beta1APIResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization v1beta1 API resources params
func (o *GetAuthorizationV1beta1APIResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization v1beta1 API resources params
func (o *GetAuthorizationV1beta1APIResourcesParams) WithHTTPClient(client *http.Client) *GetAuthorizationV1beta1APIResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization v1beta1 API resources params
func (o *GetAuthorizationV1beta1APIResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationV1beta1APIResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
