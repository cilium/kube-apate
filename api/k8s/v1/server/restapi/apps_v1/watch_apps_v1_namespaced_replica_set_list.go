// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAppsV1NamespacedReplicaSetListHandlerFunc turns a function with the right signature into a watch apps v1 namespaced replica set list handler
type WatchAppsV1NamespacedReplicaSetListHandlerFunc func(WatchAppsV1NamespacedReplicaSetListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAppsV1NamespacedReplicaSetListHandlerFunc) Handle(params WatchAppsV1NamespacedReplicaSetListParams) middleware.Responder {
	return fn(params)
}

// WatchAppsV1NamespacedReplicaSetListHandler interface for that can handle valid watch apps v1 namespaced replica set list params
type WatchAppsV1NamespacedReplicaSetListHandler interface {
	Handle(WatchAppsV1NamespacedReplicaSetListParams) middleware.Responder
}

// NewWatchAppsV1NamespacedReplicaSetList creates a new http.Handler for the watch apps v1 namespaced replica set list operation
func NewWatchAppsV1NamespacedReplicaSetList(ctx *middleware.Context, handler WatchAppsV1NamespacedReplicaSetListHandler) *WatchAppsV1NamespacedReplicaSetList {
	return &WatchAppsV1NamespacedReplicaSetList{Context: ctx, Handler: handler}
}

/*WatchAppsV1NamespacedReplicaSetList swagger:route GET /apis/apps/v1/watch/namespaces/{namespace}/replicasets apps_v1 watchAppsV1NamespacedReplicaSetList

watch individual changes to a list of ReplicaSet. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAppsV1NamespacedReplicaSetList struct {
	Context *middleware.Context
	Handler WatchAppsV1NamespacedReplicaSetListHandler
}

func (o *WatchAppsV1NamespacedReplicaSetList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAppsV1NamespacedReplicaSetListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
