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

// WatchBatchV1beta1NamespacedCronJobHandlerFunc turns a function with the right signature into a watch batch v1beta1 namespaced cron job handler
type WatchBatchV1beta1NamespacedCronJobHandlerFunc func(WatchBatchV1beta1NamespacedCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV1beta1NamespacedCronJobHandlerFunc) Handle(params WatchBatchV1beta1NamespacedCronJobParams) middleware.Responder {
	return fn(params)
}

// WatchBatchV1beta1NamespacedCronJobHandler interface for that can handle valid watch batch v1beta1 namespaced cron job params
type WatchBatchV1beta1NamespacedCronJobHandler interface {
	Handle(WatchBatchV1beta1NamespacedCronJobParams) middleware.Responder
}

// NewWatchBatchV1beta1NamespacedCronJob creates a new http.Handler for the watch batch v1beta1 namespaced cron job operation
func NewWatchBatchV1beta1NamespacedCronJob(ctx *middleware.Context, handler WatchBatchV1beta1NamespacedCronJobHandler) *WatchBatchV1beta1NamespacedCronJob {
	return &WatchBatchV1beta1NamespacedCronJob{Context: ctx, Handler: handler}
}

/*WatchBatchV1beta1NamespacedCronJob swagger:route GET /apis/batch/v1beta1/watch/namespaces/{namespace}/cronjobs/{name} batch_v1beta1 watchBatchV1beta1NamespacedCronJob

watch changes to an object of kind CronJob. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchBatchV1beta1NamespacedCronJob struct {
	Context *middleware.Context
	Handler WatchBatchV1beta1NamespacedCronJobHandler
}

func (o *WatchBatchV1beta1NamespacedCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV1beta1NamespacedCronJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
