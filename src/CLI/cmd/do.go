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
	"CLI/db"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "To remove task from the list",
	Long:  `To remove a task from the list use CLI -do {Serial Number} command`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := strconv.Atoi(strings.Join(args, ""))
		if err != nil {
			fmt.Println(err)
			return
		}
		tasks, err := db.ListAllTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		if key <= 0 || key > len(tasks) {
			fmt.Println("Wrong Id please check it again")
			return
		}
		err = db.RemoveTask(tasks[key-1].Key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Task has been removed")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
