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

// ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandlerFunc turns a function with the right signature into a replace core v1 namespaced persistent volume claim status handler
type ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandlerFunc func(ReplaceCoreV1NamespacedPersistentVolumeClaimStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandlerFunc) Handle(params ReplaceCoreV1NamespacedPersistentVolumeClaimStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandler interface for that can handle valid replace core v1 namespaced persistent volume claim status params
type ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandler interface {
	Handle(ReplaceCoreV1NamespacedPersistentVolumeClaimStatusParams) middleware.Responder
}

// NewReplaceCoreV1NamespacedPersistentVolumeClaimStatus creates a new http.Handler for the replace core v1 namespaced persistent volume claim status operation
func NewReplaceCoreV1NamespacedPersistentVolumeClaimStatus(ctx *middleware.Context, handler ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandler) *ReplaceCoreV1NamespacedPersistentVolumeClaimStatus {
	return &ReplaceCoreV1NamespacedPersistentVolumeClaimStatus{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespacedPersistentVolumeClaimStatus swagger:route PUT /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status core_v1 replaceCoreV1NamespacedPersistentVolumeClaimStatus

replace status of the specified PersistentVolumeClaim

*/
type ReplaceCoreV1NamespacedPersistentVolumeClaimStatus struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandler
}

func (o *ReplaceCoreV1NamespacedPersistentVolumeClaimStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespacedPersistentVolumeClaimStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}