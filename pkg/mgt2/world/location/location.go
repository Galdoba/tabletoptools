package location

type worldLocation struct {
	sectorCoords string
}

func NewSectorCoordinates(loc string) *worldLocation {
	wl := worldLocation{}
	wl.sectorCoords = loc
	return &wl
}

type Location interface {
	SectorCoords() string
}

func (loc *worldLocation) SectorCoords() string {
	return loc.sectorCoords
}
