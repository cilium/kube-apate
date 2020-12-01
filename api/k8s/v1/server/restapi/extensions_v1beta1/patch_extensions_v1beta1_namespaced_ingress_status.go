// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchExtensionsV1beta1NamespacedIngressStatusHandlerFunc turns a function with the right signature into a patch extensions v1beta1 namespaced ingress status handler
type PatchExtensionsV1beta1NamespacedIngressStatusHandlerFunc func(PatchExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchExtensionsV1beta1NamespacedIngressStatusHandlerFunc) Handle(params PatchExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder {
	return fn(params)
}

// PatchExtensionsV1beta1NamespacedIngressStatusHandler interface for that can handle valid patch extensions v1beta1 namespaced ingress status params
type PatchExtensionsV1beta1NamespacedIngressStatusHandler interface {
	Handle(PatchExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder
}

// NewPatchExtensionsV1beta1NamespacedIngressStatus creates a new http.Handler for the patch extensions v1beta1 namespaced ingress status operation
func NewPatchExtensionsV1beta1NamespacedIngressStatus(ctx *middleware.Context, handler PatchExtensionsV1beta1NamespacedIngressStatusHandler) *PatchExtensionsV1beta1NamespacedIngressStatus {
	return &PatchExtensionsV1beta1NamespacedIngressStatus{Context: ctx, Handler: handler}
}

/*PatchExtensionsV1beta1NamespacedIngressStatus swagger:route PATCH /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}/status extensions_v1beta1 patchExtensionsV1beta1NamespacedIngressStatus

partially update status of the specified Ingress

*/
type PatchExtensionsV1beta1NamespacedIngressStatus struct {
	Context *middleware.Context
	Handler PatchExtensionsV1beta1NamespacedIngressStatusHandler
}

func (o *PatchExtensionsV1beta1NamespacedIngressStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchExtensionsV1beta1NamespacedIngressStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
