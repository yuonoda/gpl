package main

import (
	"fmt"
	"testing"
)

func TestWordCounter2_Write(t *testing.T) {
	bytes := []byte("Hello\n World")
	var c LineCounter2
	c.Write(bytes)
	fmt.Printf("c:%v", c)

	e := 2
	if c != LineCounter2(e) {
		t.Errorf("not match")

	}
}

func TestLineCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 1},
		{"Hello World\nHello World", 2},
		{"Hello World\nHello World\n", 2},
		{"Hello World\nHello World\n\n", 3},
		{"Hello World! こんにちは　世界", 1},
		{"Hello World!\nこんにちは　世界", 2},
	}

	var c LineCounter2
	for _, d := range data {
		c = 0
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

		if c != LineCounter2(d.expected) {
			t.Errorf("Result is %d, want %d", c, d.expected)
		}

	}
}
