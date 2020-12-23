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

// IoK8sAPINetworkingV1IngressBackend IngressBackend describes all endpoints for a given service and port.
//
// swagger:model io.k8s.api.networking.v1.IngressBackend
type IoK8sAPINetworkingV1IngressBackend struct {

	// Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
	Resource *IoK8sAPICoreV1TypedLocalObjectReference `json:"resource,omitempty"`

	// Service references a Service as a Backend. This is a mutually exclusive setting with "Resource".
	Service *IoK8sAPINetworkingV1IngressServiceBackend `json:"service,omitempty"`
}

// Validate validates this io k8s api networking v1 ingress backend
func (m *IoK8sAPINetworkingV1IngressBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1IngressBackend) validateResource(formats strfmt.Registry) error {

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

func (m *IoK8sAPINetworkingV1IngressBackend) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1IngressBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1IngressBackend) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1IngressBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}