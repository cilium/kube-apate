// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NamespacedEventHandlerFunc turns a function with the right signature into a replace core v1 namespaced event handler
type ReplaceCoreV1NamespacedEventHandlerFunc func(ReplaceCoreV1NamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespacedEventHandlerFunc) Handle(params ReplaceCoreV1NamespacedEventParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NamespacedEventHandler interface for that can handle valid replace core v1 namespaced event params
type ReplaceCoreV1NamespacedEventHandler interface {
	Handle(ReplaceCoreV1NamespacedEventParams) middleware.Responder
}

// NewReplaceCoreV1NamespacedEvent creates a new http.Handler for the replace core v1 namespaced event operation
func NewReplaceCoreV1NamespacedEvent(ctx *middleware.Context, handler ReplaceCoreV1NamespacedEventHandler) *ReplaceCoreV1NamespacedEvent {
	return &ReplaceCoreV1NamespacedEvent{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespacedEvent swagger:route PUT /api/v1/namespaces/{namespace}/events/{name} core_v1 replaceCoreV1NamespacedEvent

replace the specified Event

*/
type ReplaceCoreV1NamespacedEvent struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespacedEventHandler
}

func (o *ReplaceCoreV1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
