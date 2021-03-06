// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/fanux/sealos/install"

	"github.com/spf13/cobra"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Simplest way to join your kubernets HA cluster",
	Long:  `sealos join  --node 192.168.0.5`,
	Run: func(cmd *cobra.Command, args []string) {
		beforeNodes:=install.Nodes
		c := &install.SealConfig{}
		c.Load("")
		install.BuildJoin()
		c.Nodes = append(c.Nodes,beforeNodes...)
		c.Dump("")
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
	joinCmd.Flags().StringSliceVar(&install.Nodes, "node", []string{}, "kubernetes nodes")
	joinCmd.Flags().StringSliceVar(&install.NodeIPs, "nodes", []string{}, "kubernetes multi-nodes ex. 192.168.0.5-192.168.0.5")
}
