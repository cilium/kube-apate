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

// NewPostManagementCiliumIoV2CiliumEndpointsParams creates a new PostManagementCiliumIoV2CiliumEndpointsParams object
// with the default values initialized.
func NewPostManagementCiliumIoV2CiliumEndpointsParams() *PostManagementCiliumIoV2CiliumEndpointsParams {
	var ()
	return &PostManagementCiliumIoV2CiliumEndpointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostManagementCiliumIoV2CiliumEndpointsParamsWithTimeout creates a new PostManagementCiliumIoV2CiliumEndpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostManagementCiliumIoV2CiliumEndpointsParamsWithTimeout(timeout time.Duration) *PostManagementCiliumIoV2CiliumEndpointsParams {
	var ()
	return &PostManagementCiliumIoV2CiliumEndpointsParams{

		timeout: timeout,
	}
}

// NewPostManagementCiliumIoV2CiliumEndpointsParamsWithContext creates a new PostManagementCiliumIoV2CiliumEndpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostManagementCiliumIoV2CiliumEndpointsParamsWithContext(ctx context.Context) *PostManagementCiliumIoV2CiliumEndpointsParams {
	var ()
	return &PostManagementCiliumIoV2CiliumEndpointsParams{

		Context: ctx,
	}
}

// NewPostManagementCiliumIoV2CiliumEndpointsParamsWithHTTPClient creates a new PostManagementCiliumIoV2CiliumEndpointsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostManagementCiliumIoV2CiliumEndpointsParamsWithHTTPClient(client *http.Client) *PostManagementCiliumIoV2CiliumEndpointsParams {
	var ()
	return &PostManagementCiliumIoV2CiliumEndpointsParams{
		HTTPClient: client,
	}
}

/*PostManagementCiliumIoV2CiliumEndpointsParams contains all the parameters to send to the API endpoint
for the post management cilium io v2 cilium endpoints operation typically these are written to a http.Request
*/
type PostManagementCiliumIoV2CiliumEndpointsParams struct {

	/*Options
	  Options to modify structures

	*/
	Options *models.Options

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) WithTimeout(timeout time.Duration) *PostManagementCiliumIoV2CiliumEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) WithContext(ctx context.Context) *PostManagementCiliumIoV2CiliumEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) WithHTTPClient(client *http.Client) *PostManagementCiliumIoV2CiliumEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptions adds the options to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) WithOptions(options *models.Options) *PostManagementCiliumIoV2CiliumEndpointsParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the post management cilium io v2 cilium endpoints params
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) SetOptions(options *models.Options) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *PostManagementCiliumIoV2CiliumEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
