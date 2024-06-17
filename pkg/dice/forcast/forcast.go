package forcast

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

/*
forecast.T5_Equal(7, forcast.For_2D) float64
forecast.Equal(7, difficulty.T5_Average) float64
forecast.EqualOrLess(mods, difficulty.MgT2_Average) float64
forecast.Less(7, 2 int) float64
forecast.More(7, 2 int) float64
forecast.MoreOrEqual(7, 2 int) float64
*/

func maxResults(diceNum int) int {
	max := 0
	switch diceNum {
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

func T5_Equal(tn int, diceNum int) float64 {
	max := maxResults(diceNum)
	if max == 0 {
		return -1.0
	}
	vals := setupMaps(tn, diceNum)
	return float64(vals[1]) / float64(max)
}

func T5_Equal_Or_More(tn int, diceNum int) float64 {
	max := maxResults(diceNum)
	if max == 0 {
		return -1.0
	}
	vals := setupMaps(tn, diceNum)
	return float64(vals[1]+vals[2]) / float64(max)
}

func setupMaps(tn, diceNum int) []int {
	vals := diceMaps()[diceNum][tn]
	if len(vals) == 0 {
		return []int{0, 0, 0}
	}
	return vals
}

func diceOucomes(qty, tn int) (int, int, int) {
	results := generateResults(qty)
	isTN := 0
	isLessTN := 0
	isMoreTN := 0
	for _, result := range results {
		sum := sum(result)
		if sum < tn {
			isLessTN++
		}
		if sum == tn {
			isTN++
		}
		if sum > tn {
			isMoreTN++
		}
	}
	return isLessTN, isTN, isMoreTN
}

func sum(sl []int) int {
	s := 0
	for _, v := range sl {
		s += v
	}
	return s
}

func generateResults(i int) [][]int {
	switch i {
	case 1:
		return oneDiceResults()
	case 2:
		return twoDiceResults()
	case 3:
		return threeDiceResults()
	case 4:
		return fourDiceResults()
	case 5:
		return fiveDiceResults()
	case 6:
		return sixDiceResults()
	case 7:
		return sevenDiceResults()
	case 8:
		return eightDiceResults()
	case 9:
		return nineDiceResults()
	case 10:
		return tenDiceResults()
	}
	return nil
}

func oneDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		o := []int{r1}
		outs = append(outs, o)
	}
	return outs
}

func twoDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			o := []int{r1, r2}
			outs = append(outs, o)
		}
	}
	return outs
}

func threeDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				o := []int{r1, r2, r3}
				outs = append(outs, o)
			}
		}
	}
	return outs
}

func fourDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					o := []int{r1, r2, r3, r4}
					outs = append(outs, o)
				}
			}
		}
	}
	return outs
}

func fiveDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					for _, r5 := range []int{1, 2, 3, 4, 5, 6} {
						o := []int{r1, r2, r3, r4, r5}
						outs = append(outs, o)
					}
				}
			}
		}
	}
	return outs
}

func sixDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					for _, r5 := range []int{1, 2, 3, 4, 5, 6} {
						for _, r6 := range []int{1, 2, 3, 4, 5, 6} {
							o := []int{r1, r2, r3, r4, r5, r6}
							outs = append(outs, o)
						}
					}
				}
			}
		}
	}
	return outs
}

func sevenDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					for _, r5 := range []int{1, 2, 3, 4, 5, 6} {
						for _, r6 := range []int{1, 2, 3, 4, 5, 6} {
							for _, r7 := range []int{1, 2, 3, 4, 5, 6} {
								o := []int{r1, r2, r3, r4, r5, r6, r7}
								outs = append(outs, o)
							}
						}
					}
				}
			}
		}
	}
	return outs
}

func eightDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					for _, r5 := range []int{1, 2, 3, 4, 5, 6} {
						for _, r6 := range []int{1, 2, 3, 4, 5, 6} {
							for _, r7 := range []int{1, 2, 3, 4, 5, 6} {
								for _, r8 := range []int{1, 2, 3, 4, 5, 6} {
									o := []int{r1, r2, r3, r4, r5, r6, r7, r8}
									outs = append(outs, o)
								}
							}
						}
					}
				}
			}
		}
	}
	return outs
}

func nineDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					for _, r5 := range []int{1, 2, 3, 4, 5, 6} {
						for _, r6 := range []int{1, 2, 3, 4, 5, 6} {
							for _, r7 := range []int{1, 2, 3, 4, 5, 6} {
								for _, r8 := range []int{1, 2, 3, 4, 5, 6} {
									for _, r9 := range []int{1, 2, 3, 4, 5, 6} {
										o := []int{r1, r2, r3, r4, r5, r6, r7, r8, r9}
										outs = append(outs, o)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return outs
}

func tenDiceResults() [][]int {
	outs := [][]int{}
	for _, r1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, r2 := range []int{1, 2, 3, 4, 5, 6} {
			for _, r3 := range []int{1, 2, 3, 4, 5, 6} {
				for _, r4 := range []int{1, 2, 3, 4, 5, 6} {
					for _, r5 := range []int{1, 2, 3, 4, 5, 6} {
						for _, r6 := range []int{1, 2, 3, 4, 5, 6} {
							for _, r7 := range []int{1, 2, 3, 4, 5, 6} {
								for _, r8 := range []int{1, 2, 3, 4, 5, 6} {
									for _, r9 := range []int{1, 2, 3, 4, 5, 6} {
										for _, r10 := range []int{1, 2, 3, 4, 5, 6} {
											o := []int{r1, r2, r3, r4, r5, r6, r7, r8, r9, r10}
											outs = append(outs, o)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return outs
}

func diceMaps() map[int]map[int][]int {
	maps := make(map[int]map[int][]int)
	dice1Map := make(map[int][]int)
	dice1Map[1] = []int{0, 1, 5}
	dice1Map[2] = []int{1, 1, 4}
	dice1Map[3] = []int{2, 1, 3}
	dice1Map[4] = []int{3, 1, 2}
	dice1Map[5] = []int{4, 1, 1}
	dice1Map[6] = []int{5, 1, 0}
	dice2Map := make(map[int][]int)
	dice2Map[2] = []int{0, 1, 35}
	dice2Map[3] = []int{1, 2, 33}
	dice2Map[4] = []int{3, 3, 30}
	dice2Map[5] = []int{6, 4, 26}
	dice2Map[6] = []int{10, 5, 21}
	dice2Map[7] = []int{15, 6, 15}
	dice2Map[8] = []int{21, 5, 10}
	dice2Map[9] = []int{26, 4, 6}
	dice2Map[10] = []int{30, 3, 3}
	dice2Map[11] = []int{33, 2, 1}
	dice2Map[12] = []int{35, 1, 0}
	dice3Map := make(map[int][]int)
	dice3Map[3] = []int{0, 1, 215}
	dice3Map[4] = []int{1, 3, 212}
	dice3Map[5] = []int{4, 6, 206}
	dice3Map[6] = []int{10, 10, 196}
	dice3Map[7] = []int{20, 15, 181}
	dice3Map[8] = []int{35, 21, 160}
	dice3Map[9] = []int{56, 25, 135}
	dice3Map[10] = []int{81, 27, 108}
	dice3Map[11] = []int{108, 27, 81}
	dice3Map[12] = []int{135, 25, 56}
	dice3Map[13] = []int{160, 21, 35}
	dice3Map[14] = []int{181, 15, 20}
	dice3Map[15] = []int{196, 10, 10}
	dice3Map[16] = []int{206, 6, 4}
	dice3Map[17] = []int{212, 3, 1}
	dice3Map[18] = []int{215, 1, 0}
	dice4Map := make(map[int][]int)
	dice4Map[4] = []int{0, 1, 1295}
	dice4Map[5] = []int{1, 4, 1291}
	dice4Map[6] = []int{5, 10, 1281}
	dice4Map[7] = []int{15, 20, 1261}
	dice4Map[8] = []int{35, 35, 1226}
	dice4Map[9] = []int{70, 56, 1170}
	dice4Map[10] = []int{126, 80, 1090}
	dice4Map[11] = []int{206, 104, 986}
	dice4Map[12] = []int{310, 125, 861}
	dice4Map[13] = []int{435, 140, 721}
	dice4Map[14] = []int{575, 146, 575}
	dice4Map[15] = []int{721, 140, 435}
	dice4Map[16] = []int{861, 125, 310}
	dice4Map[17] = []int{986, 104, 206}
	dice4Map[18] = []int{1090, 80, 126}
	dice4Map[19] = []int{1170, 56, 70}
	dice4Map[20] = []int{1226, 35, 35}
	dice4Map[21] = []int{1261, 20, 15}
	dice4Map[22] = []int{1281, 10, 5}
	dice4Map[23] = []int{1291, 4, 1}
	dice4Map[24] = []int{1295, 1, 0}
	dice5Map := make(map[int][]int)
	dice5Map[5] = []int{0, 1, 7775}
	dice5Map[6] = []int{1, 5, 7770}
	dice5Map[7] = []int{6, 15, 7755}
	dice5Map[8] = []int{21, 35, 7720}
	dice5Map[9] = []int{56, 70, 7650}
	dice5Map[10] = []int{126, 126, 7524}
	dice5Map[11] = []int{252, 205, 7319}
	dice5Map[12] = []int{457, 305, 7014}
	dice5Map[13] = []int{762, 420, 6594}
	dice5Map[14] = []int{1182, 540, 6054}
	dice5Map[15] = []int{1722, 651, 5403}
	dice5Map[16] = []int{2373, 735, 4668}
	dice5Map[17] = []int{3108, 780, 3888}
	dice5Map[18] = []int{3888, 780, 3108}
	dice5Map[19] = []int{4668, 735, 2373}
	dice5Map[20] = []int{5403, 651, 1722}
	dice5Map[21] = []int{6054, 540, 1182}
	dice5Map[22] = []int{6594, 420, 762}
	dice5Map[23] = []int{7014, 305, 457}
	dice5Map[24] = []int{7319, 205, 252}
	dice5Map[25] = []int{7524, 126, 126}
	dice5Map[26] = []int{7650, 70, 56}
	dice5Map[27] = []int{7720, 35, 21}
	dice5Map[28] = []int{7755, 15, 6}
	dice5Map[29] = []int{7770, 5, 1}
	dice5Map[30] = []int{7775, 1, 0}
	dice6Map := make(map[int][]int)
	dice6Map[6] = []int{0, 1, 46655}
	dice6Map[7] = []int{1, 6, 46649}
	dice6Map[8] = []int{7, 21, 46628}
	dice6Map[9] = []int{28, 56, 46572}
	dice6Map[10] = []int{84, 126, 46446}
	dice6Map[11] = []int{210, 252, 46194}
	dice6Map[12] = []int{462, 456, 45738}
	dice6Map[13] = []int{918, 756, 44982}
	dice6Map[14] = []int{1674, 1161, 43821}
	dice6Map[15] = []int{2835, 1666, 42155}
	dice6Map[16] = []int{4501, 2247, 39908}
	dice6Map[17] = []int{6748, 2856, 37052}
	dice6Map[18] = []int{9604, 3431, 33621}
	dice6Map[19] = []int{13035, 3906, 29715}
	dice6Map[20] = []int{16941, 4221, 25494}
	dice6Map[21] = []int{21162, 4332, 21162}
	dice6Map[22] = []int{25494, 4221, 16941}
	dice6Map[23] = []int{29715, 3906, 13035}
	dice6Map[24] = []int{33621, 3431, 9604}
	dice6Map[25] = []int{37052, 2856, 6748}
	dice6Map[26] = []int{39908, 2247, 4501}
	dice6Map[27] = []int{42155, 1666, 2835}
	dice6Map[28] = []int{43821, 1161, 1674}
	dice6Map[29] = []int{44982, 756, 918}
	dice6Map[30] = []int{45738, 456, 462}
	dice6Map[31] = []int{46194, 252, 210}
	dice6Map[32] = []int{46446, 126, 84}
	dice6Map[33] = []int{46572, 56, 28}
	dice6Map[34] = []int{46628, 21, 7}
	dice6Map[35] = []int{46649, 6, 1}
	dice6Map[36] = []int{46655, 1, 0}
	dice7Map := make(map[int][]int)
	dice7Map[7] = []int{0, 1, 279935}
	dice7Map[8] = []int{1, 7, 279928}
	dice7Map[9] = []int{8, 28, 279900}
	dice7Map[10] = []int{36, 84, 279816}
	dice7Map[11] = []int{120, 210, 279606}
	dice7Map[12] = []int{330, 462, 279144}
	dice7Map[13] = []int{792, 917, 278227}
	dice7Map[14] = []int{1709, 1667, 276560}
	dice7Map[15] = []int{3376, 2807, 273753}
	dice7Map[16] = []int{6183, 4417, 269336}
	dice7Map[17] = []int{10600, 6538, 262798}
	dice7Map[18] = []int{17138, 9142, 253656}
	dice7Map[19] = []int{26280, 12117, 241539}
	dice7Map[20] = []int{38397, 15267, 226272}
	dice7Map[21] = []int{53664, 18327, 207945}
	dice7Map[22] = []int{71991, 20993, 186952}
	dice7Map[23] = []int{92984, 22967, 163985}
	dice7Map[24] = []int{115951, 24017, 139968}
	dice7Map[25] = []int{139968, 24017, 115951}
	dice7Map[26] = []int{163985, 22967, 92984}
	dice7Map[27] = []int{186952, 20993, 71991}
	dice7Map[28] = []int{207945, 18327, 53664}
	dice7Map[29] = []int{226272, 15267, 38397}
	dice7Map[30] = []int{241539, 12117, 26280}
	dice7Map[31] = []int{253656, 9142, 17138}
	dice7Map[32] = []int{262798, 6538, 10600}
	dice7Map[33] = []int{269336, 4417, 6183}
	dice7Map[34] = []int{273753, 2807, 3376}
	dice7Map[35] = []int{276560, 1667, 1709}
	dice7Map[36] = []int{278227, 917, 792}
	dice7Map[37] = []int{279144, 462, 330}
	dice7Map[38] = []int{279606, 210, 120}
	dice7Map[39] = []int{279816, 84, 36}
	dice7Map[40] = []int{279900, 28, 8}
	dice7Map[41] = []int{279928, 7, 1}
	dice7Map[42] = []int{279935, 1, 0}
	dice8Map := make(map[int][]int)
	dice8Map[8] = []int{0, 1, 1679615}
	dice8Map[9] = []int{1, 8, 1679607}
	dice8Map[10] = []int{9, 36, 1679571}
	dice8Map[11] = []int{45, 120, 1679451}
	dice8Map[12] = []int{165, 330, 1679121}
	dice8Map[13] = []int{495, 792, 1678329}
	dice8Map[14] = []int{1287, 1708, 1676621}
	dice8Map[15] = []int{2995, 3368, 1673253}
	dice8Map[16] = []int{6363, 6147, 1667106}
	dice8Map[17] = []int{12510, 10480, 1656626}
	dice8Map[18] = []int{22990, 16808, 1639818}
	dice8Map[19] = []int{39798, 25488, 1614330}
	dice8Map[20] = []int{65286, 36688, 1577642}
	dice8Map[21] = []int{101974, 50288, 1527354}
	dice8Map[22] = []int{152262, 65808, 1461546}
	dice8Map[23] = []int{218070, 82384, 1379162}
	dice8Map[24] = []int{300454, 98813, 1280349}
	dice8Map[25] = []int{399267, 113688, 1166661}
	dice8Map[26] = []int{512955, 125588, 1041073}
	dice8Map[27] = []int{638543, 133288, 907785}
	dice8Map[28] = []int{771831, 135954, 771831}
	dice8Map[29] = []int{907785, 133288, 638543}
	dice8Map[30] = []int{1041073, 125588, 512955}
	dice8Map[31] = []int{1166661, 113688, 399267}
	dice8Map[32] = []int{1280349, 98813, 300454}
	dice8Map[33] = []int{1379162, 82384, 218070}
	dice8Map[34] = []int{1461546, 65808, 152262}
	dice8Map[35] = []int{1527354, 50288, 101974}
	dice8Map[36] = []int{1577642, 36688, 65286}
	dice8Map[37] = []int{1614330, 25488, 39798}
	dice8Map[38] = []int{1639818, 16808, 22990}
	dice8Map[39] = []int{1656626, 10480, 12510}
	dice8Map[40] = []int{1667106, 6147, 6363}
	dice8Map[41] = []int{1673253, 3368, 2995}
	dice8Map[42] = []int{1676621, 1708, 1287}
	dice8Map[43] = []int{1678329, 792, 495}
	dice8Map[44] = []int{1679121, 330, 165}
	dice8Map[45] = []int{1679451, 120, 45}
	dice8Map[46] = []int{1679571, 36, 9}
	dice8Map[47] = []int{1679607, 8, 1}
	dice8Map[48] = []int{1679615, 1, 0}
	dice9Map := make(map[int][]int)
	dice9Map[9] = []int{0, 1, 10077695}
	dice9Map[10] = []int{1, 9, 10077686}
	dice9Map[11] = []int{10, 45, 10077641}
	dice9Map[12] = []int{55, 165, 10077476}
	dice9Map[13] = []int{220, 495, 10076981}
	dice9Map[14] = []int{715, 1287, 10075694}
	dice9Map[15] = []int{2002, 2994, 10072700}
	dice9Map[16] = []int{4996, 6354, 10066346}
	dice9Map[17] = []int{11350, 12465, 10053881}
	dice9Map[18] = []int{23815, 22825, 10031056}
	dice9Map[19] = []int{46640, 39303, 9991753}
	dice9Map[20] = []int{85943, 63999, 9927754}
	dice9Map[21] = []int{149942, 98979, 9828775}
	dice9Map[22] = []int{248921, 145899, 9682876}
	dice9Map[23] = []int{394820, 205560, 9477316}
	dice9Map[24] = []int{600380, 277464, 9199852}
	dice9Map[25] = []int{877844, 359469, 8840383}
	dice9Map[26] = []int{1237313, 447669, 8392714}
	dice9Map[27] = []int{1684982, 536569, 7856145}
	dice9Map[28] = []int{2221551, 619569, 7236576}
	dice9Map[29] = []int{2841120, 689715, 6546861}
	dice9Map[30] = []int{3530835, 740619, 5806242}
	dice9Map[31] = []int{4271454, 767394, 5038848}
	dice9Map[32] = []int{5038848, 767394, 4271454}
	dice9Map[33] = []int{5806242, 740619, 3530835}
	dice9Map[34] = []int{6546861, 689715, 2841120}
	dice9Map[35] = []int{7236576, 619569, 2221551}
	dice9Map[36] = []int{7856145, 536569, 1684982}
	dice9Map[37] = []int{8392714, 447669, 1237313}
	dice9Map[38] = []int{8840383, 359469, 877844}
	dice9Map[39] = []int{9199852, 277464, 600380}
	dice9Map[40] = []int{9477316, 205560, 394820}
	dice9Map[41] = []int{9682876, 145899, 248921}
	dice9Map[42] = []int{9828775, 98979, 149942}
	dice9Map[43] = []int{9927754, 63999, 85943}
	dice9Map[44] = []int{9991753, 39303, 46640}
	dice9Map[45] = []int{10031056, 22825, 23815}
	dice9Map[46] = []int{10053881, 12465, 11350}
	dice9Map[47] = []int{10066346, 6354, 4996}
	dice9Map[48] = []int{10072700, 2994, 2002}
	dice9Map[49] = []int{10075694, 1287, 715}
	dice9Map[50] = []int{10076981, 495, 220}
	dice9Map[51] = []int{10077476, 165, 55}
	dice9Map[52] = []int{10077641, 45, 10}
	dice9Map[53] = []int{10077686, 9, 1}
	dice9Map[54] = []int{10077695, 1, 0}
	dice10Map := make(map[int][]int)
	dice10Map[10] = []int{0, 1, 60466175}
	dice10Map[11] = []int{1, 10, 60466165}
	dice10Map[12] = []int{11, 55, 60466110}
	dice10Map[13] = []int{66, 220, 60465890}
	dice10Map[14] = []int{286, 715, 60465175}
	dice10Map[15] = []int{1001, 2002, 60463173}
	dice10Map[16] = []int{3003, 4995, 60458178}
	dice10Map[17] = []int{7998, 11340, 60446838}
	dice10Map[18] = []int{19338, 23760, 60423078}
	dice10Map[19] = []int{43098, 46420, 60376658}
	dice10Map[20] = []int{89518, 85228, 60291430}
	dice10Map[21] = []int{174746, 147940, 60143490}
	dice10Map[22] = []int{322686, 243925, 59899565}
	dice10Map[23] = []int{566611, 383470, 59516095}
	dice10Map[24] = []int{950081, 576565, 58939530}
	dice10Map[25] = []int{1526646, 831204, 58108326}
	dice10Map[26] = []int{2357850, 1151370, 56956956}
	dice10Map[27] = []int{3509220, 1535040, 55421916}
	dice10Map[28] = []int{5044260, 1972630, 53449286}
	dice10Map[29] = []int{7016890, 2446300, 51002986}
	dice10Map[30] = []int{9463190, 2930455, 48072531}
	dice10Map[31] = []int{12393645, 3393610, 44678921}
	dice10Map[32] = []int{15787255, 3801535, 40877386}
	dice10Map[33] = []int{19588790, 4121260, 36756126}
	dice10Map[34] = []int{23710050, 4325310, 32430816}
	dice10Map[35] = []int{28035360, 4395456, 28035360}
	dice10Map[36] = []int{32430816, 4325310, 23710050}
	dice10Map[37] = []int{36756126, 4121260, 19588790}
	dice10Map[38] = []int{40877386, 3801535, 15787255}
	dice10Map[39] = []int{44678921, 3393610, 12393645}
	dice10Map[40] = []int{48072531, 2930455, 9463190}
	dice10Map[41] = []int{51002986, 2446300, 7016890}
	dice10Map[42] = []int{53449286, 1972630, 5044260}
	dice10Map[43] = []int{55421916, 1535040, 3509220}
	dice10Map[44] = []int{56956956, 1151370, 2357850}
	dice10Map[45] = []int{58108326, 831204, 1526646}
	dice10Map[46] = []int{58939530, 576565, 950081}
	dice10Map[47] = []int{59516095, 383470, 566611}
	dice10Map[48] = []int{59899565, 243925, 322686}
	dice10Map[49] = []int{60143490, 147940, 174746}
	dice10Map[50] = []int{60291430, 85228, 89518}
	dice10Map[51] = []int{60376658, 46420, 43098}
	dice10Map[52] = []int{60423078, 23760, 19338}
	dice10Map[53] = []int{60446838, 11340, 7998}
	dice10Map[54] = []int{60458178, 4995, 3003}
	dice10Map[55] = []int{60463173, 2002, 1001}
	dice10Map[56] = []int{60465175, 715, 286}
	dice10Map[57] = []int{60465890, 220, 66}
	dice10Map[58] = []int{60466110, 55, 11}
	dice10Map[59] = []int{60466165, 10, 1}
	dice10Map[60] = []int{60466175, 1, 0}
	maps[1] = dice1Map
	maps[2] = dice2Map
	maps[3] = dice3Map
	maps[4] = dice4Map
	maps[5] = dice5Map
	maps[6] = dice6Map
	maps[7] = dice7Map
	maps[8] = dice8Map
	maps[9] = dice9Map
	maps[10] = dice10Map
	return maps
}
