// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetApiregistrationV1beta1APIResourcesHandlerFunc turns a function with the right signature into a get apiregistration v1beta1 API resources handler
type GetApiregistrationV1beta1APIResourcesHandlerFunc func(GetApiregistrationV1beta1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApiregistrationV1beta1APIResourcesHandlerFunc) Handle(params GetApiregistrationV1beta1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetApiregistrationV1beta1APIResourcesHandler interface for that can handle valid get apiregistration v1beta1 API resources params
type GetApiregistrationV1beta1APIResourcesHandler interface {
	Handle(GetApiregistrationV1beta1APIResourcesParams) middleware.Responder
}

// NewGetApiregistrationV1beta1APIResources creates a new http.Handler for the get apiregistration v1beta1 API resources operation
func NewGetApiregistrationV1beta1APIResources(ctx *middleware.Context, handler GetApiregistrationV1beta1APIResourcesHandler) *GetApiregistrationV1beta1APIResources {
	return &GetApiregistrationV1beta1APIResources{Context: ctx, Handler: handler}
}

/*GetApiregistrationV1beta1APIResources swagger:route GET /apis/apiregistration.k8s.io/v1beta1/ apiregistration_v1beta1 getApiregistrationV1beta1ApiResources

get available resources

*/
type GetApiregistrationV1beta1APIResources struct {
	Context *middleware.Context
	Handler GetApiregistrationV1beta1APIResourcesHandler
}

func (o *GetApiregistrationV1beta1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetApiregistrationV1beta1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
