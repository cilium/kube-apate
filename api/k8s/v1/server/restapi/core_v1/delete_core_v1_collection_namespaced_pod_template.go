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

// DeleteCoreV1CollectionNamespacedPodTemplateHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced pod template handler
type DeleteCoreV1CollectionNamespacedPodTemplateHandlerFunc func(DeleteCoreV1CollectionNamespacedPodTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedPodTemplateHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedPodTemplateParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1CollectionNamespacedPodTemplateHandler interface for that can handle valid delete core v1 collection namespaced pod template params
type DeleteCoreV1CollectionNamespacedPodTemplateHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedPodTemplateParams) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedPodTemplate creates a new http.Handler for the delete core v1 collection namespaced pod template operation
func NewDeleteCoreV1CollectionNamespacedPodTemplate(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedPodTemplateHandler) *DeleteCoreV1CollectionNamespacedPodTemplate {
	return &DeleteCoreV1CollectionNamespacedPodTemplate{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedPodTemplate swagger:route DELETE /api/v1/namespaces/{namespace}/podtemplates core_v1 deleteCoreV1CollectionNamespacedPodTemplate

delete collection of PodTemplate

*/
type DeleteCoreV1CollectionNamespacedPodTemplate struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedPodTemplateHandler
}

func (o *DeleteCoreV1CollectionNamespacedPodTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedPodTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}