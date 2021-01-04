// Copyright 2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/cilium/kube-apate/api/management/v1/models"

	"github.com/spf13/cobra"
)

var (
	addNodes, delNodes int64
)

// nodeCMD represents the endpoint_config command
var nodeCMD = &cobra.Command{
	Use:     "node",
	Short:   "Create and delete auto-generated Nodes",
	Example: "node --add 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if addNodes != 0 || delNodes != 0 {
			addNewNodes(addNodes, delNodes, !withoutDeps)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(nodeCMD)
	nodeCMD.Flags().Int64VarP(&addNodes, "add", "", 0, "Add new Nodes")
	nodeCMD.Flags().Int64VarP(&delNodes, "del", "", 0, "Delete existing Nodes")
	nodeCMD.Flags().BoolVarP(&withoutDeps, "without-dependents", "", false, "Do not create node dependents (CiliumNodes) when auto-generating Nodes")
}

func addNewNodes(add, del int64, withDeps bool) {
	err := client.KubernetesNodesAdd(&models.Options{
		Add:            add,
		Del:            del,
		WithDependents: withDeps,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d and Deleted %d Nodes\n", add, del)
}
