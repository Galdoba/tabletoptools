package options

const (
	PD_NAME      = "Name"
	PD_AGE       = "Age"
	PD_SPECIE    = "Specie"
	PD_HOMEWORLD = "Homeworld"
	PD_TRAITS    = "Traits"
	RULESET      = "Ruleset"
	VALUE_TYPE   = "Value Type"
)

type Option struct {
	Key string
	Val interface{}
}

func New(key string, val interface{}) Option {
	opt := Option{}
	opt.Key = key
	opt.Val = val
	return opt
}
