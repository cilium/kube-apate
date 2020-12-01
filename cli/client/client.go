// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"
	"os"

	clientPkg "github.com/cilium/kube-apate/internal/client"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	apateManagementHostFlag = "host"
)

var (
	client *clientPkg.Client
)

var Cmd = &cobra.Command{
	Use:   "client",
	Short: "CLI to manage kube-apate server",
	Long:  "CLI to manage kube-apate server",
}

func init() {
	flags := Cmd.PersistentFlags()
	flags.StringP(apateManagementHostFlag, "", "tcp://localhost:8081", "URI to connect to the management console of apate's API")
	viper.BindPFlags(flags)
	Cmd.SetOut(os.Stdout)
	Cmd.SetErr(os.Stderr)

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cl, err := clientPkg.NewClient(viper.GetString(apateManagementHostFlag))
	if err != nil {
		panic(fmt.Sprintf("Error while creating client: %s\n", err))
	}
	client = cl
}
