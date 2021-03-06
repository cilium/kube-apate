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

// WatchCoreV1NamespaceListHandlerFunc turns a function with the right signature into a watch core v1 namespace list handler
type WatchCoreV1NamespaceListHandlerFunc func(WatchCoreV1NamespaceListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespaceListHandlerFunc) Handle(params WatchCoreV1NamespaceListParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespaceListHandler interface for that can handle valid watch core v1 namespace list params
type WatchCoreV1NamespaceListHandler interface {
	Handle(WatchCoreV1NamespaceListParams) middleware.Responder
}

// NewWatchCoreV1NamespaceList creates a new http.Handler for the watch core v1 namespace list operation
func NewWatchCoreV1NamespaceList(ctx *middleware.Context, handler WatchCoreV1NamespaceListHandler) *WatchCoreV1NamespaceList {
	return &WatchCoreV1NamespaceList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespaceList swagger:route GET /api/v1/watch/namespaces core_v1 watchCoreV1NamespaceList

watch individual changes to a list of Namespace. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespaceList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespaceListHandler
}

func (o *WatchCoreV1NamespaceList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespaceListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
