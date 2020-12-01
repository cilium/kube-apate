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

// ReplaceCoreV1NamespacedResourceQuotaStatusHandlerFunc turns a function with the right signature into a replace core v1 namespaced resource quota status handler
type ReplaceCoreV1NamespacedResourceQuotaStatusHandlerFunc func(ReplaceCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespacedResourceQuotaStatusHandlerFunc) Handle(params ReplaceCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NamespacedResourceQuotaStatusHandler interface for that can handle valid replace core v1 namespaced resource quota status params
type ReplaceCoreV1NamespacedResourceQuotaStatusHandler interface {
	Handle(ReplaceCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder
}

// NewReplaceCoreV1NamespacedResourceQuotaStatus creates a new http.Handler for the replace core v1 namespaced resource quota status operation
func NewReplaceCoreV1NamespacedResourceQuotaStatus(ctx *middleware.Context, handler ReplaceCoreV1NamespacedResourceQuotaStatusHandler) *ReplaceCoreV1NamespacedResourceQuotaStatus {
	return &ReplaceCoreV1NamespacedResourceQuotaStatus{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespacedResourceQuotaStatus swagger:route PUT /api/v1/namespaces/{namespace}/resourcequotas/{name}/status core_v1 replaceCoreV1NamespacedResourceQuotaStatus

replace status of the specified ResourceQuota

*/
type ReplaceCoreV1NamespacedResourceQuotaStatus struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespacedResourceQuotaStatusHandler
}

func (o *ReplaceCoreV1NamespacedResourceQuotaStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespacedResourceQuotaStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
