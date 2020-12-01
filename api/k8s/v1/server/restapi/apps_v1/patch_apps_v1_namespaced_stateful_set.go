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

// PatchAppsV1NamespacedStatefulSetHandlerFunc turns a function with the right signature into a patch apps v1 namespaced stateful set handler
type PatchAppsV1NamespacedStatefulSetHandlerFunc func(PatchAppsV1NamespacedStatefulSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedStatefulSetHandlerFunc) Handle(params PatchAppsV1NamespacedStatefulSetParams) middleware.Responder {
	return fn(params)
}

// PatchAppsV1NamespacedStatefulSetHandler interface for that can handle valid patch apps v1 namespaced stateful set params
type PatchAppsV1NamespacedStatefulSetHandler interface {
	Handle(PatchAppsV1NamespacedStatefulSetParams) middleware.Responder
}

// NewPatchAppsV1NamespacedStatefulSet creates a new http.Handler for the patch apps v1 namespaced stateful set operation
func NewPatchAppsV1NamespacedStatefulSet(ctx *middleware.Context, handler PatchAppsV1NamespacedStatefulSetHandler) *PatchAppsV1NamespacedStatefulSet {
	return &PatchAppsV1NamespacedStatefulSet{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedStatefulSet swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/statefulsets/{name} apps_v1 patchAppsV1NamespacedStatefulSet

partially update the specified StatefulSet

*/
type PatchAppsV1NamespacedStatefulSet struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedStatefulSetHandler
}

func (o *PatchAppsV1NamespacedStatefulSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedStatefulSetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
