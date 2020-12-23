// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

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

// NewReplaceNetworkingV1beta1IngressClassParams creates a new ReplaceNetworkingV1beta1IngressClassParams object
// with the default values initialized.
func NewReplaceNetworkingV1beta1IngressClassParams() *ReplaceNetworkingV1beta1IngressClassParams {
	var ()
	return &ReplaceNetworkingV1beta1IngressClassParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceNetworkingV1beta1IngressClassParamsWithTimeout creates a new ReplaceNetworkingV1beta1IngressClassParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceNetworkingV1beta1IngressClassParamsWithTimeout(timeout time.Duration) *ReplaceNetworkingV1beta1IngressClassParams {
	var ()
	return &ReplaceNetworkingV1beta1IngressClassParams{

		timeout: timeout,
	}
}

// NewReplaceNetworkingV1beta1IngressClassParamsWithContext creates a new ReplaceNetworkingV1beta1IngressClassParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceNetworkingV1beta1IngressClassParamsWithContext(ctx context.Context) *ReplaceNetworkingV1beta1IngressClassParams {
	var ()
	return &ReplaceNetworkingV1beta1IngressClassParams{

		Context: ctx,
	}
}

// NewReplaceNetworkingV1beta1IngressClassParamsWithHTTPClient creates a new ReplaceNetworkingV1beta1IngressClassParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceNetworkingV1beta1IngressClassParamsWithHTTPClient(client *http.Client) *ReplaceNetworkingV1beta1IngressClassParams {
	var ()
	return &ReplaceNetworkingV1beta1IngressClassParams{
		HTTPClient: client,
	}
}

/*ReplaceNetworkingV1beta1IngressClassParams contains all the parameters to send to the API endpoint
for the replace networking v1beta1 ingress class operation typically these are written to a http.Request
*/
type ReplaceNetworkingV1beta1IngressClassParams struct {

	/*Body*/
	Body *models.IoK8sAPINetworkingV1beta1IngressClass
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.

	*/
	FieldManager *string
	/*Name
	  name of the IngressClass

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

// WithTimeout adds the timeout to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithTimeout(timeout time.Duration) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithContext(ctx context.Context) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithHTTPClient(client *http.Client) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithBody(body *models.IoK8sAPINetworkingV1beta1IngressClass) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetBody(body *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Body = body
}

// WithDryRun adds the dryRun to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithDryRun(dryRun *string) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithFieldManager(fieldManager *string) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithName adds the name to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithName(name string) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) WithPretty(pretty *string) *ReplaceNetworkingV1beta1IngressClassParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the replace networking v1beta1 ingress class params
func (o *ReplaceNetworkingV1beta1IngressClassParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceNetworkingV1beta1IngressClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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