package pm

type Project interface {
	Name() string
	Folder() string
	Cmd() string
}

type Store interface {
	Create(Project) error
	Delete(string) error
	Get() ([]Project, error)
	GetByName(string) (Project, error)
}

type project struct {
	name   string
	folder string
	cmd    string
}

func (p *project) Name() string {
	return p.name
}

func (p *project) Folder() string {
	return p.folder
}

func (p *project) Cmd() string {
	return p.cmd
}
