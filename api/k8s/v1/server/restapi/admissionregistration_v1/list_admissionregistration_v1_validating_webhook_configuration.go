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

// ListAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc turns a function with the right signature into a list admissionregistration v1 validating webhook configuration handler
type ListAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc func(ListAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc) Handle(params ListAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// ListAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface for that can handle valid list admissionregistration v1 validating webhook configuration params
type ListAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface {
	Handle(ListAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder
}

// NewListAdmissionregistrationV1ValidatingWebhookConfiguration creates a new http.Handler for the list admissionregistration v1 validating webhook configuration operation
func NewListAdmissionregistrationV1ValidatingWebhookConfiguration(ctx *middleware.Context, handler ListAdmissionregistrationV1ValidatingWebhookConfigurationHandler) *ListAdmissionregistrationV1ValidatingWebhookConfiguration {
	return &ListAdmissionregistrationV1ValidatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*ListAdmissionregistrationV1ValidatingWebhookConfiguration swagger:route GET /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations admissionregistration_v1 listAdmissionregistrationV1ValidatingWebhookConfiguration

list or watch objects of kind ValidatingWebhookConfiguration

*/
type ListAdmissionregistrationV1ValidatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler ListAdmissionregistrationV1ValidatingWebhookConfigurationHandler
}

func (o *ListAdmissionregistrationV1ValidatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAdmissionregistrationV1ValidatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
