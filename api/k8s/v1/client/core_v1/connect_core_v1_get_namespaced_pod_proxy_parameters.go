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

// NewConnectCoreV1GetNamespacedPodProxyParams creates a new ConnectCoreV1GetNamespacedPodProxyParams object
// with the default values initialized.
func NewConnectCoreV1GetNamespacedPodProxyParams() *ConnectCoreV1GetNamespacedPodProxyParams {
	var ()
	return &ConnectCoreV1GetNamespacedPodProxyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectCoreV1GetNamespacedPodProxyParamsWithTimeout creates a new ConnectCoreV1GetNamespacedPodProxyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectCoreV1GetNamespacedPodProxyParamsWithTimeout(timeout time.Duration) *ConnectCoreV1GetNamespacedPodProxyParams {
	var ()
	return &ConnectCoreV1GetNamespacedPodProxyParams{

		timeout: timeout,
	}
}

// NewConnectCoreV1GetNamespacedPodProxyParamsWithContext creates a new ConnectCoreV1GetNamespacedPodProxyParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectCoreV1GetNamespacedPodProxyParamsWithContext(ctx context.Context) *ConnectCoreV1GetNamespacedPodProxyParams {
	var ()
	return &ConnectCoreV1GetNamespacedPodProxyParams{

		Context: ctx,
	}
}

// NewConnectCoreV1GetNamespacedPodProxyParamsWithHTTPClient creates a new ConnectCoreV1GetNamespacedPodProxyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectCoreV1GetNamespacedPodProxyParamsWithHTTPClient(client *http.Client) *ConnectCoreV1GetNamespacedPodProxyParams {
	var ()
	return &ConnectCoreV1GetNamespacedPodProxyParams{
		HTTPClient: client,
	}
}

/*ConnectCoreV1GetNamespacedPodProxyParams contains all the parameters to send to the API endpoint
for the connect core v1 get namespaced pod proxy operation typically these are written to a http.Request
*/
type ConnectCoreV1GetNamespacedPodProxyParams struct {

	/*Name
	  name of the PodProxyOptions

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*Path
	  Path is the URL path to use for the current proxy request to pod.

	*/
	Path *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WithTimeout(timeout time.Duration) *ConnectCoreV1GetNamespacedPodProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WithContext(ctx context.Context) *ConnectCoreV1GetNamespacedPodProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WithHTTPClient(client *http.Client) *ConnectCoreV1GetNamespacedPodProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WithName(name string) *ConnectCoreV1GetNamespacedPodProxyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WithNamespace(namespace string) *ConnectCoreV1GetNamespacedPodProxyParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPath adds the path to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WithPath(path *string) *ConnectCoreV1GetNamespacedPodProxyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the connect core v1 get namespaced pod proxy params
func (o *ConnectCoreV1GetNamespacedPodProxyParams) SetPath(path *string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectCoreV1GetNamespacedPodProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Path != nil {

		// query param path
		var qrPath string
		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {
			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
