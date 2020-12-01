// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadApiregistrationV1beta1APIServiceHandlerFunc turns a function with the right signature into a read apiregistration v1beta1 API service handler
type ReadApiregistrationV1beta1APIServiceHandlerFunc func(ReadApiregistrationV1beta1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadApiregistrationV1beta1APIServiceHandlerFunc) Handle(params ReadApiregistrationV1beta1APIServiceParams) middleware.Responder {
	return fn(params)
}

// ReadApiregistrationV1beta1APIServiceHandler interface for that can handle valid read apiregistration v1beta1 API service params
type ReadApiregistrationV1beta1APIServiceHandler interface {
	Handle(ReadApiregistrationV1beta1APIServiceParams) middleware.Responder
}

// NewReadApiregistrationV1beta1APIService creates a new http.Handler for the read apiregistration v1beta1 API service operation
func NewReadApiregistrationV1beta1APIService(ctx *middleware.Context, handler ReadApiregistrationV1beta1APIServiceHandler) *ReadApiregistrationV1beta1APIService {
	return &ReadApiregistrationV1beta1APIService{Context: ctx, Handler: handler}
}

/*ReadApiregistrationV1beta1APIService swagger:route GET /apis/apiregistration.k8s.io/v1beta1/apiservices/{name} apiregistration_v1beta1 readApiregistrationV1beta1ApiService

read the specified APIService

*/
type ReadApiregistrationV1beta1APIService struct {
	Context *middleware.Context
	Handler ReadApiregistrationV1beta1APIServiceHandler
}

func (o *ReadApiregistrationV1beta1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadApiregistrationV1beta1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
