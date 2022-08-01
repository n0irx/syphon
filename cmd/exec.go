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
	"github.com/noirvelle/syphon/handler"

	"github.com/spf13/cobra"
)

// execCmd represents the add command
var execCmd = &cobra.Command{
	Use:     "exec <alias>",
	Short:   "execute saved shell command.",
	Long:    "execute saved shell command from database by shell command alias or id.",
	Aliases: []string{"e"},
	Example: `  - syphon exec ssh-server-1
  - syphon e ssh-server-1`,
	Run: func(cmd *cobra.Command, args []string) {
		execShellCommand(args)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func execShellCommand(args []string) {
	handler.ExecCommand(args[0])
}
