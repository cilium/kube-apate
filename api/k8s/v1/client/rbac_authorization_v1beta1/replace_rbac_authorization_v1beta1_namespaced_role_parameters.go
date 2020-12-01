// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1beta1

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

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleParams creates a new ReplaceRbacAuthorizationV1beta1NamespacedRoleParams object
// with the default values initialized.
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleParams() *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	var ()
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleParamsWithTimeout creates a new ReplaceRbacAuthorizationV1beta1NamespacedRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleParamsWithTimeout(timeout time.Duration) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	var ()
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleParams{

		timeout: timeout,
	}
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleParamsWithContext creates a new ReplaceRbacAuthorizationV1beta1NamespacedRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleParamsWithContext(ctx context.Context) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	var ()
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleParams{

		Context: ctx,
	}
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleParamsWithHTTPClient creates a new ReplaceRbacAuthorizationV1beta1NamespacedRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleParamsWithHTTPClient(client *http.Client) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	var ()
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleParams{
		HTTPClient: client,
	}
}

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleParams contains all the parameters to send to the API endpoint
for the replace rbac authorization v1beta1 namespaced role operation typically these are written to a http.Request
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleParams struct {

	/*Body*/
	Body *models.IoK8sAPIRbacV1beta1Role
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the Role

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

// WithTimeout adds the timeout to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithTimeout(timeout time.Duration) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithContext(ctx context.Context) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithHTTPClient(client *http.Client) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithBody(body *models.IoK8sAPIRbacV1beta1Role) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetBody(body *models.IoK8sAPIRbacV1beta1Role) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithDryRun(dryRun *string) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithFieldManager(fieldManager *string) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithName(name string) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithNamespace(namespace string) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WithPretty(pretty *string) *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace rbac authorization v1beta1 namespaced role params
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.DryRun != nil {

		// query param dryRun
		var qrDryRun string
		if o.DryRun != nil {
			qrDryRun = *o.DryRun
		}
		qDryRun := qrDryRun
		if qDryRun != "" {
			if err := r.SetQueryParam("dryRun", qDryRun); err != nil {
				return err
			}
		}

	}

	if o.FieldManager != nil {

		// query param fieldManager
		var qrFieldManager string
		if o.FieldManager != nil {
			qrFieldManager = *o.FieldManager
		}
		qFieldManager := qrFieldManager
		if qFieldManager != "" {
			if err := r.SetQueryParam("fieldManager", qFieldManager); err != nil {
				return err
			}
		}

	}

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
