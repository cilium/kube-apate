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

// ListRbacAuthorizationV1ClusterRoleBindingHandlerFunc turns a function with the right signature into a list rbac authorization v1 cluster role binding handler
type ListRbacAuthorizationV1ClusterRoleBindingHandlerFunc func(ListRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbacAuthorizationV1ClusterRoleBindingHandlerFunc) Handle(params ListRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
	return fn(params)
}

// ListRbacAuthorizationV1ClusterRoleBindingHandler interface for that can handle valid list rbac authorization v1 cluster role binding params
type ListRbacAuthorizationV1ClusterRoleBindingHandler interface {
	Handle(ListRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder
}

// NewListRbacAuthorizationV1ClusterRoleBinding creates a new http.Handler for the list rbac authorization v1 cluster role binding operation
func NewListRbacAuthorizationV1ClusterRoleBinding(ctx *middleware.Context, handler ListRbacAuthorizationV1ClusterRoleBindingHandler) *ListRbacAuthorizationV1ClusterRoleBinding {
	return &ListRbacAuthorizationV1ClusterRoleBinding{Context: ctx, Handler: handler}
}

/*ListRbacAuthorizationV1ClusterRoleBinding swagger:route GET /apis/rbac.authorization.k8s.io/v1/clusterrolebindings rbacAuthorization_v1 listRbacAuthorizationV1ClusterRoleBinding

list or watch objects of kind ClusterRoleBinding

*/
type ListRbacAuthorizationV1ClusterRoleBinding struct {
	Context *middleware.Context
	Handler ListRbacAuthorizationV1ClusterRoleBindingHandler
}

func (o *ListRbacAuthorizationV1ClusterRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbacAuthorizationV1ClusterRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
