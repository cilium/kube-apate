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

// ReplaceApiregistrationV1APIServiceHandlerFunc turns a function with the right signature into a replace apiregistration v1 API service handler
type ReplaceApiregistrationV1APIServiceHandlerFunc func(ReplaceApiregistrationV1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceApiregistrationV1APIServiceHandlerFunc) Handle(params ReplaceApiregistrationV1APIServiceParams) middleware.Responder {
	return fn(params)
}

// ReplaceApiregistrationV1APIServiceHandler interface for that can handle valid replace apiregistration v1 API service params
type ReplaceApiregistrationV1APIServiceHandler interface {
	Handle(ReplaceApiregistrationV1APIServiceParams) middleware.Responder
}

// NewReplaceApiregistrationV1APIService creates a new http.Handler for the replace apiregistration v1 API service operation
func NewReplaceApiregistrationV1APIService(ctx *middleware.Context, handler ReplaceApiregistrationV1APIServiceHandler) *ReplaceApiregistrationV1APIService {
	return &ReplaceApiregistrationV1APIService{Context: ctx, Handler: handler}
}

/*ReplaceApiregistrationV1APIService swagger:route PUT /apis/apiregistration.k8s.io/v1/apiservices/{name} apiregistration_v1 replaceApiregistrationV1ApiService

replace the specified APIService

*/
type ReplaceApiregistrationV1APIService struct {
	Context *middleware.Context
	Handler ReplaceApiregistrationV1APIServiceHandler
}

func (o *ReplaceApiregistrationV1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceApiregistrationV1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
