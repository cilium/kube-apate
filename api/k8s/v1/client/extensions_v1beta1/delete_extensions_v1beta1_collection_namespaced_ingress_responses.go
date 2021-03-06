// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteExtensionsV1beta1CollectionNamespacedIngressReader is a Reader for the DeleteExtensionsV1beta1CollectionNamespacedIngress structure.
type DeleteExtensionsV1beta1CollectionNamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExtensionsV1beta1CollectionNamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExtensionsV1beta1CollectionNamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExtensionsV1beta1CollectionNamespacedIngressOK creates a DeleteExtensionsV1beta1CollectionNamespacedIngressOK with default headers values
func NewDeleteExtensionsV1beta1CollectionNamespacedIngressOK() *DeleteExtensionsV1beta1CollectionNamespacedIngressOK {
	return &DeleteExtensionsV1beta1CollectionNamespacedIngressOK{}
}

/*DeleteExtensionsV1beta1CollectionNamespacedIngressOK handles this case with default header values.

OK
*/
type DeleteExtensionsV1beta1CollectionNamespacedIngressOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteExtensionsV1beta1CollectionNamespacedIngressOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/extensions/v1beta1/namespaces/{namespace}/ingresses][%d] deleteExtensionsV1beta1CollectionNamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *DeleteExtensionsV1beta1CollectionNamespacedIngressOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteExtensionsV1beta1CollectionNamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized creates a DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized with default headers values
func NewDeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized() *DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized {
	return &DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized{}
}

/*DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized struct {
}

func (o *DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/extensions/v1beta1/namespaces/{namespace}/ingresses][%d] deleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized ", 401)
}

func (o *DeleteExtensionsV1beta1CollectionNamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
