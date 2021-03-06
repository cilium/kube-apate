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

// WatchApiregistrationV1APIServiceHandlerFunc turns a function with the right signature into a watch apiregistration v1 API service handler
type WatchApiregistrationV1APIServiceHandlerFunc func(WatchApiregistrationV1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchApiregistrationV1APIServiceHandlerFunc) Handle(params WatchApiregistrationV1APIServiceParams) middleware.Responder {
	return fn(params)
}

// WatchApiregistrationV1APIServiceHandler interface for that can handle valid watch apiregistration v1 API service params
type WatchApiregistrationV1APIServiceHandler interface {
	Handle(WatchApiregistrationV1APIServiceParams) middleware.Responder
}

// NewWatchApiregistrationV1APIService creates a new http.Handler for the watch apiregistration v1 API service operation
func NewWatchApiregistrationV1APIService(ctx *middleware.Context, handler WatchApiregistrationV1APIServiceHandler) *WatchApiregistrationV1APIService {
	return &WatchApiregistrationV1APIService{Context: ctx, Handler: handler}
}

/*WatchApiregistrationV1APIService swagger:route GET /apis/apiregistration.k8s.io/v1/watch/apiservices/{name} apiregistration_v1 watchApiregistrationV1ApiService

watch changes to an object of kind APIService. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchApiregistrationV1APIService struct {
	Context *middleware.Context
	Handler WatchApiregistrationV1APIServiceHandler
}

func (o *WatchApiregistrationV1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchApiregistrationV1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
