// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteApiregistrationV1CollectionAPIServiceReader is a Reader for the DeleteApiregistrationV1CollectionAPIService structure.
type DeleteApiregistrationV1CollectionAPIServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApiregistrationV1CollectionAPIServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApiregistrationV1CollectionAPIServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteApiregistrationV1CollectionAPIServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteApiregistrationV1CollectionAPIServiceOK creates a DeleteApiregistrationV1CollectionAPIServiceOK with default headers values
func NewDeleteApiregistrationV1CollectionAPIServiceOK() *DeleteApiregistrationV1CollectionAPIServiceOK {
	return &DeleteApiregistrationV1CollectionAPIServiceOK{}
}

/*DeleteApiregistrationV1CollectionAPIServiceOK handles this case with default header values.

OK
*/
type DeleteApiregistrationV1CollectionAPIServiceOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteApiregistrationV1CollectionAPIServiceOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/apiregistration.k8s.io/v1/apiservices][%d] deleteApiregistrationV1CollectionApiServiceOK  %+v", 200, o.Payload)
}

func (o *DeleteApiregistrationV1CollectionAPIServiceOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteApiregistrationV1CollectionAPIServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteApiregistrationV1CollectionAPIServiceUnauthorized creates a DeleteApiregistrationV1CollectionAPIServiceUnauthorized with default headers values
func NewDeleteApiregistrationV1CollectionAPIServiceUnauthorized() *DeleteApiregistrationV1CollectionAPIServiceUnauthorized {
	return &DeleteApiregistrationV1CollectionAPIServiceUnauthorized{}
}

/*DeleteApiregistrationV1CollectionAPIServiceUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteApiregistrationV1CollectionAPIServiceUnauthorized struct {
}

func (o *DeleteApiregistrationV1CollectionAPIServiceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/apiregistration.k8s.io/v1/apiservices][%d] deleteApiregistrationV1CollectionApiServiceUnauthorized ", 401)
}

func (o *DeleteApiregistrationV1CollectionAPIServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
