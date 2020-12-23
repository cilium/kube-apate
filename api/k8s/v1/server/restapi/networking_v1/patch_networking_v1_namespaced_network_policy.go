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

// PatchNetworkingV1NamespacedNetworkPolicyHandlerFunc turns a function with the right signature into a patch networking v1 namespaced network policy handler
type PatchNetworkingV1NamespacedNetworkPolicyHandlerFunc func(PatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchNetworkingV1NamespacedNetworkPolicyHandlerFunc) Handle(params PatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
	return fn(params)
}

// PatchNetworkingV1NamespacedNetworkPolicyHandler interface for that can handle valid patch networking v1 namespaced network policy params
type PatchNetworkingV1NamespacedNetworkPolicyHandler interface {
	Handle(PatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder
}

// NewPatchNetworkingV1NamespacedNetworkPolicy creates a new http.Handler for the patch networking v1 namespaced network policy operation
func NewPatchNetworkingV1NamespacedNetworkPolicy(ctx *middleware.Context, handler PatchNetworkingV1NamespacedNetworkPolicyHandler) *PatchNetworkingV1NamespacedNetworkPolicy {
	return &PatchNetworkingV1NamespacedNetworkPolicy{Context: ctx, Handler: handler}
}

/*PatchNetworkingV1NamespacedNetworkPolicy swagger:route PATCH /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies/{name} networking_v1 patchNetworkingV1NamespacedNetworkPolicy

partially update the specified NetworkPolicy

*/
type PatchNetworkingV1NamespacedNetworkPolicy struct {
	Context *middleware.Context
	Handler PatchNetworkingV1NamespacedNetworkPolicyHandler
}

func (o *PatchNetworkingV1NamespacedNetworkPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchNetworkingV1NamespacedNetworkPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}