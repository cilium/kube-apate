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

// PostManagementKubernetesIoV1NodesOKCode is the HTTP code returned for type PostManagementKubernetesIoV1NodesOK
const PostManagementKubernetesIoV1NodesOKCode int = 200

/*PostManagementKubernetesIoV1NodesOK OK

swagger:response postManagementKubernetesIoV1NodesOK
*/
type PostManagementKubernetesIoV1NodesOK struct {
}

// NewPostManagementKubernetesIoV1NodesOK creates PostManagementKubernetesIoV1NodesOK with default headers values
func NewPostManagementKubernetesIoV1NodesOK() *PostManagementKubernetesIoV1NodesOK {

	return &PostManagementKubernetesIoV1NodesOK{}
}

// WriteResponse to the client
func (o *PostManagementKubernetesIoV1NodesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostManagementKubernetesIoV1NodesAcceptedCode is the HTTP code returned for type PostManagementKubernetesIoV1NodesAccepted
const PostManagementKubernetesIoV1NodesAcceptedCode int = 202

/*PostManagementKubernetesIoV1NodesAccepted Accepted

swagger:response postManagementKubernetesIoV1NodesAccepted
*/
type PostManagementKubernetesIoV1NodesAccepted struct {
}

// NewPostManagementKubernetesIoV1NodesAccepted creates PostManagementKubernetesIoV1NodesAccepted with default headers values
func NewPostManagementKubernetesIoV1NodesAccepted() *PostManagementKubernetesIoV1NodesAccepted {

	return &PostManagementKubernetesIoV1NodesAccepted{}
}

// WriteResponse to the client
func (o *PostManagementKubernetesIoV1NodesAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}
