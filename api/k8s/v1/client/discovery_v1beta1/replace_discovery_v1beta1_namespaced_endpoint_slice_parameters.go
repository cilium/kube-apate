// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package discovery_v1beta1

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

// NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParams creates a new ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams object
// with the default values initialized.
func NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParams() *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	var ()
	return &ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParamsWithTimeout creates a new ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParamsWithTimeout(timeout time.Duration) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	var ()
	return &ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams{

		timeout: timeout,
	}
}

// NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParamsWithContext creates a new ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParamsWithContext(ctx context.Context) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	var ()
	return &ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams{

		Context: ctx,
	}
}

// NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParamsWithHTTPClient creates a new ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceDiscoveryV1beta1NamespacedEndpointSliceParamsWithHTTPClient(client *http.Client) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	var ()
	return &ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams{
		HTTPClient: client,
	}
}

/*ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams contains all the parameters to send to the API endpoint
for the replace discovery v1beta1 namespaced endpoint slice operation typically these are written to a http.Request
*/
type ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams struct {

	/*Body*/
	Body *models.IoK8sAPIDiscoveryV1beta1EndpointSlice
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the EndpointSlice

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

// WithTimeout adds the timeout to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithTimeout(timeout time.Duration) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithContext(ctx context.Context) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithHTTPClient(client *http.Client) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithBody(body *models.IoK8sAPIDiscoveryV1beta1EndpointSlice) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetBody(body *models.IoK8sAPIDiscoveryV1beta1EndpointSlice) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithDryRun(dryRun *string) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithFieldManager(fieldManager *string) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithName(name string) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithNamespace(namespace string) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WithPretty(pretty *string) *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace discovery v1beta1 namespaced endpoint slice params
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
