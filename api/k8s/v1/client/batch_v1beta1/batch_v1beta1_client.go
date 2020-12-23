// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new batch v1beta1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for batch v1beta1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBatchV1beta1NamespacedCronJob(params *CreateBatchV1beta1NamespacedCronJobParams) (*CreateBatchV1beta1NamespacedCronJobOK, *CreateBatchV1beta1NamespacedCronJobCreated, *CreateBatchV1beta1NamespacedCronJobAccepted, error)

	DeleteBatchV1beta1CollectionNamespacedCronJob(params *DeleteBatchV1beta1CollectionNamespacedCronJobParams) (*DeleteBatchV1beta1CollectionNamespacedCronJobOK, error)

	DeleteBatchV1beta1NamespacedCronJob(params *DeleteBatchV1beta1NamespacedCronJobParams) (*DeleteBatchV1beta1NamespacedCronJobOK, *DeleteBatchV1beta1NamespacedCronJobAccepted, error)

	GetBatchV1beta1APIResources(params *GetBatchV1beta1APIResourcesParams) (*GetBatchV1beta1APIResourcesOK, error)

	ListBatchV1beta1CronJobForAllNamespaces(params *ListBatchV1beta1CronJobForAllNamespacesParams) (*ListBatchV1beta1CronJobForAllNamespacesOK, error)

	ListBatchV1beta1NamespacedCronJob(params *ListBatchV1beta1NamespacedCronJobParams) (*ListBatchV1beta1NamespacedCronJobOK, error)

	PatchBatchV1beta1NamespacedCronJob(params *PatchBatchV1beta1NamespacedCronJobParams) (*PatchBatchV1beta1NamespacedCronJobOK, error)

	PatchBatchV1beta1NamespacedCronJobStatus(params *PatchBatchV1beta1NamespacedCronJobStatusParams) (*PatchBatchV1beta1NamespacedCronJobStatusOK, error)

	ReadBatchV1beta1NamespacedCronJob(params *ReadBatchV1beta1NamespacedCronJobParams) (*ReadBatchV1beta1NamespacedCronJobOK, error)

	ReadBatchV1beta1NamespacedCronJobStatus(params *ReadBatchV1beta1NamespacedCronJobStatusParams) (*ReadBatchV1beta1NamespacedCronJobStatusOK, error)

	ReplaceBatchV1beta1NamespacedCronJob(params *ReplaceBatchV1beta1NamespacedCronJobParams) (*ReplaceBatchV1beta1NamespacedCronJobOK, *ReplaceBatchV1beta1NamespacedCronJobCreated, error)

	ReplaceBatchV1beta1NamespacedCronJobStatus(params *ReplaceBatchV1beta1NamespacedCronJobStatusParams) (*ReplaceBatchV1beta1NamespacedCronJobStatusOK, *ReplaceBatchV1beta1NamespacedCronJobStatusCreated, error)

	WatchBatchV1beta1CronJobListForAllNamespaces(params *WatchBatchV1beta1CronJobListForAllNamespacesParams) (*WatchBatchV1beta1CronJobListForAllNamespacesOK, error)

	WatchBatchV1beta1NamespacedCronJob(params *WatchBatchV1beta1NamespacedCronJobParams) (*WatchBatchV1beta1NamespacedCronJobOK, error)

	WatchBatchV1beta1NamespacedCronJobList(params *WatchBatchV1beta1NamespacedCronJobListParams) (*WatchBatchV1beta1NamespacedCronJobListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBatchV1beta1NamespacedCronJob create a CronJob
*/
func (a *Client) CreateBatchV1beta1NamespacedCronJob(params *CreateBatchV1beta1NamespacedCronJobParams) (*CreateBatchV1beta1NamespacedCronJobOK, *CreateBatchV1beta1NamespacedCronJobCreated, *CreateBatchV1beta1NamespacedCronJobAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBatchV1beta1NamespacedCronJob",
		Method:             "POST",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBatchV1beta1NamespacedCronJobOK:
		return value, nil, nil, nil
	case *CreateBatchV1beta1NamespacedCronJobCreated:
		return nil, value, nil, nil
	case *CreateBatchV1beta1NamespacedCronJobAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBatchV1beta1CollectionNamespacedCronJob delete collection of CronJob
*/
func (a *Client) DeleteBatchV1beta1CollectionNamespacedCronJob(params *DeleteBatchV1beta1CollectionNamespacedCronJobParams) (*DeleteBatchV1beta1CollectionNamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBatchV1beta1CollectionNamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBatchV1beta1CollectionNamespacedCronJob",
		Method:             "DELETE",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBatchV1beta1CollectionNamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBatchV1beta1CollectionNamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBatchV1beta1CollectionNamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBatchV1beta1NamespacedCronJob delete a CronJob
*/
func (a *Client) DeleteBatchV1beta1NamespacedCronJob(params *DeleteBatchV1beta1NamespacedCronJobParams) (*DeleteBatchV1beta1NamespacedCronJobOK, *DeleteBatchV1beta1NamespacedCronJobAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBatchV1beta1NamespacedCronJob",
		Method:             "DELETE",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBatchV1beta1NamespacedCronJobOK:
		return value, nil, nil
	case *DeleteBatchV1beta1NamespacedCronJobAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBatchV1beta1APIResources get available resources
*/
func (a *Client) GetBatchV1beta1APIResources(params *GetBatchV1beta1APIResourcesParams) (*GetBatchV1beta1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchV1beta1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBatchV1beta1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBatchV1beta1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBatchV1beta1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatchV1beta1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBatchV1beta1CronJobForAllNamespaces list or watch objects of kind CronJob
*/
func (a *Client) ListBatchV1beta1CronJobForAllNamespaces(params *ListBatchV1beta1CronJobForAllNamespacesParams) (*ListBatchV1beta1CronJobForAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBatchV1beta1CronJobForAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBatchV1beta1CronJobForAllNamespaces",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBatchV1beta1CronJobForAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBatchV1beta1CronJobForAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBatchV1beta1CronJobForAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBatchV1beta1NamespacedCronJob list or watch objects of kind CronJob
*/
func (a *Client) ListBatchV1beta1NamespacedCronJob(params *ListBatchV1beta1NamespacedCronJobParams) (*ListBatchV1beta1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBatchV1beta1NamespacedCronJob",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBatchV1beta1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBatchV1beta1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchBatchV1beta1NamespacedCronJob partially update the specified CronJob
*/
func (a *Client) PatchBatchV1beta1NamespacedCronJob(params *PatchBatchV1beta1NamespacedCronJobParams) (*PatchBatchV1beta1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchBatchV1beta1NamespacedCronJob",
		Method:             "PATCH",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchBatchV1beta1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchBatchV1beta1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchBatchV1beta1NamespacedCronJobStatus partially update status of the specified CronJob
*/
func (a *Client) PatchBatchV1beta1NamespacedCronJobStatus(params *PatchBatchV1beta1NamespacedCronJobStatusParams) (*PatchBatchV1beta1NamespacedCronJobStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchBatchV1beta1NamespacedCronJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchBatchV1beta1NamespacedCronJobStatus",
		Method:             "PATCH",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchBatchV1beta1NamespacedCronJobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchBatchV1beta1NamespacedCronJobStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchBatchV1beta1NamespacedCronJobStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBatchV1beta1NamespacedCronJob read the specified CronJob
*/
func (a *Client) ReadBatchV1beta1NamespacedCronJob(params *ReadBatchV1beta1NamespacedCronJobParams) (*ReadBatchV1beta1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readBatchV1beta1NamespacedCronJob",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBatchV1beta1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readBatchV1beta1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBatchV1beta1NamespacedCronJobStatus read status of the specified CronJob
*/
func (a *Client) ReadBatchV1beta1NamespacedCronJobStatus(params *ReadBatchV1beta1NamespacedCronJobStatusParams) (*ReadBatchV1beta1NamespacedCronJobStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBatchV1beta1NamespacedCronJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readBatchV1beta1NamespacedCronJobStatus",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadBatchV1beta1NamespacedCronJobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBatchV1beta1NamespacedCronJobStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readBatchV1beta1NamespacedCronJobStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceBatchV1beta1NamespacedCronJob replace the specified CronJob
*/
func (a *Client) ReplaceBatchV1beta1NamespacedCronJob(params *ReplaceBatchV1beta1NamespacedCronJobParams) (*ReplaceBatchV1beta1NamespacedCronJobOK, *ReplaceBatchV1beta1NamespacedCronJobCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBatchV1beta1NamespacedCronJob",
		Method:             "PUT",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBatchV1beta1NamespacedCronJobOK:
		return value, nil, nil
	case *ReplaceBatchV1beta1NamespacedCronJobCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceBatchV1beta1NamespacedCronJobStatus replace status of the specified CronJob
*/
func (a *Client) ReplaceBatchV1beta1NamespacedCronJobStatus(params *ReplaceBatchV1beta1NamespacedCronJobStatusParams) (*ReplaceBatchV1beta1NamespacedCronJobStatusOK, *ReplaceBatchV1beta1NamespacedCronJobStatusCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBatchV1beta1NamespacedCronJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBatchV1beta1NamespacedCronJobStatus",
		Method:             "PUT",
		PathPattern:        "/apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBatchV1beta1NamespacedCronJobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBatchV1beta1NamespacedCronJobStatusOK:
		return value, nil, nil
	case *ReplaceBatchV1beta1NamespacedCronJobStatusCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v1beta1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchBatchV1beta1CronJobListForAllNamespaces watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchBatchV1beta1CronJobListForAllNamespaces(params *WatchBatchV1beta1CronJobListForAllNamespacesParams) (*WatchBatchV1beta1CronJobListForAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchBatchV1beta1CronJobListForAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchBatchV1beta1CronJobListForAllNamespaces",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/watch/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchBatchV1beta1CronJobListForAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchBatchV1beta1CronJobListForAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchBatchV1beta1CronJobListForAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchBatchV1beta1NamespacedCronJob watch changes to an object of kind CronJob. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.
*/
func (a *Client) WatchBatchV1beta1NamespacedCronJob(params *WatchBatchV1beta1NamespacedCronJobParams) (*WatchBatchV1beta1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchBatchV1beta1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchBatchV1beta1NamespacedCronJob",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/watch/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchBatchV1beta1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchBatchV1beta1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchBatchV1beta1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchBatchV1beta1NamespacedCronJobList watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchBatchV1beta1NamespacedCronJobList(params *WatchBatchV1beta1NamespacedCronJobListParams) (*WatchBatchV1beta1NamespacedCronJobListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchBatchV1beta1NamespacedCronJobListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchBatchV1beta1NamespacedCronJobList",
		Method:             "GET",
		PathPattern:        "/apis/batch/v1beta1/watch/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchBatchV1beta1NamespacedCronJobListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchBatchV1beta1NamespacedCronJobListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchBatchV1beta1NamespacedCronJobList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}