// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/cilium/kube-apate/api/management/v1/models"

	"github.com/spf13/cobra"
)

var (
	addPods, delPods int64
	withoutDeps      bool
)

// podCMD represents the endpoint_config command
var podCMD = &cobra.Command{
	Use:     "pod",
	Short:   "Create and delete auto-generated Pods",
	Example: "pod --add 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if addPods != 0 || delPods != 0 {
			addNewPods(addPods, delPods, !withoutDeps)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(podCMD)
	podCMD.Flags().Int64VarP(&addPods, "add", "", 0, "Add new Pods")
	podCMD.Flags().Int64VarP(&delPods, "del", "", 0, "Delete existing Pods")
	podCMD.Flags().BoolVarP(&withoutDeps, "without-dependents", "", false, "Do not create Pod dependents (CEPs) when auto-generating Pods")
}

func addNewPods(add, del int64, withDeps bool) {
	err := client.KubernetesPodsAdd(&models.Options{
		Add:            models.Add(add),
		Del:            models.Del(del),
		WithDependents: models.Dependents(withDeps),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added %d and Deleted %d Pods\n", add, del)
}
