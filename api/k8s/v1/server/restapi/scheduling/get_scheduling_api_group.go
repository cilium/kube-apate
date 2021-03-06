// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSchedulingAPIGroupHandlerFunc turns a function with the right signature into a get scheduling API group handler
type GetSchedulingAPIGroupHandlerFunc func(GetSchedulingAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSchedulingAPIGroupHandlerFunc) Handle(params GetSchedulingAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetSchedulingAPIGroupHandler interface for that can handle valid get scheduling API group params
type GetSchedulingAPIGroupHandler interface {
	Handle(GetSchedulingAPIGroupParams) middleware.Responder
}

// NewGetSchedulingAPIGroup creates a new http.Handler for the get scheduling API group operation
func NewGetSchedulingAPIGroup(ctx *middleware.Context, handler GetSchedulingAPIGroupHandler) *GetSchedulingAPIGroup {
	return &GetSchedulingAPIGroup{Context: ctx, Handler: handler}
}

/*GetSchedulingAPIGroup swagger:route GET /apis/scheduling.k8s.io/ scheduling getSchedulingApiGroup

get information of a group

*/
type GetSchedulingAPIGroup struct {
	Context *middleware.Context
	Handler GetSchedulingAPIGroupHandler
}

func (o *GetSchedulingAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSchedulingAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
