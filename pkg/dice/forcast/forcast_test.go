package forcast

import (
	"fmt"
	"testing"
)

func TestForcast(t *testing.T) {
	// f, _ := os.OpenFile("/home/galdoba/go/src/github.com/Galdoba/tabletoptools/pkg/dice/forcast/maps.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	// f.Truncate(0)
	// defer f.Close()
	// for d := 1; d <= 10; d++ {
	// 	f.Write([]byte(fmt.Sprintf("dice%vMap := make(map[int][]int)\n", d)))
	// 	for i := 1 * d; i <= 6*d; i++ {
	// 		l, e, m := diceOucomes(d, i)
	// 		fmt.Printf("dice%vMap[%v] = []int{%v, %v, %v}\n", d, i, l, e, m)

	// 		f.Write([]byte(fmt.Sprintf("dice%vMap[%v] = []int{%v, %v, %v}\n", d, i, l, e, m)))
	// 	}
	// }
	fmt.Println(T5_Equal_Or_More(14, 3))
}
