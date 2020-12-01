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

// CreateApiregistrationV1APIServiceHandlerFunc turns a function with the right signature into a create apiregistration v1 API service handler
type CreateApiregistrationV1APIServiceHandlerFunc func(CreateApiregistrationV1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateApiregistrationV1APIServiceHandlerFunc) Handle(params CreateApiregistrationV1APIServiceParams) middleware.Responder {
	return fn(params)
}

// CreateApiregistrationV1APIServiceHandler interface for that can handle valid create apiregistration v1 API service params
type CreateApiregistrationV1APIServiceHandler interface {
	Handle(CreateApiregistrationV1APIServiceParams) middleware.Responder
}

// NewCreateApiregistrationV1APIService creates a new http.Handler for the create apiregistration v1 API service operation
func NewCreateApiregistrationV1APIService(ctx *middleware.Context, handler CreateApiregistrationV1APIServiceHandler) *CreateApiregistrationV1APIService {
	return &CreateApiregistrationV1APIService{Context: ctx, Handler: handler}
}

/*CreateApiregistrationV1APIService swagger:route POST /apis/apiregistration.k8s.io/v1/apiservices apiregistration_v1 createApiregistrationV1ApiService

create an APIService

*/
type CreateApiregistrationV1APIService struct {
	Context *middleware.Context
	Handler CreateApiregistrationV1APIServiceHandler
}

func (o *CreateApiregistrationV1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateApiregistrationV1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
