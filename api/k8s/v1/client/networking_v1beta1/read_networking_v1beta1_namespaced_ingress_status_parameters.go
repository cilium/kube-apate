// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

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

// NewReadNetworkingV1beta1NamespacedIngressStatusParams creates a new ReadNetworkingV1beta1NamespacedIngressStatusParams object
// with the default values initialized.
func NewReadNetworkingV1beta1NamespacedIngressStatusParams() *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReadNetworkingV1beta1NamespacedIngressStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadNetworkingV1beta1NamespacedIngressStatusParamsWithTimeout creates a new ReadNetworkingV1beta1NamespacedIngressStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadNetworkingV1beta1NamespacedIngressStatusParamsWithTimeout(timeout time.Duration) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReadNetworkingV1beta1NamespacedIngressStatusParams{

		timeout: timeout,
	}
}

// NewReadNetworkingV1beta1NamespacedIngressStatusParamsWithContext creates a new ReadNetworkingV1beta1NamespacedIngressStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadNetworkingV1beta1NamespacedIngressStatusParamsWithContext(ctx context.Context) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReadNetworkingV1beta1NamespacedIngressStatusParams{

		Context: ctx,
	}
}

// NewReadNetworkingV1beta1NamespacedIngressStatusParamsWithHTTPClient creates a new ReadNetworkingV1beta1NamespacedIngressStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadNetworkingV1beta1NamespacedIngressStatusParamsWithHTTPClient(client *http.Client) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReadNetworkingV1beta1NamespacedIngressStatusParams{
		HTTPClient: client,
	}
}

/*ReadNetworkingV1beta1NamespacedIngressStatusParams contains all the parameters to send to the API endpoint
for the read networking v1beta1 namespaced ingress status operation typically these are written to a http.Request
*/
type ReadNetworkingV1beta1NamespacedIngressStatusParams struct {

	/*Name
	  name of the Ingress

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

// WithTimeout adds the timeout to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WithTimeout(timeout time.Duration) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WithContext(ctx context.Context) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WithHTTPClient(client *http.Client) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WithName(name string) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WithNamespace(namespace string) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WithPretty(pretty *string) *ReadNetworkingV1beta1NamespacedIngressStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read networking v1beta1 namespaced ingress status params
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadNetworkingV1beta1NamespacedIngressStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
