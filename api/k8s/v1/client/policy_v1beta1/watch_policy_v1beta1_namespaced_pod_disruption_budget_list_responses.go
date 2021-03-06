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

// WatchPolicyV1beta1NamespacedPodDisruptionBudgetListReader is a Reader for the WatchPolicyV1beta1NamespacedPodDisruptionBudgetList structure.
type WatchPolicyV1beta1NamespacedPodDisruptionBudgetListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK creates a WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK with default headers values
func NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK() *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK {
	return &WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK{}
}

/*WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK handles this case with default header values.

OK
*/
type WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK) Error() string {
	return fmt.Sprintf("[GET /apis/policy/v1beta1/watch/namespaces/{namespace}/poddisruptionbudgets][%d] watchPolicyV1beta1NamespacedPodDisruptionBudgetListOK  %+v", 200, o.Payload)
}

func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized creates a WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized with default headers values
func NewWatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized() *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized {
	return &WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized{}
}

/*WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized struct {
}

func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/policy/v1beta1/watch/namespaces/{namespace}/poddisruptionbudgets][%d] watchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized ", 401)
}

func (o *WatchPolicyV1beta1NamespacedPodDisruptionBudgetListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
