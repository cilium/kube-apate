// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetEventsAPIGroupReader is a Reader for the GetEventsAPIGroup structure.
type GetEventsAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEventsAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEventsAPIGroupOK creates a GetEventsAPIGroupOK with default headers values
func NewGetEventsAPIGroupOK() *GetEventsAPIGroupOK {
	return &GetEventsAPIGroupOK{}
}

/*GetEventsAPIGroupOK handles this case with default header values.

OK
*/
type GetEventsAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetEventsAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/events.k8s.io/][%d] getEventsApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetEventsAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetEventsAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsAPIGroupUnauthorized creates a GetEventsAPIGroupUnauthorized with default headers values
func NewGetEventsAPIGroupUnauthorized() *GetEventsAPIGroupUnauthorized {
	return &GetEventsAPIGroupUnauthorized{}
}

/*GetEventsAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEventsAPIGroupUnauthorized struct {
}

func (o *GetEventsAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/events.k8s.io/][%d] getEventsApiGroupUnauthorized ", 401)
}

func (o *GetEventsAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
