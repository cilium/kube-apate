// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewPostApisCiliumIoV2CiliumEndpointParams creates a new PostApisCiliumIoV2CiliumEndpointParams object
// no default values defined in spec.
func NewPostApisCiliumIoV2CiliumEndpointParams() PostApisCiliumIoV2CiliumEndpointParams {

	return PostApisCiliumIoV2CiliumEndpointParams{}
}

// PostApisCiliumIoV2CiliumEndpointParams contains all the bound params for the post apis cilium io v2 cilium endpoint operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostApisCiliumIoV2CiliumEndpoint
type PostApisCiliumIoV2CiliumEndpointParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*namespace of the Cilium Endpoint
	  Required: true
	  Unique: true
	  In: path
	*/
	Namespace string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostApisCiliumIoV2CiliumEndpointParams() beforehand.
func (o *PostApisCiliumIoV2CiliumEndpointParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *PostApisCiliumIoV2CiliumEndpointParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PostApisCiliumIoV2CiliumEndpointParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}
