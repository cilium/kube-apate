// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1alpha1

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

// NewGetRbacAuthorizationV1alpha1APIResourcesParams creates a new GetRbacAuthorizationV1alpha1APIResourcesParams object
// with the default values initialized.
func NewGetRbacAuthorizationV1alpha1APIResourcesParams() *GetRbacAuthorizationV1alpha1APIResourcesParams {

	return &GetRbacAuthorizationV1alpha1APIResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRbacAuthorizationV1alpha1APIResourcesParamsWithTimeout creates a new GetRbacAuthorizationV1alpha1APIResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRbacAuthorizationV1alpha1APIResourcesParamsWithTimeout(timeout time.Duration) *GetRbacAuthorizationV1alpha1APIResourcesParams {

	return &GetRbacAuthorizationV1alpha1APIResourcesParams{

		timeout: timeout,
	}
}

// NewGetRbacAuthorizationV1alpha1APIResourcesParamsWithContext creates a new GetRbacAuthorizationV1alpha1APIResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRbacAuthorizationV1alpha1APIResourcesParamsWithContext(ctx context.Context) *GetRbacAuthorizationV1alpha1APIResourcesParams {

	return &GetRbacAuthorizationV1alpha1APIResourcesParams{

		Context: ctx,
	}
}

// NewGetRbacAuthorizationV1alpha1APIResourcesParamsWithHTTPClient creates a new GetRbacAuthorizationV1alpha1APIResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRbacAuthorizationV1alpha1APIResourcesParamsWithHTTPClient(client *http.Client) *GetRbacAuthorizationV1alpha1APIResourcesParams {

	return &GetRbacAuthorizationV1alpha1APIResourcesParams{
		HTTPClient: client,
	}
}

/*GetRbacAuthorizationV1alpha1APIResourcesParams contains all the parameters to send to the API endpoint
for the get rbac authorization v1alpha1 API resources operation typically these are written to a http.Request
*/
type GetRbacAuthorizationV1alpha1APIResourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rbac authorization v1alpha1 API resources params
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) WithTimeout(timeout time.Duration) *GetRbacAuthorizationV1alpha1APIResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rbac authorization v1alpha1 API resources params
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rbac authorization v1alpha1 API resources params
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) WithContext(ctx context.Context) *GetRbacAuthorizationV1alpha1APIResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rbac authorization v1alpha1 API resources params
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rbac authorization v1alpha1 API resources params
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) WithHTTPClient(client *http.Client) *GetRbacAuthorizationV1alpha1APIResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rbac authorization v1alpha1 API resources params
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRbacAuthorizationV1alpha1APIResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
