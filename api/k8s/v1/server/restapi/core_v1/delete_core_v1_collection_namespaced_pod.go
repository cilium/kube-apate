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

// DeleteCoreV1CollectionNamespacedPodHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced pod handler
type DeleteCoreV1CollectionNamespacedPodHandlerFunc func(DeleteCoreV1CollectionNamespacedPodParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedPodHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedPodParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1CollectionNamespacedPodHandler interface for that can handle valid delete core v1 collection namespaced pod params
type DeleteCoreV1CollectionNamespacedPodHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedPodParams) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedPod creates a new http.Handler for the delete core v1 collection namespaced pod operation
func NewDeleteCoreV1CollectionNamespacedPod(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedPodHandler) *DeleteCoreV1CollectionNamespacedPod {
	return &DeleteCoreV1CollectionNamespacedPod{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedPod swagger:route DELETE /api/v1/namespaces/{namespace}/pods core_v1 deleteCoreV1CollectionNamespacedPod

delete collection of Pod

*/
type DeleteCoreV1CollectionNamespacedPod struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedPodHandler
}

func (o *DeleteCoreV1CollectionNamespacedPod) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedPodParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
