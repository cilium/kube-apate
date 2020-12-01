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

// CreateNetworkingV1NamespacedNetworkPolicyHandlerFunc turns a function with the right signature into a create networking v1 namespaced network policy handler
type CreateNetworkingV1NamespacedNetworkPolicyHandlerFunc func(CreateNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateNetworkingV1NamespacedNetworkPolicyHandlerFunc) Handle(params CreateNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
	return fn(params)
}

// CreateNetworkingV1NamespacedNetworkPolicyHandler interface for that can handle valid create networking v1 namespaced network policy params
type CreateNetworkingV1NamespacedNetworkPolicyHandler interface {
	Handle(CreateNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder
}

// NewCreateNetworkingV1NamespacedNetworkPolicy creates a new http.Handler for the create networking v1 namespaced network policy operation
func NewCreateNetworkingV1NamespacedNetworkPolicy(ctx *middleware.Context, handler CreateNetworkingV1NamespacedNetworkPolicyHandler) *CreateNetworkingV1NamespacedNetworkPolicy {
	return &CreateNetworkingV1NamespacedNetworkPolicy{Context: ctx, Handler: handler}
}

/*CreateNetworkingV1NamespacedNetworkPolicy swagger:route POST /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies networking_v1 createNetworkingV1NamespacedNetworkPolicy

create a NetworkPolicy

*/
type CreateNetworkingV1NamespacedNetworkPolicy struct {
	Context *middleware.Context
	Handler CreateNetworkingV1NamespacedNetworkPolicyHandler
}

func (o *CreateNetworkingV1NamespacedNetworkPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateNetworkingV1NamespacedNetworkPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
