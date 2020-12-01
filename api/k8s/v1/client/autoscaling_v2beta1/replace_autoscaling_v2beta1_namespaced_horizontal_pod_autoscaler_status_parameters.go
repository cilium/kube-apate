// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta1

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

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams creates a new ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams object
// with the default values initialized.
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams() *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	var ()
	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParamsWithTimeout creates a new ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParamsWithTimeout(timeout time.Duration) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	var ()
	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams{

		timeout: timeout,
	}
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParamsWithContext creates a new ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParamsWithContext(ctx context.Context) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	var ()
	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams{

		Context: ctx,
	}
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParamsWithHTTPClient creates a new ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParamsWithHTTPClient(client *http.Client) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	var ()
	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams{
		HTTPClient: client,
	}
}

/*ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams contains all the parameters to send to the API endpoint
for the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status operation typically these are written to a http.Request
*/
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams struct {

	/*Body*/
	Body *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the HorizontalPodAutoscaler

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

// WithTimeout adds the timeout to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithTimeout(timeout time.Duration) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithContext(ctx context.Context) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithHTTPClient(client *http.Client) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithBody(body *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetBody(body *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithDryRun(dryRun *string) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithFieldManager(fieldManager *string) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithName(name string) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithNamespace(namespace string) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WithPretty(pretty *string) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
