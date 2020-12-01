// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListStorageV1alpha1VolumeAttachmentHandlerFunc turns a function with the right signature into a list storage v1alpha1 volume attachment handler
type ListStorageV1alpha1VolumeAttachmentHandlerFunc func(ListStorageV1alpha1VolumeAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStorageV1alpha1VolumeAttachmentHandlerFunc) Handle(params ListStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
	return fn(params)
}

// ListStorageV1alpha1VolumeAttachmentHandler interface for that can handle valid list storage v1alpha1 volume attachment params
type ListStorageV1alpha1VolumeAttachmentHandler interface {
	Handle(ListStorageV1alpha1VolumeAttachmentParams) middleware.Responder
}

// NewListStorageV1alpha1VolumeAttachment creates a new http.Handler for the list storage v1alpha1 volume attachment operation
func NewListStorageV1alpha1VolumeAttachment(ctx *middleware.Context, handler ListStorageV1alpha1VolumeAttachmentHandler) *ListStorageV1alpha1VolumeAttachment {
	return &ListStorageV1alpha1VolumeAttachment{Context: ctx, Handler: handler}
}

/*ListStorageV1alpha1VolumeAttachment swagger:route GET /apis/storage.k8s.io/v1alpha1/volumeattachments storage_v1alpha1 listStorageV1alpha1VolumeAttachment

list or watch objects of kind VolumeAttachment

*/
type ListStorageV1alpha1VolumeAttachment struct {
	Context *middleware.Context
	Handler ListStorageV1alpha1VolumeAttachmentHandler
}

func (o *ListStorageV1alpha1VolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStorageV1alpha1VolumeAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
