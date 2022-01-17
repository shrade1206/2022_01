/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package server

import (
	"fmt"
	"zzz/controller"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// APICmd represents the API command
var APICmd = &cobra.Command{
	Use:   "API",
	Short: "G",
	Long:  `Test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("API called")
		num, err := cmd.Flags().GetString("APINum")
		if err != nil {
			log.Warn().Err(err).Caller().Msg("num err")
			return
		}
		var (
			G = "GET"
			P = "POST"
		)
		switch num {
		case "1":
			fmt.Println("API 1號")
			port := "30015?Larry=123"
			header := "123"
			code := "123"
			controller.GetAPi(G, port, header, code)
		case "2":
			fmt.Println("API 2號")
			port := "30016"
			header := "321"
			code := "321"
			controller.GetAPi(P, port, header, code)
		case "3":
			fmt.Println("API 3號")
			port := "30017"
			header := "gordan"
			code := "gordan"
			controller.GetAPi(P, port, header, code)
		}

	},
}

func init() {
	rootCmd.AddCommand(APICmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// APICmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// APICmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	APICmd.Flags().StringP("APINum", "g", "", "輸入要取得的API數字")
}
