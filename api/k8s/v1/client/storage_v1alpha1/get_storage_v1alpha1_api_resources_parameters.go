// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1alpha1

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

// NewGetStorageV1alpha1APIResourcesParams creates a new GetStorageV1alpha1APIResourcesParams object
// with the default values initialized.
func NewGetStorageV1alpha1APIResourcesParams() *GetStorageV1alpha1APIResourcesParams {

	return &GetStorageV1alpha1APIResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageV1alpha1APIResourcesParamsWithTimeout creates a new GetStorageV1alpha1APIResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStorageV1alpha1APIResourcesParamsWithTimeout(timeout time.Duration) *GetStorageV1alpha1APIResourcesParams {

	return &GetStorageV1alpha1APIResourcesParams{

		timeout: timeout,
	}
}

// NewGetStorageV1alpha1APIResourcesParamsWithContext creates a new GetStorageV1alpha1APIResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStorageV1alpha1APIResourcesParamsWithContext(ctx context.Context) *GetStorageV1alpha1APIResourcesParams {

	return &GetStorageV1alpha1APIResourcesParams{

		Context: ctx,
	}
}

// NewGetStorageV1alpha1APIResourcesParamsWithHTTPClient creates a new GetStorageV1alpha1APIResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStorageV1alpha1APIResourcesParamsWithHTTPClient(client *http.Client) *GetStorageV1alpha1APIResourcesParams {

	return &GetStorageV1alpha1APIResourcesParams{
		HTTPClient: client,
	}
}

/*GetStorageV1alpha1APIResourcesParams contains all the parameters to send to the API endpoint
for the get storage v1alpha1 API resources operation typically these are written to a http.Request
*/
type GetStorageV1alpha1APIResourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage v1alpha1 API resources params
func (o *GetStorageV1alpha1APIResourcesParams) WithTimeout(timeout time.Duration) *GetStorageV1alpha1APIResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage v1alpha1 API resources params
func (o *GetStorageV1alpha1APIResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage v1alpha1 API resources params
func (o *GetStorageV1alpha1APIResourcesParams) WithContext(ctx context.Context) *GetStorageV1alpha1APIResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage v1alpha1 API resources params
func (o *GetStorageV1alpha1APIResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage v1alpha1 API resources params
func (o *GetStorageV1alpha1APIResourcesParams) WithHTTPClient(client *http.Client) *GetStorageV1alpha1APIResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage v1alpha1 API resources params
func (o *GetStorageV1alpha1APIResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageV1alpha1APIResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
