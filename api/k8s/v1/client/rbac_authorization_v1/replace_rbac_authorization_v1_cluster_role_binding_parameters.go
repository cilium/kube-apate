// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1

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

// NewReplaceRbacAuthorizationV1ClusterRoleBindingParams creates a new ReplaceRbacAuthorizationV1ClusterRoleBindingParams object
// with the default values initialized.
func NewReplaceRbacAuthorizationV1ClusterRoleBindingParams() *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	var ()
	return &ReplaceRbacAuthorizationV1ClusterRoleBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRbacAuthorizationV1ClusterRoleBindingParamsWithTimeout creates a new ReplaceRbacAuthorizationV1ClusterRoleBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceRbacAuthorizationV1ClusterRoleBindingParamsWithTimeout(timeout time.Duration) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	var ()
	return &ReplaceRbacAuthorizationV1ClusterRoleBindingParams{

		timeout: timeout,
	}
}

// NewReplaceRbacAuthorizationV1ClusterRoleBindingParamsWithContext creates a new ReplaceRbacAuthorizationV1ClusterRoleBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceRbacAuthorizationV1ClusterRoleBindingParamsWithContext(ctx context.Context) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	var ()
	return &ReplaceRbacAuthorizationV1ClusterRoleBindingParams{

		Context: ctx,
	}
}

// NewReplaceRbacAuthorizationV1ClusterRoleBindingParamsWithHTTPClient creates a new ReplaceRbacAuthorizationV1ClusterRoleBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceRbacAuthorizationV1ClusterRoleBindingParamsWithHTTPClient(client *http.Client) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	var ()
	return &ReplaceRbacAuthorizationV1ClusterRoleBindingParams{
		HTTPClient: client,
	}
}

/*ReplaceRbacAuthorizationV1ClusterRoleBindingParams contains all the parameters to send to the API endpoint
for the replace rbac authorization v1 cluster role binding operation typically these are written to a http.Request
*/
type ReplaceRbacAuthorizationV1ClusterRoleBindingParams struct {

	/*Body*/
	Body *models.IoK8sAPIRbacV1ClusterRoleBinding
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the ClusterRoleBinding

	*/
	Name string
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithTimeout(timeout time.Duration) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithContext(ctx context.Context) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithHTTPClient(client *http.Client) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithBody(body *models.IoK8sAPIRbacV1ClusterRoleBinding) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetBody(body *models.IoK8sAPIRbacV1ClusterRoleBinding) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithDryRun(dryRun *string) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithFieldManager(fieldManager *string) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithName(name string) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WithPretty(pretty *string) *ReplaceRbacAuthorizationV1ClusterRoleBindingParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace rbac authorization v1 cluster role binding params
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
