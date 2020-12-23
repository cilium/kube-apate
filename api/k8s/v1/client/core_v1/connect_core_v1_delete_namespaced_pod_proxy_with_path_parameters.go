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

// NewConnectCoreV1DeleteNamespacedPodProxyWithPathParams creates a new ConnectCoreV1DeleteNamespacedPodProxyWithPathParams object
// with the default values initialized.
func NewConnectCoreV1DeleteNamespacedPodProxyWithPathParams() *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNamespacedPodProxyWithPathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectCoreV1DeleteNamespacedPodProxyWithPathParamsWithTimeout creates a new ConnectCoreV1DeleteNamespacedPodProxyWithPathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectCoreV1DeleteNamespacedPodProxyWithPathParamsWithTimeout(timeout time.Duration) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNamespacedPodProxyWithPathParams{

		timeout: timeout,
	}
}

// NewConnectCoreV1DeleteNamespacedPodProxyWithPathParamsWithContext creates a new ConnectCoreV1DeleteNamespacedPodProxyWithPathParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectCoreV1DeleteNamespacedPodProxyWithPathParamsWithContext(ctx context.Context) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNamespacedPodProxyWithPathParams{

		Context: ctx,
	}
}

// NewConnectCoreV1DeleteNamespacedPodProxyWithPathParamsWithHTTPClient creates a new ConnectCoreV1DeleteNamespacedPodProxyWithPathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectCoreV1DeleteNamespacedPodProxyWithPathParamsWithHTTPClient(client *http.Client) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNamespacedPodProxyWithPathParams{
		HTTPClient: client,
	}
}

/*ConnectCoreV1DeleteNamespacedPodProxyWithPathParams contains all the parameters to send to the API endpoint
for the connect core v1 delete namespaced pod proxy with path operation typically these are written to a http.Request
*/
type ConnectCoreV1DeleteNamespacedPodProxyWithPathParams struct {

	/*Name
	  name of the PodProxyOptions

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*Path
	  path to the resource

	*/
	PathPath string
	/*Path
	  Path is the URL path to use for the current proxy request to pod.

	*/
	QueryPath *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithTimeout(timeout time.Duration) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithContext(ctx context.Context) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithHTTPClient(client *http.Client) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithName(name string) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithNamespace(namespace string) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPathPath adds the path to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithPathPath(path string) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetPathPath(path)
	return o
}

// SetPathPath adds the path to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetPathPath(path string) {
	o.PathPath = path
}

// WithQueryPath adds the path to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WithQueryPath(path *string) *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams {
	o.SetQueryPath(path)
	return o
}

// SetQueryPath adds the path to the connect core v1 delete namespaced pod proxy with path params
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) SetQueryPath(path *string) {
	o.QueryPath = path
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param path
	if err := r.SetPathParam("path", o.PathPath); err != nil {
		return err
	}

	if o.QueryPath != nil {

		// query param path
		var qrPath string
		if o.QueryPath != nil {
			qrPath = *o.QueryPath
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