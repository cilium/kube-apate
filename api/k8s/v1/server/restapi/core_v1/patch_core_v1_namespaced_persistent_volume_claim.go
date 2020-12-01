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

// PatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc turns a function with the right signature into a patch core v1 namespaced persistent volume claim handler
type PatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc func(PatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc) Handle(params PatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
	return fn(params)
}

// PatchCoreV1NamespacedPersistentVolumeClaimHandler interface for that can handle valid patch core v1 namespaced persistent volume claim params
type PatchCoreV1NamespacedPersistentVolumeClaimHandler interface {
	Handle(PatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder
}

// NewPatchCoreV1NamespacedPersistentVolumeClaim creates a new http.Handler for the patch core v1 namespaced persistent volume claim operation
func NewPatchCoreV1NamespacedPersistentVolumeClaim(ctx *middleware.Context, handler PatchCoreV1NamespacedPersistentVolumeClaimHandler) *PatchCoreV1NamespacedPersistentVolumeClaim {
	return &PatchCoreV1NamespacedPersistentVolumeClaim{Context: ctx, Handler: handler}
}

/*PatchCoreV1NamespacedPersistentVolumeClaim swagger:route PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} core_v1 patchCoreV1NamespacedPersistentVolumeClaim

partially update the specified PersistentVolumeClaim

*/
type PatchCoreV1NamespacedPersistentVolumeClaim struct {
	Context *middleware.Context
	Handler PatchCoreV1NamespacedPersistentVolumeClaimHandler
}

func (o *PatchCoreV1NamespacedPersistentVolumeClaim) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchCoreV1NamespacedPersistentVolumeClaimParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
