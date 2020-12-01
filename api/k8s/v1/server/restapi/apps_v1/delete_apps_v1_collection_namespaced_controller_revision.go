// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAppsV1CollectionNamespacedControllerRevisionHandlerFunc turns a function with the right signature into a delete apps v1 collection namespaced controller revision handler
type DeleteAppsV1CollectionNamespacedControllerRevisionHandlerFunc func(DeleteAppsV1CollectionNamespacedControllerRevisionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAppsV1CollectionNamespacedControllerRevisionHandlerFunc) Handle(params DeleteAppsV1CollectionNamespacedControllerRevisionParams) middleware.Responder {
	return fn(params)
}

// DeleteAppsV1CollectionNamespacedControllerRevisionHandler interface for that can handle valid delete apps v1 collection namespaced controller revision params
type DeleteAppsV1CollectionNamespacedControllerRevisionHandler interface {
	Handle(DeleteAppsV1CollectionNamespacedControllerRevisionParams) middleware.Responder
}

// NewDeleteAppsV1CollectionNamespacedControllerRevision creates a new http.Handler for the delete apps v1 collection namespaced controller revision operation
func NewDeleteAppsV1CollectionNamespacedControllerRevision(ctx *middleware.Context, handler DeleteAppsV1CollectionNamespacedControllerRevisionHandler) *DeleteAppsV1CollectionNamespacedControllerRevision {
	return &DeleteAppsV1CollectionNamespacedControllerRevision{Context: ctx, Handler: handler}
}

/*DeleteAppsV1CollectionNamespacedControllerRevision swagger:route DELETE /apis/apps/v1/namespaces/{namespace}/controllerrevisions apps_v1 deleteAppsV1CollectionNamespacedControllerRevision

delete collection of ControllerRevision

*/
type DeleteAppsV1CollectionNamespacedControllerRevision struct {
	Context *middleware.Context
	Handler DeleteAppsV1CollectionNamespacedControllerRevisionHandler
}

func (o *DeleteAppsV1CollectionNamespacedControllerRevision) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAppsV1CollectionNamespacedControllerRevisionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
