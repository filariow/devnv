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
package main

import (
	"log"
	"os"
	"path"

	"github.com/filariow/devnv/cmd/devnv/cmd"
	"github.com/filariow/devnv/pkg/pm"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cf, err := configFolder()
	if err != nil {
		return err
	}
	js, err := pm.NewJSONStore(cf)
	if err != nil {
		return err
	}
	p := pm.New(js)

	cmd.Execute(p)
	return nil
}

func configFolder() (string, error) {
	cf := os.Getenv("PM_CONFIGFOLDER")
	if cf != "" {
		return cf, nil
	}

	h, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cf = path.Join(h, ".devenv")
	return cf, nil
}
