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

// ReadPolicyV1beta1PodSecurityPolicyHandlerFunc turns a function with the right signature into a read policy v1beta1 pod security policy handler
type ReadPolicyV1beta1PodSecurityPolicyHandlerFunc func(ReadPolicyV1beta1PodSecurityPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadPolicyV1beta1PodSecurityPolicyHandlerFunc) Handle(params ReadPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
	return fn(params)
}

// ReadPolicyV1beta1PodSecurityPolicyHandler interface for that can handle valid read policy v1beta1 pod security policy params
type ReadPolicyV1beta1PodSecurityPolicyHandler interface {
	Handle(ReadPolicyV1beta1PodSecurityPolicyParams) middleware.Responder
}

// NewReadPolicyV1beta1PodSecurityPolicy creates a new http.Handler for the read policy v1beta1 pod security policy operation
func NewReadPolicyV1beta1PodSecurityPolicy(ctx *middleware.Context, handler ReadPolicyV1beta1PodSecurityPolicyHandler) *ReadPolicyV1beta1PodSecurityPolicy {
	return &ReadPolicyV1beta1PodSecurityPolicy{Context: ctx, Handler: handler}
}

/*ReadPolicyV1beta1PodSecurityPolicy swagger:route GET /apis/policy/v1beta1/podsecuritypolicies/{name} policy_v1beta1 readPolicyV1beta1PodSecurityPolicy

read the specified PodSecurityPolicy

*/
type ReadPolicyV1beta1PodSecurityPolicy struct {
	Context *middleware.Context
	Handler ReadPolicyV1beta1PodSecurityPolicyHandler
}

func (o *ReadPolicyV1beta1PodSecurityPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadPolicyV1beta1PodSecurityPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
