// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesReader is a Reader for the WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespaces structure.
type WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK creates a WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK with default headers values
func NewWatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK() *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK {
	return &WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK{}
}

/*WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/policy/v1beta1/watch/poddisruptionbudgets][%d] watchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized creates a WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized with default headers values
func NewWatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized() *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized {
	return &WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized{}
}

/*WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized struct {
}

func (o *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/policy/v1beta1/watch/poddisruptionbudgets][%d] watchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
