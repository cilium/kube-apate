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

// ListCoreV1ResourceQuotaForAllNamespacesReader is a Reader for the ListCoreV1ResourceQuotaForAllNamespaces structure.
type ListCoreV1ResourceQuotaForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCoreV1ResourceQuotaForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCoreV1ResourceQuotaForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCoreV1ResourceQuotaForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCoreV1ResourceQuotaForAllNamespacesOK creates a ListCoreV1ResourceQuotaForAllNamespacesOK with default headers values
func NewListCoreV1ResourceQuotaForAllNamespacesOK() *ListCoreV1ResourceQuotaForAllNamespacesOK {
	return &ListCoreV1ResourceQuotaForAllNamespacesOK{}
}

/*ListCoreV1ResourceQuotaForAllNamespacesOK handles this case with default header values.

OK
*/
type ListCoreV1ResourceQuotaForAllNamespacesOK struct {
	Payload *models.IoK8sAPICoreV1ResourceQuotaList
}

func (o *ListCoreV1ResourceQuotaForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/resourcequotas][%d] listCoreV1ResourceQuotaForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListCoreV1ResourceQuotaForAllNamespacesOK) GetPayload() *models.IoK8sAPICoreV1ResourceQuotaList {
	return o.Payload
}

func (o *ListCoreV1ResourceQuotaForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ResourceQuotaList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCoreV1ResourceQuotaForAllNamespacesUnauthorized creates a ListCoreV1ResourceQuotaForAllNamespacesUnauthorized with default headers values
func NewListCoreV1ResourceQuotaForAllNamespacesUnauthorized() *ListCoreV1ResourceQuotaForAllNamespacesUnauthorized {
	return &ListCoreV1ResourceQuotaForAllNamespacesUnauthorized{}
}

/*ListCoreV1ResourceQuotaForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCoreV1ResourceQuotaForAllNamespacesUnauthorized struct {
}

func (o *ListCoreV1ResourceQuotaForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/resourcequotas][%d] listCoreV1ResourceQuotaForAllNamespacesUnauthorized ", 401)
}

func (o *ListCoreV1ResourceQuotaForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}