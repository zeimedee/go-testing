package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 = 4")
	}

}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, -2},
		{5, 10},
		{-5, -10},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted: {}, {} expected, reccieved: {}", test.input, test.expected, output)
		}
	}
}
