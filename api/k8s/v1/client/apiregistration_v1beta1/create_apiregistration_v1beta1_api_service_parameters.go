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

// NewCreateApiregistrationV1beta1APIServiceParams creates a new CreateApiregistrationV1beta1APIServiceParams object
// with the default values initialized.
func NewCreateApiregistrationV1beta1APIServiceParams() *CreateApiregistrationV1beta1APIServiceParams {
	var ()
	return &CreateApiregistrationV1beta1APIServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApiregistrationV1beta1APIServiceParamsWithTimeout creates a new CreateApiregistrationV1beta1APIServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateApiregistrationV1beta1APIServiceParamsWithTimeout(timeout time.Duration) *CreateApiregistrationV1beta1APIServiceParams {
	var ()
	return &CreateApiregistrationV1beta1APIServiceParams{

		timeout: timeout,
	}
}

// NewCreateApiregistrationV1beta1APIServiceParamsWithContext creates a new CreateApiregistrationV1beta1APIServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateApiregistrationV1beta1APIServiceParamsWithContext(ctx context.Context) *CreateApiregistrationV1beta1APIServiceParams {
	var ()
	return &CreateApiregistrationV1beta1APIServiceParams{

		Context: ctx,
	}
}

// NewCreateApiregistrationV1beta1APIServiceParamsWithHTTPClient creates a new CreateApiregistrationV1beta1APIServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateApiregistrationV1beta1APIServiceParamsWithHTTPClient(client *http.Client) *CreateApiregistrationV1beta1APIServiceParams {
	var ()
	return &CreateApiregistrationV1beta1APIServiceParams{
		HTTPClient: client,
	}
}

/*CreateApiregistrationV1beta1APIServiceParams contains all the parameters to send to the API endpoint
for the create apiregistration v1beta1 API service operation typically these are written to a http.Request
*/
type CreateApiregistrationV1beta1APIServiceParams struct {

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
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithTimeout(timeout time.Duration) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithContext(ctx context.Context) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithHTTPClient(client *http.Client) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithBody(body *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetBody(body *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithDryRun(dryRun *string) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithFieldManager(fieldManager *string) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithPretty adds the pretty to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) WithPretty(pretty *string) *CreateApiregistrationV1beta1APIServiceParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create apiregistration v1beta1 API service params
func (o *CreateApiregistrationV1beta1APIServiceParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApiregistrationV1beta1APIServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
