// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNodeV1alpha1RuntimeClassListHandlerFunc turns a function with the right signature into a watch node v1alpha1 runtime class list handler
type WatchNodeV1alpha1RuntimeClassListHandlerFunc func(WatchNodeV1alpha1RuntimeClassListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNodeV1alpha1RuntimeClassListHandlerFunc) Handle(params WatchNodeV1alpha1RuntimeClassListParams) middleware.Responder {
	return fn(params)
}

// WatchNodeV1alpha1RuntimeClassListHandler interface for that can handle valid watch node v1alpha1 runtime class list params
type WatchNodeV1alpha1RuntimeClassListHandler interface {
	Handle(WatchNodeV1alpha1RuntimeClassListParams) middleware.Responder
}

// NewWatchNodeV1alpha1RuntimeClassList creates a new http.Handler for the watch node v1alpha1 runtime class list operation
func NewWatchNodeV1alpha1RuntimeClassList(ctx *middleware.Context, handler WatchNodeV1alpha1RuntimeClassListHandler) *WatchNodeV1alpha1RuntimeClassList {
	return &WatchNodeV1alpha1RuntimeClassList{Context: ctx, Handler: handler}
}

/*WatchNodeV1alpha1RuntimeClassList swagger:route GET /apis/node.k8s.io/v1alpha1/watch/runtimeclasses node_v1alpha1 watchNodeV1alpha1RuntimeClassList

watch individual changes to a list of RuntimeClass. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchNodeV1alpha1RuntimeClassList struct {
	Context *middleware.Context
	Handler WatchNodeV1alpha1RuntimeClassListHandler
}

func (o *WatchNodeV1alpha1RuntimeClassList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNodeV1alpha1RuntimeClassListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
