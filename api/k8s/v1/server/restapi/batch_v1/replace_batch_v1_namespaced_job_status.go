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

// ReplaceBatchV1NamespacedJobStatusHandlerFunc turns a function with the right signature into a replace batch v1 namespaced job status handler
type ReplaceBatchV1NamespacedJobStatusHandlerFunc func(ReplaceBatchV1NamespacedJobStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceBatchV1NamespacedJobStatusHandlerFunc) Handle(params ReplaceBatchV1NamespacedJobStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceBatchV1NamespacedJobStatusHandler interface for that can handle valid replace batch v1 namespaced job status params
type ReplaceBatchV1NamespacedJobStatusHandler interface {
	Handle(ReplaceBatchV1NamespacedJobStatusParams) middleware.Responder
}

// NewReplaceBatchV1NamespacedJobStatus creates a new http.Handler for the replace batch v1 namespaced job status operation
func NewReplaceBatchV1NamespacedJobStatus(ctx *middleware.Context, handler ReplaceBatchV1NamespacedJobStatusHandler) *ReplaceBatchV1NamespacedJobStatus {
	return &ReplaceBatchV1NamespacedJobStatus{Context: ctx, Handler: handler}
}

/*ReplaceBatchV1NamespacedJobStatus swagger:route PUT /apis/batch/v1/namespaces/{namespace}/jobs/{name}/status batch_v1 replaceBatchV1NamespacedJobStatus

replace status of the specified Job

*/
type ReplaceBatchV1NamespacedJobStatus struct {
	Context *middleware.Context
	Handler ReplaceBatchV1NamespacedJobStatusHandler
}

func (o *ReplaceBatchV1NamespacedJobStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceBatchV1NamespacedJobStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}