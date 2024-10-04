package values

func BoundInt(i, min, max int) int {
	if min > max {
		panic("min can't be more than max")
	}
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}
