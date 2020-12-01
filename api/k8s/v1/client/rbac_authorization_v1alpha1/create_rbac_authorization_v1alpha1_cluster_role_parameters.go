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

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// NewCreateRbacAuthorizationV1alpha1ClusterRoleParams creates a new CreateRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized.
func NewCreateRbacAuthorizationV1alpha1ClusterRoleParams() *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &CreateRbacAuthorizationV1alpha1ClusterRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleParamsWithTimeout creates a new CreateRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRbacAuthorizationV1alpha1ClusterRoleParamsWithTimeout(timeout time.Duration) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &CreateRbacAuthorizationV1alpha1ClusterRoleParams{

		timeout: timeout,
	}
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleParamsWithContext creates a new CreateRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRbacAuthorizationV1alpha1ClusterRoleParamsWithContext(ctx context.Context) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &CreateRbacAuthorizationV1alpha1ClusterRoleParams{

		Context: ctx,
	}
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleParamsWithHTTPClient creates a new CreateRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRbacAuthorizationV1alpha1ClusterRoleParamsWithHTTPClient(client *http.Client) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &CreateRbacAuthorizationV1alpha1ClusterRoleParams{
		HTTPClient: client,
	}
}

/*CreateRbacAuthorizationV1alpha1ClusterRoleParams contains all the parameters to send to the API endpoint
for the create rbac authorization v1alpha1 cluster role operation typically these are written to a http.Request
*/
type CreateRbacAuthorizationV1alpha1ClusterRoleParams struct {

	/*Body*/
	Body *models.IoK8sAPIRbacV1alpha1ClusterRole
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithTimeout(timeout time.Duration) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithContext(ctx context.Context) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithHTTPClient(client *http.Client) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithBody(body *models.IoK8sAPIRbacV1alpha1ClusterRole) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetBody(body *models.IoK8sAPIRbacV1alpha1ClusterRole) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithDryRun(dryRun *string) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithFieldManager(fieldManager *string) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithPretty adds the pretty to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WithPretty(pretty *string) *CreateRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create rbac authorization v1alpha1 cluster role params
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
