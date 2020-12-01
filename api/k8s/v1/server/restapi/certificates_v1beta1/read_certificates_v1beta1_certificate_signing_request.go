// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCertificatesV1beta1CertificateSigningRequestHandlerFunc turns a function with the right signature into a read certificates v1beta1 certificate signing request handler
type ReadCertificatesV1beta1CertificateSigningRequestHandlerFunc func(ReadCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCertificatesV1beta1CertificateSigningRequestHandlerFunc) Handle(params ReadCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
	return fn(params)
}

// ReadCertificatesV1beta1CertificateSigningRequestHandler interface for that can handle valid read certificates v1beta1 certificate signing request params
type ReadCertificatesV1beta1CertificateSigningRequestHandler interface {
	Handle(ReadCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder
}

// NewReadCertificatesV1beta1CertificateSigningRequest creates a new http.Handler for the read certificates v1beta1 certificate signing request operation
func NewReadCertificatesV1beta1CertificateSigningRequest(ctx *middleware.Context, handler ReadCertificatesV1beta1CertificateSigningRequestHandler) *ReadCertificatesV1beta1CertificateSigningRequest {
	return &ReadCertificatesV1beta1CertificateSigningRequest{Context: ctx, Handler: handler}
}

/*ReadCertificatesV1beta1CertificateSigningRequest swagger:route GET /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name} certificates_v1beta1 readCertificatesV1beta1CertificateSigningRequest

read the specified CertificateSigningRequest

*/
type ReadCertificatesV1beta1CertificateSigningRequest struct {
	Context *middleware.Context
	Handler ReadCertificatesV1beta1CertificateSigningRequestHandler
}

func (o *ReadCertificatesV1beta1CertificateSigningRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCertificatesV1beta1CertificateSigningRequestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
