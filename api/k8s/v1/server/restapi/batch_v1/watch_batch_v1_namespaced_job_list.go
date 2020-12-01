// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchBatchV1NamespacedJobListHandlerFunc turns a function with the right signature into a watch batch v1 namespaced job list handler
type WatchBatchV1NamespacedJobListHandlerFunc func(WatchBatchV1NamespacedJobListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV1NamespacedJobListHandlerFunc) Handle(params WatchBatchV1NamespacedJobListParams) middleware.Responder {
	return fn(params)
}

// WatchBatchV1NamespacedJobListHandler interface for that can handle valid watch batch v1 namespaced job list params
type WatchBatchV1NamespacedJobListHandler interface {
	Handle(WatchBatchV1NamespacedJobListParams) middleware.Responder
}

// NewWatchBatchV1NamespacedJobList creates a new http.Handler for the watch batch v1 namespaced job list operation
func NewWatchBatchV1NamespacedJobList(ctx *middleware.Context, handler WatchBatchV1NamespacedJobListHandler) *WatchBatchV1NamespacedJobList {
	return &WatchBatchV1NamespacedJobList{Context: ctx, Handler: handler}
}

/*WatchBatchV1NamespacedJobList swagger:route GET /apis/batch/v1/watch/namespaces/{namespace}/jobs batch_v1 watchBatchV1NamespacedJobList

watch individual changes to a list of Job. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchBatchV1NamespacedJobList struct {
	Context *middleware.Context
	Handler WatchBatchV1NamespacedJobListHandler
}

func (o *WatchBatchV1NamespacedJobList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV1NamespacedJobListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
