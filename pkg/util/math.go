package util

func Clamp(x int, min int, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
