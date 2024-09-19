package characteristic

import "fmt"

/*
Modifier
	Mod() int
*/
func (ch *characteristic) Mod() int {
	switch ch.effectiveScore {
	case 0:
		return -3
	default:
		return (ch.effectiveScore / 3) - 2
	}
}

/*
Trainable
	Train() error
*/
func (ch *characteristic) Train() error {
	if ch.maxScore >= unmodifiedLimit(ch.creationDice) {
		return fmt.Errorf("characteristic limit reached")
	}
	ch.maxScore++
	return nil
}

func (cs *Set) Train(key string) error {
	key, ch := findChrInSet(key, cs)
	if ch == nil {
		return fmt.Errorf("no characteristic with key '%v'", key)
	}
	if ch.maxScore+1 > unmodifiedLimit(ch.creationDice) {
		return fmt.Errorf("characteristic limit reached")
	}
	ch.maxScore++
	cs.ByCode[key] = ch
	return nil
}

func findChrInSet(key string, cs *Set) (string, *characteristic) {
	for k, v := range cs.ByCode {
		if v.name == key || v.code == key || v.abb == key {
			return k, v
		}
	}
	return "", nil
}

func (cs *Set) Ensure(key string, value int) error {
	key, ch := findChrInSet(key, cs)
	if ch == nil {
		return fmt.Errorf("no characteristic with key '%v'", key)
	}
	if ch.maxScore < value {
		ch.maxScore = value
	}
	cs.ByCode[key] = ch
	return nil
}
