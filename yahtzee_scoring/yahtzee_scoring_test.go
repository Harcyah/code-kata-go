package yahtzee_scoring

import (
	"testing"
)

func TestYahtzeeScoringSimple(t *testing.T) {
	example0 := YahtzeeScoring([]int{2, 3, 5, 5, 6})
	if example0 != 10 {
		t.Fatalf("Scoring of [2, 3, 5, 5, 6] shoud be 10, found %d", example0)
	}
	example1 := YahtzeeScoring([]int{1, 1, 1, 1, 3})
	if example1 != 4 {
		t.Fatalf("Scoring of [1, 1, 1, 1, 3] shoud be 4, found %d", example1)
	}
	example2 := YahtzeeScoring([]int{1, 1, 1, 3, 3})
	if example2 != 6 {
		t.Fatalf("Scoring of [1, 1, 1, 3, 3] shoud be 6, found %d", example2)
	}
	example3 := YahtzeeScoring([]int{1, 2, 3, 4, 5})
	if example3 != 5 {
		t.Fatalf("Scoring of [1, 2, 3, 4, 5] shoud be 5, found %d", example3)
	}
	example4 := YahtzeeScoring([]int{6, 6, 6, 6, 6})
	if example4 != 30 {
		t.Fatalf("Scoring of [6, 6, 6, 6, 6] shoud be 30, found %d", example4)
	}
}

func TestYahtzeeScoringBonus(t *testing.T) {
	example := YahtzeeScoring([]int{1654, 1654, 50995, 30864, 1654, 50995, 22747, 1654, 1654, 1654, 1654, 1654, 30864, 4868, 1654, 4868, 1654, 30864, 4868, 30864})
	if example != 123456 {
		t.Fatalf("Scoring of bonus shoud be 123456, found %d", example)
	}
}
