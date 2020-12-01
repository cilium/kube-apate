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

// NewCreateApiextensionsV1beta1CustomResourceDefinitionParams creates a new CreateApiextensionsV1beta1CustomResourceDefinitionParams object
// with the default values initialized.
func NewCreateApiextensionsV1beta1CustomResourceDefinitionParams() *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	var ()
	return &CreateApiextensionsV1beta1CustomResourceDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApiextensionsV1beta1CustomResourceDefinitionParamsWithTimeout creates a new CreateApiextensionsV1beta1CustomResourceDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateApiextensionsV1beta1CustomResourceDefinitionParamsWithTimeout(timeout time.Duration) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	var ()
	return &CreateApiextensionsV1beta1CustomResourceDefinitionParams{

		timeout: timeout,
	}
}

// NewCreateApiextensionsV1beta1CustomResourceDefinitionParamsWithContext creates a new CreateApiextensionsV1beta1CustomResourceDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateApiextensionsV1beta1CustomResourceDefinitionParamsWithContext(ctx context.Context) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	var ()
	return &CreateApiextensionsV1beta1CustomResourceDefinitionParams{

		Context: ctx,
	}
}

// NewCreateApiextensionsV1beta1CustomResourceDefinitionParamsWithHTTPClient creates a new CreateApiextensionsV1beta1CustomResourceDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateApiextensionsV1beta1CustomResourceDefinitionParamsWithHTTPClient(client *http.Client) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	var ()
	return &CreateApiextensionsV1beta1CustomResourceDefinitionParams{
		HTTPClient: client,
	}
}

/*CreateApiextensionsV1beta1CustomResourceDefinitionParams contains all the parameters to send to the API endpoint
for the create apiextensions v1beta1 custom resource definition operation typically these are written to a http.Request
*/
type CreateApiextensionsV1beta1CustomResourceDefinitionParams struct {

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
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithTimeout(timeout time.Duration) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithContext(ctx context.Context) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithHTTPClient(client *http.Client) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithBody(body *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetBody(body *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithDryRun(dryRun *string) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithFieldManager(fieldManager *string) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithPretty adds the pretty to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WithPretty(pretty *string) *CreateApiextensionsV1beta1CustomResourceDefinitionParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create apiextensions v1beta1 custom resource definition params
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApiextensionsV1beta1CustomResourceDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
