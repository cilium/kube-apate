// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoordinationV1beta1NamespacedLeaseHandlerFunc turns a function with the right signature into a list coordination v1beta1 namespaced lease handler
type ListCoordinationV1beta1NamespacedLeaseHandlerFunc func(ListCoordinationV1beta1NamespacedLeaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoordinationV1beta1NamespacedLeaseHandlerFunc) Handle(params ListCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
	return fn(params)
}

// ListCoordinationV1beta1NamespacedLeaseHandler interface for that can handle valid list coordination v1beta1 namespaced lease params
type ListCoordinationV1beta1NamespacedLeaseHandler interface {
	Handle(ListCoordinationV1beta1NamespacedLeaseParams) middleware.Responder
}

// NewListCoordinationV1beta1NamespacedLease creates a new http.Handler for the list coordination v1beta1 namespaced lease operation
func NewListCoordinationV1beta1NamespacedLease(ctx *middleware.Context, handler ListCoordinationV1beta1NamespacedLeaseHandler) *ListCoordinationV1beta1NamespacedLease {
	return &ListCoordinationV1beta1NamespacedLease{Context: ctx, Handler: handler}
}

/*ListCoordinationV1beta1NamespacedLease swagger:route GET /apis/coordination.k8s.io/v1beta1/namespaces/{namespace}/leases coordination_v1beta1 listCoordinationV1beta1NamespacedLease

list or watch objects of kind Lease

*/
type ListCoordinationV1beta1NamespacedLease struct {
	Context *middleware.Context
	Handler ListCoordinationV1beta1NamespacedLeaseHandler
}

func (o *ListCoordinationV1beta1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoordinationV1beta1NamespacedLeaseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}