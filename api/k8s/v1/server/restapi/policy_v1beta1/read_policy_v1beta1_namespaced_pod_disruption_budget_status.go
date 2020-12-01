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

// ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc turns a function with the right signature into a read policy v1beta1 namespaced pod disruption budget status handler
type ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc func(ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc) Handle(params ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder {
	return fn(params)
}

// ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler interface for that can handle valid read policy v1beta1 namespaced pod disruption budget status params
type ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler interface {
	Handle(ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder
}

// NewReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus creates a new http.Handler for the read policy v1beta1 namespaced pod disruption budget status operation
func NewReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus(ctx *middleware.Context, handler ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler) *ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus {
	return &ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus{Context: ctx, Handler: handler}
}

/*ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus swagger:route GET /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status policy_v1beta1 readPolicyV1beta1NamespacedPodDisruptionBudgetStatus

read status of the specified PodDisruptionBudget

*/
type ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus struct {
	Context *middleware.Context
	Handler ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler
}

func (o *ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
