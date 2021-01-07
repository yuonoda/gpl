package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter2(t *testing.T) {
	str := "Hello World"
	w, c := CountingWriter2(os.Stdout)
	w.Write([]byte(str))
	fmt.Printf("c:%d\n", c)
	return
}

func TestCountingWriter(t *testing.T) {
	data := []string{
		"Hello World\n",
		"How are you?\n",
		"The Go Programming Language\n",
		"プログラミング言語Go\n",
	}

	w, c := CountingWriter2(os.Stdout)

	var total int64 = 0

	for _, d := range data {
		bytes := []byte(d)
		w.Write(bytes)
		total += int64(len(bytes))

		if *c != total {
			t.Errorf("Count iis %d, want %d", *c, total)
		}
	}
}
