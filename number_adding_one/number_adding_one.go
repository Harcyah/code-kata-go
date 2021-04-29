package number_adding_one

import "fmt"

func NumberAddingOne(number int) string {
	if number < 10 {
		return fmt.Sprint(number + 1)
	}

	left := NumberAddingOne(number / 10)
	right := NumberAddingOne(number % 10)

	return left + right
}
