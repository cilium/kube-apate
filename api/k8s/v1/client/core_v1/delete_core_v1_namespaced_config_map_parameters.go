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

// NewDeleteCoreV1NamespacedConfigMapParams creates a new DeleteCoreV1NamespacedConfigMapParams object
// with the default values initialized.
func NewDeleteCoreV1NamespacedConfigMapParams() *DeleteCoreV1NamespacedConfigMapParams {
	var ()
	return &DeleteCoreV1NamespacedConfigMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCoreV1NamespacedConfigMapParamsWithTimeout creates a new DeleteCoreV1NamespacedConfigMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCoreV1NamespacedConfigMapParamsWithTimeout(timeout time.Duration) *DeleteCoreV1NamespacedConfigMapParams {
	var ()
	return &DeleteCoreV1NamespacedConfigMapParams{

		timeout: timeout,
	}
}

// NewDeleteCoreV1NamespacedConfigMapParamsWithContext creates a new DeleteCoreV1NamespacedConfigMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCoreV1NamespacedConfigMapParamsWithContext(ctx context.Context) *DeleteCoreV1NamespacedConfigMapParams {
	var ()
	return &DeleteCoreV1NamespacedConfigMapParams{

		Context: ctx,
	}
}

// NewDeleteCoreV1NamespacedConfigMapParamsWithHTTPClient creates a new DeleteCoreV1NamespacedConfigMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCoreV1NamespacedConfigMapParamsWithHTTPClient(client *http.Client) *DeleteCoreV1NamespacedConfigMapParams {
	var ()
	return &DeleteCoreV1NamespacedConfigMapParams{
		HTTPClient: client,
	}
}

/*DeleteCoreV1NamespacedConfigMapParams contains all the parameters to send to the API endpoint
for the delete core v1 namespaced config map operation typically these are written to a http.Request
*/
type DeleteCoreV1NamespacedConfigMapParams struct {

	/*Body*/
	Body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*GracePeriodSeconds
	  The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.

	*/
	GracePeriodSeconds *int64
	/*Name
	  name of the ConfigMap

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*OrphanDependents
	  Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.

	*/
	OrphanDependents *bool
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string
	/*PropagationPolicy
	  Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.

	*/
	PropagationPolicy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithTimeout(timeout time.Duration) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithContext(ctx context.Context) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithHTTPClient(client *http.Client) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) {
	o.Body = body
}

// WithDryRun adds the dryRun to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithDryRun(dryRun *string) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithGracePeriodSeconds adds the gracePeriodSeconds to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithGracePeriodSeconds(gracePeriodSeconds *int64) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetGracePeriodSeconds(gracePeriodSeconds)
	return o
}

// SetGracePeriodSeconds adds the gracePeriodSeconds to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetGracePeriodSeconds(gracePeriodSeconds *int64) {
	o.GracePeriodSeconds = gracePeriodSeconds
}

// WithName adds the name to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithName(name string) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithNamespace(namespace string) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOrphanDependents adds the orphanDependents to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithOrphanDependents(orphanDependents *bool) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetOrphanDependents(orphanDependents)
	return o
}

// SetOrphanDependents adds the orphanDependents to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetOrphanDependents(orphanDependents *bool) {
	o.OrphanDependents = orphanDependents
}

// WithPretty adds the pretty to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithPretty(pretty *string) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WithPropagationPolicy adds the propagationPolicy to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) WithPropagationPolicy(propagationPolicy *string) *DeleteCoreV1NamespacedConfigMapParams {
	o.SetPropagationPolicy(propagationPolicy)
	return o
}

// SetPropagationPolicy adds the propagationPolicy to the delete core v1 namespaced config map params
func (o *DeleteCoreV1NamespacedConfigMapParams) SetPropagationPolicy(propagationPolicy *string) {
	o.PropagationPolicy = propagationPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCoreV1NamespacedConfigMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.GracePeriodSeconds != nil {

		// query param gracePeriodSeconds
		var qrGracePeriodSeconds int64
		if o.GracePeriodSeconds != nil {
			qrGracePeriodSeconds = *o.GracePeriodSeconds
		}
		qGracePeriodSeconds := swag.FormatInt64(qrGracePeriodSeconds)
		if qGracePeriodSeconds != "" {
			if err := r.SetQueryParam("gracePeriodSeconds", qGracePeriodSeconds); err != nil {
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

	if o.OrphanDependents != nil {

		// query param orphanDependents
		var qrOrphanDependents bool
		if o.OrphanDependents != nil {
			qrOrphanDependents = *o.OrphanDependents
		}
		qOrphanDependents := swag.FormatBool(qrOrphanDependents)
		if qOrphanDependents != "" {
			if err := r.SetQueryParam("orphanDependents", qOrphanDependents); err != nil {
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

	if o.PropagationPolicy != nil {

		// query param propagationPolicy
		var qrPropagationPolicy string
		if o.PropagationPolicy != nil {
			qrPropagationPolicy = *o.PropagationPolicy
		}
		qPropagationPolicy := qrPropagationPolicy
		if qPropagationPolicy != "" {
			if err := r.SetQueryParam("propagationPolicy", qPropagationPolicy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
