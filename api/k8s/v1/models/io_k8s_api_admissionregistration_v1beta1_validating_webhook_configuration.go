// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it. Deprecated in v1.16, planned for removal in v1.19. Use admissionregistration.k8s.io/v1 ValidatingWebhookConfiguration instead.
//
// swagger:model io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfiguration
type IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []*IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhook `json:"webhooks"`
}

// Validate validates this io k8s api admissionregistration v1beta1 validating webhook configuration
func (m *IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) validateWebhooks(formats strfmt.Registry) error {

	if swag.IsZero(m.Webhooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Webhooks); i++ {
		if swag.IsZero(m.Webhooks[i]) { // not required
			continue
		}

		if m.Webhooks[i] != nil {
			if err := m.Webhooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
