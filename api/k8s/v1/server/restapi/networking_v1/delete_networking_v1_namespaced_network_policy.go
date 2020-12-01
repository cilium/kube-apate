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

// DeleteNetworkingV1NamespacedNetworkPolicyHandlerFunc turns a function with the right signature into a delete networking v1 namespaced network policy handler
type DeleteNetworkingV1NamespacedNetworkPolicyHandlerFunc func(DeleteNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNetworkingV1NamespacedNetworkPolicyHandlerFunc) Handle(params DeleteNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
	return fn(params)
}

// DeleteNetworkingV1NamespacedNetworkPolicyHandler interface for that can handle valid delete networking v1 namespaced network policy params
type DeleteNetworkingV1NamespacedNetworkPolicyHandler interface {
	Handle(DeleteNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder
}

// NewDeleteNetworkingV1NamespacedNetworkPolicy creates a new http.Handler for the delete networking v1 namespaced network policy operation
func NewDeleteNetworkingV1NamespacedNetworkPolicy(ctx *middleware.Context, handler DeleteNetworkingV1NamespacedNetworkPolicyHandler) *DeleteNetworkingV1NamespacedNetworkPolicy {
	return &DeleteNetworkingV1NamespacedNetworkPolicy{Context: ctx, Handler: handler}
}

/*DeleteNetworkingV1NamespacedNetworkPolicy swagger:route DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies/{name} networking_v1 deleteNetworkingV1NamespacedNetworkPolicy

delete a NetworkPolicy

*/
type DeleteNetworkingV1NamespacedNetworkPolicy struct {
	Context *middleware.Context
	Handler DeleteNetworkingV1NamespacedNetworkPolicyHandler
}

func (o *DeleteNetworkingV1NamespacedNetworkPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNetworkingV1NamespacedNetworkPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
