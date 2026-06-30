package numtowords

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "zero"},
		{5, "five"},
		{10, "ten"},
		{11, "eleven"},
		{19, "nineteen"},
		{20, "twenty"},
		{21, "twenty one"},
		{99, "ninety nine"},
		{100, "one hundred"},
		{101, "one hundred one"},
		{110, "one hundred ten"},
		{115, "one hundred fifteen"},
		{999, "nine hundred ninety nine"},
		{1000, "one thousand"},
		{1001, "one thousand one"},
		{1010, "one thousand ten"},
		{1015, "one thousand fifteen"},
		{1100, "one thousand one hundred"},
		{1111, "one thousand one hundred eleven"},
		{9999, "nine thousand nine hundred ninety nine"},
		{-1, "Number out of range"},
		{10000, "Number out of range"},
	}

	for _, test := range tests {
		result := Convert(test.input)
		if result != test.expected {
			t.Errorf("Convert(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestConvert_InvalidNumber(t *testing.T) {
	invalidInputs := []int{-1, -100, -9999, MaxNumber + 1, 99999}
	for _, n := range invalidInputs {
		if got := Convert(n); got != "Number out of range" {
			t.Errorf("Convert(%d) = %q; want %q", n, got, "Number out of range")
		}
	}
}
