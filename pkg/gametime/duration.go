package gametime

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type DurationType int

const (
	Duration_None                 = "Same Moment"
	ABSOLUTE         DurationType = 1
	VARIABLE         DurationType = 2
	RANDOMIZED       DurationType = 3
	MINUTE           int64        = 10
	TEN_MINUTES      int64        = MINUTE * 10
	HOUR             int64        = MINUTE * 60
	DAY              int64        = HOUR * 24
	WEEK             int64        = DAY * 7
	MONTH            int64        = WEEK * 4
	YEAR             int64        = DAY * 365
	ABOUT_10_MINUTES              = -1
	ABOUT_AN_HOUR                 = -2
	ALL_DAY                       = -3
	ABOUT_A_WEEK                  = -4
	ABOUT_A_MONTH                 = -5
)

type duration struct {
	tick int64
}

type Duration interface {
	Ticks() int64
	String() string
}

func (d *duration) Ticks() int64 {
	return d.tick
}

type Roller interface {
	Flux() int
	Roll3D() int
}

func NewDuration(dType DurationType, ticks int64, r ...Roller) Duration {
	switch dType {
	case ABSOLUTE:
		return newAbsoluteDuration(ticks)
	case VARIABLE:
		return newVariableDuration(ticks, r...)
	case RANDOMIZED:
		return newRandomizedDuration(ticks, r...)
	}
	return newAbsoluteDuration(0)
}

func newDuration(tick int64) *duration {
	if tick < 0 {
		tick = 0
	}
	d := duration{tick: tick}
	return &d
}

func Between(t1, t2 *Gametime) *duration {
	ticks := t1.tick - t2.tick
	if ticks < 0 {
		ticks = ticks * -1
	}
	dur := newDuration(ticks)
	return dur
}

func (d *duration) String() string {
	y, dy, h, m, s := 0, 0, 0, 0, 0
	ys, dys, hs, ms, ss := "", "", "", "", ""
	switch d.tick < 0 {
	case false:
		y = int(d.tick / TicksPerYear)
		if y != 0 {
			ys = toString(y) + "y"
		}
		dy = int(d.tick/TicksPerDay) % 365
		if dy != 0 {
			dys = toString(dy) + "d"
		}
		h = int(d.tick/TicksPerHour) % 24
		if h != 0 {
			hs = toString(h) + "h"
		}
		m = int(d.tick/TicksPerMinute) % 60
		if m != 0 {
			ms = toString(m) + "m"
		}
		s = (int(d.tick) % 10) * 6
		if s != 0 {
			ss = toString(s) + "s"
		}

	}
	durStr := fmt.Sprintf("%v%v%v%v%v", ys, dys, hs, ms, ss)
	if durStr == "" {
		durStr = Duration_None
	}
	return durStr
}

func toString(i int) string {
	if i == 0 {
		return ""
	}
	if i < 0 {
		i = i * -1
	}
	return fmt.Sprintf("%v", i)
}

func newAbsoluteDuration(t int64) *duration {
	var base int64
	switch t {
	default:
		base = t
	case ABOUT_10_MINUTES:
		base = 10 * MINUTE
	case ABOUT_AN_HOUR:
		base = HOUR
	case ALL_DAY:
		base = DAY
	case ABOUT_A_WEEK:
		base = WEEK
	case MONTH:
		base = MONTH
	}
	return newDuration(base)
}

func flux(r ...Roller) int64 {
	flx := 0
	switch len(r) {
	case 0:
		rn := rand.New(rand.NewSource(time.Now().UnixNano()))
		flx = (rn.Intn(6) + 1) - (rn.Intn(6) + 1)
	default:
		flx = r[0].Flux()
	}
	return int64(flx)
}

func random3D(r ...Roller) int64 {
	rnd := 0
	switch len(r) {
	case 0:
		rn := rand.New(rand.NewSource(time.Now().UnixNano()))
		rnd = rn.Intn(6) + rn.Intn(6) + rn.Intn(6) + 3
	default:
		rnd = r[0].Flux()
	}
	return int64(rnd)
}

func variableDurations(t int64, r ...Roller) (int64, int64) {
	var stDur int64
	switch t {
	default:
		stDur = t
		tenth := stDur / 10
		variable := tenth * flux(r...)
		return stDur, variable
	case ABOUT_10_MINUTES:
		return MINUTE * 10, flux(r...) * MINUTE
	case ABOUT_AN_HOUR:
		return MINUTE * 60, 10 * MINUTE * flux(r...)
	case ALL_DAY:
		return 10 * HOUR, flux(r...) * HOUR
	case ABOUT_A_WEEK:
		return 6 * DAY, (flux(r...) * DAY)
	case ABOUT_A_MONTH:
		return 6 * WEEK, (flux(r...) * WEEK)
	}
}

func newVariableDuration(t int64, r ...Roller) *duration {
	base, variable := variableDurations(t, r...)
	return newDuration(base + variable)
}

func randomizedDurations(t int64, r ...Roller) int64 {
	var stDur int64
	switch t {
	default:
		stDur = t
	case ABOUT_10_MINUTES:
		stDur = MINUTE * 10
	case ABOUT_AN_HOUR:
		stDur = MINUTE * 60
	case ALL_DAY:
		stDur = HOUR * 10
	case ABOUT_A_WEEK:
		stDur = DAY * 7
	case ABOUT_A_MONTH:
		stDur = DAY * 28
	}
	tenth := stDur / 10
	randomized := tenth * random3D(r...)
	if randomized < 0 {
		randomized = 0
	}
	return randomized
}

func newRandomizedDuration(t int64, r ...Roller) *duration {
	t = randomizedDurations(t, r...)
	return newDuration(t)
}

func AddAll(durations ...*duration) *duration {
	total := newDuration(0)
	for _, d := range durations {
		total.tick += d.tick
	}
	return total
}

func ParseDuration(str string) (Duration, error) {
	data := strings.Split(str, "")
	current := ""
	var ticks int64
	v := 0
	err := fmt.Errorf("not parsed")
	for _, glyph := range data {
		switch glyph {
		case "y":
			v, err = strconv.Atoi(current)
			ticks += int64(v * TicksPerYear)
		case "d":
			v, err = strconv.Atoi(current)
			ticks += int64(v * TicksPerDay)
		case "h":
			v, err = strconv.Atoi(current)
			ticks += int64(v * TicksPerHour)
		case "m":
			v, err = strconv.Atoi(current)
			ticks += int64(v * TicksPerMinute)
		case "s":
			v, err = strconv.Atoi(current)
			ticks += int64(v / 6)
		default:
			current += glyph
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("can't parse duration from '%v': %v", str, err.Error())
		}
		current = ""
	}
	return NewDuration(ABSOLUTE, ticks), nil
}
