// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

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

// NewReadApiregistrationV1APIServiceStatusParams creates a new ReadApiregistrationV1APIServiceStatusParams object
// with the default values initialized.
func NewReadApiregistrationV1APIServiceStatusParams() *ReadApiregistrationV1APIServiceStatusParams {
	var ()
	return &ReadApiregistrationV1APIServiceStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadApiregistrationV1APIServiceStatusParamsWithTimeout creates a new ReadApiregistrationV1APIServiceStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadApiregistrationV1APIServiceStatusParamsWithTimeout(timeout time.Duration) *ReadApiregistrationV1APIServiceStatusParams {
	var ()
	return &ReadApiregistrationV1APIServiceStatusParams{

		timeout: timeout,
	}
}

// NewReadApiregistrationV1APIServiceStatusParamsWithContext creates a new ReadApiregistrationV1APIServiceStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadApiregistrationV1APIServiceStatusParamsWithContext(ctx context.Context) *ReadApiregistrationV1APIServiceStatusParams {
	var ()
	return &ReadApiregistrationV1APIServiceStatusParams{

		Context: ctx,
	}
}

// NewReadApiregistrationV1APIServiceStatusParamsWithHTTPClient creates a new ReadApiregistrationV1APIServiceStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadApiregistrationV1APIServiceStatusParamsWithHTTPClient(client *http.Client) *ReadApiregistrationV1APIServiceStatusParams {
	var ()
	return &ReadApiregistrationV1APIServiceStatusParams{
		HTTPClient: client,
	}
}

/*ReadApiregistrationV1APIServiceStatusParams contains all the parameters to send to the API endpoint
for the read apiregistration v1 API service status operation typically these are written to a http.Request
*/
type ReadApiregistrationV1APIServiceStatusParams struct {

	/*Name
	  name of the APIService

	*/
	Name string
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) WithTimeout(timeout time.Duration) *ReadApiregistrationV1APIServiceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) WithContext(ctx context.Context) *ReadApiregistrationV1APIServiceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) WithHTTPClient(client *http.Client) *ReadApiregistrationV1APIServiceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) WithName(name string) *ReadApiregistrationV1APIServiceStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) WithPretty(pretty *string) *ReadApiregistrationV1APIServiceStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read apiregistration v1 API service status params
func (o *ReadApiregistrationV1APIServiceStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadApiregistrationV1APIServiceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Pretty != nil {

		// query param pretty
		var qrPretty string
		if o.Pretty != nil {
			qrPretty = *o.Pretty
		}
		qPretty := qrPretty
		if qPretty != "" {
			if err := r.SetQueryParam("pretty", qPretty); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}