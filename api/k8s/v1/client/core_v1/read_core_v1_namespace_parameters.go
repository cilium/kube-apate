// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

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

// NewReadCoreV1NamespaceParams creates a new ReadCoreV1NamespaceParams object
// with the default values initialized.
func NewReadCoreV1NamespaceParams() *ReadCoreV1NamespaceParams {
	var ()
	return &ReadCoreV1NamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadCoreV1NamespaceParamsWithTimeout creates a new ReadCoreV1NamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadCoreV1NamespaceParamsWithTimeout(timeout time.Duration) *ReadCoreV1NamespaceParams {
	var ()
	return &ReadCoreV1NamespaceParams{

		timeout: timeout,
	}
}

// NewReadCoreV1NamespaceParamsWithContext creates a new ReadCoreV1NamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadCoreV1NamespaceParamsWithContext(ctx context.Context) *ReadCoreV1NamespaceParams {
	var ()
	return &ReadCoreV1NamespaceParams{

		Context: ctx,
	}
}

// NewReadCoreV1NamespaceParamsWithHTTPClient creates a new ReadCoreV1NamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadCoreV1NamespaceParamsWithHTTPClient(client *http.Client) *ReadCoreV1NamespaceParams {
	var ()
	return &ReadCoreV1NamespaceParams{
		HTTPClient: client,
	}
}

/*ReadCoreV1NamespaceParams contains all the parameters to send to the API endpoint
for the read core v1 namespace operation typically these are written to a http.Request
*/
type ReadCoreV1NamespaceParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the Namespace

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

// WithTimeout adds the timeout to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithTimeout(timeout time.Duration) *ReadCoreV1NamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithContext(ctx context.Context) *ReadCoreV1NamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithHTTPClient(client *http.Client) *ReadCoreV1NamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithExact(exact *bool) *ReadCoreV1NamespaceParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithExport(export *bool) *ReadCoreV1NamespaceParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithName(name string) *ReadCoreV1NamespaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) WithPretty(pretty *string) *ReadCoreV1NamespaceParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read core v1 namespace params
func (o *ReadCoreV1NamespaceParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadCoreV1NamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
