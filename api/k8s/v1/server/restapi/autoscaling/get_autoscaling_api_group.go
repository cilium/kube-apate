// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAutoscalingAPIGroupHandlerFunc turns a function with the right signature into a get autoscaling API group handler
type GetAutoscalingAPIGroupHandlerFunc func(GetAutoscalingAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAutoscalingAPIGroupHandlerFunc) Handle(params GetAutoscalingAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetAutoscalingAPIGroupHandler interface for that can handle valid get autoscaling API group params
type GetAutoscalingAPIGroupHandler interface {
	Handle(GetAutoscalingAPIGroupParams) middleware.Responder
}

// NewGetAutoscalingAPIGroup creates a new http.Handler for the get autoscaling API group operation
func NewGetAutoscalingAPIGroup(ctx *middleware.Context, handler GetAutoscalingAPIGroupHandler) *GetAutoscalingAPIGroup {
	return &GetAutoscalingAPIGroup{Context: ctx, Handler: handler}
}

/*GetAutoscalingAPIGroup swagger:route GET /apis/autoscaling/ autoscaling getAutoscalingApiGroup

get information of a group

*/
type GetAutoscalingAPIGroup struct {
	Context *middleware.Context
	Handler GetAutoscalingAPIGroupHandler
}

func (o *GetAutoscalingAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAutoscalingAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
