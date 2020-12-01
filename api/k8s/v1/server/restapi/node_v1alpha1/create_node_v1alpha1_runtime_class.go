// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateNodeV1alpha1RuntimeClassHandlerFunc turns a function with the right signature into a create node v1alpha1 runtime class handler
type CreateNodeV1alpha1RuntimeClassHandlerFunc func(CreateNodeV1alpha1RuntimeClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateNodeV1alpha1RuntimeClassHandlerFunc) Handle(params CreateNodeV1alpha1RuntimeClassParams) middleware.Responder {
	return fn(params)
}

// CreateNodeV1alpha1RuntimeClassHandler interface for that can handle valid create node v1alpha1 runtime class params
type CreateNodeV1alpha1RuntimeClassHandler interface {
	Handle(CreateNodeV1alpha1RuntimeClassParams) middleware.Responder
}

// NewCreateNodeV1alpha1RuntimeClass creates a new http.Handler for the create node v1alpha1 runtime class operation
func NewCreateNodeV1alpha1RuntimeClass(ctx *middleware.Context, handler CreateNodeV1alpha1RuntimeClassHandler) *CreateNodeV1alpha1RuntimeClass {
	return &CreateNodeV1alpha1RuntimeClass{Context: ctx, Handler: handler}
}

/*CreateNodeV1alpha1RuntimeClass swagger:route POST /apis/node.k8s.io/v1alpha1/runtimeclasses node_v1alpha1 createNodeV1alpha1RuntimeClass

create a RuntimeClass

*/
type CreateNodeV1alpha1RuntimeClass struct {
	Context *middleware.Context
	Handler CreateNodeV1alpha1RuntimeClassHandler
}

func (o *CreateNodeV1alpha1RuntimeClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateNodeV1alpha1RuntimeClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
