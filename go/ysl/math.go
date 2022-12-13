package ysl

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Compare(a, b int) int {
	return Sgn(a - b)
}

func Sgn(x int) int {
	if x == 0 {
		return 0
	}
	return x / Abs(x)
}
