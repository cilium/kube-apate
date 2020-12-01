// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNodeAPIGroupHandlerFunc turns a function with the right signature into a get node API group handler
type GetNodeAPIGroupHandlerFunc func(GetNodeAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNodeAPIGroupHandlerFunc) Handle(params GetNodeAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetNodeAPIGroupHandler interface for that can handle valid get node API group params
type GetNodeAPIGroupHandler interface {
	Handle(GetNodeAPIGroupParams) middleware.Responder
}

// NewGetNodeAPIGroup creates a new http.Handler for the get node API group operation
func NewGetNodeAPIGroup(ctx *middleware.Context, handler GetNodeAPIGroupHandler) *GetNodeAPIGroup {
	return &GetNodeAPIGroup{Context: ctx, Handler: handler}
}

/*GetNodeAPIGroup swagger:route GET /apis/node.k8s.io/ node getNodeApiGroup

get information of a group

*/
type GetNodeAPIGroup struct {
	Context *middleware.Context
	Handler GetNodeAPIGroupHandler
}

func (o *GetNodeAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNodeAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
