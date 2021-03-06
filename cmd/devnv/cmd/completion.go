/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func completionCmd(cliName string) *cobra.Command {
	return &cobra.Command{
		Use:                   "completion [bash|fish|oh-my-zsh|powershell|zsh]",
		Short:                 "Generates shell completions",
		Args:                  cobra.ExactValidArgs(1),
		ValidArgs:             []string{"bash", "fish", "oh-my-zsh", "zsh", "powershell"},
		DisableFlagsInUseLine: true,
		Long: fmt.Sprintf(`
	Outputs shell completion for the given shell (bash, fish, oh-my-zsh, or zsh)
	OS X:
		$ source $(brew --prefix)/etc/bash_completion	# for bash users
		$ source <(%[1]s completion bash)			# for bash users
		$ source <(%[1]s completion oh-my-zsh)		# for oh-my-zsh users
		$ source <(%[1]s completion zsh)			# for zsh users
	Ubuntu:
		$ source /etc/bash-completion	   	# for bash users
		$ source <(%[1]s completion bash) 	# for bash users
		$ source <(%[1]s completion oh-my-zsh) 	# for oh-my-zsh users
		$ source <(%[1]s completion zsh)  	# for zsh users
	Additionally, you may want to add this to your .bashrc/.zshrc
`, cliName),
		RunE: func(cmd *cobra.Command, args []string) error {
			rootCmd, writer := cmd.Root(), cmd.OutOrStdout()
			switch args[0] {
			case "bash":
				return rootCmd.GenBashCompletion(writer)
			case "fish":
				return rootCmd.GenFishCompletion(writer, true)
			case "powershell":
				return rootCmd.GenPowerShellCompletion(writer)
			case "oh-my-zsh":
				if err := rootCmd.GenZshCompletion(writer); err != nil {
					return err
				}
				compdef := fmt.Sprintf("\n compdef _%[1]s %[1]s\n", cliName)
				_, err := io.WriteString(writer, compdef)
				return err
			case "zsh":
				return rootCmd.GenZshCompletion(writer)
			}
			return nil
		},
	}
}
func init() {
	rootCmd.AddCommand(completionCmd(os.Args[0]))
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
