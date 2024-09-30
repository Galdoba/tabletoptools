package poi

type PointOfInterest struct {
	Name string
	Type string
}

func NewPOI(bt string) PointOfInterest {
	poi := PointOfInterest{}
	switch bt {
	case "P":
		poi.Type = "Commercial Port"
	case "M":
		poi.Type = "Military Base"
	case "N":
		poi.Type = "Naval Base"
	case "D":
		poi.Type = "Depot"
	case "C":
		poi.Type = "Corsair Base"
	case "S":
		poi.Type = "Scout Base"
	case "R":
		poi.Type = "Research Base"
	case "W":
		poi.Type = "Waystation"
	}
	return poi
}
