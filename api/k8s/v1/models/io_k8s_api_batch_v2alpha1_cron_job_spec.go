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

// IoK8sAPIBatchV2alpha1CronJobSpec CronJobSpec describes how the job execution will look like and when it will actually run.
//
// swagger:model io.k8s.api.batch.v2alpha1.CronJobSpec
type IoK8sAPIBatchV2alpha1CronJobSpec struct {

	// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	ConcurrencyPolicy string `json:"concurrencyPolicy,omitempty"`

	// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	FailedJobsHistoryLimit int32 `json:"failedJobsHistoryLimit,omitempty"`

	// Specifies the job that will be created when executing a CronJob.
	// Required: true
	JobTemplate *IoK8sAPIBatchV2alpha1JobTemplateSpec `json:"jobTemplate"`

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// Required: true
	Schedule *string `json:"schedule"`

	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds int64 `json:"startingDeadlineSeconds,omitempty"`

	// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	SuccessfulJobsHistoryLimit int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
	Suspend bool `json:"suspend,omitempty"`
}

// Validate validates this io k8s api batch v2alpha1 cron job spec
func (m *IoK8sAPIBatchV2alpha1CronJobSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIBatchV2alpha1CronJobSpec) validateJobTemplate(formats strfmt.Registry) error {

	if err := validate.Required("jobTemplate", "body", m.JobTemplate); err != nil {
		return err
	}

	if m.JobTemplate != nil {
		if err := m.JobTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIBatchV2alpha1CronJobSpec) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIBatchV2alpha1CronJobSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIBatchV2alpha1CronJobSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIBatchV2alpha1CronJobSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
