package forcast

import "fmt"

const (
	FOR_1D  = 1
	FOR_2D  = 2
	FOR_3D  = 3
	FOR_4D  = 4
	FOR_5D  = 5
	FOR_6D  = 6
	FOR_7D  = 7
	FOR_8D  = 8
	FOR_9D  = 9
	FOR_10D = 10
	one     = 6
	two     = 36
	three   = 216
	four    = 1296
	five    = 7776
	six     = 46656
	seven   = 279936
	eight   = 1679616
	nine    = 10077696
	ten     = 60466176
)

//GENERAL FUNCS

//LESS - chance of [1-10]d6 roll less than tn
func LESS(tn int, diceNum int) float64 {
	if tn > diceNum*6 {
		return 1
	}
	if tn < diceNum {
		return 0
	}
	max := maxResults(diceNum)
	vals := setupMaps(tn, diceNum)
	return float64(vals[0]) / float64(max)
}

//EQUALorLESS - chance of [1-10]d6 roll equal or less than tn
func EQUALorLESS(tn int, diceNum int) float64 {
	if tn >= diceNum*6 {
		return 1
	}
	if tn < diceNum {
		return 0
	}
	max := maxResults(diceNum)
	vals := setupMaps(tn, diceNum)
	return float64(vals[0]+vals[1]) / float64(max)
}

//EQUAL - chance of [1-10]d6 roll equal to tn
func EQUAL(tn int, diceNum int) float64 {
	max := maxResults(diceNum)
	vals := setupMaps(tn, diceNum)
	return float64(vals[1]) / float64(max)
}

//EQUALorMORE - chance of [1-10]d6 roll equal or more than tn
func EQUALorMORE(tn int, diceNum int) float64 {
	if tn <= diceNum {
		return 1
	}
	if tn > diceNum*6 {
		return 0
	}
	max := maxResults(diceNum)
	vals := setupMaps(tn, diceNum)
	return float64(vals[1]+vals[2]) / float64(max)
}

//MORE - chance of [1-10]d6 roll equal or more than tn
func MORE(tn int, diceNum int) float64 {
	if tn < diceNum {
		return 1
	}
	if tn > diceNum*6 {
		return 0
	}

	max := maxResults(diceNum)
	vals := setupMaps(tn, diceNum)
	return float64(vals[2]) / float64(max)
}

////MGT2

func MGT2_SUCCESS(dif int, mods ...int) float64 {
	tn := dif - sum(mods)
	return EQUALorMORE(tn, FOR_2D)
}

func MGT2_SUCCESS_SPECTACULAR(dif int, mods ...int) float64 {
	tn := dif - sum(mods) + 6
	return EQUALorMORE(tn, FOR_2D)
}

func MGT2_FAILURE(dif int, mods ...int) float64 {
	tn := dif - sum(mods)
	return LESS(tn, FOR_2D)
}

func MGT2_FAILURE_SPECTACULAR(dif int, mods ...int) float64 {
	tn := dif - sum(mods) - 6
	return EQUALorLESS(tn, FOR_2D)
}

/////T5

func T5_SUCCESS(dif int, assets ...int) float64 {
	tn := sum(assets)
	if tn >= dif*6 {
		return 1
	}
	if tn < dif {
		return 0
	}
	max := maxResults(dif)
	vals := setupMaps(tn, dif)
	return float64(vals[0]+vals[1]+vals[9]) / float64(max)
}

func T5_SUCCESS_SPECTACULAR(dif int, assets ...int) float64 {
	tn := sum(assets)
	if dif < 3 {
		return 0
	}
	max := maxResults(dif)
	vals := setupMaps(tn, dif)
	return float64(vals[3]+vals[6]+vals[9]) / float64(max)
}

func T5_FAILURE(dif int, assets ...int) float64 {
	tn := sum(assets)
	return MORE(tn, dif)
}

func T5_FAILURE_SPECTACULAR(dif int, assets ...int) float64 {
	tn := sum(assets)
	if dif < 3 {
		return 0
	}
	max := maxResults(dif)
	vals := setupMaps(tn, dif)
	return float64(vals[7]+vals[10]) / float64(max)
}

////////HELPERS//////////

func setupMaps(tn, diceNum int) []int {
	vals := diceMaps()[diceNum][tn]
	if len(vals) == 0 {
		return []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}
	return vals
}

func maxResults(diceNum int) int {
	max := 0
	switch diceNum {
	default:
		panic(fmt.Sprintf("dice forcast: roll results for %vd6 not precalculated", diceNum))
	case 1:
		max = one
	case 2:
		max = two
	case 3:
		max = three
	case 4:
		max = four
	case 5:
		max = five
	case 6:
		max = six
	case 7:
		max = seven
	case 8:
		max = eight
	case 9:
		max = nine
	case 10:
		max = ten
	}
	return max
}

func sum(sl []int) int {
	s := 0
	for _, n := range sl {
		s += n
	}
	return s
}
