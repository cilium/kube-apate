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

// CreateBatchV2alpha1NamespacedCronJobHandlerFunc turns a function with the right signature into a create batch v2alpha1 namespaced cron job handler
type CreateBatchV2alpha1NamespacedCronJobHandlerFunc func(CreateBatchV2alpha1NamespacedCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBatchV2alpha1NamespacedCronJobHandlerFunc) Handle(params CreateBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
	return fn(params)
}

// CreateBatchV2alpha1NamespacedCronJobHandler interface for that can handle valid create batch v2alpha1 namespaced cron job params
type CreateBatchV2alpha1NamespacedCronJobHandler interface {
	Handle(CreateBatchV2alpha1NamespacedCronJobParams) middleware.Responder
}

// NewCreateBatchV2alpha1NamespacedCronJob creates a new http.Handler for the create batch v2alpha1 namespaced cron job operation
func NewCreateBatchV2alpha1NamespacedCronJob(ctx *middleware.Context, handler CreateBatchV2alpha1NamespacedCronJobHandler) *CreateBatchV2alpha1NamespacedCronJob {
	return &CreateBatchV2alpha1NamespacedCronJob{Context: ctx, Handler: handler}
}

/*CreateBatchV2alpha1NamespacedCronJob swagger:route POST /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs batch_v2alpha1 createBatchV2alpha1NamespacedCronJob

create a CronJob

*/
type CreateBatchV2alpha1NamespacedCronJob struct {
	Context *middleware.Context
	Handler CreateBatchV2alpha1NamespacedCronJobHandler
}

func (o *CreateBatchV2alpha1NamespacedCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateBatchV2alpha1NamespacedCronJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
