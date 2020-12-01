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

// PatchCertificatesV1CertificateSigningRequestHandlerFunc turns a function with the right signature into a patch certificates v1 certificate signing request handler
type PatchCertificatesV1CertificateSigningRequestHandlerFunc func(PatchCertificatesV1CertificateSigningRequestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchCertificatesV1CertificateSigningRequestHandlerFunc) Handle(params PatchCertificatesV1CertificateSigningRequestParams) middleware.Responder {
	return fn(params)
}

// PatchCertificatesV1CertificateSigningRequestHandler interface for that can handle valid patch certificates v1 certificate signing request params
type PatchCertificatesV1CertificateSigningRequestHandler interface {
	Handle(PatchCertificatesV1CertificateSigningRequestParams) middleware.Responder
}

// NewPatchCertificatesV1CertificateSigningRequest creates a new http.Handler for the patch certificates v1 certificate signing request operation
func NewPatchCertificatesV1CertificateSigningRequest(ctx *middleware.Context, handler PatchCertificatesV1CertificateSigningRequestHandler) *PatchCertificatesV1CertificateSigningRequest {
	return &PatchCertificatesV1CertificateSigningRequest{Context: ctx, Handler: handler}
}

/*PatchCertificatesV1CertificateSigningRequest swagger:route PATCH /apis/certificates.k8s.io/v1/certificatesigningrequests/{name} certificates_v1 patchCertificatesV1CertificateSigningRequest

partially update the specified CertificateSigningRequest

*/
type PatchCertificatesV1CertificateSigningRequest struct {
	Context *middleware.Context
	Handler PatchCertificatesV1CertificateSigningRequestHandler
}

func (o *PatchCertificatesV1CertificateSigningRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchCertificatesV1CertificateSigningRequestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
