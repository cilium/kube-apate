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

// WatchStorageV1VolumeAttachmentHandlerFunc turns a function with the right signature into a watch storage v1 volume attachment handler
type WatchStorageV1VolumeAttachmentHandlerFunc func(WatchStorageV1VolumeAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchStorageV1VolumeAttachmentHandlerFunc) Handle(params WatchStorageV1VolumeAttachmentParams) middleware.Responder {
	return fn(params)
}

// WatchStorageV1VolumeAttachmentHandler interface for that can handle valid watch storage v1 volume attachment params
type WatchStorageV1VolumeAttachmentHandler interface {
	Handle(WatchStorageV1VolumeAttachmentParams) middleware.Responder
}

// NewWatchStorageV1VolumeAttachment creates a new http.Handler for the watch storage v1 volume attachment operation
func NewWatchStorageV1VolumeAttachment(ctx *middleware.Context, handler WatchStorageV1VolumeAttachmentHandler) *WatchStorageV1VolumeAttachment {
	return &WatchStorageV1VolumeAttachment{Context: ctx, Handler: handler}
}

/*WatchStorageV1VolumeAttachment swagger:route GET /apis/storage.k8s.io/v1/watch/volumeattachments/{name} storage_v1 watchStorageV1VolumeAttachment

watch changes to an object of kind VolumeAttachment. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchStorageV1VolumeAttachment struct {
	Context *middleware.Context
	Handler WatchStorageV1VolumeAttachmentHandler
}

func (o *WatchStorageV1VolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchStorageV1VolumeAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
