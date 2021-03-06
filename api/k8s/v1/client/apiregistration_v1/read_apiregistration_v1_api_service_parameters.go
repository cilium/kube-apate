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
	"github.com/go-openapi/swag"
)

// NewReadApiregistrationV1APIServiceParams creates a new ReadApiregistrationV1APIServiceParams object
// with the default values initialized.
func NewReadApiregistrationV1APIServiceParams() *ReadApiregistrationV1APIServiceParams {
	var ()
	return &ReadApiregistrationV1APIServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadApiregistrationV1APIServiceParamsWithTimeout creates a new ReadApiregistrationV1APIServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadApiregistrationV1APIServiceParamsWithTimeout(timeout time.Duration) *ReadApiregistrationV1APIServiceParams {
	var ()
	return &ReadApiregistrationV1APIServiceParams{

		timeout: timeout,
	}
}

// NewReadApiregistrationV1APIServiceParamsWithContext creates a new ReadApiregistrationV1APIServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadApiregistrationV1APIServiceParamsWithContext(ctx context.Context) *ReadApiregistrationV1APIServiceParams {
	var ()
	return &ReadApiregistrationV1APIServiceParams{

		Context: ctx,
	}
}

// NewReadApiregistrationV1APIServiceParamsWithHTTPClient creates a new ReadApiregistrationV1APIServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadApiregistrationV1APIServiceParamsWithHTTPClient(client *http.Client) *ReadApiregistrationV1APIServiceParams {
	var ()
	return &ReadApiregistrationV1APIServiceParams{
		HTTPClient: client,
	}
}

/*ReadApiregistrationV1APIServiceParams contains all the parameters to send to the API endpoint
for the read apiregistration v1 API service operation typically these are written to a http.Request
*/
type ReadApiregistrationV1APIServiceParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
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

// WithTimeout adds the timeout to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithTimeout(timeout time.Duration) *ReadApiregistrationV1APIServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithContext(ctx context.Context) *ReadApiregistrationV1APIServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithHTTPClient(client *http.Client) *ReadApiregistrationV1APIServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithExact(exact *bool) *ReadApiregistrationV1APIServiceParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithExport(export *bool) *ReadApiregistrationV1APIServiceParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithName(name string) *ReadApiregistrationV1APIServiceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) WithPretty(pretty *string) *ReadApiregistrationV1APIServiceParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read apiregistration v1 API service params
func (o *ReadApiregistrationV1APIServiceParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadApiregistrationV1APIServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exact != nil {

		// query param exact
		var qrExact bool
		if o.Exact != nil {
			qrExact = *o.Exact
		}
		qExact := swag.FormatBool(qrExact)
		if qExact != "" {
			if err := r.SetQueryParam("exact", qExact); err != nil {
				return err
			}
		}

	}

	if o.Export != nil {

		// query param export
		var qrExport bool
		if o.Export != nil {
			qrExport = *o.Export
		}
		qExport := swag.FormatBool(qrExport)
		if qExport != "" {
			if err := r.SetQueryParam("export", qExport); err != nil {
				return err
			}
		}

	}

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
