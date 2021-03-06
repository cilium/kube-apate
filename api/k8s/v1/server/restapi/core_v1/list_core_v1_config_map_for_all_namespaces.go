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

// ListCoreV1ConfigMapForAllNamespacesHandlerFunc turns a function with the right signature into a list core v1 config map for all namespaces handler
type ListCoreV1ConfigMapForAllNamespacesHandlerFunc func(ListCoreV1ConfigMapForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1ConfigMapForAllNamespacesHandlerFunc) Handle(params ListCoreV1ConfigMapForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListCoreV1ConfigMapForAllNamespacesHandler interface for that can handle valid list core v1 config map for all namespaces params
type ListCoreV1ConfigMapForAllNamespacesHandler interface {
	Handle(ListCoreV1ConfigMapForAllNamespacesParams) middleware.Responder
}

// NewListCoreV1ConfigMapForAllNamespaces creates a new http.Handler for the list core v1 config map for all namespaces operation
func NewListCoreV1ConfigMapForAllNamespaces(ctx *middleware.Context, handler ListCoreV1ConfigMapForAllNamespacesHandler) *ListCoreV1ConfigMapForAllNamespaces {
	return &ListCoreV1ConfigMapForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListCoreV1ConfigMapForAllNamespaces swagger:route GET /api/v1/configmaps core_v1 listCoreV1ConfigMapForAllNamespaces

list or watch objects of kind ConfigMap

*/
type ListCoreV1ConfigMapForAllNamespaces struct {
	Context *middleware.Context
	Handler ListCoreV1ConfigMapForAllNamespacesHandler
}

func (o *ListCoreV1ConfigMapForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1ConfigMapForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
