// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNetworkingV1IngressClassListHandlerFunc turns a function with the right signature into a watch networking v1 ingress class list handler
type WatchNetworkingV1IngressClassListHandlerFunc func(WatchNetworkingV1IngressClassListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNetworkingV1IngressClassListHandlerFunc) Handle(params WatchNetworkingV1IngressClassListParams) middleware.Responder {
	return fn(params)
}

// WatchNetworkingV1IngressClassListHandler interface for that can handle valid watch networking v1 ingress class list params
type WatchNetworkingV1IngressClassListHandler interface {
	Handle(WatchNetworkingV1IngressClassListParams) middleware.Responder
}

// NewWatchNetworkingV1IngressClassList creates a new http.Handler for the watch networking v1 ingress class list operation
func NewWatchNetworkingV1IngressClassList(ctx *middleware.Context, handler WatchNetworkingV1IngressClassListHandler) *WatchNetworkingV1IngressClassList {
	return &WatchNetworkingV1IngressClassList{Context: ctx, Handler: handler}
}

/*WatchNetworkingV1IngressClassList swagger:route GET /apis/networking.k8s.io/v1/watch/ingressclasses networking_v1 watchNetworkingV1IngressClassList

watch individual changes to a list of IngressClass. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchNetworkingV1IngressClassList struct {
	Context *middleware.Context
	Handler WatchNetworkingV1IngressClassListHandler
}

func (o *WatchNetworkingV1IngressClassList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNetworkingV1IngressClassListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
