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

// ReadCoreV1NamespacedReplicationControllerReader is a Reader for the ReadCoreV1NamespacedReplicationController structure.
type ReadCoreV1NamespacedReplicationControllerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1NamespacedReplicationControllerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1NamespacedReplicationControllerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1NamespacedReplicationControllerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1NamespacedReplicationControllerOK creates a ReadCoreV1NamespacedReplicationControllerOK with default headers values
func NewReadCoreV1NamespacedReplicationControllerOK() *ReadCoreV1NamespacedReplicationControllerOK {
	return &ReadCoreV1NamespacedReplicationControllerOK{}
}

/*ReadCoreV1NamespacedReplicationControllerOK handles this case with default header values.

OK
*/
type ReadCoreV1NamespacedReplicationControllerOK struct {
	Payload *models.IoK8sAPICoreV1ReplicationController
}

func (o *ReadCoreV1NamespacedReplicationControllerOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}][%d] readCoreV1NamespacedReplicationControllerOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1NamespacedReplicationControllerOK) GetPayload() *models.IoK8sAPICoreV1ReplicationController {
	return o.Payload
}

func (o *ReadCoreV1NamespacedReplicationControllerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ReplicationController)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1NamespacedReplicationControllerUnauthorized creates a ReadCoreV1NamespacedReplicationControllerUnauthorized with default headers values
func NewReadCoreV1NamespacedReplicationControllerUnauthorized() *ReadCoreV1NamespacedReplicationControllerUnauthorized {
	return &ReadCoreV1NamespacedReplicationControllerUnauthorized{}
}

/*ReadCoreV1NamespacedReplicationControllerUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1NamespacedReplicationControllerUnauthorized struct {
}

func (o *ReadCoreV1NamespacedReplicationControllerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}][%d] readCoreV1NamespacedReplicationControllerUnauthorized ", 401)
}

func (o *ReadCoreV1NamespacedReplicationControllerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}