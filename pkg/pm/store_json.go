package pm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
)

type jsonstore struct {
	mux      sync.RWMutex
	projects map[string]Project

	fstore       *os.File
	fpath        string
	configFolder string
}

type projectJSON struct {
	Name   string
	Folder string
	Cmd    string
}

func (s *jsonstore) mapToModel(pj projectJSON) Project {
	return &project{
		name:   pj.Name,
		folder: pj.Folder,
		cmd:    pj.Cmd,
	}
}

func (s *jsonstore) mapToJSON(p Project) projectJSON {
	return projectJSON{
		Name:   p.Name(),
		Folder: p.Folder(),
		Cmd:    p.Cmd(),
	}
}

func NewJSONStore(folder string) (Store, error) {
	fp, err := filepath.Abs(folder)
	if err != nil {
		return nil, err
	}

	sf := path.Join(folder, "jsonstore")
	os.MkdirAll(sf, 0755)

	sff := path.Join(sf, "projects.json")
	fd, err := os.Open(sff)
	toInit := false
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		fd, err = os.Create(sff)
		toInit = true
		if err != nil {
			return nil, err
		}
	}

	s := jsonstore{
		fstore:       fd,
		fpath:        sff,
		configFolder: fp,
	}

	if toInit {
		s.updateJSON()
	}
	if err := s.loadJSON(); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *jsonstore) Create(p Project) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.projects[p.Name()]; ok {
		return fmt.Errorf("a project with name %s already exists", p.Name())
	}
	s.projects[p.Name()] = p
	err := s.updateJSON()
	return err
}

func (s *jsonstore) Delete(name string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.projects[name]; !ok {
		return fmt.Errorf("a project with name %s does not exist", name)
	}
	delete(s.projects, name)
	err := s.updateJSON()
	return err
}

func (s *jsonstore) Get() ([]Project, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	i, pp := 0, make([]Project, len(s.projects))
	for _, p := range s.projects {
		pp[i] = p
		i++
	}
	return pp, nil
}

func (s *jsonstore) GetByName(name string) (Project, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	p, ok := s.projects[name]
	if !ok {
		return nil, fmt.Errorf("a project with name %s does not exist", name)
	}
	return p, nil
}

func (s *jsonstore) updateJSON() error {
	pp := make([]projectJSON, len(s.projects))
	i := 0
	for _, p := range s.projects {
		pp[i] = s.mapToJSON(p)
		i++
	}

	bb, err := json.Marshal(pp)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(s.fpath, bb, 0); err != nil {
		return err
	}
	return nil
}

func (s *jsonstore) loadJSON() error {
	bb, err := ioutil.ReadAll(s.fstore)
	if err != nil {
		return err
	}

	var pp []projectJSON
	err = json.Unmarshal(bb, &pp)
	if err != nil {
		return err
	}

	m := make(map[string]Project, len(pp))
	for _, p := range pp {
		m[p.Name] = s.mapToModel(p)
	}
	s.projects = m
	return nil
}
