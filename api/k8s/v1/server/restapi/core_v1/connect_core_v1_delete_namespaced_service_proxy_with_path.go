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

// ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 delete namespaced service proxy with path handler
type ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandlerFunc func(ConnectCoreV1DeleteNamespacedServiceProxyWithPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandlerFunc) Handle(params ConnectCoreV1DeleteNamespacedServiceProxyWithPathParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandler interface for that can handle valid connect core v1 delete namespaced service proxy with path params
type ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandler interface {
	Handle(ConnectCoreV1DeleteNamespacedServiceProxyWithPathParams) middleware.Responder
}

// NewConnectCoreV1DeleteNamespacedServiceProxyWithPath creates a new http.Handler for the connect core v1 delete namespaced service proxy with path operation
func NewConnectCoreV1DeleteNamespacedServiceProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandler) *ConnectCoreV1DeleteNamespacedServiceProxyWithPath {
	return &ConnectCoreV1DeleteNamespacedServiceProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1DeleteNamespacedServiceProxyWithPath swagger:route DELETE /api/v1/namespaces/{namespace}/services/{name}/proxy/{path} core_v1 connectCoreV1DeleteNamespacedServiceProxyWithPath

connect DELETE requests to proxy of Service

*/
type ConnectCoreV1DeleteNamespacedServiceProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandler
}

func (o *ConnectCoreV1DeleteNamespacedServiceProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1DeleteNamespacedServiceProxyWithPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
