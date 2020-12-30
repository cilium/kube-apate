// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cilium/kube-apate/api/management/v1/models"
)

var addCNs, delCNs int64

// cnCMD represents the endpoint_config command
var cnCMD = &cobra.Command{
	Use:     "cn",
	Short:   "Create and delete auto-generated CiliumNodes",
	Example: "cn --add 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if addCNs != 0 || delCNs != 0 {
			addNewCNs(addCNs, delCNs)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(cnCMD)
	cnCMD.Flags().Int64VarP(&addCNs, "add", "", 0, "Add new CiliumNodes")
	cnCMD.Flags().Int64VarP(&delCNs, "del", "", 0, "Delete existing CiliumNodes")
}

func addNewCNs(add, del int64) {
	err := client.CiliumNodesAdd(&models.Options{
		Add: add,
		Del: del,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d and Deleted %d CiliumNodes\n", add, del)
}
