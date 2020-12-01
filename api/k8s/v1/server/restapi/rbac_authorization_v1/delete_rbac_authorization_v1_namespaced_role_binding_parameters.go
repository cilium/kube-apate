// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1

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

// NewDeleteRbacAuthorizationV1NamespacedRoleBindingParams creates a new DeleteRbacAuthorizationV1NamespacedRoleBindingParams object
// no default values defined in spec.
func NewDeleteRbacAuthorizationV1NamespacedRoleBindingParams() DeleteRbacAuthorizationV1NamespacedRoleBindingParams {

	return DeleteRbacAuthorizationV1NamespacedRoleBindingParams{}
}

// DeleteRbacAuthorizationV1NamespacedRoleBindingParams contains all the bound params for the delete rbac authorization v1 namespaced role binding operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteRbacAuthorizationV1NamespacedRoleBinding
type DeleteRbacAuthorizationV1NamespacedRoleBindingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
	/*When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	  Unique: true
	  In: query
	*/
	DryRun *string
	/*The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	  Unique: true
	  In: query
	*/
	GracePeriodSeconds *int64
	/*name of the RoleBinding
	  Required: true
	  Unique: true
	  In: path
	*/
	Name string
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
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteRbacAuthorizationV1NamespacedRoleBindingParams() beforehand.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
	qDryRun, qhkDryRun, _ := qs.GetOK("dryRun")
	if err := o.bindDryRun(qDryRun, qhkDryRun, route.Formats); err != nil {
		res = append(res, err)
	}

	qGracePeriodSeconds, qhkGracePeriodSeconds, _ := qs.GetOK("gracePeriodSeconds")
	if err := o.bindGracePeriodSeconds(qGracePeriodSeconds, qhkGracePeriodSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDryRun binds and validates parameter DryRun from query.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindDryRun(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validateDryRun(formats strfmt.Registry) error {

	return nil
}

// bindGracePeriodSeconds binds and validates parameter GracePeriodSeconds from query.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindGracePeriodSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validateGracePeriodSeconds(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Name = raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindOrphanDependents binds and validates parameter OrphanDependents from query.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindOrphanDependents(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validateOrphanDependents(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validatePretty(formats strfmt.Registry) error {

	return nil
}

// bindPropagationPolicy binds and validates parameter PropagationPolicy from query.
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) bindPropagationPolicy(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteRbacAuthorizationV1NamespacedRoleBindingParams) validatePropagationPolicy(formats strfmt.Registry) error {

	return nil
}
