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

// DeleteBatchV1beta1NamespacedCronJobHandlerFunc turns a function with the right signature into a delete batch v1beta1 namespaced cron job handler
type DeleteBatchV1beta1NamespacedCronJobHandlerFunc func(DeleteBatchV1beta1NamespacedCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBatchV1beta1NamespacedCronJobHandlerFunc) Handle(params DeleteBatchV1beta1NamespacedCronJobParams) middleware.Responder {
	return fn(params)
}

// DeleteBatchV1beta1NamespacedCronJobHandler interface for that can handle valid delete batch v1beta1 namespaced cron job params
type DeleteBatchV1beta1NamespacedCronJobHandler interface {
	Handle(DeleteBatchV1beta1NamespacedCronJobParams) middleware.Responder
}

// NewDeleteBatchV1beta1NamespacedCronJob creates a new http.Handler for the delete batch v1beta1 namespaced cron job operation
func NewDeleteBatchV1beta1NamespacedCronJob(ctx *middleware.Context, handler DeleteBatchV1beta1NamespacedCronJobHandler) *DeleteBatchV1beta1NamespacedCronJob {
	return &DeleteBatchV1beta1NamespacedCronJob{Context: ctx, Handler: handler}
}

/*DeleteBatchV1beta1NamespacedCronJob swagger:route DELETE /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name} batch_v1beta1 deleteBatchV1beta1NamespacedCronJob

delete a CronJob

*/
type DeleteBatchV1beta1NamespacedCronJob struct {
	Context *middleware.Context
	Handler DeleteBatchV1beta1NamespacedCronJobHandler
}

func (o *DeleteBatchV1beta1NamespacedCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBatchV1beta1NamespacedCronJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
