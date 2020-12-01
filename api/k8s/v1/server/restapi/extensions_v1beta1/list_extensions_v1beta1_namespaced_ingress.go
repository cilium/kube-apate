// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListExtensionsV1beta1NamespacedIngressHandlerFunc turns a function with the right signature into a list extensions v1beta1 namespaced ingress handler
type ListExtensionsV1beta1NamespacedIngressHandlerFunc func(ListExtensionsV1beta1NamespacedIngressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListExtensionsV1beta1NamespacedIngressHandlerFunc) Handle(params ListExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
	return fn(params)
}

// ListExtensionsV1beta1NamespacedIngressHandler interface for that can handle valid list extensions v1beta1 namespaced ingress params
type ListExtensionsV1beta1NamespacedIngressHandler interface {
	Handle(ListExtensionsV1beta1NamespacedIngressParams) middleware.Responder
}

// NewListExtensionsV1beta1NamespacedIngress creates a new http.Handler for the list extensions v1beta1 namespaced ingress operation
func NewListExtensionsV1beta1NamespacedIngress(ctx *middleware.Context, handler ListExtensionsV1beta1NamespacedIngressHandler) *ListExtensionsV1beta1NamespacedIngress {
	return &ListExtensionsV1beta1NamespacedIngress{Context: ctx, Handler: handler}
}

/*ListExtensionsV1beta1NamespacedIngress swagger:route GET /apis/extensions/v1beta1/namespaces/{namespace}/ingresses extensions_v1beta1 listExtensionsV1beta1NamespacedIngress

list or watch objects of kind Ingress

*/
type ListExtensionsV1beta1NamespacedIngress struct {
	Context *middleware.Context
	Handler ListExtensionsV1beta1NamespacedIngressHandler
}

func (o *ListExtensionsV1beta1NamespacedIngress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListExtensionsV1beta1NamespacedIngressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
