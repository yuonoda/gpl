package main

import (
	"fmt"
	"testing"
)

func TestLineCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 1},
		{"Hello \nMy World", 2},
		{"Hello \nMy World \n", 3},
		{"Hello World\nこんにちは　世界", 2},
		{"\nHello World\nこんにちは　世界", 3},
	}

	var c LineCounter
	for _, d := range data {
		c = 0
		fmt.Println("d.s:", d.s)
		bytes := []byte(d.s)
		n, err := c.Write(bytes)

		if err != nil {
			t.Errorf("Unexpected Error: %v", err)
			continue
		}

		if n != len(bytes) {
			t.Errorf("Written bytes is %d, want %d", n, len(bytes))
			continue
		}

		if c != LineCounter(d.expected) {
			t.Errorf("Result is %d, want %d", c, d.expected)
		}

	}
}
