package upc_check_digit

import "math"

func UpcCheckDigit(input int) int {
	sum := 0

	for i := 0; i < 11; i ++ {
		digit := nthDigit(input, i)
		if i % 2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	m := sum % 10
	if m == 0 {
		return 0
	} else {
		return 10 - m
	}
}

func nthDigit(input int, nth int) int {
	rank := int(math.Pow(float64(10), float64(nth)))
	return (input / rank) % 10
}
