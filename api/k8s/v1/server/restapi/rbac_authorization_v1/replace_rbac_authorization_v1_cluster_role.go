// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceRbacAuthorizationV1ClusterRoleHandlerFunc turns a function with the right signature into a replace rbac authorization v1 cluster role handler
type ReplaceRbacAuthorizationV1ClusterRoleHandlerFunc func(ReplaceRbacAuthorizationV1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceRbacAuthorizationV1ClusterRoleHandlerFunc) Handle(params ReplaceRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// ReplaceRbacAuthorizationV1ClusterRoleHandler interface for that can handle valid replace rbac authorization v1 cluster role params
type ReplaceRbacAuthorizationV1ClusterRoleHandler interface {
	Handle(ReplaceRbacAuthorizationV1ClusterRoleParams) middleware.Responder
}

// NewReplaceRbacAuthorizationV1ClusterRole creates a new http.Handler for the replace rbac authorization v1 cluster role operation
func NewReplaceRbacAuthorizationV1ClusterRole(ctx *middleware.Context, handler ReplaceRbacAuthorizationV1ClusterRoleHandler) *ReplaceRbacAuthorizationV1ClusterRole {
	return &ReplaceRbacAuthorizationV1ClusterRole{Context: ctx, Handler: handler}
}

/*ReplaceRbacAuthorizationV1ClusterRole swagger:route PUT /apis/rbac.authorization.k8s.io/v1/clusterroles/{name} rbacAuthorization_v1 replaceRbacAuthorizationV1ClusterRole

replace the specified ClusterRole

*/
type ReplaceRbacAuthorizationV1ClusterRole struct {
	Context *middleware.Context
	Handler ReplaceRbacAuthorizationV1ClusterRoleHandler
}

func (o *ReplaceRbacAuthorizationV1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceRbacAuthorizationV1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
