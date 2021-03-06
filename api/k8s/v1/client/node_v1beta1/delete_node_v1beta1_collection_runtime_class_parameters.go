// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1beta1

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

// NewDeleteNodeV1beta1CollectionRuntimeClassParams creates a new DeleteNodeV1beta1CollectionRuntimeClassParams object
// with the default values initialized.
func NewDeleteNodeV1beta1CollectionRuntimeClassParams() *DeleteNodeV1beta1CollectionRuntimeClassParams {
	var ()
	return &DeleteNodeV1beta1CollectionRuntimeClassParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodeV1beta1CollectionRuntimeClassParamsWithTimeout creates a new DeleteNodeV1beta1CollectionRuntimeClassParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodeV1beta1CollectionRuntimeClassParamsWithTimeout(timeout time.Duration) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	var ()
	return &DeleteNodeV1beta1CollectionRuntimeClassParams{

		timeout: timeout,
	}
}

// NewDeleteNodeV1beta1CollectionRuntimeClassParamsWithContext creates a new DeleteNodeV1beta1CollectionRuntimeClassParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodeV1beta1CollectionRuntimeClassParamsWithContext(ctx context.Context) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	var ()
	return &DeleteNodeV1beta1CollectionRuntimeClassParams{

		Context: ctx,
	}
}

// NewDeleteNodeV1beta1CollectionRuntimeClassParamsWithHTTPClient creates a new DeleteNodeV1beta1CollectionRuntimeClassParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodeV1beta1CollectionRuntimeClassParamsWithHTTPClient(client *http.Client) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	var ()
	return &DeleteNodeV1beta1CollectionRuntimeClassParams{
		HTTPClient: client,
	}
}

/*DeleteNodeV1beta1CollectionRuntimeClassParams contains all the parameters to send to the API endpoint
for the delete node v1beta1 collection runtime class operation typically these are written to a http.Request
*/
type DeleteNodeV1beta1CollectionRuntimeClassParams struct {

	/*Body*/
	Body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
	/*Continue
	  The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the "next key".

	This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.

	*/
	Continue *string
	/*DryRun
	  When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed

	*/
	DryRun *string
	/*FieldSelector
	  A selector to restrict the list of returned objects by their fields. Defaults to everything.

	*/
	FieldSelector *string
	/*GracePeriodSeconds
	  The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.

	*/
	GracePeriodSeconds *int64
	/*LabelSelector
	  A selector to restrict the list of returned objects by their labels. Defaults to everything.

	*/
	LabelSelector *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.

	The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.

	*/
	Limit *int64
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
	/*ResourceVersion
	  resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset

	*/
	ResourceVersion *string
	/*ResourceVersionMatch
	  resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset

	*/
	ResourceVersionMatch *string
	/*TimeoutSeconds
	  Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.

	*/
	TimeoutSeconds *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithTimeout(timeout time.Duration) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithContext(ctx context.Context) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithHTTPClient(client *http.Client) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetBody(body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions) {
	o.Body = body
}

// WithContinue adds the continueVar to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithContinue(continueVar *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithDryRun adds the dryRun to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithDryRun(dryRun *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetDryRun(dryRun *string) {
	o.DryRun = dryRun
}

// WithFieldSelector adds the fieldSelector to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithFieldSelector(fieldSelector *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetFieldSelector(fieldSelector)
	return o
}

// SetFieldSelector adds the fieldSelector to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetFieldSelector(fieldSelector *string) {
	o.FieldSelector = fieldSelector
}

// WithGracePeriodSeconds adds the gracePeriodSeconds to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithGracePeriodSeconds(gracePeriodSeconds *int64) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetGracePeriodSeconds(gracePeriodSeconds)
	return o
}

// SetGracePeriodSeconds adds the gracePeriodSeconds to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetGracePeriodSeconds(gracePeriodSeconds *int64) {
	o.GracePeriodSeconds = gracePeriodSeconds
}

// WithLabelSelector adds the labelSelector to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithLabelSelector(labelSelector *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetLabelSelector(labelSelector)
	return o
}

// SetLabelSelector adds the labelSelector to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetLabelSelector(labelSelector *string) {
	o.LabelSelector = labelSelector
}

// WithLimit adds the limit to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithLimit(limit *int64) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOrphanDependents adds the orphanDependents to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithOrphanDependents(orphanDependents *bool) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetOrphanDependents(orphanDependents)
	return o
}

// SetOrphanDependents adds the orphanDependents to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetOrphanDependents(orphanDependents *bool) {
	o.OrphanDependents = orphanDependents
}

// WithPretty adds the pretty to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithPretty(pretty *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WithPropagationPolicy adds the propagationPolicy to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithPropagationPolicy(propagationPolicy *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetPropagationPolicy(propagationPolicy)
	return o
}

// SetPropagationPolicy adds the propagationPolicy to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetPropagationPolicy(propagationPolicy *string) {
	o.PropagationPolicy = propagationPolicy
}

// WithResourceVersion adds the resourceVersion to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithResourceVersion(resourceVersion *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetResourceVersion(resourceVersion)
	return o
}

// SetResourceVersion adds the resourceVersion to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetResourceVersion(resourceVersion *string) {
	o.ResourceVersion = resourceVersion
}

// WithResourceVersionMatch adds the resourceVersionMatch to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithResourceVersionMatch(resourceVersionMatch *string) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetResourceVersionMatch(resourceVersionMatch)
	return o
}

// SetResourceVersionMatch adds the resourceVersionMatch to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetResourceVersionMatch(resourceVersionMatch *string) {
	o.ResourceVersionMatch = resourceVersionMatch
}

// WithTimeoutSeconds adds the timeoutSeconds to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WithTimeoutSeconds(timeoutSeconds *int64) *DeleteNodeV1beta1CollectionRuntimeClassParams {
	o.SetTimeoutSeconds(timeoutSeconds)
	return o
}

// SetTimeoutSeconds adds the timeoutSeconds to the delete node v1beta1 collection runtime class params
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) SetTimeoutSeconds(timeoutSeconds *int64) {
	o.TimeoutSeconds = timeoutSeconds
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Continue != nil {

		// query param continue
		var qrContinue string
		if o.Continue != nil {
			qrContinue = *o.Continue
		}
		qContinue := qrContinue
		if qContinue != "" {
			if err := r.SetQueryParam("continue", qContinue); err != nil {
				return err
			}
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

	if o.FieldSelector != nil {

		// query param fieldSelector
		var qrFieldSelector string
		if o.FieldSelector != nil {
			qrFieldSelector = *o.FieldSelector
		}
		qFieldSelector := qrFieldSelector
		if qFieldSelector != "" {
			if err := r.SetQueryParam("fieldSelector", qFieldSelector); err != nil {
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

	if o.LabelSelector != nil {

		// query param labelSelector
		var qrLabelSelector string
		if o.LabelSelector != nil {
			qrLabelSelector = *o.LabelSelector
		}
		qLabelSelector := qrLabelSelector
		if qLabelSelector != "" {
			if err := r.SetQueryParam("labelSelector", qLabelSelector); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

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

	if o.ResourceVersion != nil {

		// query param resourceVersion
		var qrResourceVersion string
		if o.ResourceVersion != nil {
			qrResourceVersion = *o.ResourceVersion
		}
		qResourceVersion := qrResourceVersion
		if qResourceVersion != "" {
			if err := r.SetQueryParam("resourceVersion", qResourceVersion); err != nil {
				return err
			}
		}

	}

	if o.ResourceVersionMatch != nil {

		// query param resourceVersionMatch
		var qrResourceVersionMatch string
		if o.ResourceVersionMatch != nil {
			qrResourceVersionMatch = *o.ResourceVersionMatch
		}
		qResourceVersionMatch := qrResourceVersionMatch
		if qResourceVersionMatch != "" {
			if err := r.SetQueryParam("resourceVersionMatch", qResourceVersionMatch); err != nil {
				return err
			}
		}

	}

	if o.TimeoutSeconds != nil {

		// query param timeoutSeconds
		var qrTimeoutSeconds int64
		if o.TimeoutSeconds != nil {
			qrTimeoutSeconds = *o.TimeoutSeconds
		}
		qTimeoutSeconds := swag.FormatInt64(qrTimeoutSeconds)
		if qTimeoutSeconds != "" {
			if err := r.SetQueryParam("timeoutSeconds", qTimeoutSeconds); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
