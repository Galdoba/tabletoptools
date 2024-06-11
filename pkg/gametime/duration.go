package gametime

import "fmt"

type Duration struct {
	tick int64
}

func NewDuration(tick int64) *Duration {
	d := Duration{tick: tick}
	return &d
}

func (d *Duration) String() string {
	y, dy, h, m, s := 0, 0, 0, 0, 0
	switch d.tick < 0 {
	case false:
		y = int(d.tick / TicksPerYear)
		// d = int(d.tick/TicksPerDay) % 365
		// y = int(d.tick / TicksPerYear)
		// y = int(d.tick / TicksPerYear)
		// y = int(d.tick / TicksPerYear)

	}
	return fmt.Sprintf("%v%v%v%v%v", y, dy, h, m, s)
}
