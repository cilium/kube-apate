// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1beta1

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

// NewReadSchedulingV1beta1PriorityClassParams creates a new ReadSchedulingV1beta1PriorityClassParams object
// with the default values initialized.
func NewReadSchedulingV1beta1PriorityClassParams() *ReadSchedulingV1beta1PriorityClassParams {
	var ()
	return &ReadSchedulingV1beta1PriorityClassParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadSchedulingV1beta1PriorityClassParamsWithTimeout creates a new ReadSchedulingV1beta1PriorityClassParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadSchedulingV1beta1PriorityClassParamsWithTimeout(timeout time.Duration) *ReadSchedulingV1beta1PriorityClassParams {
	var ()
	return &ReadSchedulingV1beta1PriorityClassParams{

		timeout: timeout,
	}
}

// NewReadSchedulingV1beta1PriorityClassParamsWithContext creates a new ReadSchedulingV1beta1PriorityClassParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadSchedulingV1beta1PriorityClassParamsWithContext(ctx context.Context) *ReadSchedulingV1beta1PriorityClassParams {
	var ()
	return &ReadSchedulingV1beta1PriorityClassParams{

		Context: ctx,
	}
}

// NewReadSchedulingV1beta1PriorityClassParamsWithHTTPClient creates a new ReadSchedulingV1beta1PriorityClassParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadSchedulingV1beta1PriorityClassParamsWithHTTPClient(client *http.Client) *ReadSchedulingV1beta1PriorityClassParams {
	var ()
	return &ReadSchedulingV1beta1PriorityClassParams{
		HTTPClient: client,
	}
}

/*ReadSchedulingV1beta1PriorityClassParams contains all the parameters to send to the API endpoint
for the read scheduling v1beta1 priority class operation typically these are written to a http.Request
*/
type ReadSchedulingV1beta1PriorityClassParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the PriorityClass

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

// WithTimeout adds the timeout to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithTimeout(timeout time.Duration) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithContext(ctx context.Context) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithHTTPClient(client *http.Client) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithExact(exact *bool) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithExport(export *bool) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithName(name string) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) WithPretty(pretty *string) *ReadSchedulingV1beta1PriorityClassParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read scheduling v1beta1 priority class params
func (o *ReadSchedulingV1beta1PriorityClassParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadSchedulingV1beta1PriorityClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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