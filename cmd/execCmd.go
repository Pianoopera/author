/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var since string

// execCmdCmd represents the execCmd command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a shell script",
	Long:  `This command executes a specific shell script.`,
	Run: func(cmd *cobra.Command, args []string) {
		// シェルスクリプトを実行
		fmt.Printf("Executing script with args: %s\n", args[0])
		fmt.Printf("Since: %s\n", since)
		fmt.Printf("Executing script with args: %s\n", args[1])
		out, err := executeShellScriptWithArgs("./main.sh", args[0], since, args[1])
		if err != nil {
			fmt.Printf("Error executing script: %s\n", err)
			return
		}
		fmt.Printf("Script output: \n%s\n", out)
	},
}

func executeShellScriptWithArgs(scriptPath string, directory string, since string, arg2 string) (string, error) {
	cmd := exec.Command("bash", scriptPath, directory, since, arg2)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func init() {
	execCmd.Flags().StringVarP(&since, "since", "s", "3", "A description of your flag")
	rootCmd.AddCommand(execCmd)
}
