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

// WatchCoreV1ReplicationControllerListForAllNamespacesHandlerFunc turns a function with the right signature into a watch core v1 replication controller list for all namespaces handler
type WatchCoreV1ReplicationControllerListForAllNamespacesHandlerFunc func(WatchCoreV1ReplicationControllerListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1ReplicationControllerListForAllNamespacesHandlerFunc) Handle(params WatchCoreV1ReplicationControllerListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1ReplicationControllerListForAllNamespacesHandler interface for that can handle valid watch core v1 replication controller list for all namespaces params
type WatchCoreV1ReplicationControllerListForAllNamespacesHandler interface {
	Handle(WatchCoreV1ReplicationControllerListForAllNamespacesParams) middleware.Responder
}

// NewWatchCoreV1ReplicationControllerListForAllNamespaces creates a new http.Handler for the watch core v1 replication controller list for all namespaces operation
func NewWatchCoreV1ReplicationControllerListForAllNamespaces(ctx *middleware.Context, handler WatchCoreV1ReplicationControllerListForAllNamespacesHandler) *WatchCoreV1ReplicationControllerListForAllNamespaces {
	return &WatchCoreV1ReplicationControllerListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchCoreV1ReplicationControllerListForAllNamespaces swagger:route GET /api/v1/watch/replicationcontrollers core_v1 watchCoreV1ReplicationControllerListForAllNamespaces

watch individual changes to a list of ReplicationController. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1ReplicationControllerListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchCoreV1ReplicationControllerListForAllNamespacesHandler
}

func (o *WatchCoreV1ReplicationControllerListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1ReplicationControllerListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
