// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc turns a function with the right signature into a patch rbac authorization v1alpha1 namespaced role binding handler
type PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc func(PatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc) Handle(params PatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
	return fn(params)
}

// PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler interface for that can handle valid patch rbac authorization v1alpha1 namespaced role binding params
type PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler interface {
	Handle(PatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder
}

// NewPatchRbacAuthorizationV1alpha1NamespacedRoleBinding creates a new http.Handler for the patch rbac authorization v1alpha1 namespaced role binding operation
func NewPatchRbacAuthorizationV1alpha1NamespacedRoleBinding(ctx *middleware.Context, handler PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler) *PatchRbacAuthorizationV1alpha1NamespacedRoleBinding {
	return &PatchRbacAuthorizationV1alpha1NamespacedRoleBinding{Context: ctx, Handler: handler}
}

/*PatchRbacAuthorizationV1alpha1NamespacedRoleBinding swagger:route PATCH /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings/{name} rbacAuthorization_v1alpha1 patchRbacAuthorizationV1alpha1NamespacedRoleBinding

partially update the specified RoleBinding

*/
type PatchRbacAuthorizationV1alpha1NamespacedRoleBinding struct {
	Context *middleware.Context
	Handler PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler
}

func (o *PatchRbacAuthorizationV1alpha1NamespacedRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
