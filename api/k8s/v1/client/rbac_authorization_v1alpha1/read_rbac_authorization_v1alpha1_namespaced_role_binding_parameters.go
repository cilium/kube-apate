// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1alpha1

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

// NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams creates a new ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams object
// with the default values initialized.
func NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams() *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	var ()
	return &ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParamsWithTimeout creates a new ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParamsWithTimeout(timeout time.Duration) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	var ()
	return &ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams{

		timeout: timeout,
	}
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParamsWithContext creates a new ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParamsWithContext(ctx context.Context) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	var ()
	return &ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams{

		Context: ctx,
	}
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParamsWithHTTPClient creates a new ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingParamsWithHTTPClient(client *http.Client) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	var ()
	return &ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams{
		HTTPClient: client,
	}
}

/*ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams contains all the parameters to send to the API endpoint
for the read rbac authorization v1alpha1 namespaced role binding operation typically these are written to a http.Request
*/
type ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams struct {

	/*Name
	  name of the RoleBinding

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

// WithTimeout adds the timeout to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WithTimeout(timeout time.Duration) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WithContext(ctx context.Context) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WithHTTPClient(client *http.Client) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WithName(name string) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WithNamespace(namespace string) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WithPretty(pretty *string) *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the read rbac authorization v1alpha1 namespaced role binding params
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
