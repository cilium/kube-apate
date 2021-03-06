// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1

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

// NewGetBatchV1APIResourcesParams creates a new GetBatchV1APIResourcesParams object
// with the default values initialized.
func NewGetBatchV1APIResourcesParams() *GetBatchV1APIResourcesParams {

	return &GetBatchV1APIResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchV1APIResourcesParamsWithTimeout creates a new GetBatchV1APIResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBatchV1APIResourcesParamsWithTimeout(timeout time.Duration) *GetBatchV1APIResourcesParams {

	return &GetBatchV1APIResourcesParams{

		timeout: timeout,
	}
}

// NewGetBatchV1APIResourcesParamsWithContext creates a new GetBatchV1APIResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBatchV1APIResourcesParamsWithContext(ctx context.Context) *GetBatchV1APIResourcesParams {

	return &GetBatchV1APIResourcesParams{

		Context: ctx,
	}
}

// NewGetBatchV1APIResourcesParamsWithHTTPClient creates a new GetBatchV1APIResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBatchV1APIResourcesParamsWithHTTPClient(client *http.Client) *GetBatchV1APIResourcesParams {

	return &GetBatchV1APIResourcesParams{
		HTTPClient: client,
	}
}

/*GetBatchV1APIResourcesParams contains all the parameters to send to the API endpoint
for the get batch v1 API resources operation typically these are written to a http.Request
*/
type GetBatchV1APIResourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get batch v1 API resources params
func (o *GetBatchV1APIResourcesParams) WithTimeout(timeout time.Duration) *GetBatchV1APIResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batch v1 API resources params
func (o *GetBatchV1APIResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batch v1 API resources params
func (o *GetBatchV1APIResourcesParams) WithContext(ctx context.Context) *GetBatchV1APIResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batch v1 API resources params
func (o *GetBatchV1APIResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batch v1 API resources params
func (o *GetBatchV1APIResourcesParams) WithHTTPClient(client *http.Client) *GetBatchV1APIResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batch v1 API resources params
func (o *GetBatchV1APIResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchV1APIResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
