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

// DeleteCoreV1NamespacedResourceQuotaHandlerFunc turns a function with the right signature into a delete core v1 namespaced resource quota handler
type DeleteCoreV1NamespacedResourceQuotaHandlerFunc func(DeleteCoreV1NamespacedResourceQuotaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1NamespacedResourceQuotaHandlerFunc) Handle(params DeleteCoreV1NamespacedResourceQuotaParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1NamespacedResourceQuotaHandler interface for that can handle valid delete core v1 namespaced resource quota params
type DeleteCoreV1NamespacedResourceQuotaHandler interface {
	Handle(DeleteCoreV1NamespacedResourceQuotaParams) middleware.Responder
}

// NewDeleteCoreV1NamespacedResourceQuota creates a new http.Handler for the delete core v1 namespaced resource quota operation
func NewDeleteCoreV1NamespacedResourceQuota(ctx *middleware.Context, handler DeleteCoreV1NamespacedResourceQuotaHandler) *DeleteCoreV1NamespacedResourceQuota {
	return &DeleteCoreV1NamespacedResourceQuota{Context: ctx, Handler: handler}
}

/*DeleteCoreV1NamespacedResourceQuota swagger:route DELETE /api/v1/namespaces/{namespace}/resourcequotas/{name} core_v1 deleteCoreV1NamespacedResourceQuota

delete a ResourceQuota

*/
type DeleteCoreV1NamespacedResourceQuota struct {
	Context *middleware.Context
	Handler DeleteCoreV1NamespacedResourceQuotaHandler
}

func (o *DeleteCoreV1NamespacedResourceQuota) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1NamespacedResourceQuotaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
