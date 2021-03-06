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

// NewConnectCoreV1DeleteNodeProxyWithPathParams creates a new ConnectCoreV1DeleteNodeProxyWithPathParams object
// with the default values initialized.
func NewConnectCoreV1DeleteNodeProxyWithPathParams() *ConnectCoreV1DeleteNodeProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNodeProxyWithPathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectCoreV1DeleteNodeProxyWithPathParamsWithTimeout creates a new ConnectCoreV1DeleteNodeProxyWithPathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectCoreV1DeleteNodeProxyWithPathParamsWithTimeout(timeout time.Duration) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNodeProxyWithPathParams{

		timeout: timeout,
	}
}

// NewConnectCoreV1DeleteNodeProxyWithPathParamsWithContext creates a new ConnectCoreV1DeleteNodeProxyWithPathParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectCoreV1DeleteNodeProxyWithPathParamsWithContext(ctx context.Context) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNodeProxyWithPathParams{

		Context: ctx,
	}
}

// NewConnectCoreV1DeleteNodeProxyWithPathParamsWithHTTPClient creates a new ConnectCoreV1DeleteNodeProxyWithPathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectCoreV1DeleteNodeProxyWithPathParamsWithHTTPClient(client *http.Client) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	var ()
	return &ConnectCoreV1DeleteNodeProxyWithPathParams{
		HTTPClient: client,
	}
}

/*ConnectCoreV1DeleteNodeProxyWithPathParams contains all the parameters to send to the API endpoint
for the connect core v1 delete node proxy with path operation typically these are written to a http.Request
*/
type ConnectCoreV1DeleteNodeProxyWithPathParams struct {

	/*Name
	  name of the NodeProxyOptions

	*/
	Name string
	/*Path
	  path to the resource

	*/
	PathPath string
	/*Path
	  Path is the URL path to use for the current proxy request to node.

	*/
	QueryPath *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WithTimeout(timeout time.Duration) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WithContext(ctx context.Context) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WithHTTPClient(client *http.Client) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WithName(name string) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) SetName(name string) {
	o.Name = name
}

// WithPathPath adds the path to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WithPathPath(path string) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	o.SetPathPath(path)
	return o
}

// SetPathPath adds the path to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) SetPathPath(path string) {
	o.PathPath = path
}

// WithQueryPath adds the path to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WithQueryPath(path *string) *ConnectCoreV1DeleteNodeProxyWithPathParams {
	o.SetQueryPath(path)
	return o
}

// SetQueryPath adds the path to the connect core v1 delete node proxy with path params
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) SetQueryPath(path *string) {
	o.QueryPath = path
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectCoreV1DeleteNodeProxyWithPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
