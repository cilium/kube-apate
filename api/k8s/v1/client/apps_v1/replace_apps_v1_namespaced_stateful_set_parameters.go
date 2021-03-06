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

// NewReplaceAppsV1NamespacedStatefulSetParams creates a new ReplaceAppsV1NamespacedStatefulSetParams object
// with the default values initialized.
func NewReplaceAppsV1NamespacedStatefulSetParams() *ReplaceAppsV1NamespacedStatefulSetParams {
	var ()
	return &ReplaceAppsV1NamespacedStatefulSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAppsV1NamespacedStatefulSetParamsWithTimeout creates a new ReplaceAppsV1NamespacedStatefulSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceAppsV1NamespacedStatefulSetParamsWithTimeout(timeout time.Duration) *ReplaceAppsV1NamespacedStatefulSetParams {
	var ()
	return &ReplaceAppsV1NamespacedStatefulSetParams{

		timeout: timeout,
	}
}

// NewReplaceAppsV1NamespacedStatefulSetParamsWithContext creates a new ReplaceAppsV1NamespacedStatefulSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceAppsV1NamespacedStatefulSetParamsWithContext(ctx context.Context) *ReplaceAppsV1NamespacedStatefulSetParams {
	var ()
	return &ReplaceAppsV1NamespacedStatefulSetParams{

		Context: ctx,
	}
}

// NewReplaceAppsV1NamespacedStatefulSetParamsWithHTTPClient creates a new ReplaceAppsV1NamespacedStatefulSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceAppsV1NamespacedStatefulSetParamsWithHTTPClient(client *http.Client) *ReplaceAppsV1NamespacedStatefulSetParams {
	var ()
	return &ReplaceAppsV1NamespacedStatefulSetParams{
		HTTPClient: client,
	}
}

/*ReplaceAppsV1NamespacedStatefulSetParams contains all the parameters to send to the API endpoint
for the replace apps v1 namespaced stateful set operation typically these are written to a http.Request
*/
type ReplaceAppsV1NamespacedStatefulSetParams struct {

	/*Body*/
	Body *models.IoK8sAPIAppsV1StatefulSet
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the StatefulSet

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

// WithTimeout adds the timeout to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithTimeout(timeout time.Duration) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithContext(ctx context.Context) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithHTTPClient(client *http.Client) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithBody(body *models.IoK8sAPIAppsV1StatefulSet) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetBody(body *models.IoK8sAPIAppsV1StatefulSet) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithDryRun(dryRun *string) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithFieldManager(fieldManager *string) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithName(name string) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithNamespace(namespace string) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WithPretty(pretty *string) *ReplaceAppsV1NamespacedStatefulSetParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace apps v1 namespaced stateful set params
func (o *ReplaceAppsV1NamespacedStatefulSetParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAppsV1NamespacedStatefulSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
