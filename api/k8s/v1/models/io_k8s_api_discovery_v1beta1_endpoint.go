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

// IoK8sAPIDiscoveryV1beta1Endpoint Endpoint represents a single logical "backend" implementing a service.
//
// swagger:model io.k8s.api.discovery.v1beta1.Endpoint
type IoK8sAPIDiscoveryV1beta1Endpoint struct {

	// addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100.
	// Required: true
	Addresses []string `json:"addresses"`

	// conditions contains information about the current status of the endpoint.
	Conditions *IoK8sAPIDiscoveryV1beta1EndpointConditions `json:"conditions,omitempty"`

	// hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must pass DNS Label (RFC 1123) validation.
	Hostname string `json:"hostname,omitempty"`

	// targetRef is a reference to a Kubernetes object that represents this endpoint.
	TargetRef *IoK8sAPICoreV1ObjectReference `json:"targetRef,omitempty"`

	// topology contains arbitrary topology information associated with the endpoint. These key/value pairs must conform with the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: * kubernetes.io/hostname: the value indicates the hostname of the node
	//   where the endpoint is located. This should match the corresponding
	//   node label.
	// * topology.kubernetes.io/zone: the value indicates the zone where the
	//   endpoint is located. This should match the corresponding node label.
	// * topology.kubernetes.io/region: the value indicates the region where the
	//   endpoint is located. This should match the corresponding node label.
	Topology map[string]string `json:"topology,omitempty"`
}

// Validate validates this io k8s api discovery v1beta1 endpoint
func (m *IoK8sAPIDiscoveryV1beta1Endpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIDiscoveryV1beta1Endpoint) validateAddresses(formats strfmt.Registry) error {

	if err := validate.Required("addresses", "body", m.Addresses); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1beta1Endpoint) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1beta1Endpoint) validateTargetRef(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetRef) { // not required
		return nil
	}

	if m.TargetRef != nil {
		if err := m.TargetRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1beta1Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1beta1Endpoint) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIDiscoveryV1beta1Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
