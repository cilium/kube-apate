// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoreV1NamespacedResourceQuotaHandlerFunc turns a function with the right signature into a list core v1 namespaced resource quota handler
type ListCoreV1NamespacedResourceQuotaHandlerFunc func(ListCoreV1NamespacedResourceQuotaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1NamespacedResourceQuotaHandlerFunc) Handle(params ListCoreV1NamespacedResourceQuotaParams) middleware.Responder {
	return fn(params)
}

// ListCoreV1NamespacedResourceQuotaHandler interface for that can handle valid list core v1 namespaced resource quota params
type ListCoreV1NamespacedResourceQuotaHandler interface {
	Handle(ListCoreV1NamespacedResourceQuotaParams) middleware.Responder
}

// NewListCoreV1NamespacedResourceQuota creates a new http.Handler for the list core v1 namespaced resource quota operation
func NewListCoreV1NamespacedResourceQuota(ctx *middleware.Context, handler ListCoreV1NamespacedResourceQuotaHandler) *ListCoreV1NamespacedResourceQuota {
	return &ListCoreV1NamespacedResourceQuota{Context: ctx, Handler: handler}
}

/*ListCoreV1NamespacedResourceQuota swagger:route GET /api/v1/namespaces/{namespace}/resourcequotas core_v1 listCoreV1NamespacedResourceQuota

list or watch objects of kind ResourceQuota

*/
type ListCoreV1NamespacedResourceQuota struct {
	Context *middleware.Context
	Handler ListCoreV1NamespacedResourceQuotaHandler
}

func (o *ListCoreV1NamespacedResourceQuota) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1NamespacedResourceQuotaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
