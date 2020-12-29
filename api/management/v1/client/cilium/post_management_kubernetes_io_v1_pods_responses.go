// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostManagementKubernetesIoV1PodsReader is a Reader for the PostManagementKubernetesIoV1Pods structure.
type PostManagementKubernetesIoV1PodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostManagementKubernetesIoV1PodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostManagementKubernetesIoV1PodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostManagementKubernetesIoV1PodsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostManagementKubernetesIoV1PodsOK creates a PostManagementKubernetesIoV1PodsOK with default headers values
func NewPostManagementKubernetesIoV1PodsOK() *PostManagementKubernetesIoV1PodsOK {
	return &PostManagementKubernetesIoV1PodsOK{}
}

/*PostManagementKubernetesIoV1PodsOK handles this case with default header values.

OK
*/
type PostManagementKubernetesIoV1PodsOK struct {
}

func (o *PostManagementKubernetesIoV1PodsOK) Error() string {
	return fmt.Sprintf("[POST /management/kubernetes.io/v1/pods][%d] postManagementKubernetesIoV1PodsOK ", 200)
}

func (o *PostManagementKubernetesIoV1PodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostManagementKubernetesIoV1PodsAccepted creates a PostManagementKubernetesIoV1PodsAccepted with default headers values
func NewPostManagementKubernetesIoV1PodsAccepted() *PostManagementKubernetesIoV1PodsAccepted {
	return &PostManagementKubernetesIoV1PodsAccepted{}
}

/*PostManagementKubernetesIoV1PodsAccepted handles this case with default header values.

Accepted
*/
type PostManagementKubernetesIoV1PodsAccepted struct {
}

func (o *PostManagementKubernetesIoV1PodsAccepted) Error() string {
	return fmt.Sprintf("[POST /management/kubernetes.io/v1/pods][%d] postManagementKubernetesIoV1PodsAccepted ", 202)
}

func (o *PostManagementKubernetesIoV1PodsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
