package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/filariow/devnv/pkg/pm"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) == 1 {
		usage()
		return nil
	}

	cf := configFolder()
	js, err := pm.NewJSONStore(cf)
	if err != nil {
		return err
	}
	p := pm.New(js)
	p.Load(os.Args[1])

	return nil
}

func usage() {
	fmt.Printf(`
usage: %s COMMAND <COMMAND_ARGS>
	cd:		open project
	add:	add a new project
	delete:	delete a project
`, os.Args[0])
}

func configFolder() string {
	cf := os.Getenv("PM_CONFIGFOLDER")
	if cf != "" {
		return cf
	}

	h := os.Getenv("HOME")
	cf = path.Join(h, ".devenv")
	return cf
}
