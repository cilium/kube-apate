// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEventsAPIGroupHandlerFunc turns a function with the right signature into a get events API group handler
type GetEventsAPIGroupHandlerFunc func(GetEventsAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEventsAPIGroupHandlerFunc) Handle(params GetEventsAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetEventsAPIGroupHandler interface for that can handle valid get events API group params
type GetEventsAPIGroupHandler interface {
	Handle(GetEventsAPIGroupParams) middleware.Responder
}

// NewGetEventsAPIGroup creates a new http.Handler for the get events API group operation
func NewGetEventsAPIGroup(ctx *middleware.Context, handler GetEventsAPIGroupHandler) *GetEventsAPIGroup {
	return &GetEventsAPIGroup{Context: ctx, Handler: handler}
}

/*GetEventsAPIGroup swagger:route GET /apis/events.k8s.io/ events getEventsApiGroup

get information of a group

*/
type GetEventsAPIGroup struct {
	Context *middleware.Context
	Handler GetEventsAPIGroupHandler
}

func (o *GetEventsAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEventsAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
