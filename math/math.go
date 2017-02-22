package math

func Int64Abs(a, b int64) int64 {
	result := a - b
	if result > 0 {
		return result
	} else {
		return -result
	}
}

func IntAbs(a, b int) int {
	result := a - b
	if result > 0 {
		return result
	} else {
		return -result
	}
}
