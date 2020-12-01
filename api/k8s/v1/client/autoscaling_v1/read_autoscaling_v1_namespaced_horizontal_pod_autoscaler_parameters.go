// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v1

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

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams creates a new ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized.
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams() *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithTimeout creates a new ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithTimeout(timeout time.Duration) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams{

		timeout: timeout,
	}
}

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithContext creates a new ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithContext(ctx context.Context) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams{

		Context: ctx,
	}
}

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithHTTPClient creates a new ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithHTTPClient(client *http.Client) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams{
		HTTPClient: client,
	}
}

/*ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams contains all the parameters to send to the API endpoint
for the read autoscaling v1 namespaced horizontal pod autoscaler operation typically these are written to a http.Request
*/
type ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the HorizontalPodAutoscaler

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

// WithTimeout adds the timeout to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithTimeout(timeout time.Duration) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithContext(ctx context.Context) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithHTTPClient(client *http.Client) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithExact(exact *bool) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithExport(export *bool) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithName(name string) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithNamespace(namespace string) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithPretty(pretty *string) *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
