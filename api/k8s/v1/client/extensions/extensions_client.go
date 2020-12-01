// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new extensions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for extensions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetExtensionsAPIGroup(params *GetExtensionsAPIGroupParams) (*GetExtensionsAPIGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetExtensionsAPIGroup get information of a group
*/
func (a *Client) GetExtensionsAPIGroup(params *GetExtensionsAPIGroupParams) (*GetExtensionsAPIGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionsAPIGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getExtensionsAPIGroup",
		Method:             "GET",
		PathPattern:        "/apis/extensions/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetExtensionsAPIGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetExtensionsAPIGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExtensionsAPIGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
