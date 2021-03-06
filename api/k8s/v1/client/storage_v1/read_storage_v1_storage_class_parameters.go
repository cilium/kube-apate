// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

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

// NewReadStorageV1StorageClassParams creates a new ReadStorageV1StorageClassParams object
// with the default values initialized.
func NewReadStorageV1StorageClassParams() *ReadStorageV1StorageClassParams {
	var ()
	return &ReadStorageV1StorageClassParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadStorageV1StorageClassParamsWithTimeout creates a new ReadStorageV1StorageClassParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadStorageV1StorageClassParamsWithTimeout(timeout time.Duration) *ReadStorageV1StorageClassParams {
	var ()
	return &ReadStorageV1StorageClassParams{

		timeout: timeout,
	}
}

// NewReadStorageV1StorageClassParamsWithContext creates a new ReadStorageV1StorageClassParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadStorageV1StorageClassParamsWithContext(ctx context.Context) *ReadStorageV1StorageClassParams {
	var ()
	return &ReadStorageV1StorageClassParams{

		Context: ctx,
	}
}

// NewReadStorageV1StorageClassParamsWithHTTPClient creates a new ReadStorageV1StorageClassParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadStorageV1StorageClassParamsWithHTTPClient(client *http.Client) *ReadStorageV1StorageClassParams {
	var ()
	return &ReadStorageV1StorageClassParams{
		HTTPClient: client,
	}
}

/*ReadStorageV1StorageClassParams contains all the parameters to send to the API endpoint
for the read storage v1 storage class operation typically these are written to a http.Request
*/
type ReadStorageV1StorageClassParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the StorageClass

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

// WithTimeout adds the timeout to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithTimeout(timeout time.Duration) *ReadStorageV1StorageClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithContext(ctx context.Context) *ReadStorageV1StorageClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithHTTPClient(client *http.Client) *ReadStorageV1StorageClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithExact(exact *bool) *ReadStorageV1StorageClassParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithExport(export *bool) *ReadStorageV1StorageClassParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithName(name string) *ReadStorageV1StorageClassParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) WithPretty(pretty *string) *ReadStorageV1StorageClassParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read storage v1 storage class params
func (o *ReadStorageV1StorageClassParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadStorageV1StorageClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
