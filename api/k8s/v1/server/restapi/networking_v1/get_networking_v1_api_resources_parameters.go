// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetNetworkingV1APIResourcesParams creates a new GetNetworkingV1APIResourcesParams object
// no default values defined in spec.
func NewGetNetworkingV1APIResourcesParams() GetNetworkingV1APIResourcesParams {

	return GetNetworkingV1APIResourcesParams{}
}

// GetNetworkingV1APIResourcesParams contains all the bound params for the get networking v1 API resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters getNetworkingV1APIResources
type GetNetworkingV1APIResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNetworkingV1APIResourcesParams() beforehand.
func (o *GetNetworkingV1APIResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}