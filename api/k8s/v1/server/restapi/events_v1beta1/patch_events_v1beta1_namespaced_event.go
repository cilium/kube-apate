// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchEventsV1beta1NamespacedEventHandlerFunc turns a function with the right signature into a patch events v1beta1 namespaced event handler
type PatchEventsV1beta1NamespacedEventHandlerFunc func(PatchEventsV1beta1NamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchEventsV1beta1NamespacedEventHandlerFunc) Handle(params PatchEventsV1beta1NamespacedEventParams) middleware.Responder {
	return fn(params)
}

// PatchEventsV1beta1NamespacedEventHandler interface for that can handle valid patch events v1beta1 namespaced event params
type PatchEventsV1beta1NamespacedEventHandler interface {
	Handle(PatchEventsV1beta1NamespacedEventParams) middleware.Responder
}

// NewPatchEventsV1beta1NamespacedEvent creates a new http.Handler for the patch events v1beta1 namespaced event operation
func NewPatchEventsV1beta1NamespacedEvent(ctx *middleware.Context, handler PatchEventsV1beta1NamespacedEventHandler) *PatchEventsV1beta1NamespacedEvent {
	return &PatchEventsV1beta1NamespacedEvent{Context: ctx, Handler: handler}
}

/*PatchEventsV1beta1NamespacedEvent swagger:route PATCH /apis/events.k8s.io/v1beta1/namespaces/{namespace}/events/{name} events_v1beta1 patchEventsV1beta1NamespacedEvent

partially update the specified Event

*/
type PatchEventsV1beta1NamespacedEvent struct {
	Context *middleware.Context
	Handler PatchEventsV1beta1NamespacedEventHandler
}

func (o *PatchEventsV1beta1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchEventsV1beta1NamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
