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

// ReplaceBatchV1beta1NamespacedCronJobStatusHandlerFunc turns a function with the right signature into a replace batch v1beta1 namespaced cron job status handler
type ReplaceBatchV1beta1NamespacedCronJobStatusHandlerFunc func(ReplaceBatchV1beta1NamespacedCronJobStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceBatchV1beta1NamespacedCronJobStatusHandlerFunc) Handle(params ReplaceBatchV1beta1NamespacedCronJobStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceBatchV1beta1NamespacedCronJobStatusHandler interface for that can handle valid replace batch v1beta1 namespaced cron job status params
type ReplaceBatchV1beta1NamespacedCronJobStatusHandler interface {
	Handle(ReplaceBatchV1beta1NamespacedCronJobStatusParams) middleware.Responder
}

// NewReplaceBatchV1beta1NamespacedCronJobStatus creates a new http.Handler for the replace batch v1beta1 namespaced cron job status operation
func NewReplaceBatchV1beta1NamespacedCronJobStatus(ctx *middleware.Context, handler ReplaceBatchV1beta1NamespacedCronJobStatusHandler) *ReplaceBatchV1beta1NamespacedCronJobStatus {
	return &ReplaceBatchV1beta1NamespacedCronJobStatus{Context: ctx, Handler: handler}
}

/*ReplaceBatchV1beta1NamespacedCronJobStatus swagger:route PUT /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}/status batch_v1beta1 replaceBatchV1beta1NamespacedCronJobStatus

replace status of the specified CronJob

*/
type ReplaceBatchV1beta1NamespacedCronJobStatus struct {
	Context *middleware.Context
	Handler ReplaceBatchV1beta1NamespacedCronJobStatusHandler
}

func (o *ReplaceBatchV1beta1NamespacedCronJobStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceBatchV1beta1NamespacedCronJobStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
