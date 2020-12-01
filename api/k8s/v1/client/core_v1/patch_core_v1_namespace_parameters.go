// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

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
	"github.com/go-openapi/swag"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// NewPatchCoreV1NamespaceParams creates a new PatchCoreV1NamespaceParams object
// with the default values initialized.
func NewPatchCoreV1NamespaceParams() *PatchCoreV1NamespaceParams {
	var ()
	return &PatchCoreV1NamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCoreV1NamespaceParamsWithTimeout creates a new PatchCoreV1NamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCoreV1NamespaceParamsWithTimeout(timeout time.Duration) *PatchCoreV1NamespaceParams {
	var ()
	return &PatchCoreV1NamespaceParams{

		timeout: timeout,
	}
}

// NewPatchCoreV1NamespaceParamsWithContext creates a new PatchCoreV1NamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCoreV1NamespaceParamsWithContext(ctx context.Context) *PatchCoreV1NamespaceParams {
	var ()
	return &PatchCoreV1NamespaceParams{

		Context: ctx,
	}
}

// NewPatchCoreV1NamespaceParamsWithHTTPClient creates a new PatchCoreV1NamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCoreV1NamespaceParamsWithHTTPClient(client *http.Client) *PatchCoreV1NamespaceParams {
	var ()
	return &PatchCoreV1NamespaceParams{
		HTTPClient: client,
	}
}

/*PatchCoreV1NamespaceParams contains all the parameters to send to the API endpoint
for the patch core v1 namespace operation typically these are written to a http.Request
*/
type PatchCoreV1NamespaceParams struct {

	/*Body*/
	Body models.IoK8sApimachineryPkgApisMetaV1Patch
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldManager
	  fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. This field is required for apply requests (application/apply-patch) but optional for non-apply patch types (JsonPatch, MergePatch, StrategicMergePatch).

	*/
	FieldManager *string
	/*Force
	  Force is going to "force" Apply requests. It means user will re-acquire conflicting fields owned by other people. Force flag must be unset for non-apply patch requests.

	*/
	Force *bool
	/*Name
	  name of the Namespace

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

// WithTimeout adds the timeout to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithTimeout(timeout time.Duration) *PatchCoreV1NamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithContext(ctx context.Context) *PatchCoreV1NamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithHTTPClient(client *http.Client) *PatchCoreV1NamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) *PatchCoreV1NamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) {
	o.Body = body
}

// WithDryRun adds the dryRun to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithDryRun(dryRun *string) *PatchCoreV1NamespaceParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithFieldManager(fieldManager *string) *PatchCoreV1NamespaceParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithForce adds the force to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithForce(force *bool) *PatchCoreV1NamespaceParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithName(name string) *PatchCoreV1NamespaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) WithPretty(pretty *string) *PatchCoreV1NamespaceParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the patch core v1 namespace params
func (o *PatchCoreV1NamespaceParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCoreV1NamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
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
