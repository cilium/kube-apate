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

// ConnectCoreV1GetNamespacedServiceProxyHandlerFunc turns a function with the right signature into a connect core v1 get namespaced service proxy handler
type ConnectCoreV1GetNamespacedServiceProxyHandlerFunc func(ConnectCoreV1GetNamespacedServiceProxyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1GetNamespacedServiceProxyHandlerFunc) Handle(params ConnectCoreV1GetNamespacedServiceProxyParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1GetNamespacedServiceProxyHandler interface for that can handle valid connect core v1 get namespaced service proxy params
type ConnectCoreV1GetNamespacedServiceProxyHandler interface {
	Handle(ConnectCoreV1GetNamespacedServiceProxyParams) middleware.Responder
}

// NewConnectCoreV1GetNamespacedServiceProxy creates a new http.Handler for the connect core v1 get namespaced service proxy operation
func NewConnectCoreV1GetNamespacedServiceProxy(ctx *middleware.Context, handler ConnectCoreV1GetNamespacedServiceProxyHandler) *ConnectCoreV1GetNamespacedServiceProxy {
	return &ConnectCoreV1GetNamespacedServiceProxy{Context: ctx, Handler: handler}
}

/*ConnectCoreV1GetNamespacedServiceProxy swagger:route GET /api/v1/namespaces/{namespace}/services/{name}/proxy core_v1 connectCoreV1GetNamespacedServiceProxy

connect GET requests to proxy of Service

*/
type ConnectCoreV1GetNamespacedServiceProxy struct {
	Context *middleware.Context
	Handler ConnectCoreV1GetNamespacedServiceProxyHandler
}

func (o *ConnectCoreV1GetNamespacedServiceProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1GetNamespacedServiceProxyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}