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
	"github.com/go-openapi/validate"
)

// IoK8sAPIAppsV1StatefulSetSpec A StatefulSetSpec is the specification of a StatefulSet.
//
// swagger:model io.k8s.api.apps.v1.StatefulSetSpec
type IoK8sAPIAppsV1StatefulSetSpec struct {

	// podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
	PodManagementPolicy string `json:"podManagementPolicy,omitempty"`

	// replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
	Replicas int32 `json:"replicas,omitempty"`

	// revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.
	RevisionHistoryLimit int32 `json:"revisionHistoryLimit,omitempty"`

	// selector is a label query over pods that should match the replica count. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// Required: true
	Selector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"selector"`

	// serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	// Required: true
	ServiceName *string `json:"serviceName"`

	// template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.
	// Required: true
	Template *IoK8sAPICoreV1PodTemplateSpec `json:"template"`

	// updateStrategy indicates the StatefulSetUpdateStrategy that will be employed to update Pods in the StatefulSet when a revision is made to Template.
	UpdateStrategy *IoK8sAPIAppsV1StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"`

	// volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
	VolumeClaimTemplates []*IoK8sAPICoreV1PersistentVolumeClaim `json:"volumeClaimTemplates"`
}

// Validate validates this io k8s api apps v1 stateful set spec
func (m *IoK8sAPIAppsV1StatefulSetSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeClaimTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1StatefulSetSpec) validateSelector(formats strfmt.Registry) error {

	if err := validate.Required("selector", "body", m.Selector); err != nil {
		return err
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1StatefulSetSpec) validateServiceName(formats strfmt.Registry) error {

	if err := validate.Required("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAppsV1StatefulSetSpec) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1StatefulSetSpec) validateUpdateStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateStrategy) { // not required
		return nil
	}

	if m.UpdateStrategy != nil {
		if err := m.UpdateStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1StatefulSetSpec) validateVolumeClaimTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeClaimTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeClaimTemplates); i++ {
		if swag.IsZero(m.VolumeClaimTemplates[i]) { // not required
			continue
		}

		if m.VolumeClaimTemplates[i] != nil {
			if err := m.VolumeClaimTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeClaimTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1StatefulSetSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1StatefulSetSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1StatefulSetSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
