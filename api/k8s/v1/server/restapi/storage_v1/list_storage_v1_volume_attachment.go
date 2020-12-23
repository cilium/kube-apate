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

// ListStorageV1VolumeAttachmentHandlerFunc turns a function with the right signature into a list storage v1 volume attachment handler
type ListStorageV1VolumeAttachmentHandlerFunc func(ListStorageV1VolumeAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStorageV1VolumeAttachmentHandlerFunc) Handle(params ListStorageV1VolumeAttachmentParams) middleware.Responder {
	return fn(params)
}

// ListStorageV1VolumeAttachmentHandler interface for that can handle valid list storage v1 volume attachment params
type ListStorageV1VolumeAttachmentHandler interface {
	Handle(ListStorageV1VolumeAttachmentParams) middleware.Responder
}

// NewListStorageV1VolumeAttachment creates a new http.Handler for the list storage v1 volume attachment operation
func NewListStorageV1VolumeAttachment(ctx *middleware.Context, handler ListStorageV1VolumeAttachmentHandler) *ListStorageV1VolumeAttachment {
	return &ListStorageV1VolumeAttachment{Context: ctx, Handler: handler}
}

/*ListStorageV1VolumeAttachment swagger:route GET /apis/storage.k8s.io/v1/volumeattachments storage_v1 listStorageV1VolumeAttachment

list or watch objects of kind VolumeAttachment

*/
type ListStorageV1VolumeAttachment struct {
	Context *middleware.Context
	Handler ListStorageV1VolumeAttachmentHandler
}

func (o *ListStorageV1VolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStorageV1VolumeAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}