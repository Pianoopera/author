/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "author version",
	Run: func(cmd *cobra.Command, args []string) {
		// 最新のバージョンを表示する
		fmt.Printf("author version 0.1.2\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
