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

// ListAppsV1NamespacedStatefulSetHandlerFunc turns a function with the right signature into a list apps v1 namespaced stateful set handler
type ListAppsV1NamespacedStatefulSetHandlerFunc func(ListAppsV1NamespacedStatefulSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAppsV1NamespacedStatefulSetHandlerFunc) Handle(params ListAppsV1NamespacedStatefulSetParams) middleware.Responder {
	return fn(params)
}

// ListAppsV1NamespacedStatefulSetHandler interface for that can handle valid list apps v1 namespaced stateful set params
type ListAppsV1NamespacedStatefulSetHandler interface {
	Handle(ListAppsV1NamespacedStatefulSetParams) middleware.Responder
}

// NewListAppsV1NamespacedStatefulSet creates a new http.Handler for the list apps v1 namespaced stateful set operation
func NewListAppsV1NamespacedStatefulSet(ctx *middleware.Context, handler ListAppsV1NamespacedStatefulSetHandler) *ListAppsV1NamespacedStatefulSet {
	return &ListAppsV1NamespacedStatefulSet{Context: ctx, Handler: handler}
}

/*ListAppsV1NamespacedStatefulSet swagger:route GET /apis/apps/v1/namespaces/{namespace}/statefulsets apps_v1 listAppsV1NamespacedStatefulSet

list or watch objects of kind StatefulSet

*/
type ListAppsV1NamespacedStatefulSet struct {
	Context *middleware.Context
	Handler ListAppsV1NamespacedStatefulSetHandler
}

func (o *ListAppsV1NamespacedStatefulSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAppsV1NamespacedStatefulSetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
