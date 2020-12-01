// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchNetworkingV1beta1NamespacedIngressStatusReader is a Reader for the PatchNetworkingV1beta1NamespacedIngressStatus structure.
type PatchNetworkingV1beta1NamespacedIngressStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchNetworkingV1beta1NamespacedIngressStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchNetworkingV1beta1NamespacedIngressStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchNetworkingV1beta1NamespacedIngressStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchNetworkingV1beta1NamespacedIngressStatusOK creates a PatchNetworkingV1beta1NamespacedIngressStatusOK with default headers values
func NewPatchNetworkingV1beta1NamespacedIngressStatusOK() *PatchNetworkingV1beta1NamespacedIngressStatusOK {
	return &PatchNetworkingV1beta1NamespacedIngressStatusOK{}
}

/*PatchNetworkingV1beta1NamespacedIngressStatusOK handles this case with default header values.

OK
*/
type PatchNetworkingV1beta1NamespacedIngressStatusOK struct {
	Payload *models.IoK8sAPINetworkingV1beta1Ingress
}

func (o *PatchNetworkingV1beta1NamespacedIngressStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name}/status][%d] patchNetworkingV1beta1NamespacedIngressStatusOK  %+v", 200, o.Payload)
}

func (o *PatchNetworkingV1beta1NamespacedIngressStatusOK) GetPayload() *models.IoK8sAPINetworkingV1beta1Ingress {
	return o.Payload
}

func (o *PatchNetworkingV1beta1NamespacedIngressStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1beta1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNetworkingV1beta1NamespacedIngressStatusUnauthorized creates a PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized with default headers values
func NewPatchNetworkingV1beta1NamespacedIngressStatusUnauthorized() *PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized {
	return &PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized{}
}

/*PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized struct {
}

func (o *PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name}/status][%d] patchNetworkingV1beta1NamespacedIngressStatusUnauthorized ", 401)
}

func (o *PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
