// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/cilium/kube-apate/api/management/v1/models"

	"github.com/spf13/cobra"
)

var addPods, delPods int64

// podCMD represents the endpoint_config command
var podCMD = &cobra.Command{
	Use:     "pod",
	Short:   "Create and delete auto-generated Pods",
	Example: "pod --add 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if addPods != 0 || delPods != 0 {
			addNewPods(addPods, delPods)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(podCMD)
	podCMD.Flags().Int64VarP(&addPods, "add", "", 0, "Add new Pods")
	podCMD.Flags().Int64VarP(&delPods, "del", "", 0, "Delete existing Pods")
}

func addNewPods(add, del int64) {
	err := client.KubernetesPodsAdd(&models.Options{
		Add: models.Add(add),
		Del: models.Del(del),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d and Deleted %d Pods\n", add, del)
}
