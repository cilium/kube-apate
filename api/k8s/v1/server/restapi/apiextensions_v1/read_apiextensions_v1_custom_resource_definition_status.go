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

// ReadApiextensionsV1CustomResourceDefinitionStatusHandlerFunc turns a function with the right signature into a read apiextensions v1 custom resource definition status handler
type ReadApiextensionsV1CustomResourceDefinitionStatusHandlerFunc func(ReadApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadApiextensionsV1CustomResourceDefinitionStatusHandlerFunc) Handle(params ReadApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder {
	return fn(params)
}

// ReadApiextensionsV1CustomResourceDefinitionStatusHandler interface for that can handle valid read apiextensions v1 custom resource definition status params
type ReadApiextensionsV1CustomResourceDefinitionStatusHandler interface {
	Handle(ReadApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder
}

// NewReadApiextensionsV1CustomResourceDefinitionStatus creates a new http.Handler for the read apiextensions v1 custom resource definition status operation
func NewReadApiextensionsV1CustomResourceDefinitionStatus(ctx *middleware.Context, handler ReadApiextensionsV1CustomResourceDefinitionStatusHandler) *ReadApiextensionsV1CustomResourceDefinitionStatus {
	return &ReadApiextensionsV1CustomResourceDefinitionStatus{Context: ctx, Handler: handler}
}

/*ReadApiextensionsV1CustomResourceDefinitionStatus swagger:route GET /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}/status apiextensions_v1 readApiextensionsV1CustomResourceDefinitionStatus

read status of the specified CustomResourceDefinition

*/
type ReadApiextensionsV1CustomResourceDefinitionStatus struct {
	Context *middleware.Context
	Handler ReadApiextensionsV1CustomResourceDefinitionStatusHandler
}

func (o *ReadApiextensionsV1CustomResourceDefinitionStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadApiextensionsV1CustomResourceDefinitionStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}