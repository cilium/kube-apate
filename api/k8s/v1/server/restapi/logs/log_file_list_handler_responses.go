// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// LogFileListHandlerUnauthorizedCode is the HTTP code returned for type LogFileListHandlerUnauthorized
const LogFileListHandlerUnauthorizedCode int = 401

/*LogFileListHandlerUnauthorized Unauthorized

swagger:response logFileListHandlerUnauthorized
*/
type LogFileListHandlerUnauthorized struct {
}

// NewLogFileListHandlerUnauthorized creates LogFileListHandlerUnauthorized with default headers values
func NewLogFileListHandlerUnauthorized() *LogFileListHandlerUnauthorized {

	return &LogFileListHandlerUnauthorized{}
}

// WriteResponse to the client
func (o *LogFileListHandlerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
