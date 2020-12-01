// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandlerFunc turns a function with the right signature into a watch admissionregistration v1beta1 validating webhook configuration list handler
type WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandlerFunc func(WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandlerFunc) Handle(params WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListParams) middleware.Responder {
	return fn(params)
}

// WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandler interface for that can handle valid watch admissionregistration v1beta1 validating webhook configuration list params
type WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandler interface {
	Handle(WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListParams) middleware.Responder
}

// NewWatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList creates a new http.Handler for the watch admissionregistration v1beta1 validating webhook configuration list operation
func NewWatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList(ctx *middleware.Context, handler WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandler) *WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList {
	return &WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList{Context: ctx, Handler: handler}
}

/*WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList swagger:route GET /apis/admissionregistration.k8s.io/v1beta1/watch/validatingwebhookconfigurations admissionregistration_v1beta1 watchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList

watch individual changes to a list of ValidatingWebhookConfiguration. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList struct {
	Context *middleware.Context
	Handler WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandler
}

func (o *WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
