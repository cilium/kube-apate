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

// WatchRbacAuthorizationV1beta1ClusterRoleHandlerFunc turns a function with the right signature into a watch rbac authorization v1beta1 cluster role handler
type WatchRbacAuthorizationV1beta1ClusterRoleHandlerFunc func(WatchRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1beta1ClusterRoleHandlerFunc) Handle(params WatchRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// WatchRbacAuthorizationV1beta1ClusterRoleHandler interface for that can handle valid watch rbac authorization v1beta1 cluster role params
type WatchRbacAuthorizationV1beta1ClusterRoleHandler interface {
	Handle(WatchRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder
}

// NewWatchRbacAuthorizationV1beta1ClusterRole creates a new http.Handler for the watch rbac authorization v1beta1 cluster role operation
func NewWatchRbacAuthorizationV1beta1ClusterRole(ctx *middleware.Context, handler WatchRbacAuthorizationV1beta1ClusterRoleHandler) *WatchRbacAuthorizationV1beta1ClusterRole {
	return &WatchRbacAuthorizationV1beta1ClusterRole{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1beta1ClusterRole swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/watch/clusterroles/{name} rbacAuthorization_v1beta1 watchRbacAuthorizationV1beta1ClusterRole

watch changes to an object of kind ClusterRole. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchRbacAuthorizationV1beta1ClusterRole struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1beta1ClusterRoleHandler
}

func (o *WatchRbacAuthorizationV1beta1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1beta1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}