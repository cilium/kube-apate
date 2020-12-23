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

// ListPolicyV1beta1PodSecurityPolicyHandlerFunc turns a function with the right signature into a list policy v1beta1 pod security policy handler
type ListPolicyV1beta1PodSecurityPolicyHandlerFunc func(ListPolicyV1beta1PodSecurityPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPolicyV1beta1PodSecurityPolicyHandlerFunc) Handle(params ListPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
	return fn(params)
}

// ListPolicyV1beta1PodSecurityPolicyHandler interface for that can handle valid list policy v1beta1 pod security policy params
type ListPolicyV1beta1PodSecurityPolicyHandler interface {
	Handle(ListPolicyV1beta1PodSecurityPolicyParams) middleware.Responder
}

// NewListPolicyV1beta1PodSecurityPolicy creates a new http.Handler for the list policy v1beta1 pod security policy operation
func NewListPolicyV1beta1PodSecurityPolicy(ctx *middleware.Context, handler ListPolicyV1beta1PodSecurityPolicyHandler) *ListPolicyV1beta1PodSecurityPolicy {
	return &ListPolicyV1beta1PodSecurityPolicy{Context: ctx, Handler: handler}
}

/*ListPolicyV1beta1PodSecurityPolicy swagger:route GET /apis/policy/v1beta1/podsecuritypolicies policy_v1beta1 listPolicyV1beta1PodSecurityPolicy

list or watch objects of kind PodSecurityPolicy

*/
type ListPolicyV1beta1PodSecurityPolicy struct {
	Context *middleware.Context
	Handler ListPolicyV1beta1PodSecurityPolicyHandler
}

func (o *ListPolicyV1beta1PodSecurityPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListPolicyV1beta1PodSecurityPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}