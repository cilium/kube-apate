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

// ListNetworkingV1IngressForAllNamespacesHandlerFunc turns a function with the right signature into a list networking v1 ingress for all namespaces handler
type ListNetworkingV1IngressForAllNamespacesHandlerFunc func(ListNetworkingV1IngressForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListNetworkingV1IngressForAllNamespacesHandlerFunc) Handle(params ListNetworkingV1IngressForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListNetworkingV1IngressForAllNamespacesHandler interface for that can handle valid list networking v1 ingress for all namespaces params
type ListNetworkingV1IngressForAllNamespacesHandler interface {
	Handle(ListNetworkingV1IngressForAllNamespacesParams) middleware.Responder
}

// NewListNetworkingV1IngressForAllNamespaces creates a new http.Handler for the list networking v1 ingress for all namespaces operation
func NewListNetworkingV1IngressForAllNamespaces(ctx *middleware.Context, handler ListNetworkingV1IngressForAllNamespacesHandler) *ListNetworkingV1IngressForAllNamespaces {
	return &ListNetworkingV1IngressForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListNetworkingV1IngressForAllNamespaces swagger:route GET /apis/networking.k8s.io/v1/ingresses networking_v1 listNetworkingV1IngressForAllNamespaces

list or watch objects of kind Ingress

*/
type ListNetworkingV1IngressForAllNamespaces struct {
	Context *middleware.Context
	Handler ListNetworkingV1IngressForAllNamespacesHandler
}

func (o *ListNetworkingV1IngressForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListNetworkingV1IngressForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
