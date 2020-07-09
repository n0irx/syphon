/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strconv"
	"syphon/handler"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete id [--alias alias] [-a alias]",
	Short: "delete shell command from database",
	Long: `delete shell command from database, this command
can use ID or alias for identifier`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		byAlias, _ := cmd.Flags().GetBool("alias")
		deleteShellCommand(args, byAlias)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolP("alias", "a", false, "delete by alias")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteShellCommand(args []string, byAlias bool) {
	if byAlias == true {
		handler.DeleteCommandByAlias(args[0])
	} else {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please supply valid id value")
		}
		handler.DeleteCommandByID(id)
	}
}
