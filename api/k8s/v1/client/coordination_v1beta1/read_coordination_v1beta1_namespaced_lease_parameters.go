// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1beta1

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

// NewReadCoordinationV1beta1NamespacedLeaseParams creates a new ReadCoordinationV1beta1NamespacedLeaseParams object
// with the default values initialized.
func NewReadCoordinationV1beta1NamespacedLeaseParams() *ReadCoordinationV1beta1NamespacedLeaseParams {
	var ()
	return &ReadCoordinationV1beta1NamespacedLeaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadCoordinationV1beta1NamespacedLeaseParamsWithTimeout creates a new ReadCoordinationV1beta1NamespacedLeaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadCoordinationV1beta1NamespacedLeaseParamsWithTimeout(timeout time.Duration) *ReadCoordinationV1beta1NamespacedLeaseParams {
	var ()
	return &ReadCoordinationV1beta1NamespacedLeaseParams{

		timeout: timeout,
	}
}

// NewReadCoordinationV1beta1NamespacedLeaseParamsWithContext creates a new ReadCoordinationV1beta1NamespacedLeaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadCoordinationV1beta1NamespacedLeaseParamsWithContext(ctx context.Context) *ReadCoordinationV1beta1NamespacedLeaseParams {
	var ()
	return &ReadCoordinationV1beta1NamespacedLeaseParams{

		Context: ctx,
	}
}

// NewReadCoordinationV1beta1NamespacedLeaseParamsWithHTTPClient creates a new ReadCoordinationV1beta1NamespacedLeaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadCoordinationV1beta1NamespacedLeaseParamsWithHTTPClient(client *http.Client) *ReadCoordinationV1beta1NamespacedLeaseParams {
	var ()
	return &ReadCoordinationV1beta1NamespacedLeaseParams{
		HTTPClient: client,
	}
}

/*ReadCoordinationV1beta1NamespacedLeaseParams contains all the parameters to send to the API endpoint
for the read coordination v1beta1 namespaced lease operation typically these are written to a http.Request
*/
type ReadCoordinationV1beta1NamespacedLeaseParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the Lease

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithTimeout(timeout time.Duration) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithContext(ctx context.Context) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithHTTPClient(client *http.Client) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithExact(exact *bool) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithExport(export *bool) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithName(name string) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithNamespace(namespace string) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WithPretty(pretty *string) *ReadCoordinationV1beta1NamespacedLeaseParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read coordination v1beta1 namespaced lease params
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadCoordinationV1beta1NamespacedLeaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
