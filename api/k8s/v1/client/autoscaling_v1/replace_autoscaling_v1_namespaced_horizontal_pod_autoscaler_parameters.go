// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v1

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

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams creates a new ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized.
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithTimeout creates a new ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithTimeout(timeout time.Duration) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams{

		timeout: timeout,
	}
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithContext creates a new ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithContext(ctx context.Context) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams{

		Context: ctx,
	}
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithHTTPClient creates a new ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParamsWithHTTPClient(client *http.Client) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	var ()
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams{
		HTTPClient: client,
	}
}

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams contains all the parameters to send to the API endpoint
for the replace autoscaling v1 namespaced horizontal pod autoscaler operation typically these are written to a http.Request
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams struct {

	/*Body*/
	Body *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler
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

// WithTimeout adds the timeout to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithTimeout(timeout time.Duration) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithContext(ctx context.Context) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithHTTPClient(client *http.Client) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithBody(body *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetBody(body *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithDryRun(dryRun *string) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithFieldManager(fieldManager *string) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithName(name string) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithNamespace(namespace string) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WithPretty(pretty *string) *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace autoscaling v1 namespaced horizontal pod autoscaler params
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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