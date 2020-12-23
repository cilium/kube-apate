// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateEventsV1NamespacedEventHandlerFunc turns a function with the right signature into a create events v1 namespaced event handler
type CreateEventsV1NamespacedEventHandlerFunc func(CreateEventsV1NamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateEventsV1NamespacedEventHandlerFunc) Handle(params CreateEventsV1NamespacedEventParams) middleware.Responder {
	return fn(params)
}

// CreateEventsV1NamespacedEventHandler interface for that can handle valid create events v1 namespaced event params
type CreateEventsV1NamespacedEventHandler interface {
	Handle(CreateEventsV1NamespacedEventParams) middleware.Responder
}

// NewCreateEventsV1NamespacedEvent creates a new http.Handler for the create events v1 namespaced event operation
func NewCreateEventsV1NamespacedEvent(ctx *middleware.Context, handler CreateEventsV1NamespacedEventHandler) *CreateEventsV1NamespacedEvent {
	return &CreateEventsV1NamespacedEvent{Context: ctx, Handler: handler}
}

/*CreateEventsV1NamespacedEvent swagger:route POST /apis/events.k8s.io/v1/namespaces/{namespace}/events events_v1 createEventsV1NamespacedEvent

create an Event

*/
type CreateEventsV1NamespacedEvent struct {
	Context *middleware.Context
	Handler CreateEventsV1NamespacedEventHandler
}

func (o *CreateEventsV1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateEventsV1NamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}