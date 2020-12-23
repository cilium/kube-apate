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

// ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandlerFunc turns a function with the right signature into a list autoscaling v2beta2 horizontal pod autoscaler for all namespaces handler
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandlerFunc func(ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandlerFunc) Handle(params ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandler interface for that can handle valid list autoscaling v2beta2 horizontal pod autoscaler for all namespaces params
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandler interface {
	Handle(ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesParams) middleware.Responder
}

// NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces creates a new http.Handler for the list autoscaling v2beta2 horizontal pod autoscaler for all namespaces operation
func NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces(ctx *middleware.Context, handler ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandler) *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces {
	return &ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces swagger:route GET /apis/autoscaling/v2beta2/horizontalpodautoscalers autoscaling_v2beta2 listAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces

list or watch objects of kind HorizontalPodAutoscaler

*/
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces struct {
	Context *middleware.Context
	Handler ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandler
}

func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}