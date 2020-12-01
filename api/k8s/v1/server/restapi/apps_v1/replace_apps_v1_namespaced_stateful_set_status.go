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

// ReplaceAppsV1NamespacedStatefulSetStatusHandlerFunc turns a function with the right signature into a replace apps v1 namespaced stateful set status handler
type ReplaceAppsV1NamespacedStatefulSetStatusHandlerFunc func(ReplaceAppsV1NamespacedStatefulSetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAppsV1NamespacedStatefulSetStatusHandlerFunc) Handle(params ReplaceAppsV1NamespacedStatefulSetStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceAppsV1NamespacedStatefulSetStatusHandler interface for that can handle valid replace apps v1 namespaced stateful set status params
type ReplaceAppsV1NamespacedStatefulSetStatusHandler interface {
	Handle(ReplaceAppsV1NamespacedStatefulSetStatusParams) middleware.Responder
}

// NewReplaceAppsV1NamespacedStatefulSetStatus creates a new http.Handler for the replace apps v1 namespaced stateful set status operation
func NewReplaceAppsV1NamespacedStatefulSetStatus(ctx *middleware.Context, handler ReplaceAppsV1NamespacedStatefulSetStatusHandler) *ReplaceAppsV1NamespacedStatefulSetStatus {
	return &ReplaceAppsV1NamespacedStatefulSetStatus{Context: ctx, Handler: handler}
}

/*ReplaceAppsV1NamespacedStatefulSetStatus swagger:route PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status apps_v1 replaceAppsV1NamespacedStatefulSetStatus

replace status of the specified StatefulSet

*/
type ReplaceAppsV1NamespacedStatefulSetStatus struct {
	Context *middleware.Context
	Handler ReplaceAppsV1NamespacedStatefulSetStatusHandler
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAppsV1NamespacedStatefulSetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
