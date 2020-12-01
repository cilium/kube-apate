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

// ReadEventsV1beta1NamespacedEventHandlerFunc turns a function with the right signature into a read events v1beta1 namespaced event handler
type ReadEventsV1beta1NamespacedEventHandlerFunc func(ReadEventsV1beta1NamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadEventsV1beta1NamespacedEventHandlerFunc) Handle(params ReadEventsV1beta1NamespacedEventParams) middleware.Responder {
	return fn(params)
}

// ReadEventsV1beta1NamespacedEventHandler interface for that can handle valid read events v1beta1 namespaced event params
type ReadEventsV1beta1NamespacedEventHandler interface {
	Handle(ReadEventsV1beta1NamespacedEventParams) middleware.Responder
}

// NewReadEventsV1beta1NamespacedEvent creates a new http.Handler for the read events v1beta1 namespaced event operation
func NewReadEventsV1beta1NamespacedEvent(ctx *middleware.Context, handler ReadEventsV1beta1NamespacedEventHandler) *ReadEventsV1beta1NamespacedEvent {
	return &ReadEventsV1beta1NamespacedEvent{Context: ctx, Handler: handler}
}

/*ReadEventsV1beta1NamespacedEvent swagger:route GET /apis/events.k8s.io/v1beta1/namespaces/{namespace}/events/{name} events_v1beta1 readEventsV1beta1NamespacedEvent

read the specified Event

*/
type ReadEventsV1beta1NamespacedEvent struct {
	Context *middleware.Context
	Handler ReadEventsV1beta1NamespacedEventHandler
}

func (o *ReadEventsV1beta1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadEventsV1beta1NamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
