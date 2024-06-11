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
	gt, err := NewGameTime("00:02:37 125-1105 [-3]")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(gt)
	fmt.Println(*gt)
	fmt.Println(gt.Timestamp())
	fmt.Println(gt.Date())

}
