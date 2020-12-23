// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateApiextensionsV1CustomResourceDefinitionHandlerFunc turns a function with the right signature into a create apiextensions v1 custom resource definition handler
type CreateApiextensionsV1CustomResourceDefinitionHandlerFunc func(CreateApiextensionsV1CustomResourceDefinitionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateApiextensionsV1CustomResourceDefinitionHandlerFunc) Handle(params CreateApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
	return fn(params)
}

// CreateApiextensionsV1CustomResourceDefinitionHandler interface for that can handle valid create apiextensions v1 custom resource definition params
type CreateApiextensionsV1CustomResourceDefinitionHandler interface {
	Handle(CreateApiextensionsV1CustomResourceDefinitionParams) middleware.Responder
}

// NewCreateApiextensionsV1CustomResourceDefinition creates a new http.Handler for the create apiextensions v1 custom resource definition operation
func NewCreateApiextensionsV1CustomResourceDefinition(ctx *middleware.Context, handler CreateApiextensionsV1CustomResourceDefinitionHandler) *CreateApiextensionsV1CustomResourceDefinition {
	return &CreateApiextensionsV1CustomResourceDefinition{Context: ctx, Handler: handler}
}

/*CreateApiextensionsV1CustomResourceDefinition swagger:route POST /apis/apiextensions.k8s.io/v1/customresourcedefinitions apiextensions_v1 createApiextensionsV1CustomResourceDefinition

create a CustomResourceDefinition

*/
type CreateApiextensionsV1CustomResourceDefinition struct {
	Context *middleware.Context
	Handler CreateApiextensionsV1CustomResourceDefinitionHandler
}

func (o *CreateApiextensionsV1CustomResourceDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateApiextensionsV1CustomResourceDefinitionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}