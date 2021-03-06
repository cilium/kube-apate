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

// NewReplaceCoreV1NamespaceStatusParams creates a new ReplaceCoreV1NamespaceStatusParams object
// with the default values initialized.
func NewReplaceCoreV1NamespaceStatusParams() *ReplaceCoreV1NamespaceStatusParams {
	var ()
	return &ReplaceCoreV1NamespaceStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceCoreV1NamespaceStatusParamsWithTimeout creates a new ReplaceCoreV1NamespaceStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceCoreV1NamespaceStatusParamsWithTimeout(timeout time.Duration) *ReplaceCoreV1NamespaceStatusParams {
	var ()
	return &ReplaceCoreV1NamespaceStatusParams{

		timeout: timeout,
	}
}

// NewReplaceCoreV1NamespaceStatusParamsWithContext creates a new ReplaceCoreV1NamespaceStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceCoreV1NamespaceStatusParamsWithContext(ctx context.Context) *ReplaceCoreV1NamespaceStatusParams {
	var ()
	return &ReplaceCoreV1NamespaceStatusParams{

		Context: ctx,
	}
}

// NewReplaceCoreV1NamespaceStatusParamsWithHTTPClient creates a new ReplaceCoreV1NamespaceStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceCoreV1NamespaceStatusParamsWithHTTPClient(client *http.Client) *ReplaceCoreV1NamespaceStatusParams {
	var ()
	return &ReplaceCoreV1NamespaceStatusParams{
		HTTPClient: client,
	}
}

/*ReplaceCoreV1NamespaceStatusParams contains all the parameters to send to the API endpoint
for the replace core v1 namespace status operation typically these are written to a http.Request
*/
type ReplaceCoreV1NamespaceStatusParams struct {

	/*Body*/
	Body *models.IoK8sAPICoreV1Namespace
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the Namespace

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

// WithTimeout adds the timeout to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithTimeout(timeout time.Duration) *ReplaceCoreV1NamespaceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithContext(ctx context.Context) *ReplaceCoreV1NamespaceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithHTTPClient(client *http.Client) *ReplaceCoreV1NamespaceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithBody(body *models.IoK8sAPICoreV1Namespace) *ReplaceCoreV1NamespaceStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetBody(body *models.IoK8sAPICoreV1Namespace) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithDryRun(dryRun *string) *ReplaceCoreV1NamespaceStatusParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithFieldManager(fieldManager *string) *ReplaceCoreV1NamespaceStatusParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithName(name string) *ReplaceCoreV1NamespaceStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) WithPretty(pretty *string) *ReplaceCoreV1NamespaceStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace core v1 namespace status params
func (o *ReplaceCoreV1NamespaceStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceCoreV1NamespaceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
