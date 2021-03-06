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

// ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandlerFunc turns a function with the right signature into a list policy v1beta1 pod disruption budget for all namespaces handler
type ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandlerFunc func(ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandlerFunc) Handle(params ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandler interface for that can handle valid list policy v1beta1 pod disruption budget for all namespaces params
type ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandler interface {
	Handle(ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesParams) middleware.Responder
}

// NewListPolicyV1beta1PodDisruptionBudgetForAllNamespaces creates a new http.Handler for the list policy v1beta1 pod disruption budget for all namespaces operation
func NewListPolicyV1beta1PodDisruptionBudgetForAllNamespaces(ctx *middleware.Context, handler ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandler) *ListPolicyV1beta1PodDisruptionBudgetForAllNamespaces {
	return &ListPolicyV1beta1PodDisruptionBudgetForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListPolicyV1beta1PodDisruptionBudgetForAllNamespaces swagger:route GET /apis/policy/v1beta1/poddisruptionbudgets policy_v1beta1 listPolicyV1beta1PodDisruptionBudgetForAllNamespaces

list or watch objects of kind PodDisruptionBudget

*/
type ListPolicyV1beta1PodDisruptionBudgetForAllNamespaces struct {
	Context *middleware.Context
	Handler ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandler
}

func (o *ListPolicyV1beta1PodDisruptionBudgetForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListPolicyV1beta1PodDisruptionBudgetForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
