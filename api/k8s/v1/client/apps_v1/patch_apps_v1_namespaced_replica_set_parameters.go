// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

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

// NewPatchAppsV1NamespacedReplicaSetParams creates a new PatchAppsV1NamespacedReplicaSetParams object
// with the default values initialized.
func NewPatchAppsV1NamespacedReplicaSetParams() *PatchAppsV1NamespacedReplicaSetParams {
	var ()
	return &PatchAppsV1NamespacedReplicaSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAppsV1NamespacedReplicaSetParamsWithTimeout creates a new PatchAppsV1NamespacedReplicaSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAppsV1NamespacedReplicaSetParamsWithTimeout(timeout time.Duration) *PatchAppsV1NamespacedReplicaSetParams {
	var ()
	return &PatchAppsV1NamespacedReplicaSetParams{

		timeout: timeout,
	}
}

// NewPatchAppsV1NamespacedReplicaSetParamsWithContext creates a new PatchAppsV1NamespacedReplicaSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAppsV1NamespacedReplicaSetParamsWithContext(ctx context.Context) *PatchAppsV1NamespacedReplicaSetParams {
	var ()
	return &PatchAppsV1NamespacedReplicaSetParams{

		Context: ctx,
	}
}

// NewPatchAppsV1NamespacedReplicaSetParamsWithHTTPClient creates a new PatchAppsV1NamespacedReplicaSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAppsV1NamespacedReplicaSetParamsWithHTTPClient(client *http.Client) *PatchAppsV1NamespacedReplicaSetParams {
	var ()
	return &PatchAppsV1NamespacedReplicaSetParams{
		HTTPClient: client,
	}
}

/*PatchAppsV1NamespacedReplicaSetParams contains all the parameters to send to the API endpoint
for the patch apps v1 namespaced replica set operation typically these are written to a http.Request
*/
type PatchAppsV1NamespacedReplicaSetParams struct {

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
	  name of the ReplicaSet

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

// WithTimeout adds the timeout to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithTimeout(timeout time.Duration) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithContext(ctx context.Context) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithHTTPClient(client *http.Client) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetBody(body models.IoK8sApimachineryPkgApisMetaV1Patch) {
	o.Body = body
}

// WithDryRun adds the dryRun to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithDryRun(dryRun *string) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldManager adds the fieldManager to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithFieldManager(fieldManager *string) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetFieldManager(fieldManager)
	return o
}

// SetFieldManager adds the fieldManager to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetFieldManager(fieldManager *string) {
	o.FieldManager = fieldManager
}

// WithForce adds the force to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithForce(force *bool) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithName(name string) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithNamespace(namespace string) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPretty adds the pretty to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) WithPretty(pretty *string) *PatchAppsV1NamespacedReplicaSetParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the patch apps v1 namespaced replica set params
func (o *PatchAppsV1NamespacedReplicaSetParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAppsV1NamespacedReplicaSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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