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

// WatchCoreV1ServiceListForAllNamespacesHandlerFunc turns a function with the right signature into a watch core v1 service list for all namespaces handler
type WatchCoreV1ServiceListForAllNamespacesHandlerFunc func(WatchCoreV1ServiceListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1ServiceListForAllNamespacesHandlerFunc) Handle(params WatchCoreV1ServiceListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1ServiceListForAllNamespacesHandler interface for that can handle valid watch core v1 service list for all namespaces params
type WatchCoreV1ServiceListForAllNamespacesHandler interface {
	Handle(WatchCoreV1ServiceListForAllNamespacesParams) middleware.Responder
}

// NewWatchCoreV1ServiceListForAllNamespaces creates a new http.Handler for the watch core v1 service list for all namespaces operation
func NewWatchCoreV1ServiceListForAllNamespaces(ctx *middleware.Context, handler WatchCoreV1ServiceListForAllNamespacesHandler) *WatchCoreV1ServiceListForAllNamespaces {
	return &WatchCoreV1ServiceListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchCoreV1ServiceListForAllNamespaces swagger:route GET /api/v1/watch/services core_v1 watchCoreV1ServiceListForAllNamespaces

watch individual changes to a list of Service. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1ServiceListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchCoreV1ServiceListForAllNamespacesHandler
}

func (o *WatchCoreV1ServiceListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1ServiceListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
