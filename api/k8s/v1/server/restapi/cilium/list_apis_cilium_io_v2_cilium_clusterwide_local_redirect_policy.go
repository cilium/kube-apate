// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandlerFunc turns a function with the right signature into a list apis cilium io v2 cilium clusterwide local redirect policy handler
type ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandlerFunc func(ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandlerFunc) Handle(params ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyParams) middleware.Responder {
	return fn(params)
}

// ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandler interface for that can handle valid list apis cilium io v2 cilium clusterwide local redirect policy params
type ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandler interface {
	Handle(ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyParams) middleware.Responder
}

// NewListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy creates a new http.Handler for the list apis cilium io v2 cilium clusterwide local redirect policy operation
func NewListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy(ctx *middleware.Context, handler ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandler) *ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy {
	return &ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy{Context: ctx, Handler: handler}
}

/*ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy swagger:route GET /apis/cilium.io/v2/ciliumclusterwidelocalredirectpolicy cilium listApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy

ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy list apis cilium io v2 cilium clusterwide local redirect policy API

*/
type ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy struct {
	Context *middleware.Context
	Handler ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandler
}

func (o *ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
