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
)

// IoK8sAPIExtensionsV1beta1IngressBackend IngressBackend describes all endpoints for a given service and port.
//
// swagger:model io.k8s.api.extensions.v1beta1.IngressBackend
type IoK8sAPIExtensionsV1beta1IngressBackend struct {

	// Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, serviceName and servicePort must not be specified.
	Resource *IoK8sAPICoreV1TypedLocalObjectReference `json:"resource,omitempty"`

	// Specifies the name of the referenced service.
	ServiceName string `json:"serviceName,omitempty"`

	// Specifies the port of the referenced service.
	ServicePort IoK8sApimachineryPkgUtilIntstrIntOrString `json:"servicePort,omitempty"`
}

// Validate validates this io k8s api extensions v1beta1 ingress backend
func (m *IoK8sAPIExtensionsV1beta1IngressBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIExtensionsV1beta1IngressBackend) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIExtensionsV1beta1IngressBackend) validateServicePort(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicePort) { // not required
		return nil
	}

	if err := m.ServicePort.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("servicePort")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIExtensionsV1beta1IngressBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIExtensionsV1beta1IngressBackend) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIExtensionsV1beta1IngressBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
