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

// DeleteRbacAuthorizationV1CollectionNamespacedRoleHandlerFunc turns a function with the right signature into a delete rbac authorization v1 collection namespaced role handler
type DeleteRbacAuthorizationV1CollectionNamespacedRoleHandlerFunc func(DeleteRbacAuthorizationV1CollectionNamespacedRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1CollectionNamespacedRoleHandlerFunc) Handle(params DeleteRbacAuthorizationV1CollectionNamespacedRoleParams) middleware.Responder {
	return fn(params)
}

// DeleteRbacAuthorizationV1CollectionNamespacedRoleHandler interface for that can handle valid delete rbac authorization v1 collection namespaced role params
type DeleteRbacAuthorizationV1CollectionNamespacedRoleHandler interface {
	Handle(DeleteRbacAuthorizationV1CollectionNamespacedRoleParams) middleware.Responder
}

// NewDeleteRbacAuthorizationV1CollectionNamespacedRole creates a new http.Handler for the delete rbac authorization v1 collection namespaced role operation
func NewDeleteRbacAuthorizationV1CollectionNamespacedRole(ctx *middleware.Context, handler DeleteRbacAuthorizationV1CollectionNamespacedRoleHandler) *DeleteRbacAuthorizationV1CollectionNamespacedRole {
	return &DeleteRbacAuthorizationV1CollectionNamespacedRole{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1CollectionNamespacedRole swagger:route DELETE /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles rbacAuthorization_v1 deleteRbacAuthorizationV1CollectionNamespacedRole

delete collection of Role

*/
type DeleteRbacAuthorizationV1CollectionNamespacedRole struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1CollectionNamespacedRoleHandler
}

func (o *DeleteRbacAuthorizationV1CollectionNamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1CollectionNamespacedRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}