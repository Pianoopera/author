/*
Copyright © 2024 Teto

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"encoding/json"
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

type Commit struct {
	Commit   string `json:"commit"`
	Date     string `json:"date"`
	Message  string `json:"message"`
	Branches string `json:"branches"` // ブランチ名を含める
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "author",
	Short: "author is a tool to get the history of commits of a specific author",
	Long: `author is a tool to get the history of commits of a specific author
You can use it to get the history of commits of a specific author in all the repositories in a directory.`,
	Args: cobra.ExactArgs(2),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		accounts = append(accounts, args[1])

		out, err := executeShellScriptWithArgs(args[0], since, accounts)
		if err != nil {
			fmt.Printf("Error executing script: %s\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf(" Script output: %s\n", out)
	},
}

func executeShellScriptWithArgs(searchDir string, ago string, accounts []string) (string, error) {
	err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == ".git" {
			repoPath := filepath.Dir(path)
			dirName := filepath.Base(repoPath)
			fmt.Println("+" + strings.Repeat("-", len(repoPath)+13) + "+")
			// lenを自動で計算する
			fmt.Println("  Repository: " + dirName)
			fmt.Println("+" + strings.Repeat("-", len(repoPath)+13) + "+")

			for _, author := range accounts {
				// print the author color green
				fmt.Printf(" <<< Account: %s >>>\n\n", author)
				// fmt.Printf(" <<< Account: \033[32m%s\033[0m >>>\n\n", author)

				cmd := exec.Command("git", "-C", repoPath, "log", "--since=\""+ago+" months ago\"", "--pretty=format:{\"commit\": \"%H\", \"branches\": \"%d\", \"date\": \"%ad\", \"message\": \"%f\"}", "--author="+author, "--date=short", "--decorate=full", "--all")
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					panic(err)
				}
				if err := cmd.Start(); err != nil {
					panic(err)
				}
				scanner := bufio.NewScanner(stdout)
				for scanner.Scan() {
					var commit Commit
					json.Unmarshal(scanner.Bytes(), &commit)
					fmt.Printf(" %s %s %s%s \n",
						commit.Commit,
						commit.Date,
						commit.Message,
						commit.Branches)
				}
				if err := cmd.Wait(); err != nil {
					panic(err)
				}
				fmt.Printf("\n\n")
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("filepath.Walk() returned %v\n", err)
	}

	return "Done!!!!!!!!", nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&since, "since", "s", "3", "since how many months ago")
	// --accounts user1,user2,user3
	rootCmd.Flags().StringSliceVarP(&accounts, "accounts", "a", []string{}, "accounts to search for")
}
