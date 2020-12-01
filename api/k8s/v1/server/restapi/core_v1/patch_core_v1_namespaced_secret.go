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

// PatchCoreV1NamespacedSecretHandlerFunc turns a function with the right signature into a patch core v1 namespaced secret handler
type PatchCoreV1NamespacedSecretHandlerFunc func(PatchCoreV1NamespacedSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchCoreV1NamespacedSecretHandlerFunc) Handle(params PatchCoreV1NamespacedSecretParams) middleware.Responder {
	return fn(params)
}

// PatchCoreV1NamespacedSecretHandler interface for that can handle valid patch core v1 namespaced secret params
type PatchCoreV1NamespacedSecretHandler interface {
	Handle(PatchCoreV1NamespacedSecretParams) middleware.Responder
}

// NewPatchCoreV1NamespacedSecret creates a new http.Handler for the patch core v1 namespaced secret operation
func NewPatchCoreV1NamespacedSecret(ctx *middleware.Context, handler PatchCoreV1NamespacedSecretHandler) *PatchCoreV1NamespacedSecret {
	return &PatchCoreV1NamespacedSecret{Context: ctx, Handler: handler}
}

/*PatchCoreV1NamespacedSecret swagger:route PATCH /api/v1/namespaces/{namespace}/secrets/{name} core_v1 patchCoreV1NamespacedSecret

partially update the specified Secret

*/
type PatchCoreV1NamespacedSecret struct {
	Context *middleware.Context
	Handler PatchCoreV1NamespacedSecretHandler
}

func (o *PatchCoreV1NamespacedSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchCoreV1NamespacedSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
