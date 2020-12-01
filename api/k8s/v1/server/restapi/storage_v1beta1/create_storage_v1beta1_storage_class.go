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

// CreateStorageV1beta1StorageClassHandlerFunc turns a function with the right signature into a create storage v1beta1 storage class handler
type CreateStorageV1beta1StorageClassHandlerFunc func(CreateStorageV1beta1StorageClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateStorageV1beta1StorageClassHandlerFunc) Handle(params CreateStorageV1beta1StorageClassParams) middleware.Responder {
	return fn(params)
}

// CreateStorageV1beta1StorageClassHandler interface for that can handle valid create storage v1beta1 storage class params
type CreateStorageV1beta1StorageClassHandler interface {
	Handle(CreateStorageV1beta1StorageClassParams) middleware.Responder
}

// NewCreateStorageV1beta1StorageClass creates a new http.Handler for the create storage v1beta1 storage class operation
func NewCreateStorageV1beta1StorageClass(ctx *middleware.Context, handler CreateStorageV1beta1StorageClassHandler) *CreateStorageV1beta1StorageClass {
	return &CreateStorageV1beta1StorageClass{Context: ctx, Handler: handler}
}

/*CreateStorageV1beta1StorageClass swagger:route POST /apis/storage.k8s.io/v1beta1/storageclasses storage_v1beta1 createStorageV1beta1StorageClass

create a StorageClass

*/
type CreateStorageV1beta1StorageClass struct {
	Context *middleware.Context
	Handler CreateStorageV1beta1StorageClassHandler
}

func (o *CreateStorageV1beta1StorageClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateStorageV1beta1StorageClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
