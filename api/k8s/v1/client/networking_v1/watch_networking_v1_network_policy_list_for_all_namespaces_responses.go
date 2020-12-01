// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchNetworkingV1NetworkPolicyListForAllNamespacesReader is a Reader for the WatchNetworkingV1NetworkPolicyListForAllNamespaces structure.
type WatchNetworkingV1NetworkPolicyListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchNetworkingV1NetworkPolicyListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchNetworkingV1NetworkPolicyListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchNetworkingV1NetworkPolicyListForAllNamespacesOK creates a WatchNetworkingV1NetworkPolicyListForAllNamespacesOK with default headers values
func NewWatchNetworkingV1NetworkPolicyListForAllNamespacesOK() *WatchNetworkingV1NetworkPolicyListForAllNamespacesOK {
	return &WatchNetworkingV1NetworkPolicyListForAllNamespacesOK{}
}

/*WatchNetworkingV1NetworkPolicyListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchNetworkingV1NetworkPolicyListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchNetworkingV1NetworkPolicyListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/watch/networkpolicies][%d] watchNetworkingV1NetworkPolicyListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchNetworkingV1NetworkPolicyListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchNetworkingV1NetworkPolicyListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized creates a WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized with default headers values
func NewWatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized() *WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized {
	return &WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized{}
}

/*WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized struct {
}

func (o *WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/watch/networkpolicies][%d] watchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchNetworkingV1NetworkPolicyListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
