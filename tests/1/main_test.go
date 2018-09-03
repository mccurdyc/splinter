package main

import "testing"

func TestLess(t *testing.T) {
	tests := []struct {
		name     string
		input1   int
		input2   int
		expected bool
	}{
		{
			name:     "5<6",
			input1:   5,
			input2:   6,
			expected: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			actual := less(c.input1, c.input2)
			if c.expected != actual {
				t.Errorf("failed %s\n", c.name)
			}
		})
	}
}
