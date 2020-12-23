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

// NewPostManagementCiliumIoV2CiliumIdentitiesParams creates a new PostManagementCiliumIoV2CiliumIdentitiesParams object
// with the default values initialized.
func NewPostManagementCiliumIoV2CiliumIdentitiesParams() *PostManagementCiliumIoV2CiliumIdentitiesParams {
	var ()
	return &PostManagementCiliumIoV2CiliumIdentitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostManagementCiliumIoV2CiliumIdentitiesParamsWithTimeout creates a new PostManagementCiliumIoV2CiliumIdentitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostManagementCiliumIoV2CiliumIdentitiesParamsWithTimeout(timeout time.Duration) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	var ()
	return &PostManagementCiliumIoV2CiliumIdentitiesParams{

		timeout: timeout,
	}
}

// NewPostManagementCiliumIoV2CiliumIdentitiesParamsWithContext creates a new PostManagementCiliumIoV2CiliumIdentitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostManagementCiliumIoV2CiliumIdentitiesParamsWithContext(ctx context.Context) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	var ()
	return &PostManagementCiliumIoV2CiliumIdentitiesParams{

		Context: ctx,
	}
}

// NewPostManagementCiliumIoV2CiliumIdentitiesParamsWithHTTPClient creates a new PostManagementCiliumIoV2CiliumIdentitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostManagementCiliumIoV2CiliumIdentitiesParamsWithHTTPClient(client *http.Client) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	var ()
	return &PostManagementCiliumIoV2CiliumIdentitiesParams{
		HTTPClient: client,
	}
}

/*PostManagementCiliumIoV2CiliumIdentitiesParams contains all the parameters to send to the API endpoint
for the post management cilium io v2 cilium identities operation typically these are written to a http.Request
*/
type PostManagementCiliumIoV2CiliumIdentitiesParams struct {

	/*Options
	  Options to modify structures

	*/
	Options *models.Options

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) WithTimeout(timeout time.Duration) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) WithContext(ctx context.Context) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) WithHTTPClient(client *http.Client) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptions adds the options to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) WithOptions(options *models.Options) *PostManagementCiliumIoV2CiliumIdentitiesParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the post management cilium io v2 cilium identities params
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) SetOptions(options *models.Options) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *PostManagementCiliumIoV2CiliumIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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