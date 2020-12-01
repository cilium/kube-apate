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

// WatchStorageV1beta1VolumeAttachmentListHandlerFunc turns a function with the right signature into a watch storage v1beta1 volume attachment list handler
type WatchStorageV1beta1VolumeAttachmentListHandlerFunc func(WatchStorageV1beta1VolumeAttachmentListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchStorageV1beta1VolumeAttachmentListHandlerFunc) Handle(params WatchStorageV1beta1VolumeAttachmentListParams) middleware.Responder {
	return fn(params)
}

// WatchStorageV1beta1VolumeAttachmentListHandler interface for that can handle valid watch storage v1beta1 volume attachment list params
type WatchStorageV1beta1VolumeAttachmentListHandler interface {
	Handle(WatchStorageV1beta1VolumeAttachmentListParams) middleware.Responder
}

// NewWatchStorageV1beta1VolumeAttachmentList creates a new http.Handler for the watch storage v1beta1 volume attachment list operation
func NewWatchStorageV1beta1VolumeAttachmentList(ctx *middleware.Context, handler WatchStorageV1beta1VolumeAttachmentListHandler) *WatchStorageV1beta1VolumeAttachmentList {
	return &WatchStorageV1beta1VolumeAttachmentList{Context: ctx, Handler: handler}
}

/*WatchStorageV1beta1VolumeAttachmentList swagger:route GET /apis/storage.k8s.io/v1beta1/watch/volumeattachments storage_v1beta1 watchStorageV1beta1VolumeAttachmentList

watch individual changes to a list of VolumeAttachment. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchStorageV1beta1VolumeAttachmentList struct {
	Context *middleware.Context
	Handler WatchStorageV1beta1VolumeAttachmentListHandler
}

func (o *WatchStorageV1beta1VolumeAttachmentList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchStorageV1beta1VolumeAttachmentListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
