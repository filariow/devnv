package pm

import (
	"fmt"
)

type PM interface {
	CD(name string) error
	Add(name, folder string) error
	Delete(name string) error
	List() ([]Project, error)
}

type bashManager struct {
	s Store
}

func New(store Store) PM {
	bm := bashManager{s: store}
	return &bm
}

func (m *bashManager) CD(project string) error {
	p, err := m.s.GetByName(project)
	if err != nil {
		return err
	}
	f := `
#!/bin/bash

cd %s
fname=".devenv"
[ -f ${fname} ] && source ${fname}
`

	s := fmt.Sprintf(f, p.Folder())
	fmt.Printf(s)
	return nil
}

func (m *bashManager) Add(name, folder string) error {
	p := project{
		name:   name,
		folder: folder,
	}
	return m.s.Create(&p)
}

func (m *bashManager) Delete(name string) error {
	return m.s.Delete(name)
}

func (m *bashManager) List() ([]Project, error) {
	return m.s.Get()
}
