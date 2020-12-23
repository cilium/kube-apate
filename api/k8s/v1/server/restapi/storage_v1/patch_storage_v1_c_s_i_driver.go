// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchStorageV1CSIDriverHandlerFunc turns a function with the right signature into a patch storage v1 c s i driver handler
type PatchStorageV1CSIDriverHandlerFunc func(PatchStorageV1CSIDriverParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchStorageV1CSIDriverHandlerFunc) Handle(params PatchStorageV1CSIDriverParams) middleware.Responder {
	return fn(params)
}

// PatchStorageV1CSIDriverHandler interface for that can handle valid patch storage v1 c s i driver params
type PatchStorageV1CSIDriverHandler interface {
	Handle(PatchStorageV1CSIDriverParams) middleware.Responder
}

// NewPatchStorageV1CSIDriver creates a new http.Handler for the patch storage v1 c s i driver operation
func NewPatchStorageV1CSIDriver(ctx *middleware.Context, handler PatchStorageV1CSIDriverHandler) *PatchStorageV1CSIDriver {
	return &PatchStorageV1CSIDriver{Context: ctx, Handler: handler}
}

/*PatchStorageV1CSIDriver swagger:route PATCH /apis/storage.k8s.io/v1/csidrivers/{name} storage_v1 patchStorageV1CSIDriver

partially update the specified CSIDriver

*/
type PatchStorageV1CSIDriver struct {
	Context *middleware.Context
	Handler PatchStorageV1CSIDriverHandler
}

func (o *PatchStorageV1CSIDriver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchStorageV1CSIDriverParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}