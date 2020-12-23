// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetNodeV1alpha1APIResourcesReader is a Reader for the GetNodeV1alpha1APIResources structure.
type GetNodeV1alpha1APIResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeV1alpha1APIResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeV1alpha1APIResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNodeV1alpha1APIResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNodeV1alpha1APIResourcesOK creates a GetNodeV1alpha1APIResourcesOK with default headers values
func NewGetNodeV1alpha1APIResourcesOK() *GetNodeV1alpha1APIResourcesOK {
	return &GetNodeV1alpha1APIResourcesOK{}
}

/*GetNodeV1alpha1APIResourcesOK handles this case with default header values.

OK
*/
type GetNodeV1alpha1APIResourcesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList
}

func (o *GetNodeV1alpha1APIResourcesOK) Error() string {
	return fmt.Sprintf("[GET /apis/node.k8s.io/v1alpha1/][%d] getNodeV1alpha1ApiResourcesOK  %+v", 200, o.Payload)
}

func (o *GetNodeV1alpha1APIResourcesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIResourceList {
	return o.Payload
}

func (o *GetNodeV1alpha1APIResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIResourceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeV1alpha1APIResourcesUnauthorized creates a GetNodeV1alpha1APIResourcesUnauthorized with default headers values
func NewGetNodeV1alpha1APIResourcesUnauthorized() *GetNodeV1alpha1APIResourcesUnauthorized {
	return &GetNodeV1alpha1APIResourcesUnauthorized{}
}

/*GetNodeV1alpha1APIResourcesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNodeV1alpha1APIResourcesUnauthorized struct {
}

func (o *GetNodeV1alpha1APIResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/node.k8s.io/v1alpha1/][%d] getNodeV1alpha1ApiResourcesUnauthorized ", 401)
}

func (o *GetNodeV1alpha1APIResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}