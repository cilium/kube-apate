// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

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

	"github.com/cilium/kube-apate/api/management/v1/models"
)

// NewPostManagementKubernetesIoV1PodsParams creates a new PostManagementKubernetesIoV1PodsParams object
// with the default values initialized.
func NewPostManagementKubernetesIoV1PodsParams() *PostManagementKubernetesIoV1PodsParams {
	var ()
	return &PostManagementKubernetesIoV1PodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostManagementKubernetesIoV1PodsParamsWithTimeout creates a new PostManagementKubernetesIoV1PodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostManagementKubernetesIoV1PodsParamsWithTimeout(timeout time.Duration) *PostManagementKubernetesIoV1PodsParams {
	var ()
	return &PostManagementKubernetesIoV1PodsParams{

		timeout: timeout,
	}
}

// NewPostManagementKubernetesIoV1PodsParamsWithContext creates a new PostManagementKubernetesIoV1PodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostManagementKubernetesIoV1PodsParamsWithContext(ctx context.Context) *PostManagementKubernetesIoV1PodsParams {
	var ()
	return &PostManagementKubernetesIoV1PodsParams{

		Context: ctx,
	}
}

// NewPostManagementKubernetesIoV1PodsParamsWithHTTPClient creates a new PostManagementKubernetesIoV1PodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostManagementKubernetesIoV1PodsParamsWithHTTPClient(client *http.Client) *PostManagementKubernetesIoV1PodsParams {
	var ()
	return &PostManagementKubernetesIoV1PodsParams{
		HTTPClient: client,
	}
}

/*PostManagementKubernetesIoV1PodsParams contains all the parameters to send to the API endpoint
for the post management kubernetes io v1 pods operation typically these are written to a http.Request
*/
type PostManagementKubernetesIoV1PodsParams struct {

	/*Options
	  Options to modify structures

	*/
	Options *models.Options

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) WithTimeout(timeout time.Duration) *PostManagementKubernetesIoV1PodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) WithContext(ctx context.Context) *PostManagementKubernetesIoV1PodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) WithHTTPClient(client *http.Client) *PostManagementKubernetesIoV1PodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptions adds the options to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) WithOptions(options *models.Options) *PostManagementKubernetesIoV1PodsParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the post management kubernetes io v1 pods params
func (o *PostManagementKubernetesIoV1PodsParams) SetOptions(options *models.Options) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *PostManagementKubernetesIoV1PodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Options != nil {
		if err := r.SetBodyParam(o.Options); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
