// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

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

// NewReplaceExtensionsV1beta1NamespacedIngressStatusParams creates a new ReplaceExtensionsV1beta1NamespacedIngressStatusParams object
// with the default values initialized.
func NewReplaceExtensionsV1beta1NamespacedIngressStatusParams() *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReplaceExtensionsV1beta1NamespacedIngressStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatusParamsWithTimeout creates a new ReplaceExtensionsV1beta1NamespacedIngressStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceExtensionsV1beta1NamespacedIngressStatusParamsWithTimeout(timeout time.Duration) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReplaceExtensionsV1beta1NamespacedIngressStatusParams{

		timeout: timeout,
	}
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatusParamsWithContext creates a new ReplaceExtensionsV1beta1NamespacedIngressStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceExtensionsV1beta1NamespacedIngressStatusParamsWithContext(ctx context.Context) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReplaceExtensionsV1beta1NamespacedIngressStatusParams{

		Context: ctx,
	}
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatusParamsWithHTTPClient creates a new ReplaceExtensionsV1beta1NamespacedIngressStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceExtensionsV1beta1NamespacedIngressStatusParamsWithHTTPClient(client *http.Client) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	var ()
	return &ReplaceExtensionsV1beta1NamespacedIngressStatusParams{
		HTTPClient: client,
	}
}

/*ReplaceExtensionsV1beta1NamespacedIngressStatusParams contains all the parameters to send to the API endpoint
for the replace extensions v1beta1 namespaced ingress status operation typically these are written to a http.Request
*/
type ReplaceExtensionsV1beta1NamespacedIngressStatusParams struct {

	/*Body*/
	Body *models.IoK8sAPIExtensionsV1beta1Ingress
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the Ingress

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

// WithTimeout adds the timeout to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithTimeout(timeout time.Duration) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithContext(ctx context.Context) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithHTTPClient(client *http.Client) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithBody(body *models.IoK8sAPIExtensionsV1beta1Ingress) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetBody(body *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithDryRun(dryRun *string) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithFieldManager(fieldManager *string) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithName(name string) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithNamespace(namespace string) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WithPretty(pretty *string) *ReplaceExtensionsV1beta1NamespacedIngressStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace extensions v1beta1 namespaced ingress status params
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
