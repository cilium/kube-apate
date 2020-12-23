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

// CreateCoreV1NamespacedBindingHandlerFunc turns a function with the right signature into a create core v1 namespaced binding handler
type CreateCoreV1NamespacedBindingHandlerFunc func(CreateCoreV1NamespacedBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreV1NamespacedBindingHandlerFunc) Handle(params CreateCoreV1NamespacedBindingParams) middleware.Responder {
	return fn(params)
}

// CreateCoreV1NamespacedBindingHandler interface for that can handle valid create core v1 namespaced binding params
type CreateCoreV1NamespacedBindingHandler interface {
	Handle(CreateCoreV1NamespacedBindingParams) middleware.Responder
}

// NewCreateCoreV1NamespacedBinding creates a new http.Handler for the create core v1 namespaced binding operation
func NewCreateCoreV1NamespacedBinding(ctx *middleware.Context, handler CreateCoreV1NamespacedBindingHandler) *CreateCoreV1NamespacedBinding {
	return &CreateCoreV1NamespacedBinding{Context: ctx, Handler: handler}
}

/*CreateCoreV1NamespacedBinding swagger:route POST /api/v1/namespaces/{namespace}/bindings core_v1 createCoreV1NamespacedBinding

create a Binding

*/
type CreateCoreV1NamespacedBinding struct {
	Context *middleware.Context
	Handler CreateCoreV1NamespacedBindingHandler
}

func (o *CreateCoreV1NamespacedBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoreV1NamespacedBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}