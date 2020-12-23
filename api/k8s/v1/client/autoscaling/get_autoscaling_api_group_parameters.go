// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling

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

// NewGetAutoscalingAPIGroupParams creates a new GetAutoscalingAPIGroupParams object
// with the default values initialized.
func NewGetAutoscalingAPIGroupParams() *GetAutoscalingAPIGroupParams {

	return &GetAutoscalingAPIGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAutoscalingAPIGroupParamsWithTimeout creates a new GetAutoscalingAPIGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAutoscalingAPIGroupParamsWithTimeout(timeout time.Duration) *GetAutoscalingAPIGroupParams {

	return &GetAutoscalingAPIGroupParams{

		timeout: timeout,
	}
}

// NewGetAutoscalingAPIGroupParamsWithContext creates a new GetAutoscalingAPIGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAutoscalingAPIGroupParamsWithContext(ctx context.Context) *GetAutoscalingAPIGroupParams {

	return &GetAutoscalingAPIGroupParams{

		Context: ctx,
	}
}

// NewGetAutoscalingAPIGroupParamsWithHTTPClient creates a new GetAutoscalingAPIGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAutoscalingAPIGroupParamsWithHTTPClient(client *http.Client) *GetAutoscalingAPIGroupParams {

	return &GetAutoscalingAPIGroupParams{
		HTTPClient: client,
	}
}

/*GetAutoscalingAPIGroupParams contains all the parameters to send to the API endpoint
for the get autoscaling API group operation typically these are written to a http.Request
*/
type GetAutoscalingAPIGroupParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get autoscaling API group params
func (o *GetAutoscalingAPIGroupParams) WithTimeout(timeout time.Duration) *GetAutoscalingAPIGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get autoscaling API group params
func (o *GetAutoscalingAPIGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get autoscaling API group params
func (o *GetAutoscalingAPIGroupParams) WithContext(ctx context.Context) *GetAutoscalingAPIGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get autoscaling API group params
func (o *GetAutoscalingAPIGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get autoscaling API group params
func (o *GetAutoscalingAPIGroupParams) WithHTTPClient(client *http.Client) *GetAutoscalingAPIGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get autoscaling API group params
func (o *GetAutoscalingAPIGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAutoscalingAPIGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}