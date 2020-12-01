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

// ReadNetworkingV1NamespacedIngressHandlerFunc turns a function with the right signature into a read networking v1 namespaced ingress handler
type ReadNetworkingV1NamespacedIngressHandlerFunc func(ReadNetworkingV1NamespacedIngressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadNetworkingV1NamespacedIngressHandlerFunc) Handle(params ReadNetworkingV1NamespacedIngressParams) middleware.Responder {
	return fn(params)
}

// ReadNetworkingV1NamespacedIngressHandler interface for that can handle valid read networking v1 namespaced ingress params
type ReadNetworkingV1NamespacedIngressHandler interface {
	Handle(ReadNetworkingV1NamespacedIngressParams) middleware.Responder
}

// NewReadNetworkingV1NamespacedIngress creates a new http.Handler for the read networking v1 namespaced ingress operation
func NewReadNetworkingV1NamespacedIngress(ctx *middleware.Context, handler ReadNetworkingV1NamespacedIngressHandler) *ReadNetworkingV1NamespacedIngress {
	return &ReadNetworkingV1NamespacedIngress{Context: ctx, Handler: handler}
}

/*ReadNetworkingV1NamespacedIngress swagger:route GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name} networking_v1 readNetworkingV1NamespacedIngress

read the specified Ingress

*/
type ReadNetworkingV1NamespacedIngress struct {
	Context *middleware.Context
	Handler ReadNetworkingV1NamespacedIngressHandler
}

func (o *ReadNetworkingV1NamespacedIngress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadNetworkingV1NamespacedIngressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
