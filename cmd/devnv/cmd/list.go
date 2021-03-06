/*
Copyright © 2021 Francesco Ilario

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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the saved projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		pp, err := p.List()
		if err != nil {
			return err
		}

		ppo := make([]projectOut, len(pp))
		for i, p := range pp {
			ppo[i] = projectOut{
				Name:   p.Name(),
				Folder: p.Folder(),
			}
		}

		d, err := json.MarshalIndent(ppo, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(d))
		return nil
	},
	Aliases: []string{"l"},
}

type projectOut struct {
	Name   string `json:"name"`
	Folder string `json:"folder"`
}

func init() {
	rootCmd.AddCommand(listCmd)
}
