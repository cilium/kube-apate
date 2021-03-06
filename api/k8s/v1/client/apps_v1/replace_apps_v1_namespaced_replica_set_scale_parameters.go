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

// NewReplaceAppsV1NamespacedReplicaSetScaleParams creates a new ReplaceAppsV1NamespacedReplicaSetScaleParams object
// with the default values initialized.
func NewReplaceAppsV1NamespacedReplicaSetScaleParams() *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	var ()
	return &ReplaceAppsV1NamespacedReplicaSetScaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAppsV1NamespacedReplicaSetScaleParamsWithTimeout creates a new ReplaceAppsV1NamespacedReplicaSetScaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceAppsV1NamespacedReplicaSetScaleParamsWithTimeout(timeout time.Duration) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	var ()
	return &ReplaceAppsV1NamespacedReplicaSetScaleParams{

		timeout: timeout,
	}
}

// NewReplaceAppsV1NamespacedReplicaSetScaleParamsWithContext creates a new ReplaceAppsV1NamespacedReplicaSetScaleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceAppsV1NamespacedReplicaSetScaleParamsWithContext(ctx context.Context) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	var ()
	return &ReplaceAppsV1NamespacedReplicaSetScaleParams{

		Context: ctx,
	}
}

// NewReplaceAppsV1NamespacedReplicaSetScaleParamsWithHTTPClient creates a new ReplaceAppsV1NamespacedReplicaSetScaleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceAppsV1NamespacedReplicaSetScaleParamsWithHTTPClient(client *http.Client) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	var ()
	return &ReplaceAppsV1NamespacedReplicaSetScaleParams{
		HTTPClient: client,
	}
}

/*ReplaceAppsV1NamespacedReplicaSetScaleParams contains all the parameters to send to the API endpoint
for the replace apps v1 namespaced replica set scale operation typically these are written to a http.Request
*/
type ReplaceAppsV1NamespacedReplicaSetScaleParams struct {

	/*Body*/
	Body *models.IoK8sAPIAutoscalingV1Scale
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the Scale

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

// WithTimeout adds the timeout to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithTimeout(timeout time.Duration) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithContext(ctx context.Context) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithHTTPClient(client *http.Client) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithBody(body *models.IoK8sAPIAutoscalingV1Scale) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetBody(body *models.IoK8sAPIAutoscalingV1Scale) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithDryRun(dryRun *string) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithFieldManager(fieldManager *string) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithName(name string) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithNamespace(namespace string) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WithPretty(pretty *string) *ReplaceAppsV1NamespacedReplicaSetScaleParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace apps v1 namespaced replica set scale params
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAppsV1NamespacedReplicaSetScaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
