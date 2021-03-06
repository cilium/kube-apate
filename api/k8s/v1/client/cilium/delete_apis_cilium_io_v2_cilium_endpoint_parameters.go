// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

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

// NewDeleteApisCiliumIoV2CiliumEndpointParams creates a new DeleteApisCiliumIoV2CiliumEndpointParams object
// with the default values initialized.
func NewDeleteApisCiliumIoV2CiliumEndpointParams() *DeleteApisCiliumIoV2CiliumEndpointParams {
	var ()
	return &DeleteApisCiliumIoV2CiliumEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteApisCiliumIoV2CiliumEndpointParamsWithTimeout creates a new DeleteApisCiliumIoV2CiliumEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteApisCiliumIoV2CiliumEndpointParamsWithTimeout(timeout time.Duration) *DeleteApisCiliumIoV2CiliumEndpointParams {
	var ()
	return &DeleteApisCiliumIoV2CiliumEndpointParams{

		timeout: timeout,
	}
}

// NewDeleteApisCiliumIoV2CiliumEndpointParamsWithContext creates a new DeleteApisCiliumIoV2CiliumEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteApisCiliumIoV2CiliumEndpointParamsWithContext(ctx context.Context) *DeleteApisCiliumIoV2CiliumEndpointParams {
	var ()
	return &DeleteApisCiliumIoV2CiliumEndpointParams{

		Context: ctx,
	}
}

// NewDeleteApisCiliumIoV2CiliumEndpointParamsWithHTTPClient creates a new DeleteApisCiliumIoV2CiliumEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteApisCiliumIoV2CiliumEndpointParamsWithHTTPClient(client *http.Client) *DeleteApisCiliumIoV2CiliumEndpointParams {
	var ()
	return &DeleteApisCiliumIoV2CiliumEndpointParams{
		HTTPClient: client,
	}
}

/*DeleteApisCiliumIoV2CiliumEndpointParams contains all the parameters to send to the API endpoint
for the delete apis cilium io v2 cilium endpoint operation typically these are written to a http.Request
*/
type DeleteApisCiliumIoV2CiliumEndpointParams struct {

	/*Name
	  name of the Cilium Endpoint

	*/
	Name string
	/*Namespace
	  namespace of the Cilium Endpoint

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) WithTimeout(timeout time.Duration) *DeleteApisCiliumIoV2CiliumEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) WithContext(ctx context.Context) *DeleteApisCiliumIoV2CiliumEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) WithHTTPClient(client *http.Client) *DeleteApisCiliumIoV2CiliumEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) WithName(name string) *DeleteApisCiliumIoV2CiliumEndpointParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) WithNamespace(namespace string) *DeleteApisCiliumIoV2CiliumEndpointParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete apis cilium io v2 cilium endpoint params
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteApisCiliumIoV2CiliumEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
