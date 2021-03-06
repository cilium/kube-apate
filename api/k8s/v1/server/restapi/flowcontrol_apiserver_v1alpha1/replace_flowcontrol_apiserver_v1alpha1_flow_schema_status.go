// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc turns a function with the right signature into a replace flowcontrol apiserver v1alpha1 flow schema status handler
type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc func(ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc) Handle(params ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler interface for that can handle valid replace flowcontrol apiserver v1alpha1 flow schema status params
type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler interface {
	Handle(ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder
}

// NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus creates a new http.Handler for the replace flowcontrol apiserver v1alpha1 flow schema status operation
func NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus(ctx *middleware.Context, handler ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler) *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus {
	return &ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus{Context: ctx, Handler: handler}
}

/*ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus swagger:route PUT /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status flowcontrolApiserver_v1alpha1 replaceFlowcontrolApiserverV1alpha1FlowSchemaStatus

replace status of the specified FlowSchema

*/
type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus struct {
	Context *middleware.Context
	Handler ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler
}

func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
