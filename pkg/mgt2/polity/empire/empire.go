package empire

import "github.com/Galdoba/tabletoptools/pkg/mgt2/polity/empire/district"

type Empire struct {
	Name      string
	Structure string
	Region    map[string]*district.District
}
