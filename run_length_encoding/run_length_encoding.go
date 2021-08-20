package run_length_encoding

import (
	"log"
	"strconv"
)

func RunLengthEncode(input string) string {
	return doRunLengthEncode(input)
}

func doRunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	char := input[0]
	counter := 1
	for i := 1; i < len(input); i++ {
		if input[i] == char {
			counter++
		} else {
			if counter == 1 {
				return string(char) + doRunLengthEncode(input[1:])
			} else {
				return strconv.Itoa(counter) + string(char) + doRunLengthEncode(input[i:])
			}
		}
	}

	if counter == 1 {
		return string(char)
	} else {
		return strconv.Itoa(counter) + string(char)
	}
}

func RunLengthDecode(input string) string {
	return doRunLengthDecode(input)
}

func doRunLengthDecode(input string) string {
	if input == "" {
		return ""
	}

	char := input[0]

	if char >= '1' && char <= '9' {
		idx := 1
		for ; input[idx] >= '1' && input[idx] <= '9'; {
			idx++
		}

		count, err := strconv.Atoi(input[0:idx])
		if err != nil {
			log.Panicf("Unable to parse %d as number", char)
		}

		char := input[idx]
		part := make([]uint8, count)
		for i := 0; i < count; i++ {
			part[i] = char
		}

		return string(part) + doRunLengthDecode(input[idx + 1:])
	} else {
		return string(char) + doRunLengthDecode(input[1:])
	}
}
