// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAuthenticationV1APIResourcesHandlerFunc turns a function with the right signature into a get authentication v1 API resources handler
type GetAuthenticationV1APIResourcesHandlerFunc func(GetAuthenticationV1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAuthenticationV1APIResourcesHandlerFunc) Handle(params GetAuthenticationV1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetAuthenticationV1APIResourcesHandler interface for that can handle valid get authentication v1 API resources params
type GetAuthenticationV1APIResourcesHandler interface {
	Handle(GetAuthenticationV1APIResourcesParams) middleware.Responder
}

// NewGetAuthenticationV1APIResources creates a new http.Handler for the get authentication v1 API resources operation
func NewGetAuthenticationV1APIResources(ctx *middleware.Context, handler GetAuthenticationV1APIResourcesHandler) *GetAuthenticationV1APIResources {
	return &GetAuthenticationV1APIResources{Context: ctx, Handler: handler}
}

/*GetAuthenticationV1APIResources swagger:route GET /apis/authentication.k8s.io/v1/ authentication_v1 getAuthenticationV1ApiResources

get available resources

*/
type GetAuthenticationV1APIResources struct {
	Context *middleware.Context
	Handler GetAuthenticationV1APIResourcesHandler
}

func (o *GetAuthenticationV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAuthenticationV1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
