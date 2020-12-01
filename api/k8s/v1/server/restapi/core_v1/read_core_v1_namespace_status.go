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

// ReadCoreV1NamespaceStatusHandlerFunc turns a function with the right signature into a read core v1 namespace status handler
type ReadCoreV1NamespaceStatusHandlerFunc func(ReadCoreV1NamespaceStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespaceStatusHandlerFunc) Handle(params ReadCoreV1NamespaceStatusParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespaceStatusHandler interface for that can handle valid read core v1 namespace status params
type ReadCoreV1NamespaceStatusHandler interface {
	Handle(ReadCoreV1NamespaceStatusParams) middleware.Responder
}

// NewReadCoreV1NamespaceStatus creates a new http.Handler for the read core v1 namespace status operation
func NewReadCoreV1NamespaceStatus(ctx *middleware.Context, handler ReadCoreV1NamespaceStatusHandler) *ReadCoreV1NamespaceStatus {
	return &ReadCoreV1NamespaceStatus{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespaceStatus swagger:route GET /api/v1/namespaces/{name}/status core_v1 readCoreV1NamespaceStatus

read status of the specified Namespace

*/
type ReadCoreV1NamespaceStatus struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespaceStatusHandler
}

func (o *ReadCoreV1NamespaceStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespaceStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
