// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

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

// NewReadAppsV1NamespacedDeploymentStatusParams creates a new ReadAppsV1NamespacedDeploymentStatusParams object
// with the default values initialized.
func NewReadAppsV1NamespacedDeploymentStatusParams() *ReadAppsV1NamespacedDeploymentStatusParams {
	var ()
	return &ReadAppsV1NamespacedDeploymentStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadAppsV1NamespacedDeploymentStatusParamsWithTimeout creates a new ReadAppsV1NamespacedDeploymentStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadAppsV1NamespacedDeploymentStatusParamsWithTimeout(timeout time.Duration) *ReadAppsV1NamespacedDeploymentStatusParams {
	var ()
	return &ReadAppsV1NamespacedDeploymentStatusParams{

		timeout: timeout,
	}
}

// NewReadAppsV1NamespacedDeploymentStatusParamsWithContext creates a new ReadAppsV1NamespacedDeploymentStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadAppsV1NamespacedDeploymentStatusParamsWithContext(ctx context.Context) *ReadAppsV1NamespacedDeploymentStatusParams {
	var ()
	return &ReadAppsV1NamespacedDeploymentStatusParams{

		Context: ctx,
	}
}

// NewReadAppsV1NamespacedDeploymentStatusParamsWithHTTPClient creates a new ReadAppsV1NamespacedDeploymentStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadAppsV1NamespacedDeploymentStatusParamsWithHTTPClient(client *http.Client) *ReadAppsV1NamespacedDeploymentStatusParams {
	var ()
	return &ReadAppsV1NamespacedDeploymentStatusParams{
		HTTPClient: client,
	}
}

/*ReadAppsV1NamespacedDeploymentStatusParams contains all the parameters to send to the API endpoint
for the read apps v1 namespaced deployment status operation typically these are written to a http.Request
*/
type ReadAppsV1NamespacedDeploymentStatusParams struct {

	/*Name
	  name of the Deployment

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

// WithTimeout adds the timeout to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WithTimeout(timeout time.Duration) *ReadAppsV1NamespacedDeploymentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WithContext(ctx context.Context) *ReadAppsV1NamespacedDeploymentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WithHTTPClient(client *http.Client) *ReadAppsV1NamespacedDeploymentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WithName(name string) *ReadAppsV1NamespacedDeploymentStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WithNamespace(namespace string) *ReadAppsV1NamespacedDeploymentStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WithPretty(pretty *string) *ReadAppsV1NamespacedDeploymentStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read apps v1 namespaced deployment status params
func (o *ReadAppsV1NamespacedDeploymentStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadAppsV1NamespacedDeploymentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
