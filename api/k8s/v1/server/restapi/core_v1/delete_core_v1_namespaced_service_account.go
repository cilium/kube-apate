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

// DeleteCoreV1NamespacedServiceAccountHandlerFunc turns a function with the right signature into a delete core v1 namespaced service account handler
type DeleteCoreV1NamespacedServiceAccountHandlerFunc func(DeleteCoreV1NamespacedServiceAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1NamespacedServiceAccountHandlerFunc) Handle(params DeleteCoreV1NamespacedServiceAccountParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1NamespacedServiceAccountHandler interface for that can handle valid delete core v1 namespaced service account params
type DeleteCoreV1NamespacedServiceAccountHandler interface {
	Handle(DeleteCoreV1NamespacedServiceAccountParams) middleware.Responder
}

// NewDeleteCoreV1NamespacedServiceAccount creates a new http.Handler for the delete core v1 namespaced service account operation
func NewDeleteCoreV1NamespacedServiceAccount(ctx *middleware.Context, handler DeleteCoreV1NamespacedServiceAccountHandler) *DeleteCoreV1NamespacedServiceAccount {
	return &DeleteCoreV1NamespacedServiceAccount{Context: ctx, Handler: handler}
}

/*DeleteCoreV1NamespacedServiceAccount swagger:route DELETE /api/v1/namespaces/{namespace}/serviceaccounts/{name} core_v1 deleteCoreV1NamespacedServiceAccount

delete a ServiceAccount

*/
type DeleteCoreV1NamespacedServiceAccount struct {
	Context *middleware.Context
	Handler DeleteCoreV1NamespacedServiceAccountHandler
}

func (o *DeleteCoreV1NamespacedServiceAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1NamespacedServiceAccountParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
