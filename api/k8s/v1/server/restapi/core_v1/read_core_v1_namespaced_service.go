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

// ReadCoreV1NamespacedServiceHandlerFunc turns a function with the right signature into a read core v1 namespaced service handler
type ReadCoreV1NamespacedServiceHandlerFunc func(ReadCoreV1NamespacedServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedServiceHandlerFunc) Handle(params ReadCoreV1NamespacedServiceParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespacedServiceHandler interface for that can handle valid read core v1 namespaced service params
type ReadCoreV1NamespacedServiceHandler interface {
	Handle(ReadCoreV1NamespacedServiceParams) middleware.Responder
}

// NewReadCoreV1NamespacedService creates a new http.Handler for the read core v1 namespaced service operation
func NewReadCoreV1NamespacedService(ctx *middleware.Context, handler ReadCoreV1NamespacedServiceHandler) *ReadCoreV1NamespacedService {
	return &ReadCoreV1NamespacedService{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedService swagger:route GET /api/v1/namespaces/{namespace}/services/{name} core_v1 readCoreV1NamespacedService

read the specified Service

*/
type ReadCoreV1NamespacedService struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedServiceHandler
}

func (o *ReadCoreV1NamespacedService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
