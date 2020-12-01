// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListStorageV1CSINodeHandlerFunc turns a function with the right signature into a list storage v1 c s i node handler
type ListStorageV1CSINodeHandlerFunc func(ListStorageV1CSINodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStorageV1CSINodeHandlerFunc) Handle(params ListStorageV1CSINodeParams) middleware.Responder {
	return fn(params)
}

// ListStorageV1CSINodeHandler interface for that can handle valid list storage v1 c s i node params
type ListStorageV1CSINodeHandler interface {
	Handle(ListStorageV1CSINodeParams) middleware.Responder
}

// NewListStorageV1CSINode creates a new http.Handler for the list storage v1 c s i node operation
func NewListStorageV1CSINode(ctx *middleware.Context, handler ListStorageV1CSINodeHandler) *ListStorageV1CSINode {
	return &ListStorageV1CSINode{Context: ctx, Handler: handler}
}

/*ListStorageV1CSINode swagger:route GET /apis/storage.k8s.io/v1/csinodes storage_v1 listStorageV1CSINode

list or watch objects of kind CSINode

*/
type ListStorageV1CSINode struct {
	Context *middleware.Context
	Handler ListStorageV1CSINodeHandler
}

func (o *ListStorageV1CSINode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStorageV1CSINodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
