// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new coordination v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for coordination v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCoordinationV1NamespacedLease(params *CreateCoordinationV1NamespacedLeaseParams) (*CreateCoordinationV1NamespacedLeaseOK, *CreateCoordinationV1NamespacedLeaseCreated, *CreateCoordinationV1NamespacedLeaseAccepted, error)

	DeleteCoordinationV1CollectionNamespacedLease(params *DeleteCoordinationV1CollectionNamespacedLeaseParams) (*DeleteCoordinationV1CollectionNamespacedLeaseOK, error)

	DeleteCoordinationV1NamespacedLease(params *DeleteCoordinationV1NamespacedLeaseParams) (*DeleteCoordinationV1NamespacedLeaseOK, *DeleteCoordinationV1NamespacedLeaseAccepted, error)

	GetCoordinationV1APIResources(params *GetCoordinationV1APIResourcesParams) (*GetCoordinationV1APIResourcesOK, error)

	ListCoordinationV1LeaseForAllNamespaces(params *ListCoordinationV1LeaseForAllNamespacesParams) (*ListCoordinationV1LeaseForAllNamespacesOK, error)

	ListCoordinationV1NamespacedLease(params *ListCoordinationV1NamespacedLeaseParams) (*ListCoordinationV1NamespacedLeaseOK, error)

	PatchCoordinationV1NamespacedLease(params *PatchCoordinationV1NamespacedLeaseParams) (*PatchCoordinationV1NamespacedLeaseOK, error)

	ReadCoordinationV1NamespacedLease(params *ReadCoordinationV1NamespacedLeaseParams) (*ReadCoordinationV1NamespacedLeaseOK, error)

	ReplaceCoordinationV1NamespacedLease(params *ReplaceCoordinationV1NamespacedLeaseParams) (*ReplaceCoordinationV1NamespacedLeaseOK, *ReplaceCoordinationV1NamespacedLeaseCreated, error)

	WatchCoordinationV1LeaseListForAllNamespaces(params *WatchCoordinationV1LeaseListForAllNamespacesParams) (*WatchCoordinationV1LeaseListForAllNamespacesOK, error)

	WatchCoordinationV1NamespacedLease(params *WatchCoordinationV1NamespacedLeaseParams) (*WatchCoordinationV1NamespacedLeaseOK, error)

	WatchCoordinationV1NamespacedLeaseList(params *WatchCoordinationV1NamespacedLeaseListParams) (*WatchCoordinationV1NamespacedLeaseListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCoordinationV1NamespacedLease create a Lease
*/
func (a *Client) CreateCoordinationV1NamespacedLease(params *CreateCoordinationV1NamespacedLeaseParams) (*CreateCoordinationV1NamespacedLeaseOK, *CreateCoordinationV1NamespacedLeaseCreated, *CreateCoordinationV1NamespacedLeaseAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCoordinationV1NamespacedLease",
		Method:             "POST",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateCoordinationV1NamespacedLeaseOK:
		return value, nil, nil, nil
	case *CreateCoordinationV1NamespacedLeaseCreated:
		return nil, value, nil, nil
	case *CreateCoordinationV1NamespacedLeaseAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for coordination_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCoordinationV1CollectionNamespacedLease delete collection of Lease
*/
func (a *Client) DeleteCoordinationV1CollectionNamespacedLease(params *DeleteCoordinationV1CollectionNamespacedLeaseParams) (*DeleteCoordinationV1CollectionNamespacedLeaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCoordinationV1CollectionNamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCoordinationV1CollectionNamespacedLease",
		Method:             "DELETE",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCoordinationV1CollectionNamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCoordinationV1CollectionNamespacedLeaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCoordinationV1CollectionNamespacedLease: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCoordinationV1NamespacedLease delete a Lease
*/
func (a *Client) DeleteCoordinationV1NamespacedLease(params *DeleteCoordinationV1NamespacedLeaseParams) (*DeleteCoordinationV1NamespacedLeaseOK, *DeleteCoordinationV1NamespacedLeaseAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCoordinationV1NamespacedLease",
		Method:             "DELETE",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteCoordinationV1NamespacedLeaseOK:
		return value, nil, nil
	case *DeleteCoordinationV1NamespacedLeaseAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for coordination_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCoordinationV1APIResources get available resources
*/
func (a *Client) GetCoordinationV1APIResources(params *GetCoordinationV1APIResourcesParams) (*GetCoordinationV1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCoordinationV1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCoordinationV1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCoordinationV1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCoordinationV1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCoordinationV1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCoordinationV1LeaseForAllNamespaces list or watch objects of kind Lease
*/
func (a *Client) ListCoordinationV1LeaseForAllNamespaces(params *ListCoordinationV1LeaseForAllNamespacesParams) (*ListCoordinationV1LeaseForAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCoordinationV1LeaseForAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCoordinationV1LeaseForAllNamespaces",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/leases",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCoordinationV1LeaseForAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCoordinationV1LeaseForAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCoordinationV1LeaseForAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCoordinationV1NamespacedLease list or watch objects of kind Lease
*/
func (a *Client) ListCoordinationV1NamespacedLease(params *ListCoordinationV1NamespacedLeaseParams) (*ListCoordinationV1NamespacedLeaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCoordinationV1NamespacedLease",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCoordinationV1NamespacedLeaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCoordinationV1NamespacedLease: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchCoordinationV1NamespacedLease partially update the specified Lease
*/
func (a *Client) PatchCoordinationV1NamespacedLease(params *PatchCoordinationV1NamespacedLeaseParams) (*PatchCoordinationV1NamespacedLeaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCoordinationV1NamespacedLease",
		Method:             "PATCH",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchCoordinationV1NamespacedLeaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchCoordinationV1NamespacedLease: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadCoordinationV1NamespacedLease read the specified Lease
*/
func (a *Client) ReadCoordinationV1NamespacedLease(params *ReadCoordinationV1NamespacedLeaseParams) (*ReadCoordinationV1NamespacedLeaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readCoordinationV1NamespacedLease",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadCoordinationV1NamespacedLeaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readCoordinationV1NamespacedLease: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceCoordinationV1NamespacedLease replace the specified Lease
*/
func (a *Client) ReplaceCoordinationV1NamespacedLease(params *ReplaceCoordinationV1NamespacedLeaseParams) (*ReplaceCoordinationV1NamespacedLeaseOK, *ReplaceCoordinationV1NamespacedLeaseCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceCoordinationV1NamespacedLease",
		Method:             "PUT",
		PathPattern:        "/apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceCoordinationV1NamespacedLeaseOK:
		return value, nil, nil
	case *ReplaceCoordinationV1NamespacedLeaseCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for coordination_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchCoordinationV1LeaseListForAllNamespaces watch individual changes to a list of Lease. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchCoordinationV1LeaseListForAllNamespaces(params *WatchCoordinationV1LeaseListForAllNamespacesParams) (*WatchCoordinationV1LeaseListForAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchCoordinationV1LeaseListForAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchCoordinationV1LeaseListForAllNamespaces",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/watch/leases",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchCoordinationV1LeaseListForAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchCoordinationV1LeaseListForAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchCoordinationV1LeaseListForAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchCoordinationV1NamespacedLease watch changes to an object of kind Lease. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.
*/
func (a *Client) WatchCoordinationV1NamespacedLease(params *WatchCoordinationV1NamespacedLeaseParams) (*WatchCoordinationV1NamespacedLeaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchCoordinationV1NamespacedLeaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchCoordinationV1NamespacedLease",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/watch/namespaces/{namespace}/leases/{name}",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchCoordinationV1NamespacedLeaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchCoordinationV1NamespacedLeaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchCoordinationV1NamespacedLease: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchCoordinationV1NamespacedLeaseList watch individual changes to a list of Lease. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchCoordinationV1NamespacedLeaseList(params *WatchCoordinationV1NamespacedLeaseListParams) (*WatchCoordinationV1NamespacedLeaseListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchCoordinationV1NamespacedLeaseListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchCoordinationV1NamespacedLeaseList",
		Method:             "GET",
		PathPattern:        "/apis/coordination.k8s.io/v1/watch/namespaces/{namespace}/leases",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchCoordinationV1NamespacedLeaseListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchCoordinationV1NamespacedLeaseListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchCoordinationV1NamespacedLeaseList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
