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

// PatchBatchV1NamespacedJobHandlerFunc turns a function with the right signature into a patch batch v1 namespaced job handler
type PatchBatchV1NamespacedJobHandlerFunc func(PatchBatchV1NamespacedJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchBatchV1NamespacedJobHandlerFunc) Handle(params PatchBatchV1NamespacedJobParams) middleware.Responder {
	return fn(params)
}

// PatchBatchV1NamespacedJobHandler interface for that can handle valid patch batch v1 namespaced job params
type PatchBatchV1NamespacedJobHandler interface {
	Handle(PatchBatchV1NamespacedJobParams) middleware.Responder
}

// NewPatchBatchV1NamespacedJob creates a new http.Handler for the patch batch v1 namespaced job operation
func NewPatchBatchV1NamespacedJob(ctx *middleware.Context, handler PatchBatchV1NamespacedJobHandler) *PatchBatchV1NamespacedJob {
	return &PatchBatchV1NamespacedJob{Context: ctx, Handler: handler}
}

/*PatchBatchV1NamespacedJob swagger:route PATCH /apis/batch/v1/namespaces/{namespace}/jobs/{name} batch_v1 patchBatchV1NamespacedJob

partially update the specified Job

*/
type PatchBatchV1NamespacedJob struct {
	Context *middleware.Context
	Handler PatchBatchV1NamespacedJobHandler
}

func (o *PatchBatchV1NamespacedJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchBatchV1NamespacedJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}