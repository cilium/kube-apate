// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateNodeV1beta1RuntimeClassHandlerFunc turns a function with the right signature into a create node v1beta1 runtime class handler
type CreateNodeV1beta1RuntimeClassHandlerFunc func(CreateNodeV1beta1RuntimeClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateNodeV1beta1RuntimeClassHandlerFunc) Handle(params CreateNodeV1beta1RuntimeClassParams) middleware.Responder {
	return fn(params)
}

// CreateNodeV1beta1RuntimeClassHandler interface for that can handle valid create node v1beta1 runtime class params
type CreateNodeV1beta1RuntimeClassHandler interface {
	Handle(CreateNodeV1beta1RuntimeClassParams) middleware.Responder
}

// NewCreateNodeV1beta1RuntimeClass creates a new http.Handler for the create node v1beta1 runtime class operation
func NewCreateNodeV1beta1RuntimeClass(ctx *middleware.Context, handler CreateNodeV1beta1RuntimeClassHandler) *CreateNodeV1beta1RuntimeClass {
	return &CreateNodeV1beta1RuntimeClass{Context: ctx, Handler: handler}
}

/*CreateNodeV1beta1RuntimeClass swagger:route POST /apis/node.k8s.io/v1beta1/runtimeclasses node_v1beta1 createNodeV1beta1RuntimeClass

create a RuntimeClass

*/
type CreateNodeV1beta1RuntimeClass struct {
	Context *middleware.Context
	Handler CreateNodeV1beta1RuntimeClassHandler
}

func (o *CreateNodeV1beta1RuntimeClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateNodeV1beta1RuntimeClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
