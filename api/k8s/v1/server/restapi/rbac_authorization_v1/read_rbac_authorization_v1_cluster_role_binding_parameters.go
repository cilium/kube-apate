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
)

// NewReadRbacAuthorizationV1ClusterRoleBindingParams creates a new ReadRbacAuthorizationV1ClusterRoleBindingParams object
// no default values defined in spec.
func NewReadRbacAuthorizationV1ClusterRoleBindingParams() ReadRbacAuthorizationV1ClusterRoleBindingParams {

	return ReadRbacAuthorizationV1ClusterRoleBindingParams{}
}

// ReadRbacAuthorizationV1ClusterRoleBindingParams contains all the bound params for the read rbac authorization v1 cluster role binding operation
// typically these are obtained from a http.Request
//
// swagger:parameters readRbacAuthorizationV1ClusterRoleBinding
type ReadRbacAuthorizationV1ClusterRoleBindingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*name of the ClusterRoleBinding
	  Required: true
	  Unique: true
	  In: path
	*/
	Name string
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReadRbacAuthorizationV1ClusterRoleBindingParams() beforehand.
func (o *ReadRbacAuthorizationV1ClusterRoleBindingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qPretty, qhkPretty, _ := qs.GetOK("pretty")
	if err := o.bindPretty(qPretty, qhkPretty, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ReadRbacAuthorizationV1ClusterRoleBindingParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadRbacAuthorizationV1ClusterRoleBindingParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *ReadRbacAuthorizationV1ClusterRoleBindingParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadRbacAuthorizationV1ClusterRoleBindingParams) validatePretty(formats strfmt.Registry) error {

	return nil
}
