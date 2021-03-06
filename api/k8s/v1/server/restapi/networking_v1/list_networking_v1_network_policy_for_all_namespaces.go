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

// ListNetworkingV1NetworkPolicyForAllNamespacesHandlerFunc turns a function with the right signature into a list networking v1 network policy for all namespaces handler
type ListNetworkingV1NetworkPolicyForAllNamespacesHandlerFunc func(ListNetworkingV1NetworkPolicyForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListNetworkingV1NetworkPolicyForAllNamespacesHandlerFunc) Handle(params ListNetworkingV1NetworkPolicyForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListNetworkingV1NetworkPolicyForAllNamespacesHandler interface for that can handle valid list networking v1 network policy for all namespaces params
type ListNetworkingV1NetworkPolicyForAllNamespacesHandler interface {
	Handle(ListNetworkingV1NetworkPolicyForAllNamespacesParams) middleware.Responder
}

// NewListNetworkingV1NetworkPolicyForAllNamespaces creates a new http.Handler for the list networking v1 network policy for all namespaces operation
func NewListNetworkingV1NetworkPolicyForAllNamespaces(ctx *middleware.Context, handler ListNetworkingV1NetworkPolicyForAllNamespacesHandler) *ListNetworkingV1NetworkPolicyForAllNamespaces {
	return &ListNetworkingV1NetworkPolicyForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListNetworkingV1NetworkPolicyForAllNamespaces swagger:route GET /apis/networking.k8s.io/v1/networkpolicies networking_v1 listNetworkingV1NetworkPolicyForAllNamespaces

list or watch objects of kind NetworkPolicy

*/
type ListNetworkingV1NetworkPolicyForAllNamespaces struct {
	Context *middleware.Context
	Handler ListNetworkingV1NetworkPolicyForAllNamespacesHandler
}

func (o *ListNetworkingV1NetworkPolicyForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListNetworkingV1NetworkPolicyForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
