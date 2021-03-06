// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAutoscalingV2beta2APIResourcesHandlerFunc turns a function with the right signature into a get autoscaling v2beta2 API resources handler
type GetAutoscalingV2beta2APIResourcesHandlerFunc func(GetAutoscalingV2beta2APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAutoscalingV2beta2APIResourcesHandlerFunc) Handle(params GetAutoscalingV2beta2APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetAutoscalingV2beta2APIResourcesHandler interface for that can handle valid get autoscaling v2beta2 API resources params
type GetAutoscalingV2beta2APIResourcesHandler interface {
	Handle(GetAutoscalingV2beta2APIResourcesParams) middleware.Responder
}

// NewGetAutoscalingV2beta2APIResources creates a new http.Handler for the get autoscaling v2beta2 API resources operation
func NewGetAutoscalingV2beta2APIResources(ctx *middleware.Context, handler GetAutoscalingV2beta2APIResourcesHandler) *GetAutoscalingV2beta2APIResources {
	return &GetAutoscalingV2beta2APIResources{Context: ctx, Handler: handler}
}

/*GetAutoscalingV2beta2APIResources swagger:route GET /apis/autoscaling/v2beta2/ autoscaling_v2beta2 getAutoscalingV2beta2ApiResources

get available resources

*/
type GetAutoscalingV2beta2APIResources struct {
	Context *middleware.Context
	Handler GetAutoscalingV2beta2APIResourcesHandler
}

func (o *GetAutoscalingV2beta2APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAutoscalingV2beta2APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
