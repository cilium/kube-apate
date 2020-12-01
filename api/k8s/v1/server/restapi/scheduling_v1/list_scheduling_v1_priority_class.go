// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListSchedulingV1PriorityClassHandlerFunc turns a function with the right signature into a list scheduling v1 priority class handler
type ListSchedulingV1PriorityClassHandlerFunc func(ListSchedulingV1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListSchedulingV1PriorityClassHandlerFunc) Handle(params ListSchedulingV1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// ListSchedulingV1PriorityClassHandler interface for that can handle valid list scheduling v1 priority class params
type ListSchedulingV1PriorityClassHandler interface {
	Handle(ListSchedulingV1PriorityClassParams) middleware.Responder
}

// NewListSchedulingV1PriorityClass creates a new http.Handler for the list scheduling v1 priority class operation
func NewListSchedulingV1PriorityClass(ctx *middleware.Context, handler ListSchedulingV1PriorityClassHandler) *ListSchedulingV1PriorityClass {
	return &ListSchedulingV1PriorityClass{Context: ctx, Handler: handler}
}

/*ListSchedulingV1PriorityClass swagger:route GET /apis/scheduling.k8s.io/v1/priorityclasses scheduling_v1 listSchedulingV1PriorityClass

list or watch objects of kind PriorityClass

*/
type ListSchedulingV1PriorityClass struct {
	Context *middleware.Context
	Handler ListSchedulingV1PriorityClassHandler
}

func (o *ListSchedulingV1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListSchedulingV1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
