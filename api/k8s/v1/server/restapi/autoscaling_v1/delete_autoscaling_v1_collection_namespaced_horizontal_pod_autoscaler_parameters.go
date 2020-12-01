// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams creates a new DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams object
// no default values defined in spec.
func NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams() DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams {

	return DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams{}
}

// DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams contains all the bound params for the delete autoscaling v1 collection namespaced horizontal pod autoscaler operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscaler
type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
	/*The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the "next key".

	This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
	  Unique: true
	  In: query
	*/
	Continue *string
	/*When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	  Unique: true
	  In: query
	*/
	DryRun *string
	/*A selector to restrict the list of returned objects by their fields. Defaults to everything.
	  Unique: true
	  In: query
	*/
	FieldSelector *string
	/*The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	  Unique: true
	  In: query
	*/
	GracePeriodSeconds *int64
	/*A selector to restrict the list of returned objects by their labels. Defaults to everything.
	  Unique: true
	  In: query
	*/
	LabelSelector *string
	/*limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.

	The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
	  Unique: true
	  In: query
	*/
	Limit *int64
	/*object name and auth scope, such as for teams and projects
	  Required: true
	  Unique: true
	  In: path
	*/
	Namespace string
	/*Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.
	  Unique: true
	  In: query
	*/
	OrphanDependents *bool
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
	/*Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.
	  Unique: true
	  In: query
	*/
	PropagationPolicy *string
	/*resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset
	  Unique: true
	  In: query
	*/
	ResourceVersion *string
	/*resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset
	  Unique: true
	  In: query
	*/
	ResourceVersionMatch *string
	/*Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.
	  Unique: true
	  In: query
	*/
	TimeoutSeconds *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams() beforehand.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	qContinue, qhkContinue, _ := qs.GetOK("continue")
	if err := o.bindContinue(qContinue, qhkContinue, route.Formats); err != nil {
		res = append(res, err)
	}

	qDryRun, qhkDryRun, _ := qs.GetOK("dryRun")
	if err := o.bindDryRun(qDryRun, qhkDryRun, route.Formats); err != nil {
		res = append(res, err)
	}

	qFieldSelector, qhkFieldSelector, _ := qs.GetOK("fieldSelector")
	if err := o.bindFieldSelector(qFieldSelector, qhkFieldSelector, route.Formats); err != nil {
		res = append(res, err)
	}

	qGracePeriodSeconds, qhkGracePeriodSeconds, _ := qs.GetOK("gracePeriodSeconds")
	if err := o.bindGracePeriodSeconds(qGracePeriodSeconds, qhkGracePeriodSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	qLabelSelector, qhkLabelSelector, _ := qs.GetOK("labelSelector")
	if err := o.bindLabelSelector(qLabelSelector, qhkLabelSelector, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrphanDependents, qhkOrphanDependents, _ := qs.GetOK("orphanDependents")
	if err := o.bindOrphanDependents(qOrphanDependents, qhkOrphanDependents, route.Formats); err != nil {
		res = append(res, err)
	}

	qPretty, qhkPretty, _ := qs.GetOK("pretty")
	if err := o.bindPretty(qPretty, qhkPretty, route.Formats); err != nil {
		res = append(res, err)
	}

	qPropagationPolicy, qhkPropagationPolicy, _ := qs.GetOK("propagationPolicy")
	if err := o.bindPropagationPolicy(qPropagationPolicy, qhkPropagationPolicy, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceVersion, qhkResourceVersion, _ := qs.GetOK("resourceVersion")
	if err := o.bindResourceVersion(qResourceVersion, qhkResourceVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceVersionMatch, qhkResourceVersionMatch, _ := qs.GetOK("resourceVersionMatch")
	if err := o.bindResourceVersionMatch(qResourceVersionMatch, qhkResourceVersionMatch, route.Formats); err != nil {
		res = append(res, err)
	}

	qTimeoutSeconds, qhkTimeoutSeconds, _ := qs.GetOK("timeoutSeconds")
	if err := o.bindTimeoutSeconds(qTimeoutSeconds, qhkTimeoutSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindContinue binds and validates parameter Continue from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindContinue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Continue = &raw

	if err := o.validateContinue(formats); err != nil {
		return err
	}

	return nil
}

// validateContinue carries on validations for parameter Continue
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateContinue(formats strfmt.Registry) error {

	return nil
}

// bindDryRun binds and validates parameter DryRun from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindDryRun(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DryRun = &raw

	if err := o.validateDryRun(formats); err != nil {
		return err
	}

	return nil
}

// validateDryRun carries on validations for parameter DryRun
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateDryRun(formats strfmt.Registry) error {

	return nil
}

// bindFieldSelector binds and validates parameter FieldSelector from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindFieldSelector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FieldSelector = &raw

	if err := o.validateFieldSelector(formats); err != nil {
		return err
	}

	return nil
}

// validateFieldSelector carries on validations for parameter FieldSelector
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateFieldSelector(formats strfmt.Registry) error {

	return nil
}

// bindGracePeriodSeconds binds and validates parameter GracePeriodSeconds from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindGracePeriodSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("gracePeriodSeconds", "query", "int64", raw)
	}
	o.GracePeriodSeconds = &value

	if err := o.validateGracePeriodSeconds(formats); err != nil {
		return err
	}

	return nil
}

// validateGracePeriodSeconds carries on validations for parameter GracePeriodSeconds
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateGracePeriodSeconds(formats strfmt.Registry) error {

	return nil
}

// bindLabelSelector binds and validates parameter LabelSelector from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindLabelSelector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LabelSelector = &raw

	if err := o.validateLabelSelector(formats); err != nil {
		return err
	}

	return nil
}

// validateLabelSelector carries on validations for parameter LabelSelector
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateLabelSelector(formats strfmt.Registry) error {

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateLimit(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	if err := o.validateNamespace(formats); err != nil {
		return err
	}

	return nil
}

// validateNamespace carries on validations for parameter Namespace
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindOrphanDependents binds and validates parameter OrphanDependents from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindOrphanDependents(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("orphanDependents", "query", "bool", raw)
	}
	o.OrphanDependents = &value

	if err := o.validateOrphanDependents(formats); err != nil {
		return err
	}

	return nil
}

// validateOrphanDependents carries on validations for parameter OrphanDependents
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateOrphanDependents(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Pretty = &raw

	if err := o.validatePretty(formats); err != nil {
		return err
	}

	return nil
}

// validatePretty carries on validations for parameter Pretty
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validatePretty(formats strfmt.Registry) error {

	return nil
}

// bindPropagationPolicy binds and validates parameter PropagationPolicy from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindPropagationPolicy(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PropagationPolicy = &raw

	if err := o.validatePropagationPolicy(formats); err != nil {
		return err
	}

	return nil
}

// validatePropagationPolicy carries on validations for parameter PropagationPolicy
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validatePropagationPolicy(formats strfmt.Registry) error {

	return nil
}

// bindResourceVersion binds and validates parameter ResourceVersion from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindResourceVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ResourceVersion = &raw

	if err := o.validateResourceVersion(formats); err != nil {
		return err
	}

	return nil
}

// validateResourceVersion carries on validations for parameter ResourceVersion
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateResourceVersion(formats strfmt.Registry) error {

	return nil
}

// bindResourceVersionMatch binds and validates parameter ResourceVersionMatch from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindResourceVersionMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ResourceVersionMatch = &raw

	if err := o.validateResourceVersionMatch(formats); err != nil {
		return err
	}

	return nil
}

// validateResourceVersionMatch carries on validations for parameter ResourceVersionMatch
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateResourceVersionMatch(formats strfmt.Registry) error {

	return nil
}

// bindTimeoutSeconds binds and validates parameter TimeoutSeconds from query.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) bindTimeoutSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("timeoutSeconds", "query", "int64", raw)
	}
	o.TimeoutSeconds = &value

	if err := o.validateTimeoutSeconds(formats); err != nil {
		return err
	}

	return nil
}

// validateTimeoutSeconds carries on validations for parameter TimeoutSeconds
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) validateTimeoutSeconds(formats strfmt.Registry) error {

	return nil
}
