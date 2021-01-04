// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/cilium/kube-apate/api/management/v1/client/cilium"
	"github.com/cilium/kube-apate/api/management/v1/models"
)

func (c *Client) KubernetesPodsAdd(opts *models.Options) error {
	params := cilium.NewPostManagementKubernetesIoV1PodsParams().WithOptions(opts).WithTimeout(ClientTimeout)

	_, _, err := c.Cilium.PostManagementKubernetesIoV1Pods(params)
	return err
}

func (c *Client) KubernetesNodesAdd(opts *models.Options) error {
	params := cilium.NewPostManagementKubernetesIoV1NodesParams().WithOptions(opts).WithTimeout(ClientTimeout)

	_, _, err := c.Cilium.PostManagementKubernetesIoV1Nodes(params)
	return err
}
