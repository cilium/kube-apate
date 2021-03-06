// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1beta1

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

// NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams creates a new CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams object
// with the default values initialized.
func NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams() *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	var ()
	return &CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParamsWithTimeout creates a new CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParamsWithTimeout(timeout time.Duration) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	var ()
	return &CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams{

		timeout: timeout,
	}
}

// NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParamsWithContext creates a new CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParamsWithContext(ctx context.Context) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	var ()
	return &CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams{

		Context: ctx,
	}
}

// NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParamsWithHTTPClient creates a new CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParamsWithHTTPClient(client *http.Client) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	var ()
	return &CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams{
		HTTPClient: client,
	}
}

/*CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams contains all the parameters to send to the API endpoint
for the create admissionregistration v1beta1 validating webhook configuration operation typically these are written to a http.Request
*/
type CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams struct {

	/*Body*/
	Body *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration
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

// WithTimeout adds the timeout to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithTimeout(timeout time.Duration) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithContext(ctx context.Context) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithHTTPClient(client *http.Client) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithBody(body *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetBody(body *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithDryRun(dryRun *string) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithFieldManager(fieldManager *string) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithPretty adds the pretty to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WithPretty(pretty *string) *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create admissionregistration v1beta1 validating webhook configuration params
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
