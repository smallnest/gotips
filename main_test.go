package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Log("TestMain")
}

func add(a, b int) int {
	return a * b
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name           string
		a, b, expected int
	}{
		{
			name:     "add 1 and 2",
			a:        1,
			b:        2,
			expected: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := add(tc.a, tc.b); got != tc.expected {
				t.Errorf("add(%d,%d) = %d, want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
