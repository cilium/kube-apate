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

// DeleteEventsV1CollectionNamespacedEventHandlerFunc turns a function with the right signature into a delete events v1 collection namespaced event handler
type DeleteEventsV1CollectionNamespacedEventHandlerFunc func(DeleteEventsV1CollectionNamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEventsV1CollectionNamespacedEventHandlerFunc) Handle(params DeleteEventsV1CollectionNamespacedEventParams) middleware.Responder {
	return fn(params)
}

// DeleteEventsV1CollectionNamespacedEventHandler interface for that can handle valid delete events v1 collection namespaced event params
type DeleteEventsV1CollectionNamespacedEventHandler interface {
	Handle(DeleteEventsV1CollectionNamespacedEventParams) middleware.Responder
}

// NewDeleteEventsV1CollectionNamespacedEvent creates a new http.Handler for the delete events v1 collection namespaced event operation
func NewDeleteEventsV1CollectionNamespacedEvent(ctx *middleware.Context, handler DeleteEventsV1CollectionNamespacedEventHandler) *DeleteEventsV1CollectionNamespacedEvent {
	return &DeleteEventsV1CollectionNamespacedEvent{Context: ctx, Handler: handler}
}

/*DeleteEventsV1CollectionNamespacedEvent swagger:route DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events events_v1 deleteEventsV1CollectionNamespacedEvent

delete collection of Event

*/
type DeleteEventsV1CollectionNamespacedEvent struct {
	Context *middleware.Context
	Handler DeleteEventsV1CollectionNamespacedEventHandler
}

func (o *DeleteEventsV1CollectionNamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteEventsV1CollectionNamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
