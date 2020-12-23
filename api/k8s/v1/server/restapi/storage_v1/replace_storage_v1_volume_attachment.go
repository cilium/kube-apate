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

// ReplaceStorageV1VolumeAttachmentHandlerFunc turns a function with the right signature into a replace storage v1 volume attachment handler
type ReplaceStorageV1VolumeAttachmentHandlerFunc func(ReplaceStorageV1VolumeAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStorageV1VolumeAttachmentHandlerFunc) Handle(params ReplaceStorageV1VolumeAttachmentParams) middleware.Responder {
	return fn(params)
}

// ReplaceStorageV1VolumeAttachmentHandler interface for that can handle valid replace storage v1 volume attachment params
type ReplaceStorageV1VolumeAttachmentHandler interface {
	Handle(ReplaceStorageV1VolumeAttachmentParams) middleware.Responder
}

// NewReplaceStorageV1VolumeAttachment creates a new http.Handler for the replace storage v1 volume attachment operation
func NewReplaceStorageV1VolumeAttachment(ctx *middleware.Context, handler ReplaceStorageV1VolumeAttachmentHandler) *ReplaceStorageV1VolumeAttachment {
	return &ReplaceStorageV1VolumeAttachment{Context: ctx, Handler: handler}
}

/*ReplaceStorageV1VolumeAttachment swagger:route PUT /apis/storage.k8s.io/v1/volumeattachments/{name} storage_v1 replaceStorageV1VolumeAttachment

replace the specified VolumeAttachment

*/
type ReplaceStorageV1VolumeAttachment struct {
	Context *middleware.Context
	Handler ReplaceStorageV1VolumeAttachmentHandler
}

func (o *ReplaceStorageV1VolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStorageV1VolumeAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}