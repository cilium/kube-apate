// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authorization v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authorization v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAuthorizationV1NamespacedLocalSubjectAccessReview(params *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewParams) (*CreateAuthorizationV1NamespacedLocalSubjectAccessReviewOK, *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewCreated, *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewAccepted, error)

	CreateAuthorizationV1SelfSubjectAccessReview(params *CreateAuthorizationV1SelfSubjectAccessReviewParams) (*CreateAuthorizationV1SelfSubjectAccessReviewOK, *CreateAuthorizationV1SelfSubjectAccessReviewCreated, *CreateAuthorizationV1SelfSubjectAccessReviewAccepted, error)

	CreateAuthorizationV1SelfSubjectRulesReview(params *CreateAuthorizationV1SelfSubjectRulesReviewParams) (*CreateAuthorizationV1SelfSubjectRulesReviewOK, *CreateAuthorizationV1SelfSubjectRulesReviewCreated, *CreateAuthorizationV1SelfSubjectRulesReviewAccepted, error)

	CreateAuthorizationV1SubjectAccessReview(params *CreateAuthorizationV1SubjectAccessReviewParams) (*CreateAuthorizationV1SubjectAccessReviewOK, *CreateAuthorizationV1SubjectAccessReviewCreated, *CreateAuthorizationV1SubjectAccessReviewAccepted, error)

	GetAuthorizationV1APIResources(params *GetAuthorizationV1APIResourcesParams) (*GetAuthorizationV1APIResourcesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAuthorizationV1NamespacedLocalSubjectAccessReview create a LocalSubjectAccessReview
*/
func (a *Client) CreateAuthorizationV1NamespacedLocalSubjectAccessReview(params *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewParams) (*CreateAuthorizationV1NamespacedLocalSubjectAccessReviewOK, *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewCreated, *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthorizationV1NamespacedLocalSubjectAccessReviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthorizationV1NamespacedLocalSubjectAccessReview",
		Method:             "POST",
		PathPattern:        "/apis/authorization.k8s.io/v1/namespaces/{namespace}/localsubjectaccessreviews",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAuthorizationV1NamespacedLocalSubjectAccessReviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewOK:
		return value, nil, nil, nil
	case *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewCreated:
		return nil, value, nil, nil
	case *CreateAuthorizationV1NamespacedLocalSubjectAccessReviewAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authorization_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAuthorizationV1SelfSubjectAccessReview create a SelfSubjectAccessReview
*/
func (a *Client) CreateAuthorizationV1SelfSubjectAccessReview(params *CreateAuthorizationV1SelfSubjectAccessReviewParams) (*CreateAuthorizationV1SelfSubjectAccessReviewOK, *CreateAuthorizationV1SelfSubjectAccessReviewCreated, *CreateAuthorizationV1SelfSubjectAccessReviewAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthorizationV1SelfSubjectAccessReviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthorizationV1SelfSubjectAccessReview",
		Method:             "POST",
		PathPattern:        "/apis/authorization.k8s.io/v1/selfsubjectaccessreviews",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAuthorizationV1SelfSubjectAccessReviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAuthorizationV1SelfSubjectAccessReviewOK:
		return value, nil, nil, nil
	case *CreateAuthorizationV1SelfSubjectAccessReviewCreated:
		return nil, value, nil, nil
	case *CreateAuthorizationV1SelfSubjectAccessReviewAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authorization_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAuthorizationV1SelfSubjectRulesReview create a SelfSubjectRulesReview
*/
func (a *Client) CreateAuthorizationV1SelfSubjectRulesReview(params *CreateAuthorizationV1SelfSubjectRulesReviewParams) (*CreateAuthorizationV1SelfSubjectRulesReviewOK, *CreateAuthorizationV1SelfSubjectRulesReviewCreated, *CreateAuthorizationV1SelfSubjectRulesReviewAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthorizationV1SelfSubjectRulesReviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthorizationV1SelfSubjectRulesReview",
		Method:             "POST",
		PathPattern:        "/apis/authorization.k8s.io/v1/selfsubjectrulesreviews",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAuthorizationV1SelfSubjectRulesReviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAuthorizationV1SelfSubjectRulesReviewOK:
		return value, nil, nil, nil
	case *CreateAuthorizationV1SelfSubjectRulesReviewCreated:
		return nil, value, nil, nil
	case *CreateAuthorizationV1SelfSubjectRulesReviewAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authorization_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAuthorizationV1SubjectAccessReview create a SubjectAccessReview
*/
func (a *Client) CreateAuthorizationV1SubjectAccessReview(params *CreateAuthorizationV1SubjectAccessReviewParams) (*CreateAuthorizationV1SubjectAccessReviewOK, *CreateAuthorizationV1SubjectAccessReviewCreated, *CreateAuthorizationV1SubjectAccessReviewAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthorizationV1SubjectAccessReviewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthorizationV1SubjectAccessReview",
		Method:             "POST",
		PathPattern:        "/apis/authorization.k8s.io/v1/subjectaccessreviews",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAuthorizationV1SubjectAccessReviewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateAuthorizationV1SubjectAccessReviewOK:
		return value, nil, nil, nil
	case *CreateAuthorizationV1SubjectAccessReviewCreated:
		return nil, value, nil, nil
	case *CreateAuthorizationV1SubjectAccessReviewAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authorization_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAuthorizationV1APIResources get available resources
*/
func (a *Client) GetAuthorizationV1APIResources(params *GetAuthorizationV1APIResourcesParams) (*GetAuthorizationV1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthorizationV1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuthorizationV1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/authorization.k8s.io/v1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuthorizationV1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthorizationV1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAuthorizationV1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
