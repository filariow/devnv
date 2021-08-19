package pm

import (
	"fmt"
)

type PM interface {
	CD(name string) error
	Add(name, folder string) error
	Delete(name string) error
	List() ([]Project, error)
	Get(name string) (Project, error)
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
fname=".devnv"
[ -f ${fname} ] && [ -z ${DEVNV_NOSOURCE} ] && source ${fname} || true
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

func (m *bashManager) Get(name string) (Project, error) {
	return m.s.GetByName(name)
}
