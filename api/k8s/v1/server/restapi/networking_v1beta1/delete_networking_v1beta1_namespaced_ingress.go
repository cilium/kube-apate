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

// DeleteNetworkingV1beta1NamespacedIngressHandlerFunc turns a function with the right signature into a delete networking v1beta1 namespaced ingress handler
type DeleteNetworkingV1beta1NamespacedIngressHandlerFunc func(DeleteNetworkingV1beta1NamespacedIngressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNetworkingV1beta1NamespacedIngressHandlerFunc) Handle(params DeleteNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
	return fn(params)
}

// DeleteNetworkingV1beta1NamespacedIngressHandler interface for that can handle valid delete networking v1beta1 namespaced ingress params
type DeleteNetworkingV1beta1NamespacedIngressHandler interface {
	Handle(DeleteNetworkingV1beta1NamespacedIngressParams) middleware.Responder
}

// NewDeleteNetworkingV1beta1NamespacedIngress creates a new http.Handler for the delete networking v1beta1 namespaced ingress operation
func NewDeleteNetworkingV1beta1NamespacedIngress(ctx *middleware.Context, handler DeleteNetworkingV1beta1NamespacedIngressHandler) *DeleteNetworkingV1beta1NamespacedIngress {
	return &DeleteNetworkingV1beta1NamespacedIngress{Context: ctx, Handler: handler}
}

/*DeleteNetworkingV1beta1NamespacedIngress swagger:route DELETE /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name} networking_v1beta1 deleteNetworkingV1beta1NamespacedIngress

delete an Ingress

*/
type DeleteNetworkingV1beta1NamespacedIngress struct {
	Context *middleware.Context
	Handler DeleteNetworkingV1beta1NamespacedIngressHandler
}

func (o *DeleteNetworkingV1beta1NamespacedIngress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNetworkingV1beta1NamespacedIngressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
