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

// WatchFlowcontrolApiserverV1alpha1FlowSchemaReader is a Reader for the WatchFlowcontrolApiserverV1alpha1FlowSchema structure.
type WatchFlowcontrolApiserverV1alpha1FlowSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchFlowcontrolApiserverV1alpha1FlowSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchFlowcontrolApiserverV1alpha1FlowSchemaOK creates a WatchFlowcontrolApiserverV1alpha1FlowSchemaOK with default headers values
func NewWatchFlowcontrolApiserverV1alpha1FlowSchemaOK() *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK {
	return &WatchFlowcontrolApiserverV1alpha1FlowSchemaOK{}
}

/*WatchFlowcontrolApiserverV1alpha1FlowSchemaOK handles this case with default header values.

OK
*/
type WatchFlowcontrolApiserverV1alpha1FlowSchemaOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK) Error() string {
	return fmt.Sprintf("[GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/watch/flowschemas/{name}][%d] watchFlowcontrolApiserverV1alpha1FlowSchemaOK  %+v", 200, o.Payload)
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized creates a WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized with default headers values
func NewWatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized() *WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized {
	return &WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized{}
}

/*WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized struct {
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/watch/flowschemas/{name}][%d] watchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized ", 401)
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}