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

// CreateApiregistrationV1beta1APIServiceHandlerFunc turns a function with the right signature into a create apiregistration v1beta1 API service handler
type CreateApiregistrationV1beta1APIServiceHandlerFunc func(CreateApiregistrationV1beta1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateApiregistrationV1beta1APIServiceHandlerFunc) Handle(params CreateApiregistrationV1beta1APIServiceParams) middleware.Responder {
	return fn(params)
}

// CreateApiregistrationV1beta1APIServiceHandler interface for that can handle valid create apiregistration v1beta1 API service params
type CreateApiregistrationV1beta1APIServiceHandler interface {
	Handle(CreateApiregistrationV1beta1APIServiceParams) middleware.Responder
}

// NewCreateApiregistrationV1beta1APIService creates a new http.Handler for the create apiregistration v1beta1 API service operation
func NewCreateApiregistrationV1beta1APIService(ctx *middleware.Context, handler CreateApiregistrationV1beta1APIServiceHandler) *CreateApiregistrationV1beta1APIService {
	return &CreateApiregistrationV1beta1APIService{Context: ctx, Handler: handler}
}

/*CreateApiregistrationV1beta1APIService swagger:route POST /apis/apiregistration.k8s.io/v1beta1/apiservices apiregistration_v1beta1 createApiregistrationV1beta1ApiService

create an APIService

*/
type CreateApiregistrationV1beta1APIService struct {
	Context *middleware.Context
	Handler CreateApiregistrationV1beta1APIServiceHandler
}

func (o *CreateApiregistrationV1beta1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateApiregistrationV1beta1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
