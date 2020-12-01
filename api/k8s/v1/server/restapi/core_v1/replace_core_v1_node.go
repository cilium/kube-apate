// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NodeHandlerFunc turns a function with the right signature into a replace core v1 node handler
type ReplaceCoreV1NodeHandlerFunc func(ReplaceCoreV1NodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NodeHandlerFunc) Handle(params ReplaceCoreV1NodeParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NodeHandler interface for that can handle valid replace core v1 node params
type ReplaceCoreV1NodeHandler interface {
	Handle(ReplaceCoreV1NodeParams) middleware.Responder
}

// NewReplaceCoreV1Node creates a new http.Handler for the replace core v1 node operation
func NewReplaceCoreV1Node(ctx *middleware.Context, handler ReplaceCoreV1NodeHandler) *ReplaceCoreV1Node {
	return &ReplaceCoreV1Node{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1Node swagger:route PUT /api/v1/nodes/{name} core_v1 replaceCoreV1Node

replace the specified Node

*/
type ReplaceCoreV1Node struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NodeHandler
}

func (o *ReplaceCoreV1Node) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
