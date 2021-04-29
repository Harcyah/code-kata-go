package necklace_matching

import (
	"strings"
)

func SameNecklace(left string, right string) bool {
	if len(left) != len(right) {
		return false
	}

	return strings.Contains(left+left, right)
}
