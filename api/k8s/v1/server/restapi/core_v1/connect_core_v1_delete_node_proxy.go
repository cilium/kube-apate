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

// ConnectCoreV1DeleteNodeProxyHandlerFunc turns a function with the right signature into a connect core v1 delete node proxy handler
type ConnectCoreV1DeleteNodeProxyHandlerFunc func(ConnectCoreV1DeleteNodeProxyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1DeleteNodeProxyHandlerFunc) Handle(params ConnectCoreV1DeleteNodeProxyParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1DeleteNodeProxyHandler interface for that can handle valid connect core v1 delete node proxy params
type ConnectCoreV1DeleteNodeProxyHandler interface {
	Handle(ConnectCoreV1DeleteNodeProxyParams) middleware.Responder
}

// NewConnectCoreV1DeleteNodeProxy creates a new http.Handler for the connect core v1 delete node proxy operation
func NewConnectCoreV1DeleteNodeProxy(ctx *middleware.Context, handler ConnectCoreV1DeleteNodeProxyHandler) *ConnectCoreV1DeleteNodeProxy {
	return &ConnectCoreV1DeleteNodeProxy{Context: ctx, Handler: handler}
}

/*ConnectCoreV1DeleteNodeProxy swagger:route DELETE /api/v1/nodes/{name}/proxy core_v1 connectCoreV1DeleteNodeProxy

connect DELETE requests to proxy of Node

*/
type ConnectCoreV1DeleteNodeProxy struct {
	Context *middleware.Context
	Handler ConnectCoreV1DeleteNodeProxyHandler
}

func (o *ConnectCoreV1DeleteNodeProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1DeleteNodeProxyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}