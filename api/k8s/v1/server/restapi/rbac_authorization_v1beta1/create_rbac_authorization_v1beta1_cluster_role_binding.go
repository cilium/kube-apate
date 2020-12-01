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

// CreateRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc turns a function with the right signature into a create rbac authorization v1beta1 cluster role binding handler
type CreateRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc func(CreateRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc) Handle(params CreateRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
	return fn(params)
}

// CreateRbacAuthorizationV1beta1ClusterRoleBindingHandler interface for that can handle valid create rbac authorization v1beta1 cluster role binding params
type CreateRbacAuthorizationV1beta1ClusterRoleBindingHandler interface {
	Handle(CreateRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBinding creates a new http.Handler for the create rbac authorization v1beta1 cluster role binding operation
func NewCreateRbacAuthorizationV1beta1ClusterRoleBinding(ctx *middleware.Context, handler CreateRbacAuthorizationV1beta1ClusterRoleBindingHandler) *CreateRbacAuthorizationV1beta1ClusterRoleBinding {
	return &CreateRbacAuthorizationV1beta1ClusterRoleBinding{Context: ctx, Handler: handler}
}

/*CreateRbacAuthorizationV1beta1ClusterRoleBinding swagger:route POST /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings rbacAuthorization_v1beta1 createRbacAuthorizationV1beta1ClusterRoleBinding

create a ClusterRoleBinding

*/
type CreateRbacAuthorizationV1beta1ClusterRoleBinding struct {
	Context *middleware.Context
	Handler CreateRbacAuthorizationV1beta1ClusterRoleBindingHandler
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRbacAuthorizationV1beta1ClusterRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
