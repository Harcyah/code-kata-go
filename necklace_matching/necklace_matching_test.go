package necklace_matching

import (
    "testing"
)

func TestSameNecklace(t *testing.T) {
	if SameNecklace("nicole", "icolen") != true {
		t.Fatalf(`SameNecklace "nicole", "icolen" should be true`)
	}
	if SameNecklace("nicole", "lenico") != true {
		t.Fatalf(`SameNecklace "nicole", "lenico" should be true`)
	}
	if SameNecklace("nicole", "coneli") != false {
		t.Fatalf(`SameNecklace "nicole", "coneli" should be false`)
	}
	if SameNecklace("aabaaaaabaab", "aabaabaabaaa") != true {
		t.Fatalf(`SameNecklace "aabaaaaabaab", "aabaabaabaaa" should be true`)
	}
	if SameNecklace("abc", "cba") != false {
		t.Fatalf(`SameNecklace "abc", "cba" should be false`)
	}
	if SameNecklace("xxyyy", "xxxyy") != false {
		t.Fatalf(`SameNecklace "xxyyy", "xxxyy" should be false`)
	}
	if SameNecklace("xyxxz", "xxyxz") != false {
		t.Fatalf(`SameNecklace "xyxxz", "xxyxz" should be false`)
	}
	if SameNecklace("x", "x") != true {
		t.Fatalf(`SameNecklace "x", "x" should be true`)
	}
	if SameNecklace("x", "xx") != false {
		t.Fatalf(`SameNecklace "x", "xx" should be false`)
	}
	if SameNecklace("x", "") != false {
		t.Fatalf(`SameNecklace "x", "" should be false`)
	}
	if SameNecklace("", "") != true {
		t.Fatalf(`SameNecklace "", "" should be true`)
	}
}
