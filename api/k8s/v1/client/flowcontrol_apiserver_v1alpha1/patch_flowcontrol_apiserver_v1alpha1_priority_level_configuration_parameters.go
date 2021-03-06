// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package flowcontrol_apiserver_v1alpha1

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

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams creates a new PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams object
// with the default values initialized.
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams() *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	var ()
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParamsWithTimeout creates a new PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParamsWithTimeout(timeout time.Duration) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	var ()
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams{

		timeout: timeout,
	}
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParamsWithContext creates a new PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParamsWithContext(ctx context.Context) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	var ()
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams{

		Context: ctx,
	}
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParamsWithHTTPClient creates a new PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParamsWithHTTPClient(client *http.Client) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	var ()
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams{
		HTTPClient: client,
	}
}

/*PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams contains all the parameters to send to the API endpoint
for the patch flowcontrol apiserver v1alpha1 priority level configuration operation typically these are written to a http.Request
*/
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams struct {

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
	  name of the PriorityLevelConfiguration

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

// WithTimeout adds the timeout to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithTimeout(timeout time.Duration) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithContext(ctx context.Context) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithHTTPClient(client *http.Client) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) {
	o.Body = body
}

// WithDryRun adds the dryRun to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithDryRun(dryRun *string) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithFieldManager(fieldManager *string) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithForce adds the force to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithForce(force *bool) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithName(name string) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetName(name string) {
	o.Name = name
}

// WithPretty adds the pretty to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WithPretty(pretty *string) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the patch flowcontrol apiserver v1alpha1 priority level configuration params
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
