// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandlerFunc turns a function with the right signature into a watch policy v1beta1 namespaced pod disruption budget list handler
type WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandlerFunc func(WatchPolicyV1beta1NamespacedPodDisruptionBudgetListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandlerFunc) Handle(params WatchPolicyV1beta1NamespacedPodDisruptionBudgetListParams) middleware.Responder {
	return fn(params)
}

// WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandler interface for that can handle valid watch policy v1beta1 namespaced pod disruption budget list params
type WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandler interface {
	Handle(WatchPolicyV1beta1NamespacedPodDisruptionBudgetListParams) middleware.Responder
}

// NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetList creates a new http.Handler for the watch policy v1beta1 namespaced pod disruption budget list operation
func NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetList(ctx *middleware.Context, handler WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandler) *WatchPolicyV1beta1NamespacedPodDisruptionBudgetList {
	return &WatchPolicyV1beta1NamespacedPodDisruptionBudgetList{Context: ctx, Handler: handler}
}

/*WatchPolicyV1beta1NamespacedPodDisruptionBudgetList swagger:route GET /apis/policy/v1beta1/watch/namespaces/{namespace}/poddisruptionbudgets policy_v1beta1 watchPolicyV1beta1NamespacedPodDisruptionBudgetList

watch individual changes to a list of PodDisruptionBudget. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchPolicyV1beta1NamespacedPodDisruptionBudgetList struct {
	Context *middleware.Context
	Handler WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandler
}

func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}