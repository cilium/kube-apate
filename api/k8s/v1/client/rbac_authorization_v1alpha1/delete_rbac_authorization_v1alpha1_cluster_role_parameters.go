// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1alpha1

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

// NewDeleteRbacAuthorizationV1alpha1ClusterRoleParams creates a new DeleteRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized.
func NewDeleteRbacAuthorizationV1alpha1ClusterRoleParams() *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &DeleteRbacAuthorizationV1alpha1ClusterRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRbacAuthorizationV1alpha1ClusterRoleParamsWithTimeout creates a new DeleteRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRbacAuthorizationV1alpha1ClusterRoleParamsWithTimeout(timeout time.Duration) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &DeleteRbacAuthorizationV1alpha1ClusterRoleParams{

		timeout: timeout,
	}
}

// NewDeleteRbacAuthorizationV1alpha1ClusterRoleParamsWithContext creates a new DeleteRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRbacAuthorizationV1alpha1ClusterRoleParamsWithContext(ctx context.Context) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &DeleteRbacAuthorizationV1alpha1ClusterRoleParams{

		Context: ctx,
	}
}

// NewDeleteRbacAuthorizationV1alpha1ClusterRoleParamsWithHTTPClient creates a new DeleteRbacAuthorizationV1alpha1ClusterRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRbacAuthorizationV1alpha1ClusterRoleParamsWithHTTPClient(client *http.Client) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	var ()
	return &DeleteRbacAuthorizationV1alpha1ClusterRoleParams{
		HTTPClient: client,
	}
}

/*DeleteRbacAuthorizationV1alpha1ClusterRoleParams contains all the parameters to send to the API endpoint
for the delete rbac authorization v1alpha1 cluster role operation typically these are written to a http.Request
*/
type DeleteRbacAuthorizationV1alpha1ClusterRoleParams struct {

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
	  name of the ClusterRole

	*/
	Name string
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

// WithTimeout adds the timeout to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithTimeout(timeout time.Duration) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithContext(ctx context.Context) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithHTTPClient(client *http.Client) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) {
	o.Body = body
}

// WithDryRun adds the dryRun to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithDryRun(dryRun *string) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithGracePeriodSeconds adds the gracePeriodSeconds to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithGracePeriodSeconds(gracePeriodSeconds *int64) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetGracePeriodSeconds(gracePeriodSeconds)
	return o
}

// SetGracePeriodSeconds adds the gracePeriodSeconds to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetGracePeriodSeconds(gracePeriodSeconds *int64) {
	o.GracePeriodSeconds = gracePeriodSeconds
}

// WithName adds the name to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithName(name string) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetName(name string) {
	o.Name = name
}

// WithOrphanDependents adds the orphanDependents to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithOrphanDependents(orphanDependents *bool) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetOrphanDependents(orphanDependents)
	return o
}

// SetOrphanDependents adds the orphanDependents to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetOrphanDependents(orphanDependents *bool) {
	o.OrphanDependents = orphanDependents
}

// WithPretty adds the pretty to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithPretty(pretty *string) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WithPropagationPolicy adds the propagationPolicy to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WithPropagationPolicy(propagationPolicy *string) *DeleteRbacAuthorizationV1alpha1ClusterRoleParams {
	o.SetPropagationPolicy(propagationPolicy)
	return o
}

// SetPropagationPolicy adds the propagationPolicy to the delete rbac authorization v1alpha1 cluster role params
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) SetPropagationPolicy(propagationPolicy *string) {
	o.PropagationPolicy = propagationPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
