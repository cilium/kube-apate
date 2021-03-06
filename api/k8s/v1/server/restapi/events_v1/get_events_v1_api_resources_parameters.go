// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetEventsV1APIResourcesParams creates a new GetEventsV1APIResourcesParams object
// no default values defined in spec.
func NewGetEventsV1APIResourcesParams() GetEventsV1APIResourcesParams {

	return GetEventsV1APIResourcesParams{}
}

// GetEventsV1APIResourcesParams contains all the bound params for the get events v1 API resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters getEventsV1APIResources
type GetEventsV1APIResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEventsV1APIResourcesParams() beforehand.
func (o *GetEventsV1APIResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
