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

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// NewCreateAppsV1NamespacedControllerRevisionParams creates a new CreateAppsV1NamespacedControllerRevisionParams object
// with the default values initialized.
func NewCreateAppsV1NamespacedControllerRevisionParams() *CreateAppsV1NamespacedControllerRevisionParams {
	var ()
	return &CreateAppsV1NamespacedControllerRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppsV1NamespacedControllerRevisionParamsWithTimeout creates a new CreateAppsV1NamespacedControllerRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAppsV1NamespacedControllerRevisionParamsWithTimeout(timeout time.Duration) *CreateAppsV1NamespacedControllerRevisionParams {
	var ()
	return &CreateAppsV1NamespacedControllerRevisionParams{

		timeout: timeout,
	}
}

// NewCreateAppsV1NamespacedControllerRevisionParamsWithContext creates a new CreateAppsV1NamespacedControllerRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAppsV1NamespacedControllerRevisionParamsWithContext(ctx context.Context) *CreateAppsV1NamespacedControllerRevisionParams {
	var ()
	return &CreateAppsV1NamespacedControllerRevisionParams{

		Context: ctx,
	}
}

// NewCreateAppsV1NamespacedControllerRevisionParamsWithHTTPClient creates a new CreateAppsV1NamespacedControllerRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAppsV1NamespacedControllerRevisionParamsWithHTTPClient(client *http.Client) *CreateAppsV1NamespacedControllerRevisionParams {
	var ()
	return &CreateAppsV1NamespacedControllerRevisionParams{
		HTTPClient: client,
	}
}

/*CreateAppsV1NamespacedControllerRevisionParams contains all the parameters to send to the API endpoint
for the create apps v1 namespaced controller revision operation typically these are written to a http.Request
*/
type CreateAppsV1NamespacedControllerRevisionParams struct {

	/*Body*/
	Body *models.IoK8sAPIAppsV1ControllerRevision
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
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

// WithTimeout adds the timeout to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithTimeout(timeout time.Duration) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithContext(ctx context.Context) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithHTTPClient(client *http.Client) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithBody(body *models.IoK8sAPIAppsV1ControllerRevision) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetBody(body *models.IoK8sAPIAppsV1ControllerRevision) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithDryRun(dryRun *string) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithFieldManager(fieldManager *string) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithNamespace adds the namespace to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithNamespace(namespace string) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) WithPretty(pretty *string) *CreateAppsV1NamespacedControllerRevisionParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create apps v1 namespaced controller revision params
func (o *CreateAppsV1NamespacedControllerRevisionParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppsV1NamespacedControllerRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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