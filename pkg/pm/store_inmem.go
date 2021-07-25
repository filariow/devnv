package pm

import "fmt"

type inmemstore struct {
	projects map[string]Project
}

func (s *inmemstore) Create(p Project) error {
	if _, ok := s.projects[p.Name()]; ok {
		return fmt.Errorf("a project with name %s already exists", p.Name())
	}
	s.projects[p.Name()] = p
	return nil
}

func (s *inmemstore) Delete(p Project) error {
	if _, ok := s.projects[p.Name()]; !ok {
		return fmt.Errorf("a project with name %s does not exist", p.Name())
	}
	delete(s.projects, p.Name())
	return nil
}

func (s *inmemstore) Get() ([]Project, error) {
	i, pp := 0, make([]Project, len(s.projects))
	for _, p := range s.projects {
		pp[i] = p
		i++
	}
	return pp, nil
}

func (s *inmemstore) GetByName(name string) (Project, error) {
	p, ok := s.projects[name]
	if !ok {
		return nil, fmt.Errorf("a project with name %s does not exist", name)
	}
	return p, nil
}
