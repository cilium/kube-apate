// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1

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

// NewReplaceCertificatesV1CertificateSigningRequestApprovalParams creates a new ReplaceCertificatesV1CertificateSigningRequestApprovalParams object
// with the default values initialized.
func NewReplaceCertificatesV1CertificateSigningRequestApprovalParams() *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	var ()
	return &ReplaceCertificatesV1CertificateSigningRequestApprovalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceCertificatesV1CertificateSigningRequestApprovalParamsWithTimeout creates a new ReplaceCertificatesV1CertificateSigningRequestApprovalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceCertificatesV1CertificateSigningRequestApprovalParamsWithTimeout(timeout time.Duration) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	var ()
	return &ReplaceCertificatesV1CertificateSigningRequestApprovalParams{

		timeout: timeout,
	}
}

// NewReplaceCertificatesV1CertificateSigningRequestApprovalParamsWithContext creates a new ReplaceCertificatesV1CertificateSigningRequestApprovalParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceCertificatesV1CertificateSigningRequestApprovalParamsWithContext(ctx context.Context) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	var ()
	return &ReplaceCertificatesV1CertificateSigningRequestApprovalParams{

		Context: ctx,
	}
}

// NewReplaceCertificatesV1CertificateSigningRequestApprovalParamsWithHTTPClient creates a new ReplaceCertificatesV1CertificateSigningRequestApprovalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceCertificatesV1CertificateSigningRequestApprovalParamsWithHTTPClient(client *http.Client) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	var ()
	return &ReplaceCertificatesV1CertificateSigningRequestApprovalParams{
		HTTPClient: client,
	}
}

/*ReplaceCertificatesV1CertificateSigningRequestApprovalParams contains all the parameters to send to the API endpoint
for the replace certificates v1 certificate signing request approval operation typically these are written to a http.Request
*/
type ReplaceCertificatesV1CertificateSigningRequestApprovalParams struct {

	/*Body*/
	Body *models.IoK8sAPICertificatesV1CertificateSigningRequest
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the CertificateSigningRequest

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

// WithTimeout adds the timeout to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithTimeout(timeout time.Duration) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithContext(ctx context.Context) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithHTTPClient(client *http.Client) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithBody(body *models.IoK8sAPICertificatesV1CertificateSigningRequest) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetBody(body *models.IoK8sAPICertificatesV1CertificateSigningRequest) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithDryRun(dryRun *string) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithFieldManager(fieldManager *string) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithName(name string) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WithPretty(pretty *string) *ReplaceCertificatesV1CertificateSigningRequestApprovalParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace certificates v1 certificate signing request approval params
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceCertificatesV1CertificateSigningRequestApprovalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
