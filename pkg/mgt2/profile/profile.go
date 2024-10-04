package v2

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/ehex"
)

type profile struct {
	fields map[string]string
}

type Profile interface {
	Inject(string, string) error
	Field(string) string
	Format(string) string
	MustEhex(string) (string, int)
	Ehex(string) (string, int, error)
}

func New() *profile {
	pr := profile{}
	pr.fields = make(map[string]string)
	return &pr
}

func (pr *profile) Inject(key, val string) error {
	if _, ok := pr.fields[key]; ok {
		return fmt.Errorf("key '%v' is present", key)
	}
	pr.fields[key] = val
	return nil
}

func (pr *profile) Field(key string) string {
	return pr.fields[key]
}

func (pr *profile) Format(key string) string {
	keys := formatKeys(key)
	output := ""
	for _, k := range keys {
		switch k {
		case SEP0, SEP1, SEP2, SEP3:
			output += k
			continue
		default:
			output += pr.fields[k]
		}
	}
	return output
}

func (pr *profile) Ehex(key string) (string, int, error) {
	val := pr.fields[key]
	if val == "" {
		return "", 0, fmt.Errorf("field[%v] is absent", key)
	}
	v := ehex.ValueOf(val)
	c := ehex.ToCode(v)
	if val != c {
		return c, v, fmt.Errorf("field[%v] is not ehex: code='%v'; value='%v'", key, c, v)
	}
	return c, v, nil
}

func (pr *profile) MustEhex(key string) (string, int) {
	val := pr.fields[key]
	v := ehex.ValueOf(val)
	c := ehex.ToCode(v)
	if val != c {
		panic(fmt.Sprintf("field[%v] is not ehex: value='%v' '%v' '%v'", key, val, v, c))
	}
	return c, v
}
