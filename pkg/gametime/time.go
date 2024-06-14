package gametime

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Undefined = iota
	Wonday
	Tuday
	Thirday
	Forday
	Fiday
	Sixday
	Senday
	Holiday

	TicksPerMinute = 10
	TicksPerHour   = TicksPerMinute * 60
	TicksPerDay    = TicksPerHour * 24
	TicksPerYear   = TicksPerDay * 365
	TicksPerEon    = TicksPerYear * 10000
)

type Gametime struct {
	eon     int
	seconds int
	minutes int
	hours   int
	day     int //0-364
	year    int //0000-9999
	week    int //0-51
	month   int //0-11
	dow     int //0-7
	tick    int64
}

func NewGameTime(format string) (*Gametime, error) {
	return parse(format)
}

func NewByTick(tick int64) *Gametime {
	gt := Gametime{}
	gt.tick = tick
	normalizeTime(&gt)
	return &gt
}

func (gt *Gametime) Tick() int64 {
	t := int64(gt.seconds / 6)
	t += int64(gt.minutes * TicksPerMinute)
	t += int64(gt.hours * TicksPerHour)
	t += int64(gt.day * TicksPerDay)
	t += int64(gt.year * TicksPerYear)

	return t
}

func (gt *Gametime) MatchTick(t int64) {
	y := 0
	for t >= TicksPerYear {
		y++
		t -= TicksPerYear
	}
	y++

	d := 0
	for t >= TicksPerDay {
		d++
		t -= TicksPerDay
	}
	d++
	if d == 366 {
		d = 1
		y++
	}

}

func parse(format string) (*Gametime, error) {
	timeParsed := false
	dateParsed := false
	eonParsed := false
	gt := Gametime{}
	sets := strings.Split(format, " ")
	for _, set := range sets {
		if strings.Contains(set, "[") {
			if eonParsed {
				return nil, fmt.Errorf("second eonset not allowed")
			}
			set = strings.TrimPrefix(set, "[")
			set = strings.TrimSuffix(set, "]")

			val, err := strconv.Atoi(set)
			if err != nil {
				return nil, fmt.Errorf("set '%v': %v", set, err.Error())
			}
			gt.eon = val
			eonParsed = true
			continue
		}
		if strings.Contains(set, "-") {
			if dateParsed {
				return nil, fmt.Errorf("second dateset not allowed")
			}
			dateparts := strings.Split(set, "-")
			if len(dateparts) != 2 {
				return nil, fmt.Errorf("dateset '%v': bad formating", set)
			}
			for i, tp := range dateparts {
				val, err := strconv.Atoi(tp)
				if err != nil {
					return nil, fmt.Errorf("set '%v': %v", set, err.Error())
				}
				switch i {
				case 0:
					gt.day = val
				case 1:
					gt.year = val
				}
			}
			dateParsed = true
			continue
		}
		if strings.Contains(set, ":") {
			if timeParsed {
				return nil, fmt.Errorf("second timeset not allowed")
			}
			timeparts := strings.Split(set, ":")
			for i, tp := range timeparts {
				if i > 2 {
					continue
				}
				val, err := strconv.Atoi(tp)
				if err != nil {
					return nil, fmt.Errorf("set '%v': %v", set, err.Error())
				}
				switch i {
				case 0:
					gt.hours = val
				case 1:
					gt.minutes = val
				case 2:
					gt.seconds = val
				}
			}
			timeParsed = true
			continue
		}

		return nil, fmt.Errorf("set '%v' have no time/date/eon markers", set)
	}
	normalizeTick(&gt)
	normalizeTime(&gt)
	return &gt, nil
}

func normalizeTime(gt *Gametime) {
	gt.eon = Eons(gt.tick)
	gt.year = Years(gt.tick)
	gt.day = Days(gt.tick)
	gt.hours = Hours(gt.tick)
	gt.minutes = Minutes(gt.tick)
	gt.seconds = Seconds(gt.tick)
	gt.dow = weekday(gt.day)
	gt.week = week(gt.day)
	gt.month = month(gt.week)
}

func normalizeTick(gt *Gametime) {
	var t int64
	t += int64(gt.seconds / 6)
	t += int64(gt.minutes * TicksPerMinute)
	t += int64(gt.hours * TicksPerHour)
	t += int64(gt.day * TicksPerDay)
	t += int64(gt.year * TicksPerYear)
	t += int64(gt.eon * TicksPerEon)
	gt.tick = t
}

// func (gt *Gametime) Eon() int {
// 	eon := 0
// 	switch gt.tick < 0 {
// 	case true:
// 		tk := gt.tick
// 		for tk < 0 {
// 			eon--
// 			tk += TicksPerEon
// 		}
// 	case false:
// 		tk := gt.tick
// 		for tk > TicksPerEon {
// 			eon++
// 			tk -= TicksPerEon
// 		}
// 	}
// 	return eon
// }

func Eons(t int64) int {
	eon := 0
	switch t < 0 {
	case true:
		tk := t
		for tk < 0 {
			eon--
			tk += TicksPerEon
		}
	case false:
		tk := t
		for tk > TicksPerEon {
			eon++
			tk -= TicksPerEon
		}
	}
	return eon
}

func eonTick(t int64) int64 {
	tick := t
	for tick < 0 {
		tick += TicksPerEon
	}
	for tick > TicksPerEon {
		tick -= TicksPerEon
	}
	if tick < 0 {
		panic(0)
	}
	return tick
}

// func (gt *Gametime) Year() int {
// 	return int(eonTick(gt.tick)/TicksPerYear) + 1
// }

func Years(t int64) int {
	return int(eonTick(t)/TicksPerYear) + 1
}

// func (gt *Gametime) Day() int {
// 	return int(eonTick(gt.tick)/TicksPerDay)%365 + 1
// }

func Days(t int64) int {
	return int(eonTick(t)/TicksPerDay)%365 + 1
}

// func (gt *Gametime) Hour() int {
// 	return int(eonTick(gt.tick)/TicksPerHour) % 24
// }

func Hours(t int64) int {
	return int(eonTick(t)/TicksPerHour) % 24
}

// func (gt *Gametime) Minute() int {
// 	return int(eonTick(gt.tick)/TicksPerMinute) % 60
// }
func Minutes(t int64) int {
	return int(eonTick(t)/TicksPerMinute) % 60
}

func (gt *Gametime) Second() int {
	return int(eonTick(gt.tick)*6) % 60
}

func Seconds(t int64) int {
	return int(eonTick(t)*6) % 60
}

func weekday(d int) int {
	switch d {
	case 1:
		return Holiday
	default:
		dow := (d - 1) % 7
		if dow == 0 {
			dow = Senday
		}
		return dow
	}
}

func week(d int) int {
	switch d {
	case 1, 2, 3, 4, 5, 6, 7, 8:
		return 1
	default:
		return ((d - 2) / 7) + 1
	}
}

func month(w int) int {
	return ((w + 3) / 4)
}

func (gt *Gametime) String() string {
	h := fmt.Sprintf("%v", gt.hours)
	for len(h) < 2 {
		h = "0" + h
	}
	m := fmt.Sprintf("%v", gt.minutes)
	for len(m) < 2 {
		m = "0" + m
	}
	s := fmt.Sprintf("%v", gt.seconds)
	for len(s) < 2 {
		s = "0" + s
	}
	d := fmt.Sprintf("%v", gt.day)
	for len(d) < 3 {
		d = "0" + d
	}
	y := ""
	switch gt.year < 0 {
	case false:
		y = fmt.Sprintf("%v", gt.year)
		for len(y) < 4 {
			y = "0" + y
		}
		y = "-" + y
	case true:
		y = fmt.Sprintf("%v", gt.year*-1)
		for len(y) < 4 {
			y = "0" + y
		}
		y = "-" + y
	}
	mn := monthStr(gt.month)
	eo := eonStr(gt.eon)
	return fmt.Sprintf("%v:%v:%v %v%v%v%v", h, m, s, d, mn, y, eo)
}

func eonStr(e int) string {
	if e == 0 {
		return ""
	}
	if e > 0 {
		return fmt.Sprintf(" [+%v]", e)
	}
	return fmt.Sprintf(" [%v]", e)
}

func (gt *Gametime) Timestamp() string {
	h := fmt.Sprintf("%v", gt.hours)
	for len(h) < 2 {
		h = "0" + h
	}
	m := fmt.Sprintf("%v", gt.minutes)
	for len(m) < 2 {
		m = "0" + m
	}
	s := fmt.Sprintf("%v", gt.seconds)
	for len(s) < 2 {
		s = "0" + s
	}
	return fmt.Sprintf("%v:%v:%v", h, m, s)
}

func (gt *Gametime) Date() string {
	d := fmt.Sprintf("%v", gt.day)
	for len(d) < 3 {
		d = "0" + d
	}
	y := ""
	switch gt.year < 0 {
	case false:
		y = fmt.Sprintf("%v", gt.year)
		for len(y) < 4 {
			y = "0" + y
		}
		y = "-" + y
	case true:
		y = fmt.Sprintf("%v", gt.year*-1)
		for len(y) < 4 {
			y = "0" + y
		}
		y = "-" + y + "*"
	}
	return fmt.Sprintf("%v%v", d, y)
}

func monthStr(i int) string {
	switch i {
	case 1:
		return "-I"
	case 2:
		return "-II"
	case 3:
		return "-III"
	case 4:
		return "-IV"
	case 5:
		return "-V"
	case 6:
		return "-VI"
	case 7:
		return "-VII"
	case 8:
		return "-VIII"
	case 9:
		return "-IX"
	case 10:
		return "-X"
	case 11:
		return "-XI"
	case 12:
		return "-XII"
	case 13:
		return "-XIII"
	default:
		return ""
	}
}

func (gt *Gametime) After(d Duration) *Gametime {
	gt2 := NewByTick(gt.Tick())
	gt2.tick += d.Ticks()
	normalizeTime(gt2)
	return gt2
}

func (gt *Gametime) Pass(d Duration) {
	gt.tick += d.Ticks()
	normalizeTime(gt)
}
