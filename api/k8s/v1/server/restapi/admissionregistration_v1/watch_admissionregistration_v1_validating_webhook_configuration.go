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

// WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc turns a function with the right signature into a watch admissionregistration v1 validating webhook configuration handler
type WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc func(WatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc) Handle(params WatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface for that can handle valid watch admissionregistration v1 validating webhook configuration params
type WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler interface {
	Handle(WatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder
}

// NewWatchAdmissionregistrationV1ValidatingWebhookConfiguration creates a new http.Handler for the watch admissionregistration v1 validating webhook configuration operation
func NewWatchAdmissionregistrationV1ValidatingWebhookConfiguration(ctx *middleware.Context, handler WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler) *WatchAdmissionregistrationV1ValidatingWebhookConfiguration {
	return &WatchAdmissionregistrationV1ValidatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*WatchAdmissionregistrationV1ValidatingWebhookConfiguration swagger:route GET /apis/admissionregistration.k8s.io/v1/watch/validatingwebhookconfigurations/{name} admissionregistration_v1 watchAdmissionregistrationV1ValidatingWebhookConfiguration

watch changes to an object of kind ValidatingWebhookConfiguration. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchAdmissionregistrationV1ValidatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler
}

func (o *WatchAdmissionregistrationV1ValidatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAdmissionregistrationV1ValidatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}