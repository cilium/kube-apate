// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1

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

// NewCreateAuthorizationV1SubjectAccessReviewParams creates a new CreateAuthorizationV1SubjectAccessReviewParams object
// with the default values initialized.
func NewCreateAuthorizationV1SubjectAccessReviewParams() *CreateAuthorizationV1SubjectAccessReviewParams {
	var ()
	return &CreateAuthorizationV1SubjectAccessReviewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthorizationV1SubjectAccessReviewParamsWithTimeout creates a new CreateAuthorizationV1SubjectAccessReviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAuthorizationV1SubjectAccessReviewParamsWithTimeout(timeout time.Duration) *CreateAuthorizationV1SubjectAccessReviewParams {
	var ()
	return &CreateAuthorizationV1SubjectAccessReviewParams{

		timeout: timeout,
	}
}

// NewCreateAuthorizationV1SubjectAccessReviewParamsWithContext creates a new CreateAuthorizationV1SubjectAccessReviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAuthorizationV1SubjectAccessReviewParamsWithContext(ctx context.Context) *CreateAuthorizationV1SubjectAccessReviewParams {
	var ()
	return &CreateAuthorizationV1SubjectAccessReviewParams{

		Context: ctx,
	}
}

// NewCreateAuthorizationV1SubjectAccessReviewParamsWithHTTPClient creates a new CreateAuthorizationV1SubjectAccessReviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAuthorizationV1SubjectAccessReviewParamsWithHTTPClient(client *http.Client) *CreateAuthorizationV1SubjectAccessReviewParams {
	var ()
	return &CreateAuthorizationV1SubjectAccessReviewParams{
		HTTPClient: client,
	}
}

/*CreateAuthorizationV1SubjectAccessReviewParams contains all the parameters to send to the API endpoint
for the create authorization v1 subject access review operation typically these are written to a http.Request
*/
type CreateAuthorizationV1SubjectAccessReviewParams struct {

	/*Body*/
	Body *models.IoK8sAPIAuthorizationV1SubjectAccessReview
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

// WithTimeout adds the timeout to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithTimeout(timeout time.Duration) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithContext(ctx context.Context) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithHTTPClient(client *http.Client) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithBody(body *models.IoK8sAPIAuthorizationV1SubjectAccessReview) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetBody(body *models.IoK8sAPIAuthorizationV1SubjectAccessReview) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithDryRun(dryRun *string) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithFieldManager(fieldManager *string) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithPretty adds the pretty to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WithPretty(pretty *string) *CreateAuthorizationV1SubjectAccessReviewParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create authorization v1 subject access review params
func (o *CreateAuthorizationV1SubjectAccessReviewParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthorizationV1SubjectAccessReviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
