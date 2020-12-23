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

// DeleteAppsV1NamespacedDaemonSetHandlerFunc turns a function with the right signature into a delete apps v1 namespaced daemon set handler
type DeleteAppsV1NamespacedDaemonSetHandlerFunc func(DeleteAppsV1NamespacedDaemonSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAppsV1NamespacedDaemonSetHandlerFunc) Handle(params DeleteAppsV1NamespacedDaemonSetParams) middleware.Responder {
	return fn(params)
}

// DeleteAppsV1NamespacedDaemonSetHandler interface for that can handle valid delete apps v1 namespaced daemon set params
type DeleteAppsV1NamespacedDaemonSetHandler interface {
	Handle(DeleteAppsV1NamespacedDaemonSetParams) middleware.Responder
}

// NewDeleteAppsV1NamespacedDaemonSet creates a new http.Handler for the delete apps v1 namespaced daemon set operation
func NewDeleteAppsV1NamespacedDaemonSet(ctx *middleware.Context, handler DeleteAppsV1NamespacedDaemonSetHandler) *DeleteAppsV1NamespacedDaemonSet {
	return &DeleteAppsV1NamespacedDaemonSet{Context: ctx, Handler: handler}
}

/*DeleteAppsV1NamespacedDaemonSet swagger:route DELETE /apis/apps/v1/namespaces/{namespace}/daemonsets/{name} apps_v1 deleteAppsV1NamespacedDaemonSet

delete a DaemonSet

*/
type DeleteAppsV1NamespacedDaemonSet struct {
	Context *middleware.Context
	Handler DeleteAppsV1NamespacedDaemonSetHandler
}

func (o *DeleteAppsV1NamespacedDaemonSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAppsV1NamespacedDaemonSetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}