// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

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

// NewReplaceCoreV1NamespacedServiceParams creates a new ReplaceCoreV1NamespacedServiceParams object
// with the default values initialized.
func NewReplaceCoreV1NamespacedServiceParams() *ReplaceCoreV1NamespacedServiceParams {
	var ()
	return &ReplaceCoreV1NamespacedServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceCoreV1NamespacedServiceParamsWithTimeout creates a new ReplaceCoreV1NamespacedServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceCoreV1NamespacedServiceParamsWithTimeout(timeout time.Duration) *ReplaceCoreV1NamespacedServiceParams {
	var ()
	return &ReplaceCoreV1NamespacedServiceParams{

		timeout: timeout,
	}
}

// NewReplaceCoreV1NamespacedServiceParamsWithContext creates a new ReplaceCoreV1NamespacedServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceCoreV1NamespacedServiceParamsWithContext(ctx context.Context) *ReplaceCoreV1NamespacedServiceParams {
	var ()
	return &ReplaceCoreV1NamespacedServiceParams{

		Context: ctx,
	}
}

// NewReplaceCoreV1NamespacedServiceParamsWithHTTPClient creates a new ReplaceCoreV1NamespacedServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceCoreV1NamespacedServiceParamsWithHTTPClient(client *http.Client) *ReplaceCoreV1NamespacedServiceParams {
	var ()
	return &ReplaceCoreV1NamespacedServiceParams{
		HTTPClient: client,
	}
}

/*ReplaceCoreV1NamespacedServiceParams contains all the parameters to send to the API endpoint
for the replace core v1 namespaced service operation typically these are written to a http.Request
*/
type ReplaceCoreV1NamespacedServiceParams struct {

	/*Body*/
	Body *models.IoK8sAPICoreV1Service
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the Service

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

// WithTimeout adds the timeout to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithTimeout(timeout time.Duration) *ReplaceCoreV1NamespacedServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithContext(ctx context.Context) *ReplaceCoreV1NamespacedServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithHTTPClient(client *http.Client) *ReplaceCoreV1NamespacedServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithBody(body *models.IoK8sAPICoreV1Service) *ReplaceCoreV1NamespacedServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetBody(body *models.IoK8sAPICoreV1Service) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithDryRun(dryRun *string) *ReplaceCoreV1NamespacedServiceParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithFieldManager(fieldManager *string) *ReplaceCoreV1NamespacedServiceParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithName(name string) *ReplaceCoreV1NamespacedServiceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithNamespace(namespace string) *ReplaceCoreV1NamespacedServiceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) WithPretty(pretty *string) *ReplaceCoreV1NamespacedServiceParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace core v1 namespaced service params
func (o *ReplaceCoreV1NamespacedServiceParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceCoreV1NamespacedServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
