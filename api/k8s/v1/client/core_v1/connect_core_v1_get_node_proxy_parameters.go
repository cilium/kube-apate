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

// NewConnectCoreV1GetNodeProxyParams creates a new ConnectCoreV1GetNodeProxyParams object
// with the default values initialized.
func NewConnectCoreV1GetNodeProxyParams() *ConnectCoreV1GetNodeProxyParams {
	var ()
	return &ConnectCoreV1GetNodeProxyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectCoreV1GetNodeProxyParamsWithTimeout creates a new ConnectCoreV1GetNodeProxyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectCoreV1GetNodeProxyParamsWithTimeout(timeout time.Duration) *ConnectCoreV1GetNodeProxyParams {
	var ()
	return &ConnectCoreV1GetNodeProxyParams{

		timeout: timeout,
	}
}

// NewConnectCoreV1GetNodeProxyParamsWithContext creates a new ConnectCoreV1GetNodeProxyParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectCoreV1GetNodeProxyParamsWithContext(ctx context.Context) *ConnectCoreV1GetNodeProxyParams {
	var ()
	return &ConnectCoreV1GetNodeProxyParams{

		Context: ctx,
	}
}

// NewConnectCoreV1GetNodeProxyParamsWithHTTPClient creates a new ConnectCoreV1GetNodeProxyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectCoreV1GetNodeProxyParamsWithHTTPClient(client *http.Client) *ConnectCoreV1GetNodeProxyParams {
	var ()
	return &ConnectCoreV1GetNodeProxyParams{
		HTTPClient: client,
	}
}

/*ConnectCoreV1GetNodeProxyParams contains all the parameters to send to the API endpoint
for the connect core v1 get node proxy operation typically these are written to a http.Request
*/
type ConnectCoreV1GetNodeProxyParams struct {

	/*Name
	  name of the NodeProxyOptions

	*/
	Name string
	/*Path
	  Path is the URL path to use for the current proxy request to node.

	*/
	Path *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) WithTimeout(timeout time.Duration) *ConnectCoreV1GetNodeProxyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) WithContext(ctx context.Context) *ConnectCoreV1GetNodeProxyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) WithHTTPClient(client *http.Client) *ConnectCoreV1GetNodeProxyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) WithName(name string) *ConnectCoreV1GetNodeProxyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) SetName(name string) {
	o.Name = name
}

// WithPath adds the path to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) WithPath(path *string) *ConnectCoreV1GetNodeProxyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the connect core v1 get node proxy params
func (o *ConnectCoreV1GetNodeProxyParams) SetPath(path *string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectCoreV1GetNodeProxyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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