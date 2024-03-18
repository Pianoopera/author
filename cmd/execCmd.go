/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var since string
var accounts []string

// execCmdCmd represents the execCmd command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a shell script",
	Long:  `This command executes a specific shell script.`,
	Run: func(cmd *cobra.Command, args []string) {
		accounts = append(accounts, args[1])
		out, err := executeShellScriptWithArgs(args[0], since, accounts)
		if err != nil {
			fmt.Printf("Error executing script: %s\n", err)
			return
		}
		fmt.Printf("Script output: %s\n", out)
	},
}

func executeShellScriptWithArgs(searchDir string, ago string, accounts []string) (string, error) {
	err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == ".git" {
			repoPath := filepath.Dir(path)
			fmt.Println("+" + strings.Repeat("-", len(repoPath)+13) + "+")
			fmt.Println("| Repository: " + repoPath + " |")
			fmt.Println("+" + strings.Repeat("-", len(repoPath)+13) + "+")

			for _, author := range accounts {
				fmt.Printf(" <<< Account: %s >>>\n\n", author)
				cmd := exec.Command("git", "-C", repoPath, "log", "--since=\""+ago+" months ago\"", "--pretty=format:\"%h %cd [%Creset%s%C(yellow)%d%C(reset)]\"", "--author="+author, "--graph", "--date=short", "--decorate", "--all")
				out, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				fmt.Printf("%s\n\n", out)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("filepath.Walk() returned %v\n", err)
	}

	return "Done!!!!!!!!", nil
}

func init() {
	execCmd.Flags().StringVarP(&since, "since", "s", "3", "A description of your flag")
	rootCmd.AddCommand(execCmd)
}
