// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAuthorizationV1SelfSubjectRulesReviewHandlerFunc turns a function with the right signature into a create authorization v1 self subject rules review handler
type CreateAuthorizationV1SelfSubjectRulesReviewHandlerFunc func(CreateAuthorizationV1SelfSubjectRulesReviewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuthorizationV1SelfSubjectRulesReviewHandlerFunc) Handle(params CreateAuthorizationV1SelfSubjectRulesReviewParams) middleware.Responder {
	return fn(params)
}

// CreateAuthorizationV1SelfSubjectRulesReviewHandler interface for that can handle valid create authorization v1 self subject rules review params
type CreateAuthorizationV1SelfSubjectRulesReviewHandler interface {
	Handle(CreateAuthorizationV1SelfSubjectRulesReviewParams) middleware.Responder
}

// NewCreateAuthorizationV1SelfSubjectRulesReview creates a new http.Handler for the create authorization v1 self subject rules review operation
func NewCreateAuthorizationV1SelfSubjectRulesReview(ctx *middleware.Context, handler CreateAuthorizationV1SelfSubjectRulesReviewHandler) *CreateAuthorizationV1SelfSubjectRulesReview {
	return &CreateAuthorizationV1SelfSubjectRulesReview{Context: ctx, Handler: handler}
}

/*CreateAuthorizationV1SelfSubjectRulesReview swagger:route POST /apis/authorization.k8s.io/v1/selfsubjectrulesreviews authorization_v1 createAuthorizationV1SelfSubjectRulesReview

create a SelfSubjectRulesReview

*/
type CreateAuthorizationV1SelfSubjectRulesReview struct {
	Context *middleware.Context
	Handler CreateAuthorizationV1SelfSubjectRulesReviewHandler
}

func (o *CreateAuthorizationV1SelfSubjectRulesReview) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAuthorizationV1SelfSubjectRulesReviewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
