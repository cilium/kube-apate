// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication v1beta1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication v1beta1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAuthenticationV1beta1TokenReview(params *CreateAuthenticationV1beta1TokenReviewParams) (*CreateAuthenticationV1beta1TokenReviewOK, *CreateAuthenticationV1beta1TokenReviewCreated, *CreateAuthenticationV1beta1TokenReviewAccepted, error)

	GetAuthenticationV1beta1APIResources(params *GetAuthenticationV1beta1APIResourcesParams) (*GetAuthenticationV1beta1APIResourcesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAuthenticationV1beta1TokenReview create a TokenReview
*/
func (a *Client) CreateAuthenticationV1beta1TokenReview(params *CreateAuthenticationV1beta1TokenReviewParams) (*CreateAuthenticationV1beta1TokenReviewOK, *CreateAuthenticationV1beta1TokenReviewCreated, *CreateAuthenticationV1beta1TokenReviewAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthenticationV1beta1TokenReviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthenticationV1beta1TokenReview",
		Method:             "POST",
		PathPattern:        "/apis/authentication.k8s.io/v1beta1/tokenreviews",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAuthenticationV1beta1TokenReviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAuthenticationV1beta1TokenReviewOK:
		return value, nil, nil, nil
	case *CreateAuthenticationV1beta1TokenReviewCreated:
		return nil, value, nil, nil
	case *CreateAuthenticationV1beta1TokenReviewAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authentication_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAuthenticationV1beta1APIResources get available resources
*/
func (a *Client) GetAuthenticationV1beta1APIResources(params *GetAuthenticationV1beta1APIResourcesParams) (*GetAuthenticationV1beta1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthenticationV1beta1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuthenticationV1beta1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/authentication.k8s.io/v1beta1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuthenticationV1beta1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthenticationV1beta1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAuthenticationV1beta1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
