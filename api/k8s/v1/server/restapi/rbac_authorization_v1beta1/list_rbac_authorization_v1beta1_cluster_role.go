// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRbacAuthorizationV1beta1ClusterRoleHandlerFunc turns a function with the right signature into a list rbac authorization v1beta1 cluster role handler
type ListRbacAuthorizationV1beta1ClusterRoleHandlerFunc func(ListRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbacAuthorizationV1beta1ClusterRoleHandlerFunc) Handle(params ListRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// ListRbacAuthorizationV1beta1ClusterRoleHandler interface for that can handle valid list rbac authorization v1beta1 cluster role params
type ListRbacAuthorizationV1beta1ClusterRoleHandler interface {
	Handle(ListRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder
}

// NewListRbacAuthorizationV1beta1ClusterRole creates a new http.Handler for the list rbac authorization v1beta1 cluster role operation
func NewListRbacAuthorizationV1beta1ClusterRole(ctx *middleware.Context, handler ListRbacAuthorizationV1beta1ClusterRoleHandler) *ListRbacAuthorizationV1beta1ClusterRole {
	return &ListRbacAuthorizationV1beta1ClusterRole{Context: ctx, Handler: handler}
}

/*ListRbacAuthorizationV1beta1ClusterRole swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/clusterroles rbacAuthorization_v1beta1 listRbacAuthorizationV1beta1ClusterRole

list or watch objects of kind ClusterRole

*/
type ListRbacAuthorizationV1beta1ClusterRole struct {
	Context *middleware.Context
	Handler ListRbacAuthorizationV1beta1ClusterRoleHandler
}

func (o *ListRbacAuthorizationV1beta1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbacAuthorizationV1beta1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
