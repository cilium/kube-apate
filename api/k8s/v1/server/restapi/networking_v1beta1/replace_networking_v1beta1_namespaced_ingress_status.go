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

// ReplaceNetworkingV1beta1NamespacedIngressStatusHandlerFunc turns a function with the right signature into a replace networking v1beta1 namespaced ingress status handler
type ReplaceNetworkingV1beta1NamespacedIngressStatusHandlerFunc func(ReplaceNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceNetworkingV1beta1NamespacedIngressStatusHandlerFunc) Handle(params ReplaceNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceNetworkingV1beta1NamespacedIngressStatusHandler interface for that can handle valid replace networking v1beta1 namespaced ingress status params
type ReplaceNetworkingV1beta1NamespacedIngressStatusHandler interface {
	Handle(ReplaceNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder
}

// NewReplaceNetworkingV1beta1NamespacedIngressStatus creates a new http.Handler for the replace networking v1beta1 namespaced ingress status operation
func NewReplaceNetworkingV1beta1NamespacedIngressStatus(ctx *middleware.Context, handler ReplaceNetworkingV1beta1NamespacedIngressStatusHandler) *ReplaceNetworkingV1beta1NamespacedIngressStatus {
	return &ReplaceNetworkingV1beta1NamespacedIngressStatus{Context: ctx, Handler: handler}
}

/*ReplaceNetworkingV1beta1NamespacedIngressStatus swagger:route PUT /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name}/status networking_v1beta1 replaceNetworkingV1beta1NamespacedIngressStatus

replace status of the specified Ingress

*/
type ReplaceNetworkingV1beta1NamespacedIngressStatus struct {
	Context *middleware.Context
	Handler ReplaceNetworkingV1beta1NamespacedIngressStatusHandler
}

func (o *ReplaceNetworkingV1beta1NamespacedIngressStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceNetworkingV1beta1NamespacedIngressStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
