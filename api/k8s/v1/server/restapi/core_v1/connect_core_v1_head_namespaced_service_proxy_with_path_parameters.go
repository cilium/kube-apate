// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewConnectCoreV1HeadNamespacedServiceProxyWithPathParams creates a new ConnectCoreV1HeadNamespacedServiceProxyWithPathParams object
// no default values defined in spec.
func NewConnectCoreV1HeadNamespacedServiceProxyWithPathParams() ConnectCoreV1HeadNamespacedServiceProxyWithPathParams {

	return ConnectCoreV1HeadNamespacedServiceProxyWithPathParams{}
}

// ConnectCoreV1HeadNamespacedServiceProxyWithPathParams contains all the bound params for the connect core v1 head namespaced service proxy with path operation
// typically these are obtained from a http.Request
//
// swagger:parameters connectCoreV1HeadNamespacedServiceProxyWithPath
type ConnectCoreV1HeadNamespacedServiceProxyWithPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*name of the ServiceProxyOptions
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
	/*path to the resource
	  Required: true
	  Unique: true
	  In: path
	*/
	PathPath string
	/*Path is the part of URLs that include service endpoints, suffixes, and parameters to use for the current proxy request to service. For example, the whole request URL is http://localhost/api/v1/namespaces/kube-system/services/elasticsearch-logging/_search?q=user:kimchy. Path is _search?q=user:kimchy.
	  Unique: true
	  In: query
	*/
	QueryPath *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewConnectCoreV1HeadNamespacedServiceProxyWithPathParams() beforehand.
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPathPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	qPath, qhkPath, _ := qs.GetOK("path")
	if err := o.bindQueryPath(qPath, qhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindPathPath binds and validates parameter PathPath from path.
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) bindPathPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PathPath = raw

	if err := o.validatePathPath(formats); err != nil {
		return err
	}

	return nil
}

// validatePathPath carries on validations for parameter PathPath
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) validatePathPath(formats strfmt.Registry) error {

	return nil
}

// bindQueryPath binds and validates parameter QueryPath from query.
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) bindQueryPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.QueryPath = &raw

	if err := o.validateQueryPath(formats); err != nil {
		return err
	}

	return nil
}

// validateQueryPath carries on validations for parameter QueryPath
func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) validateQueryPath(formats strfmt.Registry) error {

	return nil
}
