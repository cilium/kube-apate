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

// IoK8sAPIEventsV1beta1Event Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
//
// swagger:model io.k8s.api.events.v1beta1.Event
type IoK8sAPIEventsV1beta1Event struct {

	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
	Action string `json:"action,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount int32 `json:"deprecatedCount,omitempty"`

	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	// Format: date-time
	DeprecatedFirstTimestamp IoK8sApimachineryPkgApisMetaV1Time `json:"deprecatedFirstTimestamp,omitempty"`

	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	// Format: date-time
	DeprecatedLastTimestamp IoK8sApimachineryPkgApisMetaV1Time `json:"deprecatedLastTimestamp,omitempty"`

	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource *IoK8sAPICoreV1EventSource `json:"deprecatedSource,omitempty"`

	// eventTime is the time when this Event was first observed. It is required.
	// Required: true
	// Format: date-time
	EventTime IoK8sApimachineryPkgApisMetaV1MicroTime `json:"eventTime"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// metadata
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note string `json:"note,omitempty"`

	// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
	Reason string `json:"reason,omitempty"`

	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *IoK8sAPICoreV1ObjectReference `json:"regarding,omitempty"`

	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *IoK8sAPICoreV1ObjectReference `json:"related,omitempty"`

	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController string `json:"reportingController,omitempty"`

	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance string `json:"reportingInstance,omitempty"`

	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series *IoK8sAPIEventsV1beta1EventSeries `json:"series,omitempty"`

	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
	Type string `json:"type,omitempty"`
}

// Validate validates this io k8s api events v1beta1 event
func (m *IoK8sAPIEventsV1beta1Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeprecatedFirstTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeprecatedLastTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeprecatedSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegarding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateDeprecatedFirstTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.DeprecatedFirstTimestamp) { // not required
		return nil
	}

	if err := m.DeprecatedFirstTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deprecatedFirstTimestamp")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateDeprecatedLastTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.DeprecatedLastTimestamp) { // not required
		return nil
	}

	if err := m.DeprecatedLastTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deprecatedLastTimestamp")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateDeprecatedSource(formats strfmt.Registry) error {

	if swag.IsZero(m.DeprecatedSource) { // not required
		return nil
	}

	if m.DeprecatedSource != nil {
		if err := m.DeprecatedSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deprecatedSource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateEventTime(formats strfmt.Registry) error {

	if err := m.EventTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eventTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateMetadata(formats strfmt.Registry) error {

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

func (m *IoK8sAPIEventsV1beta1Event) validateRegarding(formats strfmt.Registry) error {

	if swag.IsZero(m.Regarding) { // not required
		return nil
	}

	if m.Regarding != nil {
		if err := m.Regarding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regarding")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateRelated(formats strfmt.Registry) error {

	if swag.IsZero(m.Related) { // not required
		return nil
	}

	if m.Related != nil {
		if err := m.Related.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("related")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1Event) validateSeries(formats strfmt.Registry) error {

	if swag.IsZero(m.Series) { // not required
		return nil
	}

	if m.Series != nil {
		if err := m.Series.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("series")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIEventsV1beta1Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIEventsV1beta1Event) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIEventsV1beta1Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
