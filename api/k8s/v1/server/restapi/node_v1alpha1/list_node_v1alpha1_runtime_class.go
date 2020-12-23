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

// ListNodeV1alpha1RuntimeClassHandlerFunc turns a function with the right signature into a list node v1alpha1 runtime class handler
type ListNodeV1alpha1RuntimeClassHandlerFunc func(ListNodeV1alpha1RuntimeClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListNodeV1alpha1RuntimeClassHandlerFunc) Handle(params ListNodeV1alpha1RuntimeClassParams) middleware.Responder {
	return fn(params)
}

// ListNodeV1alpha1RuntimeClassHandler interface for that can handle valid list node v1alpha1 runtime class params
type ListNodeV1alpha1RuntimeClassHandler interface {
	Handle(ListNodeV1alpha1RuntimeClassParams) middleware.Responder
}

// NewListNodeV1alpha1RuntimeClass creates a new http.Handler for the list node v1alpha1 runtime class operation
func NewListNodeV1alpha1RuntimeClass(ctx *middleware.Context, handler ListNodeV1alpha1RuntimeClassHandler) *ListNodeV1alpha1RuntimeClass {
	return &ListNodeV1alpha1RuntimeClass{Context: ctx, Handler: handler}
}

/*ListNodeV1alpha1RuntimeClass swagger:route GET /apis/node.k8s.io/v1alpha1/runtimeclasses node_v1alpha1 listNodeV1alpha1RuntimeClass

list or watch objects of kind RuntimeClass

*/
type ListNodeV1alpha1RuntimeClass struct {
	Context *middleware.Context
	Handler ListNodeV1alpha1RuntimeClassHandler
}

func (o *ListNodeV1alpha1RuntimeClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListNodeV1alpha1RuntimeClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}