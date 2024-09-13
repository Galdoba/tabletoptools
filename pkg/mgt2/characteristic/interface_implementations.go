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
		return (ch.effectiveScore % 3) - 2
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
