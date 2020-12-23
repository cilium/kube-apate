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

// DeleteCoreV1CollectionNamespacedServiceAccountHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced service account handler
type DeleteCoreV1CollectionNamespacedServiceAccountHandlerFunc func(DeleteCoreV1CollectionNamespacedServiceAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedServiceAccountHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedServiceAccountParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1CollectionNamespacedServiceAccountHandler interface for that can handle valid delete core v1 collection namespaced service account params
type DeleteCoreV1CollectionNamespacedServiceAccountHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedServiceAccountParams) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedServiceAccount creates a new http.Handler for the delete core v1 collection namespaced service account operation
func NewDeleteCoreV1CollectionNamespacedServiceAccount(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedServiceAccountHandler) *DeleteCoreV1CollectionNamespacedServiceAccount {
	return &DeleteCoreV1CollectionNamespacedServiceAccount{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedServiceAccount swagger:route DELETE /api/v1/namespaces/{namespace}/serviceaccounts core_v1 deleteCoreV1CollectionNamespacedServiceAccount

delete collection of ServiceAccount

*/
type DeleteCoreV1CollectionNamespacedServiceAccount struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedServiceAccountHandler
}

func (o *DeleteCoreV1CollectionNamespacedServiceAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedServiceAccountParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}