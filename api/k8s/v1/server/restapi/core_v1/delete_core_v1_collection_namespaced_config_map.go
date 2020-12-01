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

// DeleteCoreV1CollectionNamespacedConfigMapHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced config map handler
type DeleteCoreV1CollectionNamespacedConfigMapHandlerFunc func(DeleteCoreV1CollectionNamespacedConfigMapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedConfigMapHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedConfigMapParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1CollectionNamespacedConfigMapHandler interface for that can handle valid delete core v1 collection namespaced config map params
type DeleteCoreV1CollectionNamespacedConfigMapHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedConfigMapParams) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedConfigMap creates a new http.Handler for the delete core v1 collection namespaced config map operation
func NewDeleteCoreV1CollectionNamespacedConfigMap(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedConfigMapHandler) *DeleteCoreV1CollectionNamespacedConfigMap {
	return &DeleteCoreV1CollectionNamespacedConfigMap{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedConfigMap swagger:route DELETE /api/v1/namespaces/{namespace}/configmaps core_v1 deleteCoreV1CollectionNamespacedConfigMap

delete collection of ConfigMap

*/
type DeleteCoreV1CollectionNamespacedConfigMap struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedConfigMapHandler
}

func (o *DeleteCoreV1CollectionNamespacedConfigMap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedConfigMapParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
