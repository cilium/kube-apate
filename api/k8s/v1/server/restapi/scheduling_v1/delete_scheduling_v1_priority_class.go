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

// DeleteSchedulingV1PriorityClassHandlerFunc turns a function with the right signature into a delete scheduling v1 priority class handler
type DeleteSchedulingV1PriorityClassHandlerFunc func(DeleteSchedulingV1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSchedulingV1PriorityClassHandlerFunc) Handle(params DeleteSchedulingV1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// DeleteSchedulingV1PriorityClassHandler interface for that can handle valid delete scheduling v1 priority class params
type DeleteSchedulingV1PriorityClassHandler interface {
	Handle(DeleteSchedulingV1PriorityClassParams) middleware.Responder
}

// NewDeleteSchedulingV1PriorityClass creates a new http.Handler for the delete scheduling v1 priority class operation
func NewDeleteSchedulingV1PriorityClass(ctx *middleware.Context, handler DeleteSchedulingV1PriorityClassHandler) *DeleteSchedulingV1PriorityClass {
	return &DeleteSchedulingV1PriorityClass{Context: ctx, Handler: handler}
}

/*DeleteSchedulingV1PriorityClass swagger:route DELETE /apis/scheduling.k8s.io/v1/priorityclasses/{name} scheduling_v1 deleteSchedulingV1PriorityClass

delete a PriorityClass

*/
type DeleteSchedulingV1PriorityClass struct {
	Context *middleware.Context
	Handler DeleteSchedulingV1PriorityClassHandler
}

func (o *DeleteSchedulingV1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteSchedulingV1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}