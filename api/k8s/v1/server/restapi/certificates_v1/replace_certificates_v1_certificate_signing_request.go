// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCertificatesV1CertificateSigningRequestHandlerFunc turns a function with the right signature into a replace certificates v1 certificate signing request handler
type ReplaceCertificatesV1CertificateSigningRequestHandlerFunc func(ReplaceCertificatesV1CertificateSigningRequestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCertificatesV1CertificateSigningRequestHandlerFunc) Handle(params ReplaceCertificatesV1CertificateSigningRequestParams) middleware.Responder {
	return fn(params)
}

// ReplaceCertificatesV1CertificateSigningRequestHandler interface for that can handle valid replace certificates v1 certificate signing request params
type ReplaceCertificatesV1CertificateSigningRequestHandler interface {
	Handle(ReplaceCertificatesV1CertificateSigningRequestParams) middleware.Responder
}

// NewReplaceCertificatesV1CertificateSigningRequest creates a new http.Handler for the replace certificates v1 certificate signing request operation
func NewReplaceCertificatesV1CertificateSigningRequest(ctx *middleware.Context, handler ReplaceCertificatesV1CertificateSigningRequestHandler) *ReplaceCertificatesV1CertificateSigningRequest {
	return &ReplaceCertificatesV1CertificateSigningRequest{Context: ctx, Handler: handler}
}

/*ReplaceCertificatesV1CertificateSigningRequest swagger:route PUT /apis/certificates.k8s.io/v1/certificatesigningrequests/{name} certificates_v1 replaceCertificatesV1CertificateSigningRequest

replace the specified CertificateSigningRequest

*/
type ReplaceCertificatesV1CertificateSigningRequest struct {
	Context *middleware.Context
	Handler ReplaceCertificatesV1CertificateSigningRequestHandler
}

func (o *ReplaceCertificatesV1CertificateSigningRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCertificatesV1CertificateSigningRequestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}