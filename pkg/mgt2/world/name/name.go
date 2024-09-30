package name

type worldName struct {
	name     string
	oldNames []string
}

type WorldName interface {
	Name() string
}

func New(name string) *worldName {
	wn := worldName{}
	wn.name = name
	return &wn
}

func (wn *worldName) Name() string {
	return wn.name
}
