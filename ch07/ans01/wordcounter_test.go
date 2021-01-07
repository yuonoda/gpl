package main

import (
	"fmt"
	"testing"
)

func TestWrite2(t *testing.T) {
	var c WordCounter2
	c.Write([]byte("Hello world"))
	fmt.Println("c:", c)
	e := 2
	if c != WordCounter2(e) {
		t.Errorf("Expected %d, got %d", e, c)
	}
}

func TestWordCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 2},
		{"Hello My World", 3},
		{"Hello My World ", 3},
		{"Hello World こんにちは　世界", 4},
		{"Hello World\nこんにちは　世界", 4},
	}

	var c WordCounter
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

		if c != WordCounter(d.expected) {
			t.Errorf("Result is %d, want %d", c, d.expected)
		}

	}
}
