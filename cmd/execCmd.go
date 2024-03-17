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

// execCmdCmd represents the execCmd command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a shell script",
	Long:  `This command executes a specific shell script.`,
	Run: func(cmd *cobra.Command, args []string) {
		// シェルスクリプトを実行
		out, err := executeShellScriptWithArgs("./main.sh", "/Users/morikawafutoshito/git/Git-Author-Log-History/", "Pianoopera")
		if err != nil {
			fmt.Printf("Error executing script: %s\n", err)
			return
		}
		fmt.Printf("Script output: \n%s\n", out)
	},
}

func executeShellScriptWithArgs(scriptPath string, arg1 string, arg2 string) (string, error) {
	cmd := exec.Command("bash", scriptPath, arg1, arg2)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func init() {
	rootCmd.AddCommand(execCmd)
}
