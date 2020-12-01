// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCoreV1NamespacedEndpointsHandlerFunc turns a function with the right signature into a read core v1 namespaced endpoints handler
type ReadCoreV1NamespacedEndpointsHandlerFunc func(ReadCoreV1NamespacedEndpointsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedEndpointsHandlerFunc) Handle(params ReadCoreV1NamespacedEndpointsParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespacedEndpointsHandler interface for that can handle valid read core v1 namespaced endpoints params
type ReadCoreV1NamespacedEndpointsHandler interface {
	Handle(ReadCoreV1NamespacedEndpointsParams) middleware.Responder
}

// NewReadCoreV1NamespacedEndpoints creates a new http.Handler for the read core v1 namespaced endpoints operation
func NewReadCoreV1NamespacedEndpoints(ctx *middleware.Context, handler ReadCoreV1NamespacedEndpointsHandler) *ReadCoreV1NamespacedEndpoints {
	return &ReadCoreV1NamespacedEndpoints{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedEndpoints swagger:route GET /api/v1/namespaces/{namespace}/endpoints/{name} core_v1 readCoreV1NamespacedEndpoints

read the specified Endpoints

*/
type ReadCoreV1NamespacedEndpoints struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedEndpointsHandler
}

func (o *ReadCoreV1NamespacedEndpoints) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedEndpointsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
