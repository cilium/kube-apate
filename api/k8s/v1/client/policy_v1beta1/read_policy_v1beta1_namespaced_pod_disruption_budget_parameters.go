// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy_v1beta1

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

// NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParams creates a new ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams object
// with the default values initialized.
func NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParams() *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	var ()
	return &ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParamsWithTimeout creates a new ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParamsWithTimeout(timeout time.Duration) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	var ()
	return &ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams{

		timeout: timeout,
	}
}

// NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParamsWithContext creates a new ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParamsWithContext(ctx context.Context) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	var ()
	return &ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams{

		Context: ctx,
	}
}

// NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParamsWithHTTPClient creates a new ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadPolicyV1beta1NamespacedPodDisruptionBudgetParamsWithHTTPClient(client *http.Client) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	var ()
	return &ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams{
		HTTPClient: client,
	}
}

/*ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams contains all the parameters to send to the API endpoint
for the read policy v1beta1 namespaced pod disruption budget operation typically these are written to a http.Request
*/
type ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams struct {

	/*Exact
	  Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.

	*/
	Exact *bool
	/*Export
	  Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.

	*/
	Export *bool
	/*Name
	  name of the PodDisruptionBudget

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

// WithTimeout adds the timeout to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithTimeout(timeout time.Duration) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithContext(ctx context.Context) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithHTTPClient(client *http.Client) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExact adds the exact to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithExact(exact *bool) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetExact(exact)
	return o
}

// SetExact adds the exact to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetExact(exact *bool) {
	o.Exact = exact
}

// WithExport adds the export to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithExport(export *bool) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetExport(export)
	return o
}

// SetExport adds the export to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetExport(export *bool) {
	o.Export = export
}

// WithName adds the name to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithName(name string) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithNamespace(namespace string) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WithPretty(pretty *string) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read policy v1beta1 namespaced pod disruption budget params
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
