// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteStorageV1beta1StorageClassHandlerFunc turns a function with the right signature into a delete storage v1beta1 storage class handler
type DeleteStorageV1beta1StorageClassHandlerFunc func(DeleteStorageV1beta1StorageClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageV1beta1StorageClassHandlerFunc) Handle(params DeleteStorageV1beta1StorageClassParams) middleware.Responder {
	return fn(params)
}

// DeleteStorageV1beta1StorageClassHandler interface for that can handle valid delete storage v1beta1 storage class params
type DeleteStorageV1beta1StorageClassHandler interface {
	Handle(DeleteStorageV1beta1StorageClassParams) middleware.Responder
}

// NewDeleteStorageV1beta1StorageClass creates a new http.Handler for the delete storage v1beta1 storage class operation
func NewDeleteStorageV1beta1StorageClass(ctx *middleware.Context, handler DeleteStorageV1beta1StorageClassHandler) *DeleteStorageV1beta1StorageClass {
	return &DeleteStorageV1beta1StorageClass{Context: ctx, Handler: handler}
}

/*DeleteStorageV1beta1StorageClass swagger:route DELETE /apis/storage.k8s.io/v1beta1/storageclasses/{name} storage_v1beta1 deleteStorageV1beta1StorageClass

delete a StorageClass

*/
type DeleteStorageV1beta1StorageClass struct {
	Context *middleware.Context
	Handler DeleteStorageV1beta1StorageClassHandler
}

func (o *DeleteStorageV1beta1StorageClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageV1beta1StorageClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
