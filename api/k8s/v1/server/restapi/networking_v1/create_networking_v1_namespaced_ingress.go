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

// CreateNetworkingV1NamespacedIngressHandlerFunc turns a function with the right signature into a create networking v1 namespaced ingress handler
type CreateNetworkingV1NamespacedIngressHandlerFunc func(CreateNetworkingV1NamespacedIngressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateNetworkingV1NamespacedIngressHandlerFunc) Handle(params CreateNetworkingV1NamespacedIngressParams) middleware.Responder {
	return fn(params)
}

// CreateNetworkingV1NamespacedIngressHandler interface for that can handle valid create networking v1 namespaced ingress params
type CreateNetworkingV1NamespacedIngressHandler interface {
	Handle(CreateNetworkingV1NamespacedIngressParams) middleware.Responder
}

// NewCreateNetworkingV1NamespacedIngress creates a new http.Handler for the create networking v1 namespaced ingress operation
func NewCreateNetworkingV1NamespacedIngress(ctx *middleware.Context, handler CreateNetworkingV1NamespacedIngressHandler) *CreateNetworkingV1NamespacedIngress {
	return &CreateNetworkingV1NamespacedIngress{Context: ctx, Handler: handler}
}

/*CreateNetworkingV1NamespacedIngress swagger:route POST /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses networking_v1 createNetworkingV1NamespacedIngress

create an Ingress

*/
type CreateNetworkingV1NamespacedIngress struct {
	Context *middleware.Context
	Handler CreateNetworkingV1NamespacedIngressHandler
}

func (o *CreateNetworkingV1NamespacedIngress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateNetworkingV1NamespacedIngressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
