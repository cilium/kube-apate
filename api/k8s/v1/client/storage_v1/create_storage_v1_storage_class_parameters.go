// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

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

// NewCreateStorageV1StorageClassParams creates a new CreateStorageV1StorageClassParams object
// with the default values initialized.
func NewCreateStorageV1StorageClassParams() *CreateStorageV1StorageClassParams {
	var ()
	return &CreateStorageV1StorageClassParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStorageV1StorageClassParamsWithTimeout creates a new CreateStorageV1StorageClassParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateStorageV1StorageClassParamsWithTimeout(timeout time.Duration) *CreateStorageV1StorageClassParams {
	var ()
	return &CreateStorageV1StorageClassParams{

		timeout: timeout,
	}
}

// NewCreateStorageV1StorageClassParamsWithContext creates a new CreateStorageV1StorageClassParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateStorageV1StorageClassParamsWithContext(ctx context.Context) *CreateStorageV1StorageClassParams {
	var ()
	return &CreateStorageV1StorageClassParams{

		Context: ctx,
	}
}

// NewCreateStorageV1StorageClassParamsWithHTTPClient creates a new CreateStorageV1StorageClassParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateStorageV1StorageClassParamsWithHTTPClient(client *http.Client) *CreateStorageV1StorageClassParams {
	var ()
	return &CreateStorageV1StorageClassParams{
		HTTPClient: client,
	}
}

/*CreateStorageV1StorageClassParams contains all the parameters to send to the API endpoint
for the create storage v1 storage class operation typically these are written to a http.Request
*/
type CreateStorageV1StorageClassParams struct {

	/*Body*/
	Body *models.IoK8sAPIStorageV1StorageClass
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

// WithTimeout adds the timeout to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithTimeout(timeout time.Duration) *CreateStorageV1StorageClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithContext(ctx context.Context) *CreateStorageV1StorageClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithHTTPClient(client *http.Client) *CreateStorageV1StorageClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithBody(body *models.IoK8sAPIStorageV1StorageClass) *CreateStorageV1StorageClassParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetBody(body *models.IoK8sAPIStorageV1StorageClass) {
	o.Body = body
}

// WithDryRun adds the dryRun to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithDryRun(dryRun *string) *CreateStorageV1StorageClassParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithFieldManager(fieldManager *string) *CreateStorageV1StorageClassParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithPretty adds the pretty to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) WithPretty(pretty *string) *CreateStorageV1StorageClassParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the create storage v1 storage class params
func (o *CreateStorageV1StorageClassParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStorageV1StorageClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
