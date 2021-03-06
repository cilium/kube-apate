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

// PatchAppsV1NamespacedDeploymentScaleHandlerFunc turns a function with the right signature into a patch apps v1 namespaced deployment scale handler
type PatchAppsV1NamespacedDeploymentScaleHandlerFunc func(PatchAppsV1NamespacedDeploymentScaleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedDeploymentScaleHandlerFunc) Handle(params PatchAppsV1NamespacedDeploymentScaleParams) middleware.Responder {
	return fn(params)
}

// PatchAppsV1NamespacedDeploymentScaleHandler interface for that can handle valid patch apps v1 namespaced deployment scale params
type PatchAppsV1NamespacedDeploymentScaleHandler interface {
	Handle(PatchAppsV1NamespacedDeploymentScaleParams) middleware.Responder
}

// NewPatchAppsV1NamespacedDeploymentScale creates a new http.Handler for the patch apps v1 namespaced deployment scale operation
func NewPatchAppsV1NamespacedDeploymentScale(ctx *middleware.Context, handler PatchAppsV1NamespacedDeploymentScaleHandler) *PatchAppsV1NamespacedDeploymentScale {
	return &PatchAppsV1NamespacedDeploymentScale{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedDeploymentScale swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/deployments/{name}/scale apps_v1 patchAppsV1NamespacedDeploymentScale

partially update scale of the specified Deployment

*/
type PatchAppsV1NamespacedDeploymentScale struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedDeploymentScaleHandler
}

func (o *PatchAppsV1NamespacedDeploymentScale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedDeploymentScaleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
