// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new storage v1alpha1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage v1alpha1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateStorageV1alpha1VolumeAttachment(params *CreateStorageV1alpha1VolumeAttachmentParams) (*CreateStorageV1alpha1VolumeAttachmentOK, *CreateStorageV1alpha1VolumeAttachmentCreated, *CreateStorageV1alpha1VolumeAttachmentAccepted, error)

	DeleteStorageV1alpha1CollectionVolumeAttachment(params *DeleteStorageV1alpha1CollectionVolumeAttachmentParams) (*DeleteStorageV1alpha1CollectionVolumeAttachmentOK, error)

	DeleteStorageV1alpha1VolumeAttachment(params *DeleteStorageV1alpha1VolumeAttachmentParams) (*DeleteStorageV1alpha1VolumeAttachmentOK, *DeleteStorageV1alpha1VolumeAttachmentAccepted, error)

	GetStorageV1alpha1APIResources(params *GetStorageV1alpha1APIResourcesParams) (*GetStorageV1alpha1APIResourcesOK, error)

	ListStorageV1alpha1VolumeAttachment(params *ListStorageV1alpha1VolumeAttachmentParams) (*ListStorageV1alpha1VolumeAttachmentOK, error)

	PatchStorageV1alpha1VolumeAttachment(params *PatchStorageV1alpha1VolumeAttachmentParams) (*PatchStorageV1alpha1VolumeAttachmentOK, error)

	ReadStorageV1alpha1VolumeAttachment(params *ReadStorageV1alpha1VolumeAttachmentParams) (*ReadStorageV1alpha1VolumeAttachmentOK, error)

	ReplaceStorageV1alpha1VolumeAttachment(params *ReplaceStorageV1alpha1VolumeAttachmentParams) (*ReplaceStorageV1alpha1VolumeAttachmentOK, *ReplaceStorageV1alpha1VolumeAttachmentCreated, error)

	WatchStorageV1alpha1VolumeAttachment(params *WatchStorageV1alpha1VolumeAttachmentParams) (*WatchStorageV1alpha1VolumeAttachmentOK, error)

	WatchStorageV1alpha1VolumeAttachmentList(params *WatchStorageV1alpha1VolumeAttachmentListParams) (*WatchStorageV1alpha1VolumeAttachmentListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateStorageV1alpha1VolumeAttachment create a VolumeAttachment
*/
func (a *Client) CreateStorageV1alpha1VolumeAttachment(params *CreateStorageV1alpha1VolumeAttachmentParams) (*CreateStorageV1alpha1VolumeAttachmentOK, *CreateStorageV1alpha1VolumeAttachmentCreated, *CreateStorageV1alpha1VolumeAttachmentAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStorageV1alpha1VolumeAttachment",
		Method:             "POST",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateStorageV1alpha1VolumeAttachmentOK:
		return value, nil, nil, nil
	case *CreateStorageV1alpha1VolumeAttachmentCreated:
		return nil, value, nil, nil
	case *CreateStorageV1alpha1VolumeAttachmentAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storage_v1alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteStorageV1alpha1CollectionVolumeAttachment delete collection of VolumeAttachment
*/
func (a *Client) DeleteStorageV1alpha1CollectionVolumeAttachment(params *DeleteStorageV1alpha1CollectionVolumeAttachmentParams) (*DeleteStorageV1alpha1CollectionVolumeAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageV1alpha1CollectionVolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStorageV1alpha1CollectionVolumeAttachment",
		Method:             "DELETE",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteStorageV1alpha1CollectionVolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteStorageV1alpha1CollectionVolumeAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteStorageV1alpha1CollectionVolumeAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteStorageV1alpha1VolumeAttachment delete a VolumeAttachment
*/
func (a *Client) DeleteStorageV1alpha1VolumeAttachment(params *DeleteStorageV1alpha1VolumeAttachmentParams) (*DeleteStorageV1alpha1VolumeAttachmentOK, *DeleteStorageV1alpha1VolumeAttachmentAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStorageV1alpha1VolumeAttachment",
		Method:             "DELETE",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteStorageV1alpha1VolumeAttachmentOK:
		return value, nil, nil
	case *DeleteStorageV1alpha1VolumeAttachmentAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storage_v1alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStorageV1alpha1APIResources get available resources
*/
func (a *Client) GetStorageV1alpha1APIResources(params *GetStorageV1alpha1APIResourcesParams) (*GetStorageV1alpha1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageV1alpha1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStorageV1alpha1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStorageV1alpha1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStorageV1alpha1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStorageV1alpha1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListStorageV1alpha1VolumeAttachment list or watch objects of kind VolumeAttachment
*/
func (a *Client) ListStorageV1alpha1VolumeAttachment(params *ListStorageV1alpha1VolumeAttachmentParams) (*ListStorageV1alpha1VolumeAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listStorageV1alpha1VolumeAttachment",
		Method:             "GET",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListStorageV1alpha1VolumeAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listStorageV1alpha1VolumeAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchStorageV1alpha1VolumeAttachment partially update the specified VolumeAttachment
*/
func (a *Client) PatchStorageV1alpha1VolumeAttachment(params *PatchStorageV1alpha1VolumeAttachmentParams) (*PatchStorageV1alpha1VolumeAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchStorageV1alpha1VolumeAttachment",
		Method:             "PATCH",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchStorageV1alpha1VolumeAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchStorageV1alpha1VolumeAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadStorageV1alpha1VolumeAttachment read the specified VolumeAttachment
*/
func (a *Client) ReadStorageV1alpha1VolumeAttachment(params *ReadStorageV1alpha1VolumeAttachmentParams) (*ReadStorageV1alpha1VolumeAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readStorageV1alpha1VolumeAttachment",
		Method:             "GET",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadStorageV1alpha1VolumeAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readStorageV1alpha1VolumeAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceStorageV1alpha1VolumeAttachment replace the specified VolumeAttachment
*/
func (a *Client) ReplaceStorageV1alpha1VolumeAttachment(params *ReplaceStorageV1alpha1VolumeAttachmentParams) (*ReplaceStorageV1alpha1VolumeAttachmentOK, *ReplaceStorageV1alpha1VolumeAttachmentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceStorageV1alpha1VolumeAttachment",
		Method:             "PUT",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/volumeattachments/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceStorageV1alpha1VolumeAttachmentOK:
		return value, nil, nil
	case *ReplaceStorageV1alpha1VolumeAttachmentCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for storage_v1alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchStorageV1alpha1VolumeAttachment watch changes to an object of kind VolumeAttachment. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.
*/
func (a *Client) WatchStorageV1alpha1VolumeAttachment(params *WatchStorageV1alpha1VolumeAttachmentParams) (*WatchStorageV1alpha1VolumeAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchStorageV1alpha1VolumeAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchStorageV1alpha1VolumeAttachment",
		Method:             "GET",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/watch/volumeattachments/{name}",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchStorageV1alpha1VolumeAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchStorageV1alpha1VolumeAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchStorageV1alpha1VolumeAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchStorageV1alpha1VolumeAttachmentList watch individual changes to a list of VolumeAttachment. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchStorageV1alpha1VolumeAttachmentList(params *WatchStorageV1alpha1VolumeAttachmentListParams) (*WatchStorageV1alpha1VolumeAttachmentListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchStorageV1alpha1VolumeAttachmentListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchStorageV1alpha1VolumeAttachmentList",
		Method:             "GET",
		PathPattern:        "/apis/storage.k8s.io/v1alpha1/watch/volumeattachments",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchStorageV1alpha1VolumeAttachmentListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchStorageV1alpha1VolumeAttachmentListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchStorageV1alpha1VolumeAttachmentList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
