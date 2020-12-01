// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new node v1alpha1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node v1alpha1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNodeV1alpha1RuntimeClass(params *CreateNodeV1alpha1RuntimeClassParams) (*CreateNodeV1alpha1RuntimeClassOK, *CreateNodeV1alpha1RuntimeClassCreated, *CreateNodeV1alpha1RuntimeClassAccepted, error)

	DeleteNodeV1alpha1CollectionRuntimeClass(params *DeleteNodeV1alpha1CollectionRuntimeClassParams) (*DeleteNodeV1alpha1CollectionRuntimeClassOK, error)

	DeleteNodeV1alpha1RuntimeClass(params *DeleteNodeV1alpha1RuntimeClassParams) (*DeleteNodeV1alpha1RuntimeClassOK, *DeleteNodeV1alpha1RuntimeClassAccepted, error)

	GetNodeV1alpha1APIResources(params *GetNodeV1alpha1APIResourcesParams) (*GetNodeV1alpha1APIResourcesOK, error)

	ListNodeV1alpha1RuntimeClass(params *ListNodeV1alpha1RuntimeClassParams) (*ListNodeV1alpha1RuntimeClassOK, error)

	PatchNodeV1alpha1RuntimeClass(params *PatchNodeV1alpha1RuntimeClassParams) (*PatchNodeV1alpha1RuntimeClassOK, error)

	ReadNodeV1alpha1RuntimeClass(params *ReadNodeV1alpha1RuntimeClassParams) (*ReadNodeV1alpha1RuntimeClassOK, error)

	ReplaceNodeV1alpha1RuntimeClass(params *ReplaceNodeV1alpha1RuntimeClassParams) (*ReplaceNodeV1alpha1RuntimeClassOK, *ReplaceNodeV1alpha1RuntimeClassCreated, error)

	WatchNodeV1alpha1RuntimeClass(params *WatchNodeV1alpha1RuntimeClassParams) (*WatchNodeV1alpha1RuntimeClassOK, error)

	WatchNodeV1alpha1RuntimeClassList(params *WatchNodeV1alpha1RuntimeClassListParams) (*WatchNodeV1alpha1RuntimeClassListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNodeV1alpha1RuntimeClass create a RuntimeClass
*/
func (a *Client) CreateNodeV1alpha1RuntimeClass(params *CreateNodeV1alpha1RuntimeClassParams) (*CreateNodeV1alpha1RuntimeClassOK, *CreateNodeV1alpha1RuntimeClassCreated, *CreateNodeV1alpha1RuntimeClassAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNodeV1alpha1RuntimeClass",
		Method:             "POST",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateNodeV1alpha1RuntimeClassOK:
		return value, nil, nil, nil
	case *CreateNodeV1alpha1RuntimeClassCreated:
		return nil, value, nil, nil
	case *CreateNodeV1alpha1RuntimeClassAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for node_v1alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNodeV1alpha1CollectionRuntimeClass delete collection of RuntimeClass
*/
func (a *Client) DeleteNodeV1alpha1CollectionRuntimeClass(params *DeleteNodeV1alpha1CollectionRuntimeClassParams) (*DeleteNodeV1alpha1CollectionRuntimeClassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodeV1alpha1CollectionRuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodeV1alpha1CollectionRuntimeClass",
		Method:             "DELETE",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNodeV1alpha1CollectionRuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNodeV1alpha1CollectionRuntimeClassOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNodeV1alpha1CollectionRuntimeClass: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNodeV1alpha1RuntimeClass delete a RuntimeClass
*/
func (a *Client) DeleteNodeV1alpha1RuntimeClass(params *DeleteNodeV1alpha1RuntimeClassParams) (*DeleteNodeV1alpha1RuntimeClassOK, *DeleteNodeV1alpha1RuntimeClassAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodeV1alpha1RuntimeClass",
		Method:             "DELETE",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteNodeV1alpha1RuntimeClassOK:
		return value, nil, nil
	case *DeleteNodeV1alpha1RuntimeClassAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for node_v1alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNodeV1alpha1APIResources get available resources
*/
func (a *Client) GetNodeV1alpha1APIResources(params *GetNodeV1alpha1APIResourcesParams) (*GetNodeV1alpha1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeV1alpha1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodeV1alpha1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNodeV1alpha1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNodeV1alpha1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNodeV1alpha1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNodeV1alpha1RuntimeClass list or watch objects of kind RuntimeClass
*/
func (a *Client) ListNodeV1alpha1RuntimeClass(params *ListNodeV1alpha1RuntimeClassParams) (*ListNodeV1alpha1RuntimeClassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodeV1alpha1RuntimeClass",
		Method:             "GET",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListNodeV1alpha1RuntimeClassOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listNodeV1alpha1RuntimeClass: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchNodeV1alpha1RuntimeClass partially update the specified RuntimeClass
*/
func (a *Client) PatchNodeV1alpha1RuntimeClass(params *PatchNodeV1alpha1RuntimeClassParams) (*PatchNodeV1alpha1RuntimeClassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchNodeV1alpha1RuntimeClass",
		Method:             "PATCH",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchNodeV1alpha1RuntimeClassOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchNodeV1alpha1RuntimeClass: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadNodeV1alpha1RuntimeClass read the specified RuntimeClass
*/
func (a *Client) ReadNodeV1alpha1RuntimeClass(params *ReadNodeV1alpha1RuntimeClassParams) (*ReadNodeV1alpha1RuntimeClassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readNodeV1alpha1RuntimeClass",
		Method:             "GET",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadNodeV1alpha1RuntimeClassOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readNodeV1alpha1RuntimeClass: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceNodeV1alpha1RuntimeClass replace the specified RuntimeClass
*/
func (a *Client) ReplaceNodeV1alpha1RuntimeClass(params *ReplaceNodeV1alpha1RuntimeClassParams) (*ReplaceNodeV1alpha1RuntimeClassOK, *ReplaceNodeV1alpha1RuntimeClassCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceNodeV1alpha1RuntimeClass",
		Method:             "PUT",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/runtimeclasses/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceNodeV1alpha1RuntimeClassOK:
		return value, nil, nil
	case *ReplaceNodeV1alpha1RuntimeClassCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for node_v1alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchNodeV1alpha1RuntimeClass watch changes to an object of kind RuntimeClass. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.
*/
func (a *Client) WatchNodeV1alpha1RuntimeClass(params *WatchNodeV1alpha1RuntimeClassParams) (*WatchNodeV1alpha1RuntimeClassOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchNodeV1alpha1RuntimeClassParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchNodeV1alpha1RuntimeClass",
		Method:             "GET",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/watch/runtimeclasses/{name}",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchNodeV1alpha1RuntimeClassReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchNodeV1alpha1RuntimeClassOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchNodeV1alpha1RuntimeClass: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchNodeV1alpha1RuntimeClassList watch individual changes to a list of RuntimeClass. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchNodeV1alpha1RuntimeClassList(params *WatchNodeV1alpha1RuntimeClassListParams) (*WatchNodeV1alpha1RuntimeClassListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchNodeV1alpha1RuntimeClassListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchNodeV1alpha1RuntimeClassList",
		Method:             "GET",
		PathPattern:        "/apis/node.k8s.io/v1alpha1/watch/runtimeclasses",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchNodeV1alpha1RuntimeClassListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchNodeV1alpha1RuntimeClassListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchNodeV1alpha1RuntimeClassList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
