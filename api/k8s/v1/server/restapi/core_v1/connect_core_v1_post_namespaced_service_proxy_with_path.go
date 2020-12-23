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

// ConnectCoreV1PostNamespacedServiceProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 post namespaced service proxy with path handler
type ConnectCoreV1PostNamespacedServiceProxyWithPathHandlerFunc func(ConnectCoreV1PostNamespacedServiceProxyWithPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1PostNamespacedServiceProxyWithPathHandlerFunc) Handle(params ConnectCoreV1PostNamespacedServiceProxyWithPathParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1PostNamespacedServiceProxyWithPathHandler interface for that can handle valid connect core v1 post namespaced service proxy with path params
type ConnectCoreV1PostNamespacedServiceProxyWithPathHandler interface {
	Handle(ConnectCoreV1PostNamespacedServiceProxyWithPathParams) middleware.Responder
}

// NewConnectCoreV1PostNamespacedServiceProxyWithPath creates a new http.Handler for the connect core v1 post namespaced service proxy with path operation
func NewConnectCoreV1PostNamespacedServiceProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1PostNamespacedServiceProxyWithPathHandler) *ConnectCoreV1PostNamespacedServiceProxyWithPath {
	return &ConnectCoreV1PostNamespacedServiceProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1PostNamespacedServiceProxyWithPath swagger:route POST /api/v1/namespaces/{namespace}/services/{name}/proxy/{path} core_v1 connectCoreV1PostNamespacedServiceProxyWithPath

connect POST requests to proxy of Service

*/
type ConnectCoreV1PostNamespacedServiceProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1PostNamespacedServiceProxyWithPathHandler
}

func (o *ConnectCoreV1PostNamespacedServiceProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1PostNamespacedServiceProxyWithPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}