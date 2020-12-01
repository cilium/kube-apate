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
	"github.com/go-openapi/swag"
)

// NewConnectCoreV1GetNamespacedPodExecParams creates a new ConnectCoreV1GetNamespacedPodExecParams object
// no default values defined in spec.
func NewConnectCoreV1GetNamespacedPodExecParams() ConnectCoreV1GetNamespacedPodExecParams {

	return ConnectCoreV1GetNamespacedPodExecParams{}
}

// ConnectCoreV1GetNamespacedPodExecParams contains all the bound params for the connect core v1 get namespaced pod exec operation
// typically these are obtained from a http.Request
//
// swagger:parameters connectCoreV1GetNamespacedPodExec
type ConnectCoreV1GetNamespacedPodExecParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Command is the remote command to execute. argv array. Not executed within a shell.
	  Unique: true
	  In: query
	*/
	Command *string
	/*Container in which to execute the command. Defaults to only container if there is only one container in the pod.
	  Unique: true
	  In: query
	*/
	Container *string
	/*name of the PodExecOptions
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
	/*Redirect the standard error stream of the pod for this call. Defaults to true.
	  Unique: true
	  In: query
	*/
	Stderr *bool
	/*Redirect the standard input stream of the pod for this call. Defaults to false.
	  Unique: true
	  In: query
	*/
	Stdin *bool
	/*Redirect the standard output stream of the pod for this call. Defaults to true.
	  Unique: true
	  In: query
	*/
	Stdout *bool
	/*TTY if true indicates that a tty will be allocated for the exec call. Defaults to false.
	  Unique: true
	  In: query
	*/
	Tty *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewConnectCoreV1GetNamespacedPodExecParams() beforehand.
func (o *ConnectCoreV1GetNamespacedPodExecParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCommand, qhkCommand, _ := qs.GetOK("command")
	if err := o.bindCommand(qCommand, qhkCommand, route.Formats); err != nil {
		res = append(res, err)
	}

	qContainer, qhkContainer, _ := qs.GetOK("container")
	if err := o.bindContainer(qContainer, qhkContainer, route.Formats); err != nil {
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

	qStderr, qhkStderr, _ := qs.GetOK("stderr")
	if err := o.bindStderr(qStderr, qhkStderr, route.Formats); err != nil {
		res = append(res, err)
	}

	qStdin, qhkStdin, _ := qs.GetOK("stdin")
	if err := o.bindStdin(qStdin, qhkStdin, route.Formats); err != nil {
		res = append(res, err)
	}

	qStdout, qhkStdout, _ := qs.GetOK("stdout")
	if err := o.bindStdout(qStdout, qhkStdout, route.Formats); err != nil {
		res = append(res, err)
	}

	qTty, qhkTty, _ := qs.GetOK("tty")
	if err := o.bindTty(qTty, qhkTty, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCommand binds and validates parameter Command from query.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindCommand(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Command = &raw

	if err := o.validateCommand(formats); err != nil {
		return err
	}

	return nil
}

// validateCommand carries on validations for parameter Command
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateCommand(formats strfmt.Registry) error {

	return nil
}

// bindContainer binds and validates parameter Container from query.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindContainer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Container = &raw

	if err := o.validateContainer(formats); err != nil {
		return err
	}

	return nil
}

// validateContainer carries on validations for parameter Container
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateContainer(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindStderr binds and validates parameter Stderr from query.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindStderr(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("stderr", "query", "bool", raw)
	}
	o.Stderr = &value

	if err := o.validateStderr(formats); err != nil {
		return err
	}

	return nil
}

// validateStderr carries on validations for parameter Stderr
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateStderr(formats strfmt.Registry) error {

	return nil
}

// bindStdin binds and validates parameter Stdin from query.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindStdin(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("stdin", "query", "bool", raw)
	}
	o.Stdin = &value

	if err := o.validateStdin(formats); err != nil {
		return err
	}

	return nil
}

// validateStdin carries on validations for parameter Stdin
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateStdin(formats strfmt.Registry) error {

	return nil
}

// bindStdout binds and validates parameter Stdout from query.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindStdout(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("stdout", "query", "bool", raw)
	}
	o.Stdout = &value

	if err := o.validateStdout(formats); err != nil {
		return err
	}

	return nil
}

// validateStdout carries on validations for parameter Stdout
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateStdout(formats strfmt.Registry) error {

	return nil
}

// bindTty binds and validates parameter Tty from query.
func (o *ConnectCoreV1GetNamespacedPodExecParams) bindTty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("tty", "query", "bool", raw)
	}
	o.Tty = &value

	if err := o.validateTty(formats); err != nil {
		return err
	}

	return nil
}

// validateTty carries on validations for parameter Tty
func (o *ConnectCoreV1GetNamespacedPodExecParams) validateTty(formats strfmt.Registry) error {

	return nil
}
