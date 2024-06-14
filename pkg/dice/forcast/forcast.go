package forcast

import "fmt"

const (
	one   = 6
	two   = 36
	three = 216
	four  = 1296
	five  = 7776
	six   = 46656
	seven = 279936
	eight = 1679616
	nine  = 10077696
	ten   = 60466176
	// two   = 6 * 6
	// two   = 6 * 6
	// two   = 6 * 6
	// two   = 6 * 6
	// two   = 6 * 6
)

func SumD6(qty int, tn int) float64 {
	if tn <= qty {
		return 1.0
	}
	if tn > qty*6 {
		return 0.0
	}
	return 2
}

func TenDiceOucomes(n int) int {
	for i := 10; i <= 6*6*6*6*6*6*6*6*6*6; i++ {
		if n == i {
			fmt.Println(n)
			break
		}
	}
	return n
}
