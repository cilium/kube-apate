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

// PostManagementCiliumIoV2CiliumIdentitiesReader is a Reader for the PostManagementCiliumIoV2CiliumIdentities structure.
type PostManagementCiliumIoV2CiliumIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostManagementCiliumIoV2CiliumIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostManagementCiliumIoV2CiliumIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostManagementCiliumIoV2CiliumIdentitiesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostManagementCiliumIoV2CiliumIdentitiesOK creates a PostManagementCiliumIoV2CiliumIdentitiesOK with default headers values
func NewPostManagementCiliumIoV2CiliumIdentitiesOK() *PostManagementCiliumIoV2CiliumIdentitiesOK {
	return &PostManagementCiliumIoV2CiliumIdentitiesOK{}
}

/*PostManagementCiliumIoV2CiliumIdentitiesOK handles this case with default header values.

OK
*/
type PostManagementCiliumIoV2CiliumIdentitiesOK struct {
}

func (o *PostManagementCiliumIoV2CiliumIdentitiesOK) Error() string {
	return fmt.Sprintf("[POST /management/cilium.io/v2/ciliumidentities][%d] postManagementCiliumIoV2CiliumIdentitiesOK ", 200)
}

func (o *PostManagementCiliumIoV2CiliumIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostManagementCiliumIoV2CiliumIdentitiesAccepted creates a PostManagementCiliumIoV2CiliumIdentitiesAccepted with default headers values
func NewPostManagementCiliumIoV2CiliumIdentitiesAccepted() *PostManagementCiliumIoV2CiliumIdentitiesAccepted {
	return &PostManagementCiliumIoV2CiliumIdentitiesAccepted{}
}

/*PostManagementCiliumIoV2CiliumIdentitiesAccepted handles this case with default header values.

Accepted
*/
type PostManagementCiliumIoV2CiliumIdentitiesAccepted struct {
}

func (o *PostManagementCiliumIoV2CiliumIdentitiesAccepted) Error() string {
	return fmt.Sprintf("[POST /management/cilium.io/v2/ciliumidentities][%d] postManagementCiliumIoV2CiliumIdentitiesAccepted ", 202)
}

func (o *PostManagementCiliumIoV2CiliumIdentitiesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
