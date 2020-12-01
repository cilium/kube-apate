// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1beta1

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

// NewReplaceApiregistrationV1beta1APIServiceStatusParams creates a new ReplaceApiregistrationV1beta1APIServiceStatusParams object
// with the default values initialized.
func NewReplaceApiregistrationV1beta1APIServiceStatusParams() *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	var ()
	return &ReplaceApiregistrationV1beta1APIServiceStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceApiregistrationV1beta1APIServiceStatusParamsWithTimeout creates a new ReplaceApiregistrationV1beta1APIServiceStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceApiregistrationV1beta1APIServiceStatusParamsWithTimeout(timeout time.Duration) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	var ()
	return &ReplaceApiregistrationV1beta1APIServiceStatusParams{

		timeout: timeout,
	}
}

// NewReplaceApiregistrationV1beta1APIServiceStatusParamsWithContext creates a new ReplaceApiregistrationV1beta1APIServiceStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceApiregistrationV1beta1APIServiceStatusParamsWithContext(ctx context.Context) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	var ()
	return &ReplaceApiregistrationV1beta1APIServiceStatusParams{

		Context: ctx,
	}
}

// NewReplaceApiregistrationV1beta1APIServiceStatusParamsWithHTTPClient creates a new ReplaceApiregistrationV1beta1APIServiceStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceApiregistrationV1beta1APIServiceStatusParamsWithHTTPClient(client *http.Client) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	var ()
	return &ReplaceApiregistrationV1beta1APIServiceStatusParams{
		HTTPClient: client,
	}
}

/*ReplaceApiregistrationV1beta1APIServiceStatusParams contains all the parameters to send to the API endpoint
for the replace apiregistration v1beta1 API service status operation typically these are written to a http.Request
*/
type ReplaceApiregistrationV1beta1APIServiceStatusParams struct {

	/*Body*/
	Body *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the APIService

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

// WithTimeout adds the timeout to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithTimeout(timeout time.Duration) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithContext(ctx context.Context) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithHTTPClient(client *http.Client) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithBody(body *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetBody(body *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithDryRun(dryRun *string) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithFieldManager(fieldManager *string) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithName(name string) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WithPretty(pretty *string) *ReplaceApiregistrationV1beta1APIServiceStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace apiregistration v1beta1 API service status params
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceApiregistrationV1beta1APIServiceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
