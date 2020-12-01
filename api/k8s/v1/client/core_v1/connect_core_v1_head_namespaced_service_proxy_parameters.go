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

// NewConnectCoreV1HeadNamespacedServiceProxyParams creates a new ConnectCoreV1HeadNamespacedServiceProxyParams object
// with the default values initialized.
func NewConnectCoreV1HeadNamespacedServiceProxyParams() *ConnectCoreV1HeadNamespacedServiceProxyParams {
	var ()
	return &ConnectCoreV1HeadNamespacedServiceProxyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectCoreV1HeadNamespacedServiceProxyParamsWithTimeout creates a new ConnectCoreV1HeadNamespacedServiceProxyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectCoreV1HeadNamespacedServiceProxyParamsWithTimeout(timeout time.Duration) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	var ()
	return &ConnectCoreV1HeadNamespacedServiceProxyParams{

		timeout: timeout,
	}
}

// NewConnectCoreV1HeadNamespacedServiceProxyParamsWithContext creates a new ConnectCoreV1HeadNamespacedServiceProxyParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectCoreV1HeadNamespacedServiceProxyParamsWithContext(ctx context.Context) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	var ()
	return &ConnectCoreV1HeadNamespacedServiceProxyParams{

		Context: ctx,
	}
}

// NewConnectCoreV1HeadNamespacedServiceProxyParamsWithHTTPClient creates a new ConnectCoreV1HeadNamespacedServiceProxyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectCoreV1HeadNamespacedServiceProxyParamsWithHTTPClient(client *http.Client) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	var ()
	return &ConnectCoreV1HeadNamespacedServiceProxyParams{
		HTTPClient: client,
	}
}

/*ConnectCoreV1HeadNamespacedServiceProxyParams contains all the parameters to send to the API endpoint
for the connect core v1 head namespaced service proxy operation typically these are written to a http.Request
*/
type ConnectCoreV1HeadNamespacedServiceProxyParams struct {

	/*Name
	  name of the ServiceProxyOptions

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*Path
	  Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.

	*/
	Path *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WithTimeout(timeout time.Duration) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WithContext(ctx context.Context) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WithHTTPClient(client *http.Client) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WithName(name string) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WithNamespace(namespace string) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPath adds the path to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WithPath(path *string) *ConnectCoreV1HeadNamespacedServiceProxyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the connect core v1 head namespaced service proxy params
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) SetPath(path *string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectCoreV1HeadNamespacedServiceProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
