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

// ReadCoreV1NamespacedReplicationControllerStatusHandlerFunc turns a function with the right signature into a read core v1 namespaced replication controller status handler
type ReadCoreV1NamespacedReplicationControllerStatusHandlerFunc func(ReadCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedReplicationControllerStatusHandlerFunc) Handle(params ReadCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespacedReplicationControllerStatusHandler interface for that can handle valid read core v1 namespaced replication controller status params
type ReadCoreV1NamespacedReplicationControllerStatusHandler interface {
	Handle(ReadCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder
}

// NewReadCoreV1NamespacedReplicationControllerStatus creates a new http.Handler for the read core v1 namespaced replication controller status operation
func NewReadCoreV1NamespacedReplicationControllerStatus(ctx *middleware.Context, handler ReadCoreV1NamespacedReplicationControllerStatusHandler) *ReadCoreV1NamespacedReplicationControllerStatus {
	return &ReadCoreV1NamespacedReplicationControllerStatus{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedReplicationControllerStatus swagger:route GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status core_v1 readCoreV1NamespacedReplicationControllerStatus

read status of the specified ReplicationController

*/
type ReadCoreV1NamespacedReplicationControllerStatus struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedReplicationControllerStatusHandler
}

func (o *ReadCoreV1NamespacedReplicationControllerStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedReplicationControllerStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
