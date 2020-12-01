// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1beta1

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

// NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams creates a new ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams object
// with the default values initialized.
func NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams() *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	var ()
	return &ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParamsWithTimeout creates a new ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParamsWithTimeout(timeout time.Duration) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	var ()
	return &ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams{

		timeout: timeout,
	}
}

// NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParamsWithContext creates a new ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParamsWithContext(ctx context.Context) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	var ()
	return &ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams{

		Context: ctx,
	}
}

// NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParamsWithHTTPClient creates a new ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParamsWithHTTPClient(client *http.Client) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	var ()
	return &ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams{
		HTTPClient: client,
	}
}

/*ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams contains all the parameters to send to the API endpoint
for the replace apiextensions v1beta1 custom resource definition status operation typically these are written to a http.Request
*/
type ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams struct {

	/*Body*/
	Body *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the CustomResourceDefinition

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

// WithTimeout adds the timeout to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithTimeout(timeout time.Duration) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithContext(ctx context.Context) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithHTTPClient(client *http.Client) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithBody(body *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetBody(body *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithDryRun(dryRun *string) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithFieldManager(fieldManager *string) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithName(name string) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WithPretty(pretty *string) *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace apiextensions v1beta1 custom resource definition status params
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
