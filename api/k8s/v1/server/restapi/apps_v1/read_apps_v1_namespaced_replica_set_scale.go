// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAppsV1NamespacedReplicaSetScaleHandlerFunc turns a function with the right signature into a read apps v1 namespaced replica set scale handler
type ReadAppsV1NamespacedReplicaSetScaleHandlerFunc func(ReadAppsV1NamespacedReplicaSetScaleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAppsV1NamespacedReplicaSetScaleHandlerFunc) Handle(params ReadAppsV1NamespacedReplicaSetScaleParams) middleware.Responder {
	return fn(params)
}

// ReadAppsV1NamespacedReplicaSetScaleHandler interface for that can handle valid read apps v1 namespaced replica set scale params
type ReadAppsV1NamespacedReplicaSetScaleHandler interface {
	Handle(ReadAppsV1NamespacedReplicaSetScaleParams) middleware.Responder
}

// NewReadAppsV1NamespacedReplicaSetScale creates a new http.Handler for the read apps v1 namespaced replica set scale operation
func NewReadAppsV1NamespacedReplicaSetScale(ctx *middleware.Context, handler ReadAppsV1NamespacedReplicaSetScaleHandler) *ReadAppsV1NamespacedReplicaSetScale {
	return &ReadAppsV1NamespacedReplicaSetScale{Context: ctx, Handler: handler}
}

/*ReadAppsV1NamespacedReplicaSetScale swagger:route GET /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/scale apps_v1 readAppsV1NamespacedReplicaSetScale

read scale of the specified ReplicaSet

*/
type ReadAppsV1NamespacedReplicaSetScale struct {
	Context *middleware.Context
	Handler ReadAppsV1NamespacedReplicaSetScaleHandler
}

func (o *ReadAppsV1NamespacedReplicaSetScale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAppsV1NamespacedReplicaSetScaleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
