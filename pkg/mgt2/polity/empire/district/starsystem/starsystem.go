package starsystem

type System struct {
	Name  string
	World map[string]World
}

type World interface {
	Name() string
}
