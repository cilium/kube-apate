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

// IoK8sAPIStorageV1beta1VolumeAttachmentStatus VolumeAttachmentStatus is the status of a VolumeAttachment request.
//
// swagger:model io.k8s.api.storage.v1beta1.VolumeAttachmentStatus
type IoK8sAPIStorageV1beta1VolumeAttachmentStatus struct {

	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachError *IoK8sAPIStorageV1beta1VolumeError `json:"attachError,omitempty"`

	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	// Required: true
	Attached *bool `json:"attached"`

	// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata map[string]string `json:"attachmentMetadata,omitempty"`

	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
	DetachError *IoK8sAPIStorageV1beta1VolumeError `json:"detachError,omitempty"`
}

// Validate validates this io k8s api storage v1beta1 volume attachment status
func (m *IoK8sAPIStorageV1beta1VolumeAttachmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttached(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetachError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1beta1VolumeAttachmentStatus) validateAttachError(formats strfmt.Registry) error {

	if swag.IsZero(m.AttachError) { // not required
		return nil
	}

	if m.AttachError != nil {
		if err := m.AttachError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachError")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIStorageV1beta1VolumeAttachmentStatus) validateAttached(formats strfmt.Registry) error {

	if err := validate.Required("attached", "body", m.Attached); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIStorageV1beta1VolumeAttachmentStatus) validateDetachError(formats strfmt.Registry) error {

	if swag.IsZero(m.DetachError) { // not required
		return nil
	}

	if m.DetachError != nil {
		if err := m.DetachError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("detachError")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeAttachmentStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeAttachmentStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1beta1VolumeAttachmentStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
