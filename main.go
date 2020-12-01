// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cilium/kube-apate/cli/client"
	"github.com/cilium/kube-apate/cli/server"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apate",
	Short: "CLI to manage kube-apate server",
	Long:  "CLI to manage kube-apate server",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	rootCmd.AddCommand(server.Cmd, client.Cmd)
}

func main() {
	Execute()
}
