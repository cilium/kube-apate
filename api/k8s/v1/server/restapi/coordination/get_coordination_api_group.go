// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCoordinationAPIGroupHandlerFunc turns a function with the right signature into a get coordination API group handler
type GetCoordinationAPIGroupHandlerFunc func(GetCoordinationAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCoordinationAPIGroupHandlerFunc) Handle(params GetCoordinationAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetCoordinationAPIGroupHandler interface for that can handle valid get coordination API group params
type GetCoordinationAPIGroupHandler interface {
	Handle(GetCoordinationAPIGroupParams) middleware.Responder
}

// NewGetCoordinationAPIGroup creates a new http.Handler for the get coordination API group operation
func NewGetCoordinationAPIGroup(ctx *middleware.Context, handler GetCoordinationAPIGroupHandler) *GetCoordinationAPIGroup {
	return &GetCoordinationAPIGroup{Context: ctx, Handler: handler}
}

/*GetCoordinationAPIGroup swagger:route GET /apis/coordination.k8s.io/ coordination getCoordinationApiGroup

get information of a group

*/
type GetCoordinationAPIGroup struct {
	Context *middleware.Context
	Handler GetCoordinationAPIGroupHandler
}

func (o *GetCoordinationAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCoordinationAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}