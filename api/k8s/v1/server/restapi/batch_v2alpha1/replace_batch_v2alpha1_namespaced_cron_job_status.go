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

// ReplaceBatchV2alpha1NamespacedCronJobStatusHandlerFunc turns a function with the right signature into a replace batch v2alpha1 namespaced cron job status handler
type ReplaceBatchV2alpha1NamespacedCronJobStatusHandlerFunc func(ReplaceBatchV2alpha1NamespacedCronJobStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceBatchV2alpha1NamespacedCronJobStatusHandlerFunc) Handle(params ReplaceBatchV2alpha1NamespacedCronJobStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceBatchV2alpha1NamespacedCronJobStatusHandler interface for that can handle valid replace batch v2alpha1 namespaced cron job status params
type ReplaceBatchV2alpha1NamespacedCronJobStatusHandler interface {
	Handle(ReplaceBatchV2alpha1NamespacedCronJobStatusParams) middleware.Responder
}

// NewReplaceBatchV2alpha1NamespacedCronJobStatus creates a new http.Handler for the replace batch v2alpha1 namespaced cron job status operation
func NewReplaceBatchV2alpha1NamespacedCronJobStatus(ctx *middleware.Context, handler ReplaceBatchV2alpha1NamespacedCronJobStatusHandler) *ReplaceBatchV2alpha1NamespacedCronJobStatus {
	return &ReplaceBatchV2alpha1NamespacedCronJobStatus{Context: ctx, Handler: handler}
}

/*ReplaceBatchV2alpha1NamespacedCronJobStatus swagger:route PUT /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status batch_v2alpha1 replaceBatchV2alpha1NamespacedCronJobStatus

replace status of the specified CronJob

*/
type ReplaceBatchV2alpha1NamespacedCronJobStatus struct {
	Context *middleware.Context
	Handler ReplaceBatchV2alpha1NamespacedCronJobStatusHandler
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceBatchV2alpha1NamespacedCronJobStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}