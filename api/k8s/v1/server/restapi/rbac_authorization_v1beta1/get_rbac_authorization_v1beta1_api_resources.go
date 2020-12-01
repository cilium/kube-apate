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

// GetRbacAuthorizationV1beta1APIResourcesHandlerFunc turns a function with the right signature into a get rbac authorization v1beta1 API resources handler
type GetRbacAuthorizationV1beta1APIResourcesHandlerFunc func(GetRbacAuthorizationV1beta1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRbacAuthorizationV1beta1APIResourcesHandlerFunc) Handle(params GetRbacAuthorizationV1beta1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetRbacAuthorizationV1beta1APIResourcesHandler interface for that can handle valid get rbac authorization v1beta1 API resources params
type GetRbacAuthorizationV1beta1APIResourcesHandler interface {
	Handle(GetRbacAuthorizationV1beta1APIResourcesParams) middleware.Responder
}

// NewGetRbacAuthorizationV1beta1APIResources creates a new http.Handler for the get rbac authorization v1beta1 API resources operation
func NewGetRbacAuthorizationV1beta1APIResources(ctx *middleware.Context, handler GetRbacAuthorizationV1beta1APIResourcesHandler) *GetRbacAuthorizationV1beta1APIResources {
	return &GetRbacAuthorizationV1beta1APIResources{Context: ctx, Handler: handler}
}

/*GetRbacAuthorizationV1beta1APIResources swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/ rbacAuthorization_v1beta1 getRbacAuthorizationV1beta1ApiResources

get available resources

*/
type GetRbacAuthorizationV1beta1APIResources struct {
	Context *middleware.Context
	Handler GetRbacAuthorizationV1beta1APIResourcesHandler
}

func (o *GetRbacAuthorizationV1beta1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRbacAuthorizationV1beta1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
