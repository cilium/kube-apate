// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusReader is a Reader for the ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus structure.
type ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK creates a ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK with default headers values
func NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK() *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK {
	return &ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK{}
}

/*ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK handles this case with default header values.

OK
*/
type ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK struct {
	Payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema
}

func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK) Error() string {
	return fmt.Sprintf("[GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status][%d] readFlowcontrolApiserverV1alpha1FlowSchemaStatusOK  %+v", 200, o.Payload)
}

func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK) GetPayload() *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema {
	return o.Payload
}

func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIFlowcontrolV1alpha1FlowSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized creates a ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized with default headers values
func NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized() *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized {
	return &ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized{}
}

/*ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized struct {
}

func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status][%d] readFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized ", 401)
}

func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
