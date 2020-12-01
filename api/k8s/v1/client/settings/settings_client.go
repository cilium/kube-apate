// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSettingsAPIGroup(params *GetSettingsAPIGroupParams) (*GetSettingsAPIGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSettingsAPIGroup get information of a group
*/
func (a *Client) GetSettingsAPIGroup(params *GetSettingsAPIGroupParams) (*GetSettingsAPIGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsAPIGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSettingsAPIGroup",
		Method:             "GET",
		PathPattern:        "/apis/settings.k8s.io/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSettingsAPIGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSettingsAPIGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSettingsAPIGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
