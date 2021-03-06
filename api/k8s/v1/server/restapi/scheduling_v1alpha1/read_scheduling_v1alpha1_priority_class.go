// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadSchedulingV1alpha1PriorityClassHandlerFunc turns a function with the right signature into a read scheduling v1alpha1 priority class handler
type ReadSchedulingV1alpha1PriorityClassHandlerFunc func(ReadSchedulingV1alpha1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadSchedulingV1alpha1PriorityClassHandlerFunc) Handle(params ReadSchedulingV1alpha1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// ReadSchedulingV1alpha1PriorityClassHandler interface for that can handle valid read scheduling v1alpha1 priority class params
type ReadSchedulingV1alpha1PriorityClassHandler interface {
	Handle(ReadSchedulingV1alpha1PriorityClassParams) middleware.Responder
}

// NewReadSchedulingV1alpha1PriorityClass creates a new http.Handler for the read scheduling v1alpha1 priority class operation
func NewReadSchedulingV1alpha1PriorityClass(ctx *middleware.Context, handler ReadSchedulingV1alpha1PriorityClassHandler) *ReadSchedulingV1alpha1PriorityClass {
	return &ReadSchedulingV1alpha1PriorityClass{Context: ctx, Handler: handler}
}

/*ReadSchedulingV1alpha1PriorityClass swagger:route GET /apis/scheduling.k8s.io/v1alpha1/priorityclasses/{name} scheduling_v1alpha1 readSchedulingV1alpha1PriorityClass

read the specified PriorityClass

*/
type ReadSchedulingV1alpha1PriorityClass struct {
	Context *middleware.Context
	Handler ReadSchedulingV1alpha1PriorityClassHandler
}

func (o *ReadSchedulingV1alpha1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadSchedulingV1alpha1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
