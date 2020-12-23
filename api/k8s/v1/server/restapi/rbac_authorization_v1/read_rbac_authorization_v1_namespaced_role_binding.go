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

// ReadRbacAuthorizationV1NamespacedRoleBindingHandlerFunc turns a function with the right signature into a read rbac authorization v1 namespaced role binding handler
type ReadRbacAuthorizationV1NamespacedRoleBindingHandlerFunc func(ReadRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadRbacAuthorizationV1NamespacedRoleBindingHandlerFunc) Handle(params ReadRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
	return fn(params)
}

// ReadRbacAuthorizationV1NamespacedRoleBindingHandler interface for that can handle valid read rbac authorization v1 namespaced role binding params
type ReadRbacAuthorizationV1NamespacedRoleBindingHandler interface {
	Handle(ReadRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder
}

// NewReadRbacAuthorizationV1NamespacedRoleBinding creates a new http.Handler for the read rbac authorization v1 namespaced role binding operation
func NewReadRbacAuthorizationV1NamespacedRoleBinding(ctx *middleware.Context, handler ReadRbacAuthorizationV1NamespacedRoleBindingHandler) *ReadRbacAuthorizationV1NamespacedRoleBinding {
	return &ReadRbacAuthorizationV1NamespacedRoleBinding{Context: ctx, Handler: handler}
}

/*ReadRbacAuthorizationV1NamespacedRoleBinding swagger:route GET /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings/{name} rbacAuthorization_v1 readRbacAuthorizationV1NamespacedRoleBinding

read the specified RoleBinding

*/
type ReadRbacAuthorizationV1NamespacedRoleBinding struct {
	Context *middleware.Context
	Handler ReadRbacAuthorizationV1NamespacedRoleBindingHandler
}

func (o *ReadRbacAuthorizationV1NamespacedRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadRbacAuthorizationV1NamespacedRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}