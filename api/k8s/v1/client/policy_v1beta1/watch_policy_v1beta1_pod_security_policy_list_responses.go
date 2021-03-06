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

// WatchPolicyV1beta1PodSecurityPolicyListReader is a Reader for the WatchPolicyV1beta1PodSecurityPolicyList structure.
type WatchPolicyV1beta1PodSecurityPolicyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchPolicyV1beta1PodSecurityPolicyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchPolicyV1beta1PodSecurityPolicyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchPolicyV1beta1PodSecurityPolicyListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchPolicyV1beta1PodSecurityPolicyListOK creates a WatchPolicyV1beta1PodSecurityPolicyListOK with default headers values
func NewWatchPolicyV1beta1PodSecurityPolicyListOK() *WatchPolicyV1beta1PodSecurityPolicyListOK {
	return &WatchPolicyV1beta1PodSecurityPolicyListOK{}
}

/*WatchPolicyV1beta1PodSecurityPolicyListOK handles this case with default header values.

OK
*/
type WatchPolicyV1beta1PodSecurityPolicyListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchPolicyV1beta1PodSecurityPolicyListOK) Error() string {
	return fmt.Sprintf("[GET /apis/policy/v1beta1/watch/podsecuritypolicies][%d] watchPolicyV1beta1PodSecurityPolicyListOK  %+v", 200, o.Payload)
}

func (o *WatchPolicyV1beta1PodSecurityPolicyListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchPolicyV1beta1PodSecurityPolicyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchPolicyV1beta1PodSecurityPolicyListUnauthorized creates a WatchPolicyV1beta1PodSecurityPolicyListUnauthorized with default headers values
func NewWatchPolicyV1beta1PodSecurityPolicyListUnauthorized() *WatchPolicyV1beta1PodSecurityPolicyListUnauthorized {
	return &WatchPolicyV1beta1PodSecurityPolicyListUnauthorized{}
}

/*WatchPolicyV1beta1PodSecurityPolicyListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchPolicyV1beta1PodSecurityPolicyListUnauthorized struct {
}

func (o *WatchPolicyV1beta1PodSecurityPolicyListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/policy/v1beta1/watch/podsecuritypolicies][%d] watchPolicyV1beta1PodSecurityPolicyListUnauthorized ", 401)
}

func (o *WatchPolicyV1beta1PodSecurityPolicyListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
