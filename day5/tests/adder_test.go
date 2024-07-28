package main

import "testing"

type testlet struct {
	input1, input2 int
	expected       int
}

func TestAdd(t *testing.T) {
	tests := []testlet{
		{
			input1:   1,
			input2:   10,
			expected: 11,
		},
		{
			input1:   2,
			input2:   6,
			expected: 8,
		},
		{
			input1:   7,
			input2:   -6,
			expected: 1,
		},
		{
			input1:   6,
			input2:   -6,
			expected: 0,
		},
	}

	for index, test := range tests {
		actual := Add(test.input1, test.input2)
		if actual != test.expected {
			t.Errorf("test %d, expected: %d, got: %d\n", index, test.expected, actual)
		}
	}
}
