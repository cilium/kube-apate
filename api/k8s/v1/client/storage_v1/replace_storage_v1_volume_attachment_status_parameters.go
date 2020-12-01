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

// NewReplaceStorageV1VolumeAttachmentStatusParams creates a new ReplaceStorageV1VolumeAttachmentStatusParams object
// with the default values initialized.
func NewReplaceStorageV1VolumeAttachmentStatusParams() *ReplaceStorageV1VolumeAttachmentStatusParams {
	var ()
	return &ReplaceStorageV1VolumeAttachmentStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStorageV1VolumeAttachmentStatusParamsWithTimeout creates a new ReplaceStorageV1VolumeAttachmentStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceStorageV1VolumeAttachmentStatusParamsWithTimeout(timeout time.Duration) *ReplaceStorageV1VolumeAttachmentStatusParams {
	var ()
	return &ReplaceStorageV1VolumeAttachmentStatusParams{

		timeout: timeout,
	}
}

// NewReplaceStorageV1VolumeAttachmentStatusParamsWithContext creates a new ReplaceStorageV1VolumeAttachmentStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceStorageV1VolumeAttachmentStatusParamsWithContext(ctx context.Context) *ReplaceStorageV1VolumeAttachmentStatusParams {
	var ()
	return &ReplaceStorageV1VolumeAttachmentStatusParams{

		Context: ctx,
	}
}

// NewReplaceStorageV1VolumeAttachmentStatusParamsWithHTTPClient creates a new ReplaceStorageV1VolumeAttachmentStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceStorageV1VolumeAttachmentStatusParamsWithHTTPClient(client *http.Client) *ReplaceStorageV1VolumeAttachmentStatusParams {
	var ()
	return &ReplaceStorageV1VolumeAttachmentStatusParams{
		HTTPClient: client,
	}
}

/*ReplaceStorageV1VolumeAttachmentStatusParams contains all the parameters to send to the API endpoint
for the replace storage v1 volume attachment status operation typically these are written to a http.Request
*/
type ReplaceStorageV1VolumeAttachmentStatusParams struct {

	/*Body*/
	Body *models.IoK8sAPIStorageV1VolumeAttachment
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the VolumeAttachment

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

// WithTimeout adds the timeout to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithTimeout(timeout time.Duration) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithContext(ctx context.Context) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithHTTPClient(client *http.Client) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithBody(body *models.IoK8sAPIStorageV1VolumeAttachment) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetBody(body *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithDryRun(dryRun *string) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithFieldManager(fieldManager *string) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithName(name string) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WithPretty(pretty *string) *ReplaceStorageV1VolumeAttachmentStatusParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace storage v1 volume attachment status params
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStorageV1VolumeAttachmentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
