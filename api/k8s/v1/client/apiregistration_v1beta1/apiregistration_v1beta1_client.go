// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new apiregistration v1beta1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for apiregistration v1beta1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateApiregistrationV1beta1APIService(params *CreateApiregistrationV1beta1APIServiceParams) (*CreateApiregistrationV1beta1APIServiceOK, *CreateApiregistrationV1beta1APIServiceCreated, *CreateApiregistrationV1beta1APIServiceAccepted, error)

	DeleteApiregistrationV1beta1APIService(params *DeleteApiregistrationV1beta1APIServiceParams) (*DeleteApiregistrationV1beta1APIServiceOK, *DeleteApiregistrationV1beta1APIServiceAccepted, error)

	DeleteApiregistrationV1beta1CollectionAPIService(params *DeleteApiregistrationV1beta1CollectionAPIServiceParams) (*DeleteApiregistrationV1beta1CollectionAPIServiceOK, error)

	GetApiregistrationV1beta1APIResources(params *GetApiregistrationV1beta1APIResourcesParams) (*GetApiregistrationV1beta1APIResourcesOK, error)

	ListApiregistrationV1beta1APIService(params *ListApiregistrationV1beta1APIServiceParams) (*ListApiregistrationV1beta1APIServiceOK, error)

	PatchApiregistrationV1beta1APIService(params *PatchApiregistrationV1beta1APIServiceParams) (*PatchApiregistrationV1beta1APIServiceOK, error)

	PatchApiregistrationV1beta1APIServiceStatus(params *PatchApiregistrationV1beta1APIServiceStatusParams) (*PatchApiregistrationV1beta1APIServiceStatusOK, error)

	ReadApiregistrationV1beta1APIService(params *ReadApiregistrationV1beta1APIServiceParams) (*ReadApiregistrationV1beta1APIServiceOK, error)

	ReadApiregistrationV1beta1APIServiceStatus(params *ReadApiregistrationV1beta1APIServiceStatusParams) (*ReadApiregistrationV1beta1APIServiceStatusOK, error)

	ReplaceApiregistrationV1beta1APIService(params *ReplaceApiregistrationV1beta1APIServiceParams) (*ReplaceApiregistrationV1beta1APIServiceOK, *ReplaceApiregistrationV1beta1APIServiceCreated, error)

	ReplaceApiregistrationV1beta1APIServiceStatus(params *ReplaceApiregistrationV1beta1APIServiceStatusParams) (*ReplaceApiregistrationV1beta1APIServiceStatusOK, *ReplaceApiregistrationV1beta1APIServiceStatusCreated, error)

	WatchApiregistrationV1beta1APIService(params *WatchApiregistrationV1beta1APIServiceParams) (*WatchApiregistrationV1beta1APIServiceOK, error)

	WatchApiregistrationV1beta1APIServiceList(params *WatchApiregistrationV1beta1APIServiceListParams) (*WatchApiregistrationV1beta1APIServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateApiregistrationV1beta1APIService create an APIService
*/
func (a *Client) CreateApiregistrationV1beta1APIService(params *CreateApiregistrationV1beta1APIServiceParams) (*CreateApiregistrationV1beta1APIServiceOK, *CreateApiregistrationV1beta1APIServiceCreated, *CreateApiregistrationV1beta1APIServiceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createApiregistrationV1beta1APIService",
		Method:             "POST",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateApiregistrationV1beta1APIServiceOK:
		return value, nil, nil, nil
	case *CreateApiregistrationV1beta1APIServiceCreated:
		return nil, value, nil, nil
	case *CreateApiregistrationV1beta1APIServiceAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for apiregistration_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteApiregistrationV1beta1APIService delete an APIService
*/
func (a *Client) DeleteApiregistrationV1beta1APIService(params *DeleteApiregistrationV1beta1APIServiceParams) (*DeleteApiregistrationV1beta1APIServiceOK, *DeleteApiregistrationV1beta1APIServiceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteApiregistrationV1beta1APIService",
		Method:             "DELETE",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteApiregistrationV1beta1APIServiceOK:
		return value, nil, nil
	case *DeleteApiregistrationV1beta1APIServiceAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for apiregistration_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteApiregistrationV1beta1CollectionAPIService delete collection of APIService
*/
func (a *Client) DeleteApiregistrationV1beta1CollectionAPIService(params *DeleteApiregistrationV1beta1CollectionAPIServiceParams) (*DeleteApiregistrationV1beta1CollectionAPIServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApiregistrationV1beta1CollectionAPIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteApiregistrationV1beta1CollectionAPIService",
		Method:             "DELETE",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteApiregistrationV1beta1CollectionAPIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteApiregistrationV1beta1CollectionAPIServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteApiregistrationV1beta1CollectionAPIService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApiregistrationV1beta1APIResources get available resources
*/
func (a *Client) GetApiregistrationV1beta1APIResources(params *GetApiregistrationV1beta1APIResourcesParams) (*GetApiregistrationV1beta1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApiregistrationV1beta1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApiregistrationV1beta1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetApiregistrationV1beta1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApiregistrationV1beta1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApiregistrationV1beta1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListApiregistrationV1beta1APIService list or watch objects of kind APIService
*/
func (a *Client) ListApiregistrationV1beta1APIService(params *ListApiregistrationV1beta1APIServiceParams) (*ListApiregistrationV1beta1APIServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listApiregistrationV1beta1APIService",
		Method:             "GET",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListApiregistrationV1beta1APIServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listApiregistrationV1beta1APIService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchApiregistrationV1beta1APIService partially update the specified APIService
*/
func (a *Client) PatchApiregistrationV1beta1APIService(params *PatchApiregistrationV1beta1APIServiceParams) (*PatchApiregistrationV1beta1APIServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchApiregistrationV1beta1APIService",
		Method:             "PATCH",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchApiregistrationV1beta1APIServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchApiregistrationV1beta1APIService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchApiregistrationV1beta1APIServiceStatus partially update status of the specified APIService
*/
func (a *Client) PatchApiregistrationV1beta1APIServiceStatus(params *PatchApiregistrationV1beta1APIServiceStatusParams) (*PatchApiregistrationV1beta1APIServiceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchApiregistrationV1beta1APIServiceStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchApiregistrationV1beta1APIServiceStatus",
		Method:             "PATCH",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchApiregistrationV1beta1APIServiceStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchApiregistrationV1beta1APIServiceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchApiregistrationV1beta1APIServiceStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadApiregistrationV1beta1APIService read the specified APIService
*/
func (a *Client) ReadApiregistrationV1beta1APIService(params *ReadApiregistrationV1beta1APIServiceParams) (*ReadApiregistrationV1beta1APIServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readApiregistrationV1beta1APIService",
		Method:             "GET",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadApiregistrationV1beta1APIServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readApiregistrationV1beta1APIService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadApiregistrationV1beta1APIServiceStatus read status of the specified APIService
*/
func (a *Client) ReadApiregistrationV1beta1APIServiceStatus(params *ReadApiregistrationV1beta1APIServiceStatusParams) (*ReadApiregistrationV1beta1APIServiceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadApiregistrationV1beta1APIServiceStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readApiregistrationV1beta1APIServiceStatus",
		Method:             "GET",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadApiregistrationV1beta1APIServiceStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadApiregistrationV1beta1APIServiceStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readApiregistrationV1beta1APIServiceStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceApiregistrationV1beta1APIService replace the specified APIService
*/
func (a *Client) ReplaceApiregistrationV1beta1APIService(params *ReplaceApiregistrationV1beta1APIServiceParams) (*ReplaceApiregistrationV1beta1APIServiceOK, *ReplaceApiregistrationV1beta1APIServiceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceApiregistrationV1beta1APIService",
		Method:             "PUT",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceApiregistrationV1beta1APIServiceOK:
		return value, nil, nil
	case *ReplaceApiregistrationV1beta1APIServiceCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for apiregistration_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceApiregistrationV1beta1APIServiceStatus replace status of the specified APIService
*/
func (a *Client) ReplaceApiregistrationV1beta1APIServiceStatus(params *ReplaceApiregistrationV1beta1APIServiceStatusParams) (*ReplaceApiregistrationV1beta1APIServiceStatusOK, *ReplaceApiregistrationV1beta1APIServiceStatusCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceApiregistrationV1beta1APIServiceStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceApiregistrationV1beta1APIServiceStatus",
		Method:             "PUT",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/apiservices/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceApiregistrationV1beta1APIServiceStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceApiregistrationV1beta1APIServiceStatusOK:
		return value, nil, nil
	case *ReplaceApiregistrationV1beta1APIServiceStatusCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for apiregistration_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchApiregistrationV1beta1APIService watch changes to an object of kind APIService. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.
*/
func (a *Client) WatchApiregistrationV1beta1APIService(params *WatchApiregistrationV1beta1APIServiceParams) (*WatchApiregistrationV1beta1APIServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchApiregistrationV1beta1APIServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchApiregistrationV1beta1APIService",
		Method:             "GET",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/watch/apiservices/{name}",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchApiregistrationV1beta1APIServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchApiregistrationV1beta1APIServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchApiregistrationV1beta1APIService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchApiregistrationV1beta1APIServiceList watch individual changes to a list of APIService. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchApiregistrationV1beta1APIServiceList(params *WatchApiregistrationV1beta1APIServiceListParams) (*WatchApiregistrationV1beta1APIServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchApiregistrationV1beta1APIServiceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchApiregistrationV1beta1APIServiceList",
		Method:             "GET",
		PathPattern:        "/apis/apiregistration.k8s.io/v1beta1/watch/apiservices",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchApiregistrationV1beta1APIServiceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchApiregistrationV1beta1APIServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchApiregistrationV1beta1APIServiceList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
