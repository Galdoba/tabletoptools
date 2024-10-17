package district

import "github.com/Galdoba/tabletoptools/pkg/mgt2/polity/empire/district/starsystem"

type District struct {
	Name   string
	System map[string]*starsystem.System
}
