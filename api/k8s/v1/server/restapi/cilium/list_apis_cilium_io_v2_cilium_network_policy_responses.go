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

// ListApisCiliumIoV2CiliumNetworkPolicyOKCode is the HTTP code returned for type ListApisCiliumIoV2CiliumNetworkPolicyOK
const ListApisCiliumIoV2CiliumNetworkPolicyOKCode int = 200

/*ListApisCiliumIoV2CiliumNetworkPolicyOK OK

swagger:response listApisCiliumIoV2CiliumNetworkPolicyOK
*/
type ListApisCiliumIoV2CiliumNetworkPolicyOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListApisCiliumIoV2CiliumNetworkPolicyOK creates ListApisCiliumIoV2CiliumNetworkPolicyOK with default headers values
func NewListApisCiliumIoV2CiliumNetworkPolicyOK() *ListApisCiliumIoV2CiliumNetworkPolicyOK {

	return &ListApisCiliumIoV2CiliumNetworkPolicyOK{}
}

// WithPayload adds the payload to the list apis cilium io v2 cilium network policy o k response
func (o *ListApisCiliumIoV2CiliumNetworkPolicyOK) WithPayload(payload interface{}) *ListApisCiliumIoV2CiliumNetworkPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list apis cilium io v2 cilium network policy o k response
func (o *ListApisCiliumIoV2CiliumNetworkPolicyOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListApisCiliumIoV2CiliumNetworkPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListApisCiliumIoV2CiliumNetworkPolicyUnauthorizedCode is the HTTP code returned for type ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized
const ListApisCiliumIoV2CiliumNetworkPolicyUnauthorizedCode int = 401

/*ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized Unauthorized

swagger:response listApisCiliumIoV2CiliumNetworkPolicyUnauthorized
*/
type ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized struct {
}

// NewListApisCiliumIoV2CiliumNetworkPolicyUnauthorized creates ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized with default headers values
func NewListApisCiliumIoV2CiliumNetworkPolicyUnauthorized() *ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized {

	return &ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *ListApisCiliumIoV2CiliumNetworkPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
