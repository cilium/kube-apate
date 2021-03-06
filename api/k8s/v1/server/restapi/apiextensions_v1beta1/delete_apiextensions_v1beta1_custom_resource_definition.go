// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteApiextensionsV1beta1CustomResourceDefinitionHandlerFunc turns a function with the right signature into a delete apiextensions v1beta1 custom resource definition handler
type DeleteApiextensionsV1beta1CustomResourceDefinitionHandlerFunc func(DeleteApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteApiextensionsV1beta1CustomResourceDefinitionHandlerFunc) Handle(params DeleteApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
	return fn(params)
}

// DeleteApiextensionsV1beta1CustomResourceDefinitionHandler interface for that can handle valid delete apiextensions v1beta1 custom resource definition params
type DeleteApiextensionsV1beta1CustomResourceDefinitionHandler interface {
	Handle(DeleteApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder
}

// NewDeleteApiextensionsV1beta1CustomResourceDefinition creates a new http.Handler for the delete apiextensions v1beta1 custom resource definition operation
func NewDeleteApiextensionsV1beta1CustomResourceDefinition(ctx *middleware.Context, handler DeleteApiextensionsV1beta1CustomResourceDefinitionHandler) *DeleteApiextensionsV1beta1CustomResourceDefinition {
	return &DeleteApiextensionsV1beta1CustomResourceDefinition{Context: ctx, Handler: handler}
}

/*DeleteApiextensionsV1beta1CustomResourceDefinition swagger:route DELETE /apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions/{name} apiextensions_v1beta1 deleteApiextensionsV1beta1CustomResourceDefinition

delete a CustomResourceDefinition

*/
type DeleteApiextensionsV1beta1CustomResourceDefinition struct {
	Context *middleware.Context
	Handler DeleteApiextensionsV1beta1CustomResourceDefinitionHandler
}

func (o *DeleteApiextensionsV1beta1CustomResourceDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteApiextensionsV1beta1CustomResourceDefinitionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
