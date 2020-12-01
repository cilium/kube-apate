// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetNodeAPIGroupParams creates a new GetNodeAPIGroupParams object
// no default values defined in spec.
func NewGetNodeAPIGroupParams() GetNodeAPIGroupParams {

	return GetNodeAPIGroupParams{}
}

// GetNodeAPIGroupParams contains all the bound params for the get node API group operation
// typically these are obtained from a http.Request
//
// swagger:parameters getNodeAPIGroup
type GetNodeAPIGroupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNodeAPIGroupParams() beforehand.
func (o *GetNodeAPIGroupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
