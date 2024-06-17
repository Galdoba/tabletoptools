package forcast

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/definition/difficulty"
)

func TestForcast(t *testing.T) {
	// f, _ := os.OpenFile(`c:\Users\pemaltynov\go\src\github.com\Galdoba\tabletoptools\pkg\dice\forcast\maps.go`, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	// f.Truncate(0)
	// f.Write([]byte(fmt.Sprintf("package forcast\n\n")))
	// f.Write([]byte(fmt.Sprintf("func diceMaps() map[int]map[int][]int {\n")))
	// f.Write([]byte(fmt.Sprintf("maps := make(map[int]map[int][]int)\n")))

	// defer f.Close()
	// for d := 1; d <= 10; d++ {
	// 	f.Write([]byte(fmt.Sprintf("dice%vMap := make(map[int][]int)\n", d)))
	// 	for i := 1 * d; i <= 6*d; i++ {
	// 		l, e, m, ol, sl, il, o, s, in, om, sm, im := diceOucomes(d, i)
	// 		fmt.Printf("dice%vMap[%v] = []int{%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v}\n", d, i, l, e, m, ol, sl, il, o, s, in, om, sm, im)

	// 		f.Write([]byte(fmt.Sprintf("dice%vMap[%v] = []int{%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v}\n", d, i, l, e, m, ol, sl, il, o, s, in, om, sm, im)))
	// 	}
	// 	f.Write([]byte(fmt.Sprintf("maps[%v] = dice%vMap\n", d, d)))

	// }
	// f.Write([]byte(fmt.Sprintf("return maps\n")))
	// f.Write([]byte(fmt.Sprintf("}\n")))
	////////////////////
	fmt.Println(T5_SUCCESS(difficulty.T5_AVERAGE, 7))
	fmt.Println(MGT2_SUCCESS(difficulty.MGT2_AVERAGE, 2))
	fmt.Println(MGT2_SUCCESS_SPECTACULAR(difficulty.MGT2_AVERAGE, 2))

}
