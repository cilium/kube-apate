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

// DeleteAppsV1CollectionNamespacedDeploymentHandlerFunc turns a function with the right signature into a delete apps v1 collection namespaced deployment handler
type DeleteAppsV1CollectionNamespacedDeploymentHandlerFunc func(DeleteAppsV1CollectionNamespacedDeploymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAppsV1CollectionNamespacedDeploymentHandlerFunc) Handle(params DeleteAppsV1CollectionNamespacedDeploymentParams) middleware.Responder {
	return fn(params)
}

// DeleteAppsV1CollectionNamespacedDeploymentHandler interface for that can handle valid delete apps v1 collection namespaced deployment params
type DeleteAppsV1CollectionNamespacedDeploymentHandler interface {
	Handle(DeleteAppsV1CollectionNamespacedDeploymentParams) middleware.Responder
}

// NewDeleteAppsV1CollectionNamespacedDeployment creates a new http.Handler for the delete apps v1 collection namespaced deployment operation
func NewDeleteAppsV1CollectionNamespacedDeployment(ctx *middleware.Context, handler DeleteAppsV1CollectionNamespacedDeploymentHandler) *DeleteAppsV1CollectionNamespacedDeployment {
	return &DeleteAppsV1CollectionNamespacedDeployment{Context: ctx, Handler: handler}
}

/*DeleteAppsV1CollectionNamespacedDeployment swagger:route DELETE /apis/apps/v1/namespaces/{namespace}/deployments apps_v1 deleteAppsV1CollectionNamespacedDeployment

delete collection of Deployment

*/
type DeleteAppsV1CollectionNamespacedDeployment struct {
	Context *middleware.Context
	Handler DeleteAppsV1CollectionNamespacedDeploymentHandler
}

func (o *DeleteAppsV1CollectionNamespacedDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAppsV1CollectionNamespacedDeploymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
