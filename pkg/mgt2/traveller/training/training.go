package training

type Trainable interface {
	Train(string) error
	Ensure(string, int) error
}
