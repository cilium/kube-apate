// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new batch v2alpha1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for batch v2alpha1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBatchV2alpha1NamespacedCronJob(params *CreateBatchV2alpha1NamespacedCronJobParams) (*CreateBatchV2alpha1NamespacedCronJobOK, *CreateBatchV2alpha1NamespacedCronJobCreated, *CreateBatchV2alpha1NamespacedCronJobAccepted, error)

	DeleteBatchV2alpha1CollectionNamespacedCronJob(params *DeleteBatchV2alpha1CollectionNamespacedCronJobParams) (*DeleteBatchV2alpha1CollectionNamespacedCronJobOK, error)

	DeleteBatchV2alpha1NamespacedCronJob(params *DeleteBatchV2alpha1NamespacedCronJobParams) (*DeleteBatchV2alpha1NamespacedCronJobOK, *DeleteBatchV2alpha1NamespacedCronJobAccepted, error)

	GetBatchV2alpha1APIResources(params *GetBatchV2alpha1APIResourcesParams) (*GetBatchV2alpha1APIResourcesOK, error)

	ListBatchV2alpha1CronJobForAllNamespaces(params *ListBatchV2alpha1CronJobForAllNamespacesParams) (*ListBatchV2alpha1CronJobForAllNamespacesOK, error)

	ListBatchV2alpha1NamespacedCronJob(params *ListBatchV2alpha1NamespacedCronJobParams) (*ListBatchV2alpha1NamespacedCronJobOK, error)

	PatchBatchV2alpha1NamespacedCronJob(params *PatchBatchV2alpha1NamespacedCronJobParams) (*PatchBatchV2alpha1NamespacedCronJobOK, error)

	PatchBatchV2alpha1NamespacedCronJobStatus(params *PatchBatchV2alpha1NamespacedCronJobStatusParams) (*PatchBatchV2alpha1NamespacedCronJobStatusOK, error)

	ReadBatchV2alpha1NamespacedCronJob(params *ReadBatchV2alpha1NamespacedCronJobParams) (*ReadBatchV2alpha1NamespacedCronJobOK, error)

	ReadBatchV2alpha1NamespacedCronJobStatus(params *ReadBatchV2alpha1NamespacedCronJobStatusParams) (*ReadBatchV2alpha1NamespacedCronJobStatusOK, error)

	ReplaceBatchV2alpha1NamespacedCronJob(params *ReplaceBatchV2alpha1NamespacedCronJobParams) (*ReplaceBatchV2alpha1NamespacedCronJobOK, *ReplaceBatchV2alpha1NamespacedCronJobCreated, error)

	ReplaceBatchV2alpha1NamespacedCronJobStatus(params *ReplaceBatchV2alpha1NamespacedCronJobStatusParams) (*ReplaceBatchV2alpha1NamespacedCronJobStatusOK, *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated, error)

	WatchBatchV2alpha1CronJobListForAllNamespaces(params *WatchBatchV2alpha1CronJobListForAllNamespacesParams) (*WatchBatchV2alpha1CronJobListForAllNamespacesOK, error)

	WatchBatchV2alpha1NamespacedCronJob(params *WatchBatchV2alpha1NamespacedCronJobParams) (*WatchBatchV2alpha1NamespacedCronJobOK, error)

	WatchBatchV2alpha1NamespacedCronJobList(params *WatchBatchV2alpha1NamespacedCronJobListParams) (*WatchBatchV2alpha1NamespacedCronJobListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBatchV2alpha1NamespacedCronJob create a CronJob
*/
func (a *Client) CreateBatchV2alpha1NamespacedCronJob(params *CreateBatchV2alpha1NamespacedCronJobParams) (*CreateBatchV2alpha1NamespacedCronJobOK, *CreateBatchV2alpha1NamespacedCronJobCreated, *CreateBatchV2alpha1NamespacedCronJobAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBatchV2alpha1NamespacedCronJob",
		Method:             "POST",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBatchV2alpha1NamespacedCronJobOK:
		return value, nil, nil, nil
	case *CreateBatchV2alpha1NamespacedCronJobCreated:
		return nil, value, nil, nil
	case *CreateBatchV2alpha1NamespacedCronJobAccepted:
		return nil, nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v2alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBatchV2alpha1CollectionNamespacedCronJob delete collection of CronJob
*/
func (a *Client) DeleteBatchV2alpha1CollectionNamespacedCronJob(params *DeleteBatchV2alpha1CollectionNamespacedCronJobParams) (*DeleteBatchV2alpha1CollectionNamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBatchV2alpha1CollectionNamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBatchV2alpha1CollectionNamespacedCronJob",
		Method:             "DELETE",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBatchV2alpha1CollectionNamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBatchV2alpha1CollectionNamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBatchV2alpha1CollectionNamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBatchV2alpha1NamespacedCronJob delete a CronJob
*/
func (a *Client) DeleteBatchV2alpha1NamespacedCronJob(params *DeleteBatchV2alpha1NamespacedCronJobParams) (*DeleteBatchV2alpha1NamespacedCronJobOK, *DeleteBatchV2alpha1NamespacedCronJobAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBatchV2alpha1NamespacedCronJob",
		Method:             "DELETE",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBatchV2alpha1NamespacedCronJobOK:
		return value, nil, nil
	case *DeleteBatchV2alpha1NamespacedCronJobAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v2alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBatchV2alpha1APIResources get available resources
*/
func (a *Client) GetBatchV2alpha1APIResources(params *GetBatchV2alpha1APIResourcesParams) (*GetBatchV2alpha1APIResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchV2alpha1APIResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBatchV2alpha1APIResources",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBatchV2alpha1APIResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBatchV2alpha1APIResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatchV2alpha1APIResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBatchV2alpha1CronJobForAllNamespaces list or watch objects of kind CronJob
*/
func (a *Client) ListBatchV2alpha1CronJobForAllNamespaces(params *ListBatchV2alpha1CronJobForAllNamespacesParams) (*ListBatchV2alpha1CronJobForAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBatchV2alpha1CronJobForAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBatchV2alpha1CronJobForAllNamespaces",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBatchV2alpha1CronJobForAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBatchV2alpha1CronJobForAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBatchV2alpha1CronJobForAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBatchV2alpha1NamespacedCronJob list or watch objects of kind CronJob
*/
func (a *Client) ListBatchV2alpha1NamespacedCronJob(params *ListBatchV2alpha1NamespacedCronJobParams) (*ListBatchV2alpha1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBatchV2alpha1NamespacedCronJob",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBatchV2alpha1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBatchV2alpha1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchBatchV2alpha1NamespacedCronJob partially update the specified CronJob
*/
func (a *Client) PatchBatchV2alpha1NamespacedCronJob(params *PatchBatchV2alpha1NamespacedCronJobParams) (*PatchBatchV2alpha1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchBatchV2alpha1NamespacedCronJob",
		Method:             "PATCH",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchBatchV2alpha1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchBatchV2alpha1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchBatchV2alpha1NamespacedCronJobStatus partially update status of the specified CronJob
*/
func (a *Client) PatchBatchV2alpha1NamespacedCronJobStatus(params *PatchBatchV2alpha1NamespacedCronJobStatusParams) (*PatchBatchV2alpha1NamespacedCronJobStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchBatchV2alpha1NamespacedCronJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchBatchV2alpha1NamespacedCronJobStatus",
		Method:             "PATCH",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"application/apply-patch+yaml", "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchBatchV2alpha1NamespacedCronJobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchBatchV2alpha1NamespacedCronJobStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchBatchV2alpha1NamespacedCronJobStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBatchV2alpha1NamespacedCronJob read the specified CronJob
*/
func (a *Client) ReadBatchV2alpha1NamespacedCronJob(params *ReadBatchV2alpha1NamespacedCronJobParams) (*ReadBatchV2alpha1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readBatchV2alpha1NamespacedCronJob",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBatchV2alpha1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readBatchV2alpha1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBatchV2alpha1NamespacedCronJobStatus read status of the specified CronJob
*/
func (a *Client) ReadBatchV2alpha1NamespacedCronJobStatus(params *ReadBatchV2alpha1NamespacedCronJobStatusParams) (*ReadBatchV2alpha1NamespacedCronJobStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBatchV2alpha1NamespacedCronJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readBatchV2alpha1NamespacedCronJobStatus",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadBatchV2alpha1NamespacedCronJobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBatchV2alpha1NamespacedCronJobStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readBatchV2alpha1NamespacedCronJobStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceBatchV2alpha1NamespacedCronJob replace the specified CronJob
*/
func (a *Client) ReplaceBatchV2alpha1NamespacedCronJob(params *ReplaceBatchV2alpha1NamespacedCronJobParams) (*ReplaceBatchV2alpha1NamespacedCronJobOK, *ReplaceBatchV2alpha1NamespacedCronJobCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBatchV2alpha1NamespacedCronJob",
		Method:             "PUT",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBatchV2alpha1NamespacedCronJobOK:
		return value, nil, nil
	case *ReplaceBatchV2alpha1NamespacedCronJobCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v2alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReplaceBatchV2alpha1NamespacedCronJobStatus replace status of the specified CronJob
*/
func (a *Client) ReplaceBatchV2alpha1NamespacedCronJobStatus(params *ReplaceBatchV2alpha1NamespacedCronJobStatusParams) (*ReplaceBatchV2alpha1NamespacedCronJobStatusOK, *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBatchV2alpha1NamespacedCronJobStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBatchV2alpha1NamespacedCronJobStatus",
		Method:             "PUT",
		PathPattern:        "/apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status",
		ProducesMediaTypes: []string{"application/json", "application/vnd.kubernetes.protobuf", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBatchV2alpha1NamespacedCronJobStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBatchV2alpha1NamespacedCronJobStatusOK:
		return value, nil, nil
	case *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for batch_v2alpha1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchBatchV2alpha1CronJobListForAllNamespaces watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchBatchV2alpha1CronJobListForAllNamespaces(params *WatchBatchV2alpha1CronJobListForAllNamespacesParams) (*WatchBatchV2alpha1CronJobListForAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchBatchV2alpha1CronJobListForAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchBatchV2alpha1CronJobListForAllNamespaces",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/watch/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchBatchV2alpha1CronJobListForAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchBatchV2alpha1CronJobListForAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchBatchV2alpha1CronJobListForAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchBatchV2alpha1NamespacedCronJob watch changes to an object of kind CronJob. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.
*/
func (a *Client) WatchBatchV2alpha1NamespacedCronJob(params *WatchBatchV2alpha1NamespacedCronJobParams) (*WatchBatchV2alpha1NamespacedCronJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchBatchV2alpha1NamespacedCronJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchBatchV2alpha1NamespacedCronJob",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/watch/namespaces/{namespace}/cronjobs/{name}",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchBatchV2alpha1NamespacedCronJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchBatchV2alpha1NamespacedCronJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchBatchV2alpha1NamespacedCronJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WatchBatchV2alpha1NamespacedCronJobList watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.
*/
func (a *Client) WatchBatchV2alpha1NamespacedCronJobList(params *WatchBatchV2alpha1NamespacedCronJobListParams) (*WatchBatchV2alpha1NamespacedCronJobListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWatchBatchV2alpha1NamespacedCronJobListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "watchBatchV2alpha1NamespacedCronJobList",
		Method:             "GET",
		PathPattern:        "/apis/batch/v2alpha1/watch/namespaces/{namespace}/cronjobs",
		ProducesMediaTypes: []string{"application/json", "application/json;stream=watch", "application/vnd.kubernetes.protobuf", "application/vnd.kubernetes.protobuf;stream=watch", "application/yaml"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WatchBatchV2alpha1NamespacedCronJobListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WatchBatchV2alpha1NamespacedCronJobListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for watchBatchV2alpha1NamespacedCronJobList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
