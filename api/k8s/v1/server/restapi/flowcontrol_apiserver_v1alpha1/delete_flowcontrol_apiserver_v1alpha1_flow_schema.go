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

// DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc turns a function with the right signature into a delete flowcontrol apiserver v1alpha1 flow schema handler
type DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc func(DeleteFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc) Handle(params DeleteFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
	return fn(params)
}

// DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandler interface for that can handle valid delete flowcontrol apiserver v1alpha1 flow schema params
type DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandler interface {
	Handle(DeleteFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder
}

// NewDeleteFlowcontrolApiserverV1alpha1FlowSchema creates a new http.Handler for the delete flowcontrol apiserver v1alpha1 flow schema operation
func NewDeleteFlowcontrolApiserverV1alpha1FlowSchema(ctx *middleware.Context, handler DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandler) *DeleteFlowcontrolApiserverV1alpha1FlowSchema {
	return &DeleteFlowcontrolApiserverV1alpha1FlowSchema{Context: ctx, Handler: handler}
}

/*DeleteFlowcontrolApiserverV1alpha1FlowSchema swagger:route DELETE /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name} flowcontrolApiserver_v1alpha1 deleteFlowcontrolApiserverV1alpha1FlowSchema

delete a FlowSchema

*/
type DeleteFlowcontrolApiserverV1alpha1FlowSchema struct {
	Context *middleware.Context
	Handler DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandler
}

func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
