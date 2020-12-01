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

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandlerFunc turns a function with the right signature into a replace autoscaling v2beta2 namespaced horizontal pod autoscaler status handler
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandlerFunc func(ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandlerFunc) Handle(params ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler interface for that can handle valid replace autoscaling v2beta2 namespaced horizontal pod autoscaler status params
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler interface {
	Handle(ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder
}

// NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus creates a new http.Handler for the replace autoscaling v2beta2 namespaced horizontal pod autoscaler status operation
func NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus(ctx *middleware.Context, handler ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler) *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus {
	return &ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus{Context: ctx, Handler: handler}
}

/*ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus swagger:route PUT /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers/{name}/status autoscaling_v2beta2 replaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus

replace status of the specified HorizontalPodAutoscaler

*/
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus struct {
	Context *middleware.Context
	Handler ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler
}

func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
