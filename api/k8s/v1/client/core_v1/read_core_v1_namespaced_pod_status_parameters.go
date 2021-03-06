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
)

// NewReadCoreV1NamespacedPodStatusParams creates a new ReadCoreV1NamespacedPodStatusParams object
// with the default values initialized.
func NewReadCoreV1NamespacedPodStatusParams() *ReadCoreV1NamespacedPodStatusParams {
	var ()
	return &ReadCoreV1NamespacedPodStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadCoreV1NamespacedPodStatusParamsWithTimeout creates a new ReadCoreV1NamespacedPodStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadCoreV1NamespacedPodStatusParamsWithTimeout(timeout time.Duration) *ReadCoreV1NamespacedPodStatusParams {
	var ()
	return &ReadCoreV1NamespacedPodStatusParams{

		timeout: timeout,
	}
}

// NewReadCoreV1NamespacedPodStatusParamsWithContext creates a new ReadCoreV1NamespacedPodStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadCoreV1NamespacedPodStatusParamsWithContext(ctx context.Context) *ReadCoreV1NamespacedPodStatusParams {
	var ()
	return &ReadCoreV1NamespacedPodStatusParams{

		Context: ctx,
	}
}

// NewReadCoreV1NamespacedPodStatusParamsWithHTTPClient creates a new ReadCoreV1NamespacedPodStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadCoreV1NamespacedPodStatusParamsWithHTTPClient(client *http.Client) *ReadCoreV1NamespacedPodStatusParams {
	var ()
	return &ReadCoreV1NamespacedPodStatusParams{
		HTTPClient: client,
	}
}

/*ReadCoreV1NamespacedPodStatusParams contains all the parameters to send to the API endpoint
for the read core v1 namespaced pod status operation typically these are written to a http.Request
*/
type ReadCoreV1NamespacedPodStatusParams struct {

	/*Name
	  name of the Pod

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

// WithTimeout adds the timeout to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) WithTimeout(timeout time.Duration) *ReadCoreV1NamespacedPodStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) WithContext(ctx context.Context) *ReadCoreV1NamespacedPodStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) WithHTTPClient(client *http.Client) *ReadCoreV1NamespacedPodStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) WithName(name string) *ReadCoreV1NamespacedPodStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) WithNamespace(namespace string) *ReadCoreV1NamespacedPodStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) WithPretty(pretty *string) *ReadCoreV1NamespacedPodStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read core v1 namespaced pod status params
func (o *ReadCoreV1NamespacedPodStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadCoreV1NamespacedPodStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
