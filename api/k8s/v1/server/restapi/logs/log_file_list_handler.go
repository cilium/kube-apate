// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LogFileListHandlerHandlerFunc turns a function with the right signature into a log file list handler handler
type LogFileListHandlerHandlerFunc func(LogFileListHandlerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LogFileListHandlerHandlerFunc) Handle(params LogFileListHandlerParams) middleware.Responder {
	return fn(params)
}

// LogFileListHandlerHandler interface for that can handle valid log file list handler params
type LogFileListHandlerHandler interface {
	Handle(LogFileListHandlerParams) middleware.Responder
}

// NewLogFileListHandler creates a new http.Handler for the log file list handler operation
func NewLogFileListHandler(ctx *middleware.Context, handler LogFileListHandlerHandler) *LogFileListHandler {
	return &LogFileListHandler{Context: ctx, Handler: handler}
}

/*LogFileListHandler swagger:route GET /logs/ logs logFileListHandler

LogFileListHandler log file list handler API

*/
type LogFileListHandler struct {
	Context *middleware.Context
	Handler LogFileListHandlerHandler
}

func (o *LogFileListHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLogFileListHandlerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}