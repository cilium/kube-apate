// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostManagementCiliumIoV2CiliumIdentitiesOKCode is the HTTP code returned for type PostManagementCiliumIoV2CiliumIdentitiesOK
const PostManagementCiliumIoV2CiliumIdentitiesOKCode int = 200

/*PostManagementCiliumIoV2CiliumIdentitiesOK OK

swagger:response postManagementCiliumIoV2CiliumIdentitiesOK
*/
type PostManagementCiliumIoV2CiliumIdentitiesOK struct {
}

// NewPostManagementCiliumIoV2CiliumIdentitiesOK creates PostManagementCiliumIoV2CiliumIdentitiesOK with default headers values
func NewPostManagementCiliumIoV2CiliumIdentitiesOK() *PostManagementCiliumIoV2CiliumIdentitiesOK {

	return &PostManagementCiliumIoV2CiliumIdentitiesOK{}
}

// WriteResponse to the client
func (o *PostManagementCiliumIoV2CiliumIdentitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostManagementCiliumIoV2CiliumIdentitiesAcceptedCode is the HTTP code returned for type PostManagementCiliumIoV2CiliumIdentitiesAccepted
const PostManagementCiliumIoV2CiliumIdentitiesAcceptedCode int = 202

/*PostManagementCiliumIoV2CiliumIdentitiesAccepted Accepted

swagger:response postManagementCiliumIoV2CiliumIdentitiesAccepted
*/
type PostManagementCiliumIoV2CiliumIdentitiesAccepted struct {
}

// NewPostManagementCiliumIoV2CiliumIdentitiesAccepted creates PostManagementCiliumIoV2CiliumIdentitiesAccepted with default headers values
func NewPostManagementCiliumIoV2CiliumIdentitiesAccepted() *PostManagementCiliumIoV2CiliumIdentitiesAccepted {

	return &PostManagementCiliumIoV2CiliumIdentitiesAccepted{}
}

// WriteResponse to the client
func (o *PostManagementCiliumIoV2CiliumIdentitiesAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}
