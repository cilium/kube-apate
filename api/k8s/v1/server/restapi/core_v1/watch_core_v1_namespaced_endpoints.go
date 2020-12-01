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

// WatchCoreV1NamespacedEndpointsHandlerFunc turns a function with the right signature into a watch core v1 namespaced endpoints handler
type WatchCoreV1NamespacedEndpointsHandlerFunc func(WatchCoreV1NamespacedEndpointsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedEndpointsHandlerFunc) Handle(params WatchCoreV1NamespacedEndpointsParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedEndpointsHandler interface for that can handle valid watch core v1 namespaced endpoints params
type WatchCoreV1NamespacedEndpointsHandler interface {
	Handle(WatchCoreV1NamespacedEndpointsParams) middleware.Responder
}

// NewWatchCoreV1NamespacedEndpoints creates a new http.Handler for the watch core v1 namespaced endpoints operation
func NewWatchCoreV1NamespacedEndpoints(ctx *middleware.Context, handler WatchCoreV1NamespacedEndpointsHandler) *WatchCoreV1NamespacedEndpoints {
	return &WatchCoreV1NamespacedEndpoints{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedEndpoints swagger:route GET /api/v1/watch/namespaces/{namespace}/endpoints/{name} core_v1 watchCoreV1NamespacedEndpoints

watch changes to an object of kind Endpoints. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoreV1NamespacedEndpoints struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedEndpointsHandler
}

func (o *WatchCoreV1NamespacedEndpoints) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedEndpointsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
