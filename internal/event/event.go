package event

import "fmt"

type Event struct {
	source       string
	initialDescr string //
	decidionOpts []string
}

type Decidion struct {
	Descriptor     string
	TestParameters map[string]string
	SuccessDescr   string
	SuccessConseq  *action
	DefaultDescr   string
	DefaultConseq  *action
}

type action struct {
	subject interface{}
	objects []interface{}
	fun     func(interface{}, ...interface{}) error
}

func NewAction(sub interface{}, fn func(interface{}, ...interface{}) error, obj ...interface{}) *action {
	ac := action{}
	ac.subject = sub
	ac.objects = obj
	ac.fun = fn
	return &ac
}

func (act *action) Exec() error {
	return act.fun(act.subject, act.objects...)
}

type subjStr struct {
	tell string
	hear string
}

type Speaker interface {
	Speak() string
}

type Listener interface {
	Listen(string) error
}

type Barker interface {
	Bark() string
}

func (s *subjStr) Speak() string {
	return s.tell
}

func (s *subjStr) Listen(word string) error {
	s.hear += " heard `" + word + "`"
	return nil
}
func (s *subjStr) Bark() string {
	return "bark"
}

func speak(spkr interface{}, lsnrs ...interface{}) error {
	var s Speaker
	switch spkr.(type) {
	default:
		return fmt.Errorf("bad speaker type")
	case Speaker:
		s = spkr.(Speaker)
	}
	var lr []Listener
	for i := range lsnrs {
		switch lsnrs[i].(type) {
		default:
			return fmt.Errorf("bad speaker type")
		case Listener:
			lr = append(lr, lsnrs[i].(Listener))
		}
	}
	for _, l := range lr {
		if err := l.Listen(s.Speak()); err != nil {
			return err
		}
	}
	return nil
}

func bark(spkr interface{}, lsnrs ...interface{}) error {
	fmt.Println("go bark")
	var s Barker
	switch spkr.(type) {
	default:
		return fmt.Errorf("bad speaker type")
	case Barker:
		s = spkr.(Barker)
	}
	var lr []Listener
	for i := range lsnrs {
		switch lsnrs[i].(type) {
		default:
			return fmt.Errorf("bad speaker type")
		case Listener:
			lr = append(lr, lsnrs[i].(Listener))
		}
	}
	for _, l := range lr {
		if err := l.Listen(s.Bark()); err != nil {
			return err
		}
	}
	return nil
}
