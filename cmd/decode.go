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

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]
		for shiftNum := int32(0); shiftNum < 26; shiftNum++ {
			fmt.Printf("%s => ", word)
			for _, c := range word {
				// c is int32 or rune type
				switch {
				case 'A' <= c && c <= 'Z':
					fmt.Printf(string((c-'A'+shiftNum)%26 + 'A'))
				case 'a' <= c && c <= 'z':
					fmt.Printf(string((c-'a'+shiftNum)%26 + 'a'))
				default:
					fmt.Printf(string(c))
				}
			}
			fmt.Printf(" ( %d shift )\n", shiftNum)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// decodeCmd.Flags().Int32VarP(&shiftNum, "shift", "s", 0, "shift numbers")
}
