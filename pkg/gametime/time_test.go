package gametime

import (
	"fmt"
	"testing"
	"time"
)

func TestGametime(t *testing.T) {
	var i int64
	i = TicksPerYear - 10
	for i < 1000000 {
		gt := NewByTick(i)
		fmt.Printf("%v: %v       \n", i, gt)
		time.Sleep(time.Millisecond * 25)
		i = i + TicksPerDay
	}
	gt, err := NewGameTime("00:02:00 125-1105")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(gt)
	gt2, err := NewGameTime("07:04:06 124-1106")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(gt2)
	dur := Between(gt, gt2)
	fmt.Println(dur)
	for i := 0; i < 10; i++ {
		d := NewDuration(ABSOLUTE, 88888889)

		fmt.Printf("%v: %v         \n", i, d)

	}

}
