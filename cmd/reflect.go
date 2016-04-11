// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

	"github.com/exponent-io/gots/generate"
	"github.com/spf13/cobra"
)

// reflectCmd represents the reflect command
var reflectCmd = &cobra.Command{
	Use:   "reflect",
	Short: "Generate a lookup table for allocating (without pkg reflect).",
	Run: func(cmd *cobra.Command, args []string) {
		pkg, err := cmd.PersistentFlags().GetString("pkg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err.Error())
			return
		}
		err = generate.GenerateReflectTable(os.Stdout, pkg, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(reflectCmd)

	reflectCmd.PersistentFlags().String("pkg", "main", "Name of the package declaration to generate.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reflectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reflectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
