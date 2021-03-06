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

// WatchCoreV1EndpointsListForAllNamespacesHandlerFunc turns a function with the right signature into a watch core v1 endpoints list for all namespaces handler
type WatchCoreV1EndpointsListForAllNamespacesHandlerFunc func(WatchCoreV1EndpointsListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1EndpointsListForAllNamespacesHandlerFunc) Handle(params WatchCoreV1EndpointsListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1EndpointsListForAllNamespacesHandler interface for that can handle valid watch core v1 endpoints list for all namespaces params
type WatchCoreV1EndpointsListForAllNamespacesHandler interface {
	Handle(WatchCoreV1EndpointsListForAllNamespacesParams) middleware.Responder
}

// NewWatchCoreV1EndpointsListForAllNamespaces creates a new http.Handler for the watch core v1 endpoints list for all namespaces operation
func NewWatchCoreV1EndpointsListForAllNamespaces(ctx *middleware.Context, handler WatchCoreV1EndpointsListForAllNamespacesHandler) *WatchCoreV1EndpointsListForAllNamespaces {
	return &WatchCoreV1EndpointsListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchCoreV1EndpointsListForAllNamespaces swagger:route GET /api/v1/watch/endpoints core_v1 watchCoreV1EndpointsListForAllNamespaces

watch individual changes to a list of Endpoints. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1EndpointsListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchCoreV1EndpointsListForAllNamespacesHandler
}

func (o *WatchCoreV1EndpointsListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1EndpointsListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
