// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchBatchV1beta1NamespacedCronJobListHandlerFunc turns a function with the right signature into a watch batch v1beta1 namespaced cron job list handler
type WatchBatchV1beta1NamespacedCronJobListHandlerFunc func(WatchBatchV1beta1NamespacedCronJobListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV1beta1NamespacedCronJobListHandlerFunc) Handle(params WatchBatchV1beta1NamespacedCronJobListParams) middleware.Responder {
	return fn(params)
}

// WatchBatchV1beta1NamespacedCronJobListHandler interface for that can handle valid watch batch v1beta1 namespaced cron job list params
type WatchBatchV1beta1NamespacedCronJobListHandler interface {
	Handle(WatchBatchV1beta1NamespacedCronJobListParams) middleware.Responder
}

// NewWatchBatchV1beta1NamespacedCronJobList creates a new http.Handler for the watch batch v1beta1 namespaced cron job list operation
func NewWatchBatchV1beta1NamespacedCronJobList(ctx *middleware.Context, handler WatchBatchV1beta1NamespacedCronJobListHandler) *WatchBatchV1beta1NamespacedCronJobList {
	return &WatchBatchV1beta1NamespacedCronJobList{Context: ctx, Handler: handler}
}

/*WatchBatchV1beta1NamespacedCronJobList swagger:route GET /apis/batch/v1beta1/watch/namespaces/{namespace}/cronjobs batch_v1beta1 watchBatchV1beta1NamespacedCronJobList

watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchBatchV1beta1NamespacedCronJobList struct {
	Context *middleware.Context
	Handler WatchBatchV1beta1NamespacedCronJobListHandler
}

func (o *WatchBatchV1beta1NamespacedCronJobList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV1beta1NamespacedCronJobListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
