// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceNetworkingV1IngressClassHandlerFunc turns a function with the right signature into a replace networking v1 ingress class handler
type ReplaceNetworkingV1IngressClassHandlerFunc func(ReplaceNetworkingV1IngressClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceNetworkingV1IngressClassHandlerFunc) Handle(params ReplaceNetworkingV1IngressClassParams) middleware.Responder {
	return fn(params)
}

// ReplaceNetworkingV1IngressClassHandler interface for that can handle valid replace networking v1 ingress class params
type ReplaceNetworkingV1IngressClassHandler interface {
	Handle(ReplaceNetworkingV1IngressClassParams) middleware.Responder
}

// NewReplaceNetworkingV1IngressClass creates a new http.Handler for the replace networking v1 ingress class operation
func NewReplaceNetworkingV1IngressClass(ctx *middleware.Context, handler ReplaceNetworkingV1IngressClassHandler) *ReplaceNetworkingV1IngressClass {
	return &ReplaceNetworkingV1IngressClass{Context: ctx, Handler: handler}
}

/*ReplaceNetworkingV1IngressClass swagger:route PUT /apis/networking.k8s.io/v1/ingressclasses/{name} networking_v1 replaceNetworkingV1IngressClass

replace the specified IngressClass

*/
type ReplaceNetworkingV1IngressClass struct {
	Context *middleware.Context
	Handler ReplaceNetworkingV1IngressClassHandler
}

func (o *ReplaceNetworkingV1IngressClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceNetworkingV1IngressClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
