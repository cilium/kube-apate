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

// DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc turns a function with the right signature into a delete admissionregistration v1 validating webhook configuration handler
type DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc func(DeleteAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc) Handle(params DeleteAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface for that can handle valid delete admissionregistration v1 validating webhook configuration params
type DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface {
	Handle(DeleteAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder
}

// NewDeleteAdmissionregistrationV1ValidatingWebhookConfiguration creates a new http.Handler for the delete admissionregistration v1 validating webhook configuration operation
func NewDeleteAdmissionregistrationV1ValidatingWebhookConfiguration(ctx *middleware.Context, handler DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandler) *DeleteAdmissionregistrationV1ValidatingWebhookConfiguration {
	return &DeleteAdmissionregistrationV1ValidatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*DeleteAdmissionregistrationV1ValidatingWebhookConfiguration swagger:route DELETE /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name} admissionregistration_v1 deleteAdmissionregistrationV1ValidatingWebhookConfiguration

delete a ValidatingWebhookConfiguration

*/
type DeleteAdmissionregistrationV1ValidatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandler
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}