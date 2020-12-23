// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

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

// NewReadStorageV1beta1CSIDriverParams creates a new ReadStorageV1beta1CSIDriverParams object
// with the default values initialized.
func NewReadStorageV1beta1CSIDriverParams() *ReadStorageV1beta1CSIDriverParams {
	var ()
	return &ReadStorageV1beta1CSIDriverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadStorageV1beta1CSIDriverParamsWithTimeout creates a new ReadStorageV1beta1CSIDriverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadStorageV1beta1CSIDriverParamsWithTimeout(timeout time.Duration) *ReadStorageV1beta1CSIDriverParams {
	var ()
	return &ReadStorageV1beta1CSIDriverParams{

		timeout: timeout,
	}
}

// NewReadStorageV1beta1CSIDriverParamsWithContext creates a new ReadStorageV1beta1CSIDriverParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadStorageV1beta1CSIDriverParamsWithContext(ctx context.Context) *ReadStorageV1beta1CSIDriverParams {
	var ()
	return &ReadStorageV1beta1CSIDriverParams{

		Context: ctx,
	}
}

// NewReadStorageV1beta1CSIDriverParamsWithHTTPClient creates a new ReadStorageV1beta1CSIDriverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadStorageV1beta1CSIDriverParamsWithHTTPClient(client *http.Client) *ReadStorageV1beta1CSIDriverParams {
	var ()
	return &ReadStorageV1beta1CSIDriverParams{
		HTTPClient: client,
	}
}

/*ReadStorageV1beta1CSIDriverParams contains all the parameters to send to the API endpoint
for the read storage v1beta1 c s i driver operation typically these are written to a http.Request
*/
type ReadStorageV1beta1CSIDriverParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the CSIDriver

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

// WithTimeout adds the timeout to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithTimeout(timeout time.Duration) *ReadStorageV1beta1CSIDriverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithContext(ctx context.Context) *ReadStorageV1beta1CSIDriverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithHTTPClient(client *http.Client) *ReadStorageV1beta1CSIDriverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithExact(exact *bool) *ReadStorageV1beta1CSIDriverParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithExport(export *bool) *ReadStorageV1beta1CSIDriverParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithName(name string) *ReadStorageV1beta1CSIDriverParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) WithPretty(pretty *string) *ReadStorageV1beta1CSIDriverParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read storage v1beta1 c s i driver params
func (o *ReadStorageV1beta1CSIDriverParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadStorageV1beta1CSIDriverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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