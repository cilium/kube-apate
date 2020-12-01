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

// PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc turns a function with the right signature into a patch admissionregistration v1 validating webhook configuration handler
type PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc func(PatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc) Handle(params PatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface for that can handle valid patch admissionregistration v1 validating webhook configuration params
type PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface {
	Handle(PatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder
}

// NewPatchAdmissionregistrationV1ValidatingWebhookConfiguration creates a new http.Handler for the patch admissionregistration v1 validating webhook configuration operation
func NewPatchAdmissionregistrationV1ValidatingWebhookConfiguration(ctx *middleware.Context, handler PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler) *PatchAdmissionregistrationV1ValidatingWebhookConfiguration {
	return &PatchAdmissionregistrationV1ValidatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*PatchAdmissionregistrationV1ValidatingWebhookConfiguration swagger:route PATCH /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name} admissionregistration_v1 patchAdmissionregistrationV1ValidatingWebhookConfiguration

partially update the specified ValidatingWebhookConfiguration

*/
type PatchAdmissionregistrationV1ValidatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler
}

func (o *PatchAdmissionregistrationV1ValidatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAdmissionregistrationV1ValidatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
