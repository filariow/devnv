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
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// scriptCmd represents the script command
var scriptCmd = &cobra.Command{
	Use:                   "script",
	Short:                 "Reads the script to be executed on change directory",
	Args:                  cobra.ExactValidArgs(1),
	ValidArgsFunction:     validArgs,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		pj, err := p.Get(args[0])
		if err != nil {
			return err
		}

		sf := path.Join(pj.Folder(), ".devnv")
		fi, err := os.Stat(sf)
		if os.IsNotExist(err) {
			fmt.Printf("'%s' file does not exists", sf)
			return nil
		}

		if fi.IsDir() {
			m := fmt.Sprintf("'%s' is a directory", sf)
			fmt.Println(m)
			return nil
		}

		f, err := ioutil.ReadFile(sf)
		if err != nil {
			fmt.Printf("Error reading file: %s", err)
			return nil
		}

		fmt.Println(string(f))
		return nil
	},
	Aliases: []string{"s"},
}

func init() {
	rootCmd.AddCommand(scriptCmd)
}
