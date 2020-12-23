// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

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

// NewReadAppsV1NamespacedControllerRevisionParams creates a new ReadAppsV1NamespacedControllerRevisionParams object
// with the default values initialized.
func NewReadAppsV1NamespacedControllerRevisionParams() *ReadAppsV1NamespacedControllerRevisionParams {
	var ()
	return &ReadAppsV1NamespacedControllerRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadAppsV1NamespacedControllerRevisionParamsWithTimeout creates a new ReadAppsV1NamespacedControllerRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadAppsV1NamespacedControllerRevisionParamsWithTimeout(timeout time.Duration) *ReadAppsV1NamespacedControllerRevisionParams {
	var ()
	return &ReadAppsV1NamespacedControllerRevisionParams{

		timeout: timeout,
	}
}

// NewReadAppsV1NamespacedControllerRevisionParamsWithContext creates a new ReadAppsV1NamespacedControllerRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadAppsV1NamespacedControllerRevisionParamsWithContext(ctx context.Context) *ReadAppsV1NamespacedControllerRevisionParams {
	var ()
	return &ReadAppsV1NamespacedControllerRevisionParams{

		Context: ctx,
	}
}

// NewReadAppsV1NamespacedControllerRevisionParamsWithHTTPClient creates a new ReadAppsV1NamespacedControllerRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadAppsV1NamespacedControllerRevisionParamsWithHTTPClient(client *http.Client) *ReadAppsV1NamespacedControllerRevisionParams {
	var ()
	return &ReadAppsV1NamespacedControllerRevisionParams{
		HTTPClient: client,
	}
}

/*ReadAppsV1NamespacedControllerRevisionParams contains all the parameters to send to the API endpoint
for the read apps v1 namespaced controller revision operation typically these are written to a http.Request
*/
type ReadAppsV1NamespacedControllerRevisionParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the ControllerRevision

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

// WithTimeout adds the timeout to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithTimeout(timeout time.Duration) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithContext(ctx context.Context) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithHTTPClient(client *http.Client) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithExact(exact *bool) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithExport(export *bool) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithName(name string) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithNamespace(namespace string) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) WithPretty(pretty *string) *ReadAppsV1NamespacedControllerRevisionParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read apps v1 namespaced controller revision params
func (o *ReadAppsV1NamespacedControllerRevisionParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAppsV1NamespacedControllerRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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