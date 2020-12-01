// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchBatchV2alpha1CronJobListForAllNamespacesHandlerFunc turns a function with the right signature into a watch batch v2alpha1 cron job list for all namespaces handler
type WatchBatchV2alpha1CronJobListForAllNamespacesHandlerFunc func(WatchBatchV2alpha1CronJobListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV2alpha1CronJobListForAllNamespacesHandlerFunc) Handle(params WatchBatchV2alpha1CronJobListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchBatchV2alpha1CronJobListForAllNamespacesHandler interface for that can handle valid watch batch v2alpha1 cron job list for all namespaces params
type WatchBatchV2alpha1CronJobListForAllNamespacesHandler interface {
	Handle(WatchBatchV2alpha1CronJobListForAllNamespacesParams) middleware.Responder
}

// NewWatchBatchV2alpha1CronJobListForAllNamespaces creates a new http.Handler for the watch batch v2alpha1 cron job list for all namespaces operation
func NewWatchBatchV2alpha1CronJobListForAllNamespaces(ctx *middleware.Context, handler WatchBatchV2alpha1CronJobListForAllNamespacesHandler) *WatchBatchV2alpha1CronJobListForAllNamespaces {
	return &WatchBatchV2alpha1CronJobListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchBatchV2alpha1CronJobListForAllNamespaces swagger:route GET /apis/batch/v2alpha1/watch/cronjobs batch_v2alpha1 watchBatchV2alpha1CronJobListForAllNamespaces

watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchBatchV2alpha1CronJobListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchBatchV2alpha1CronJobListForAllNamespacesHandler
}

func (o *WatchBatchV2alpha1CronJobListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV2alpha1CronJobListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
