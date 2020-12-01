// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListNetworkingV1beta1IngressClassHandlerFunc turns a function with the right signature into a list networking v1beta1 ingress class handler
type ListNetworkingV1beta1IngressClassHandlerFunc func(ListNetworkingV1beta1IngressClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListNetworkingV1beta1IngressClassHandlerFunc) Handle(params ListNetworkingV1beta1IngressClassParams) middleware.Responder {
	return fn(params)
}

// ListNetworkingV1beta1IngressClassHandler interface for that can handle valid list networking v1beta1 ingress class params
type ListNetworkingV1beta1IngressClassHandler interface {
	Handle(ListNetworkingV1beta1IngressClassParams) middleware.Responder
}

// NewListNetworkingV1beta1IngressClass creates a new http.Handler for the list networking v1beta1 ingress class operation
func NewListNetworkingV1beta1IngressClass(ctx *middleware.Context, handler ListNetworkingV1beta1IngressClassHandler) *ListNetworkingV1beta1IngressClass {
	return &ListNetworkingV1beta1IngressClass{Context: ctx, Handler: handler}
}

/*ListNetworkingV1beta1IngressClass swagger:route GET /apis/networking.k8s.io/v1beta1/ingressclasses networking_v1beta1 listNetworkingV1beta1IngressClass

list or watch objects of kind IngressClass

*/
type ListNetworkingV1beta1IngressClass struct {
	Context *middleware.Context
	Handler ListNetworkingV1beta1IngressClassHandler
}

func (o *ListNetworkingV1beta1IngressClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListNetworkingV1beta1IngressClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
