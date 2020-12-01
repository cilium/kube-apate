// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchApiregistrationV1APIServiceHandlerFunc turns a function with the right signature into a patch apiregistration v1 API service handler
type PatchApiregistrationV1APIServiceHandlerFunc func(PatchApiregistrationV1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchApiregistrationV1APIServiceHandlerFunc) Handle(params PatchApiregistrationV1APIServiceParams) middleware.Responder {
	return fn(params)
}

// PatchApiregistrationV1APIServiceHandler interface for that can handle valid patch apiregistration v1 API service params
type PatchApiregistrationV1APIServiceHandler interface {
	Handle(PatchApiregistrationV1APIServiceParams) middleware.Responder
}

// NewPatchApiregistrationV1APIService creates a new http.Handler for the patch apiregistration v1 API service operation
func NewPatchApiregistrationV1APIService(ctx *middleware.Context, handler PatchApiregistrationV1APIServiceHandler) *PatchApiregistrationV1APIService {
	return &PatchApiregistrationV1APIService{Context: ctx, Handler: handler}
}

/*PatchApiregistrationV1APIService swagger:route PATCH /apis/apiregistration.k8s.io/v1/apiservices/{name} apiregistration_v1 patchApiregistrationV1ApiService

partially update the specified APIService

*/
type PatchApiregistrationV1APIService struct {
	Context *middleware.Context
	Handler PatchApiregistrationV1APIServiceHandler
}

func (o *PatchApiregistrationV1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchApiregistrationV1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
