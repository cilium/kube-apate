// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cilium/kube-apate/api/management/v1/models"
)

var addCEPs, delCEPs int64

// cepCMD represents the endpoint_config command
var cepCMD = &cobra.Command{
	Use:     "cep",
	Short:   "Create and delete auto-generated CEPs",
	Example: "cep --add 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if addCEPs != 0 || delCEPs != 0 {
			addNewCEPs(addCEPs, delCEPs)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(cepCMD)
	cepCMD.Flags().Int64VarP(&addCEPs, "add", "", 0, "Add new CEPs")
	cepCMD.Flags().Int64VarP(&delCEPs, "del", "", 0, "Delete existing CEPs")
}

func addNewCEPs(add, del int64) {
	err := client.CiliumEndpointsAdd(&models.Options{
		Add: add,
		Del: del,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d and Deleted %d CEPs\n", add, del)
}
