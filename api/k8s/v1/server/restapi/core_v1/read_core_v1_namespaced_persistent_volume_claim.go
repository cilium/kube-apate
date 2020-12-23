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

// ReadCoreV1NamespacedPersistentVolumeClaimHandlerFunc turns a function with the right signature into a read core v1 namespaced persistent volume claim handler
type ReadCoreV1NamespacedPersistentVolumeClaimHandlerFunc func(ReadCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedPersistentVolumeClaimHandlerFunc) Handle(params ReadCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespacedPersistentVolumeClaimHandler interface for that can handle valid read core v1 namespaced persistent volume claim params
type ReadCoreV1NamespacedPersistentVolumeClaimHandler interface {
	Handle(ReadCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder
}

// NewReadCoreV1NamespacedPersistentVolumeClaim creates a new http.Handler for the read core v1 namespaced persistent volume claim operation
func NewReadCoreV1NamespacedPersistentVolumeClaim(ctx *middleware.Context, handler ReadCoreV1NamespacedPersistentVolumeClaimHandler) *ReadCoreV1NamespacedPersistentVolumeClaim {
	return &ReadCoreV1NamespacedPersistentVolumeClaim{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedPersistentVolumeClaim swagger:route GET /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} core_v1 readCoreV1NamespacedPersistentVolumeClaim

read the specified PersistentVolumeClaim

*/
type ReadCoreV1NamespacedPersistentVolumeClaim struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedPersistentVolumeClaimHandler
}

func (o *ReadCoreV1NamespacedPersistentVolumeClaim) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedPersistentVolumeClaimParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}