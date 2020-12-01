// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1beta1

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

// NewReadBatchV1beta1NamespacedCronJobParams creates a new ReadBatchV1beta1NamespacedCronJobParams object
// with the default values initialized.
func NewReadBatchV1beta1NamespacedCronJobParams() *ReadBatchV1beta1NamespacedCronJobParams {
	var ()
	return &ReadBatchV1beta1NamespacedCronJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadBatchV1beta1NamespacedCronJobParamsWithTimeout creates a new ReadBatchV1beta1NamespacedCronJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadBatchV1beta1NamespacedCronJobParamsWithTimeout(timeout time.Duration) *ReadBatchV1beta1NamespacedCronJobParams {
	var ()
	return &ReadBatchV1beta1NamespacedCronJobParams{

		timeout: timeout,
	}
}

// NewReadBatchV1beta1NamespacedCronJobParamsWithContext creates a new ReadBatchV1beta1NamespacedCronJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadBatchV1beta1NamespacedCronJobParamsWithContext(ctx context.Context) *ReadBatchV1beta1NamespacedCronJobParams {
	var ()
	return &ReadBatchV1beta1NamespacedCronJobParams{

		Context: ctx,
	}
}

// NewReadBatchV1beta1NamespacedCronJobParamsWithHTTPClient creates a new ReadBatchV1beta1NamespacedCronJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadBatchV1beta1NamespacedCronJobParamsWithHTTPClient(client *http.Client) *ReadBatchV1beta1NamespacedCronJobParams {
	var ()
	return &ReadBatchV1beta1NamespacedCronJobParams{
		HTTPClient: client,
	}
}

/*ReadBatchV1beta1NamespacedCronJobParams contains all the parameters to send to the API endpoint
for the read batch v1beta1 namespaced cron job operation typically these are written to a http.Request
*/
type ReadBatchV1beta1NamespacedCronJobParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the CronJob

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

// WithTimeout adds the timeout to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithTimeout(timeout time.Duration) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithContext(ctx context.Context) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithHTTPClient(client *http.Client) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithExact(exact *bool) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithExport(export *bool) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithName(name string) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithNamespace(namespace string) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) WithPretty(pretty *string) *ReadBatchV1beta1NamespacedCronJobParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read batch v1beta1 namespaced cron job params
func (o *ReadBatchV1beta1NamespacedCronJobParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadBatchV1beta1NamespacedCronJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
