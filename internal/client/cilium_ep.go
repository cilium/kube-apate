// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/cilium/kube-apate/api/management/v1/client/cilium"
	"github.com/cilium/kube-apate/api/management/v1/models"
)

func (c *Client) CiliumEndpointsAdd(opts *models.Options) error {
	params := cilium.NewPostManagementCiliumIoV2CiliumEndpointsParams().WithOptions(opts).WithTimeout(ClientTimeout)

	_, _, err := c.Cilium.PostManagementCiliumIoV2CiliumEndpoints(params)
	return err
}

func (c *Client) CiliumNodesAdd(opts *models.Options) error {
	params := cilium.NewPostManagementCiliumIoV2CiliumNodesParams().WithOptions(opts).WithTimeout(ClientTimeout)

	_, _, err := c.Cilium.PostManagementCiliumIoV2CiliumNodes(params)
	return err
}
