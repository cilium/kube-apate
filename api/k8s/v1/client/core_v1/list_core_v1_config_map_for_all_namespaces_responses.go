// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListCoreV1ConfigMapForAllNamespacesReader is a Reader for the ListCoreV1ConfigMapForAllNamespaces structure.
type ListCoreV1ConfigMapForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCoreV1ConfigMapForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCoreV1ConfigMapForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCoreV1ConfigMapForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCoreV1ConfigMapForAllNamespacesOK creates a ListCoreV1ConfigMapForAllNamespacesOK with default headers values
func NewListCoreV1ConfigMapForAllNamespacesOK() *ListCoreV1ConfigMapForAllNamespacesOK {
	return &ListCoreV1ConfigMapForAllNamespacesOK{}
}

/*ListCoreV1ConfigMapForAllNamespacesOK handles this case with default header values.

OK
*/
type ListCoreV1ConfigMapForAllNamespacesOK struct {
	Payload *models.IoK8sAPICoreV1ConfigMapList
}

func (o *ListCoreV1ConfigMapForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/configmaps][%d] listCoreV1ConfigMapForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListCoreV1ConfigMapForAllNamespacesOK) GetPayload() *models.IoK8sAPICoreV1ConfigMapList {
	return o.Payload
}

func (o *ListCoreV1ConfigMapForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ConfigMapList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCoreV1ConfigMapForAllNamespacesUnauthorized creates a ListCoreV1ConfigMapForAllNamespacesUnauthorized with default headers values
func NewListCoreV1ConfigMapForAllNamespacesUnauthorized() *ListCoreV1ConfigMapForAllNamespacesUnauthorized {
	return &ListCoreV1ConfigMapForAllNamespacesUnauthorized{}
}

/*ListCoreV1ConfigMapForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCoreV1ConfigMapForAllNamespacesUnauthorized struct {
}

func (o *ListCoreV1ConfigMapForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/configmaps][%d] listCoreV1ConfigMapForAllNamespacesUnauthorized ", 401)
}

func (o *ListCoreV1ConfigMapForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
