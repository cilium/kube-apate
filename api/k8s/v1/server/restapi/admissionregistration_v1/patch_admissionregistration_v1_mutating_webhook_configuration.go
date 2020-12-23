// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a patch admissionregistration v1 mutating webhook configuration handler
type PatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc func(PatchAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc) Handle(params PatchAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// PatchAdmissionregistrationV1MutatingWebhookConfigurationHandler interface for that can handle valid patch admissionregistration v1 mutating webhook configuration params
type PatchAdmissionregistrationV1MutatingWebhookConfigurationHandler interface {
	Handle(PatchAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder
}

// NewPatchAdmissionregistrationV1MutatingWebhookConfiguration creates a new http.Handler for the patch admissionregistration v1 mutating webhook configuration operation
func NewPatchAdmissionregistrationV1MutatingWebhookConfiguration(ctx *middleware.Context, handler PatchAdmissionregistrationV1MutatingWebhookConfigurationHandler) *PatchAdmissionregistrationV1MutatingWebhookConfiguration {
	return &PatchAdmissionregistrationV1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*PatchAdmissionregistrationV1MutatingWebhookConfiguration swagger:route PATCH /apis/admissionregistration.k8s.io/v1/mutatingwebhookconfigurations/{name} admissionregistration_v1 patchAdmissionregistrationV1MutatingWebhookConfiguration

partially update the specified MutatingWebhookConfiguration

*/
type PatchAdmissionregistrationV1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler PatchAdmissionregistrationV1MutatingWebhookConfigurationHandler
}

func (o *PatchAdmissionregistrationV1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAdmissionregistrationV1MutatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}