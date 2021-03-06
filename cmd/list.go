/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"todo/data"

	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks",
	Run: func(cmd *cobra.Command, args []string) {

		rdb := data.Client()
		done, err := cmd.Flags().GetBool("done")
		if err != nil {
			fmt.Println("Could not list", err)
		}
		data.List(context.Background(), rdb, done)
	},
}

func init() {
	listCmd.PersistentFlags().BoolVarP(&Verbose, "done", "d", false, "list of done tasks")
	rootCmd.AddCommand(listCmd)
}
