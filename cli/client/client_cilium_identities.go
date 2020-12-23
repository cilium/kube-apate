// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cilium/kube-apate/api/management/v1/models"
)

var addCIs, delCIs int64

// ciCMD represents the endpoint_config command
var ciCMD = &cobra.Command{
	Use:     "ci",
	Short:   "Create and delete auto-generated CiliumIdentities",
	Example: "ci --add 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if addCIs != 0 || delCIs != 0 {
			addNewCIs(addCIs, delCIs)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(ciCMD)
	ciCMD.Flags().Int64VarP(&addCIs, "add", "", 0, "Add new CiliumIdentities")
	ciCMD.Flags().Int64VarP(&delCIs, "del", "", 0, "Delete existing CiliumIdentities")
}

func addNewCIs(add, del int64) {
	err := client.CiliumIdentitiesAdd(&models.Options{
		Add: models.Add(add),
		Del: models.Del(del),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d and Deleted %d CiliumIdentities\n", add, del)
}
