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

// IoK8sAPICoreV1PodStatus PodStatus represents information about the status of a pod. Status may trail the actual state of a system, especially if the node that hosts the pod cannot contact the control plane.
//
// swagger:model io.k8s.api.core.v1.PodStatus
type IoK8sAPICoreV1PodStatus struct {

	// Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Conditions []*IoK8sAPICoreV1PodCondition `json:"conditions"`

	// The list has one entry per container in the manifest. Each entry is currently the output of `docker inspect`. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	ContainerStatuses []*IoK8sAPICoreV1ContainerStatus `json:"containerStatuses"`

	// Status for any ephemeral containers that have run in this pod. This field is alpha-level and is only populated by servers that enable the EphemeralContainers feature.
	EphemeralContainerStatuses []*IoK8sAPICoreV1ContainerStatus `json:"ephemeralContainerStatuses"`

	// IP address of the host to which the pod is assigned. Empty if not yet scheduled.
	HostIP string `json:"hostIP,omitempty"`

	// The list has one entry per init container in the manifest. The most recent successful init container will have ready = true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	InitContainerStatuses []*IoK8sAPICoreV1ContainerStatus `json:"initContainerStatuses"`

	// A human readable message indicating details about why the pod is in this condition.
	Message string `json:"message,omitempty"`

	// nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled.
	NominatedNodeName string `json:"nominatedNodeName,omitempty"`

	// The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod's status. There are five possible phase values:
	//
	// Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.
	//
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
	Phase string `json:"phase,omitempty"`

	// IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated.
	PodIP string `json:"podIP,omitempty"`

	// podIPs holds the IP addresses allocated to the pod. If this field is specified, the 0th entry must match the podIP field. Pods may be allocated at most 1 value for each of IPv4 and IPv6. This list is empty if no IPs have been allocated yet.
	PodIPs []*IoK8sAPICoreV1PodIP `json:"podIPs"`

	// The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md
	QosClass string `json:"qosClass,omitempty"`

	// A brief CamelCase message indicating details about why the pod is in this state. e.g. 'Evicted'
	Reason string `json:"reason,omitempty"`

	// RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod.
	// Format: date-time
	StartTime IoK8sApimachineryPkgApisMetaV1Time `json:"startTime,omitempty"`
}

// Validate validates this io k8s api core v1 pod status
func (m *IoK8sAPICoreV1PodStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeralContainerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitContainerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PodStatus) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodStatus) validateContainerStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.ContainerStatuses); i++ {
		if swag.IsZero(m.ContainerStatuses[i]) { // not required
			continue
		}

		if m.ContainerStatuses[i] != nil {
			if err := m.ContainerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodStatus) validateEphemeralContainerStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.EphemeralContainerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.EphemeralContainerStatuses); i++ {
		if swag.IsZero(m.EphemeralContainerStatuses[i]) { // not required
			continue
		}

		if m.EphemeralContainerStatuses[i] != nil {
			if err := m.EphemeralContainerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ephemeralContainerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodStatus) validateInitContainerStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.InitContainerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.InitContainerStatuses); i++ {
		if swag.IsZero(m.InitContainerStatuses[i]) { // not required
			continue
		}

		if m.InitContainerStatuses[i] != nil {
			if err := m.InitContainerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initContainerStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodStatus) validatePodIPs(formats strfmt.Registry) error {

	if swag.IsZero(m.PodIPs) { // not required
		return nil
	}

	for i := 0; i < len(m.PodIPs); i++ {
		if swag.IsZero(m.PodIPs[i]) { // not required
			continue
		}

		if m.PodIPs[i] != nil {
			if err := m.PodIPs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podIPs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodStatus) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := m.StartTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
