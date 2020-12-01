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

// NewPatchCoreV1NamespacedPodTemplateParams creates a new PatchCoreV1NamespacedPodTemplateParams object
// with the default values initialized.
func NewPatchCoreV1NamespacedPodTemplateParams() *PatchCoreV1NamespacedPodTemplateParams {
	var ()
	return &PatchCoreV1NamespacedPodTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCoreV1NamespacedPodTemplateParamsWithTimeout creates a new PatchCoreV1NamespacedPodTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCoreV1NamespacedPodTemplateParamsWithTimeout(timeout time.Duration) *PatchCoreV1NamespacedPodTemplateParams {
	var ()
	return &PatchCoreV1NamespacedPodTemplateParams{

		timeout: timeout,
	}
}

// NewPatchCoreV1NamespacedPodTemplateParamsWithContext creates a new PatchCoreV1NamespacedPodTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCoreV1NamespacedPodTemplateParamsWithContext(ctx context.Context) *PatchCoreV1NamespacedPodTemplateParams {
	var ()
	return &PatchCoreV1NamespacedPodTemplateParams{

		Context: ctx,
	}
}

// NewPatchCoreV1NamespacedPodTemplateParamsWithHTTPClient creates a new PatchCoreV1NamespacedPodTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCoreV1NamespacedPodTemplateParamsWithHTTPClient(client *http.Client) *PatchCoreV1NamespacedPodTemplateParams {
	var ()
	return &PatchCoreV1NamespacedPodTemplateParams{
		HTTPClient: client,
	}
}

/*PatchCoreV1NamespacedPodTemplateParams contains all the parameters to send to the API endpoint
for the patch core v1 namespaced pod template operation typically these are written to a http.Request
*/
type PatchCoreV1NamespacedPodTemplateParams struct {

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
	  name of the PodTemplate

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithTimeout(timeout time.Duration) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithContext(ctx context.Context) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithHTTPClient(client *http.Client) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) {
	o.Body = body
}

// WithDryRun adds the dryRun to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithDryRun(dryRun *string) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithFieldManager(fieldManager *string) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithForce adds the force to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithForce(force *bool) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithName(name string) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithNamespace(namespace string) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) WithPretty(pretty *string) *PatchCoreV1NamespacedPodTemplateParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the patch core v1 namespaced pod template params
func (o *PatchCoreV1NamespacedPodTemplateParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCoreV1NamespacedPodTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
