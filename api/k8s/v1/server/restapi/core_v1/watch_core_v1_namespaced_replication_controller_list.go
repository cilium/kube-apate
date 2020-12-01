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

// WatchCoreV1NamespacedReplicationControllerListHandlerFunc turns a function with the right signature into a watch core v1 namespaced replication controller list handler
type WatchCoreV1NamespacedReplicationControllerListHandlerFunc func(WatchCoreV1NamespacedReplicationControllerListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedReplicationControllerListHandlerFunc) Handle(params WatchCoreV1NamespacedReplicationControllerListParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedReplicationControllerListHandler interface for that can handle valid watch core v1 namespaced replication controller list params
type WatchCoreV1NamespacedReplicationControllerListHandler interface {
	Handle(WatchCoreV1NamespacedReplicationControllerListParams) middleware.Responder
}

// NewWatchCoreV1NamespacedReplicationControllerList creates a new http.Handler for the watch core v1 namespaced replication controller list operation
func NewWatchCoreV1NamespacedReplicationControllerList(ctx *middleware.Context, handler WatchCoreV1NamespacedReplicationControllerListHandler) *WatchCoreV1NamespacedReplicationControllerList {
	return &WatchCoreV1NamespacedReplicationControllerList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedReplicationControllerList swagger:route GET /api/v1/watch/namespaces/{namespace}/replicationcontrollers core_v1 watchCoreV1NamespacedReplicationControllerList

watch individual changes to a list of ReplicationController. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespacedReplicationControllerList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedReplicationControllerListHandler
}

func (o *WatchCoreV1NamespacedReplicationControllerList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedReplicationControllerListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
