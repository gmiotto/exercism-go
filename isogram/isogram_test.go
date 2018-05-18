package isogram

import (
	"fmt"
	"testing"
)

func TestIsIsogram(t *testing.T) {
	for _, c := range testCases {
		fmt.Println(c.input)
		if IsIsogram(c.input) != c.expected {
			t.Fatalf("FAIL: %s\nWord %q, expected %t, got %t", c.description, c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range testCases {
			IsIsogram(c.input)
		}

	}
}
