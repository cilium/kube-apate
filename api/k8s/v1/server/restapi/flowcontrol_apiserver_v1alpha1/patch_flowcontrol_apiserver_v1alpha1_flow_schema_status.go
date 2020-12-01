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

// PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc turns a function with the right signature into a patch flowcontrol apiserver v1alpha1 flow schema status handler
type PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc func(PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc) Handle(params PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder {
	return fn(params)
}

// PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler interface for that can handle valid patch flowcontrol apiserver v1alpha1 flow schema status params
type PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler interface {
	Handle(PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder
}

// NewPatchFlowcontrolApiserverV1alpha1FlowSchemaStatus creates a new http.Handler for the patch flowcontrol apiserver v1alpha1 flow schema status operation
func NewPatchFlowcontrolApiserverV1alpha1FlowSchemaStatus(ctx *middleware.Context, handler PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler) *PatchFlowcontrolApiserverV1alpha1FlowSchemaStatus {
	return &PatchFlowcontrolApiserverV1alpha1FlowSchemaStatus{Context: ctx, Handler: handler}
}

/*PatchFlowcontrolApiserverV1alpha1FlowSchemaStatus swagger:route PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status flowcontrolApiserver_v1alpha1 patchFlowcontrolApiserverV1alpha1FlowSchemaStatus

partially update status of the specified FlowSchema

*/
type PatchFlowcontrolApiserverV1alpha1FlowSchemaStatus struct {
	Context *middleware.Context
	Handler PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler
}

func (o *PatchFlowcontrolApiserverV1alpha1FlowSchemaStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchFlowcontrolApiserverV1alpha1FlowSchemaStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
