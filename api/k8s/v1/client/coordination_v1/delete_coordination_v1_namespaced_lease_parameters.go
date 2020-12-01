// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

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

// NewDeleteCoordinationV1NamespacedLeaseParams creates a new DeleteCoordinationV1NamespacedLeaseParams object
// with the default values initialized.
func NewDeleteCoordinationV1NamespacedLeaseParams() *DeleteCoordinationV1NamespacedLeaseParams {
	var ()
	return &DeleteCoordinationV1NamespacedLeaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCoordinationV1NamespacedLeaseParamsWithTimeout creates a new DeleteCoordinationV1NamespacedLeaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCoordinationV1NamespacedLeaseParamsWithTimeout(timeout time.Duration) *DeleteCoordinationV1NamespacedLeaseParams {
	var ()
	return &DeleteCoordinationV1NamespacedLeaseParams{

		timeout: timeout,
	}
}

// NewDeleteCoordinationV1NamespacedLeaseParamsWithContext creates a new DeleteCoordinationV1NamespacedLeaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCoordinationV1NamespacedLeaseParamsWithContext(ctx context.Context) *DeleteCoordinationV1NamespacedLeaseParams {
	var ()
	return &DeleteCoordinationV1NamespacedLeaseParams{

		Context: ctx,
	}
}

// NewDeleteCoordinationV1NamespacedLeaseParamsWithHTTPClient creates a new DeleteCoordinationV1NamespacedLeaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCoordinationV1NamespacedLeaseParamsWithHTTPClient(client *http.Client) *DeleteCoordinationV1NamespacedLeaseParams {
	var ()
	return &DeleteCoordinationV1NamespacedLeaseParams{
		HTTPClient: client,
	}
}

/*DeleteCoordinationV1NamespacedLeaseParams contains all the parameters to send to the API endpoint
for the delete coordination v1 namespaced lease operation typically these are written to a http.Request
*/
type DeleteCoordinationV1NamespacedLeaseParams struct {

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
	  name of the Lease

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

// WithTimeout adds the timeout to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithTimeout(timeout time.Duration) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithContext(ctx context.Context) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithHTTPClient(client *http.Client) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) {
	o.Body = body
}

// WithDryRun adds the dryRun to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithDryRun(dryRun *string) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithGracePeriodSeconds adds the gracePeriodSeconds to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithGracePeriodSeconds(gracePeriodSeconds *int64) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetGracePeriodSeconds(gracePeriodSeconds)
	return o
}

// SetGracePeriodSeconds adds the gracePeriodSeconds to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetGracePeriodSeconds(gracePeriodSeconds *int64) {
	o.GracePeriodSeconds = gracePeriodSeconds
}

// WithName adds the name to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithName(name string) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithNamespace(namespace string) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOrphanDependents adds the orphanDependents to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithOrphanDependents(orphanDependents *bool) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetOrphanDependents(orphanDependents)
	return o
}

// SetOrphanDependents adds the orphanDependents to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetOrphanDependents(orphanDependents *bool) {
	o.OrphanDependents = orphanDependents
}

// WithPretty adds the pretty to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithPretty(pretty *string) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WithPropagationPolicy adds the propagationPolicy to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) WithPropagationPolicy(propagationPolicy *string) *DeleteCoordinationV1NamespacedLeaseParams {
	o.SetPropagationPolicy(propagationPolicy)
	return o
}

// SetPropagationPolicy adds the propagationPolicy to the delete coordination v1 namespaced lease params
func (o *DeleteCoordinationV1NamespacedLeaseParams) SetPropagationPolicy(propagationPolicy *string) {
	o.PropagationPolicy = propagationPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCoordinationV1NamespacedLeaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
