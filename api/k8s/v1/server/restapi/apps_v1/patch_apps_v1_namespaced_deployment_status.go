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

// PatchAppsV1NamespacedDeploymentStatusHandlerFunc turns a function with the right signature into a patch apps v1 namespaced deployment status handler
type PatchAppsV1NamespacedDeploymentStatusHandlerFunc func(PatchAppsV1NamespacedDeploymentStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedDeploymentStatusHandlerFunc) Handle(params PatchAppsV1NamespacedDeploymentStatusParams) middleware.Responder {
	return fn(params)
}

// PatchAppsV1NamespacedDeploymentStatusHandler interface for that can handle valid patch apps v1 namespaced deployment status params
type PatchAppsV1NamespacedDeploymentStatusHandler interface {
	Handle(PatchAppsV1NamespacedDeploymentStatusParams) middleware.Responder
}

// NewPatchAppsV1NamespacedDeploymentStatus creates a new http.Handler for the patch apps v1 namespaced deployment status operation
func NewPatchAppsV1NamespacedDeploymentStatus(ctx *middleware.Context, handler PatchAppsV1NamespacedDeploymentStatusHandler) *PatchAppsV1NamespacedDeploymentStatus {
	return &PatchAppsV1NamespacedDeploymentStatus{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedDeploymentStatus swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status apps_v1 patchAppsV1NamespacedDeploymentStatus

partially update status of the specified Deployment

*/
type PatchAppsV1NamespacedDeploymentStatus struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedDeploymentStatusHandler
}

func (o *PatchAppsV1NamespacedDeploymentStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedDeploymentStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
