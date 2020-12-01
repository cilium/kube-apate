// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cilium/kube-apate/internal/server"
)

const (
	apatePortFlag           = "apate-port"
	apateManagementPortFlag = "management-port"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Run kube-apate to mock kube-apiserver requests",
	Long:  "Run kube-apate to mock kube-apiserver requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		apatePortNumber := viper.GetInt(apatePortFlag)
		apateManagementNumber := viper.GetInt(apateManagementPortFlag)
		go server.StartApate(apatePortNumber)
		server.StartManagement(apateManagementNumber)
		return nil
	},
}

func init() {
	flags := Cmd.PersistentFlags()
	flags.IntP(apatePortFlag, "", 8080, "Port to mock kube-apiserver requests")
	flags.IntP(apateManagementPortFlag, "", 8081, "Port to manage kube-apate")
	viper.BindPFlags(flags)
}
