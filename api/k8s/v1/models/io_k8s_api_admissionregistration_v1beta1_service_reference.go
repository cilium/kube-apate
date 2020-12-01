// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPIAdmissionregistrationV1beta1ServiceReference ServiceReference holds a reference to Service.legacy.k8s.io
//
// swagger:model io.k8s.api.admissionregistration.v1beta1.ServiceReference
type IoK8sAPIAdmissionregistrationV1beta1ServiceReference struct {

	// `name` is the name of the service. Required
	// Required: true
	Name *string `json:"name"`

	// `namespace` is the namespace of the service. Required
	// Required: true
	Namespace *string `json:"namespace"`

	// `path` is an optional URL path which will be sent in any request to this service.
	Path string `json:"path,omitempty"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port int32 `json:"port,omitempty"`
}

// Validate validates this io k8s api admissionregistration v1beta1 service reference
func (m *IoK8sAPIAdmissionregistrationV1beta1ServiceReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1beta1ServiceReference) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1beta1ServiceReference) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAdmissionregistrationV1beta1ServiceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAdmissionregistrationV1beta1ServiceReference) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAdmissionregistrationV1beta1ServiceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
