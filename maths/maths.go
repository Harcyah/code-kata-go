package maths

func Min(left int, right int) int {
	if left <= right {
		return left
	} else {
		return right
	}
}

func Max(left int, right int) int {
	if left >= right {
		return left
	} else {
		return right
	}
}

func Abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}
