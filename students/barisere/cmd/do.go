// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gophercises/task/students/barisere/data"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do id...",
	Short: "Mark a task on your TODO list as complete",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := data.NewStorage()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create db connection", err)
			os.Exit(1)
		}
		defer storage.Close()
		for _, v := range args {
			id, err := strconv.ParseUint(v, 10, 0)
			if err != nil {
				continue
			}
			todo := data.NewTodo(id, data.Pending, "")
			if err = storage.Do(todo); err != nil {
				fmt.Fprintln(os.Stderr, "Failed to save Todo\n", err)
				os.Exit(1)
			}
		}
		fmt.Printf("Todo item #%s removed from your todo list\n", strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
