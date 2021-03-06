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

	"github.com/gophercises/task/students/barisere/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := data.NewStorage()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create db connection", err)
			os.Exit(1)
		}
		defer storage.Close()
		todos, err := storage.List()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to retrieve Todo list\n", err)
			os.Exit(1)
		}
		for _, v := range todos {
			if v.Status == data.Pending {
				fmt.Print("[ ] ")
			} else if v.Status == data.Done {
				fmt.Print("[X] ")
			}
			fmt.Fprintf(os.Stdout, "%d.\t%s\n", v.ID, v.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
