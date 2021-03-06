// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAuthenticationV1TokenReview(params *CreateAuthenticationV1TokenReviewParams) (*CreateAuthenticationV1TokenReviewOK, *CreateAuthenticationV1TokenReviewCreated, *CreateAuthenticationV1TokenReviewAccepted, error)

	GetAuthenticationV1APIResources(params *GetAuthenticationV1APIResourcesParams) (*GetAuthenticationV1APIResourcesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAuthenticationV1TokenReview create a TokenReview
*/
func (a *Client) CreateAuthenticationV1TokenReview(params *CreateAuthenticationV1TokenReviewParams) (*CreateAuthenticationV1TokenReviewOK, *CreateAuthenticationV1TokenReviewCreated, *CreateAuthenticationV1TokenReviewAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthenticationV1TokenReviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthenticationV1TokenReview",
		Method:             "POST",
		PathPattern:        "/apis/authentication.k8s.io/v1/tokenreviews",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAuthenticationV1TokenReviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAuthenticationV1TokenReviewOK:
		return value, nil, nil, nil
	case *CreateAuthenticationV1TokenReviewCreated:
		return nil, value, nil, nil
	case *CreateAuthenticationV1TokenReviewAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAuthenticationV1APIResources get available resources
*/
func (a *Client) GetAuthenticationV1APIResources(params *GetAuthenticationV1APIResourcesParams) (*GetAuthenticationV1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthenticationV1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuthenticationV1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/authentication.k8s.io/v1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuthenticationV1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthenticationV1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAuthenticationV1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
