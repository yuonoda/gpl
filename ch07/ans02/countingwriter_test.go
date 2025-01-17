package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter3(t *testing.T) {
	str := "Hello World"
	w, c := CountingWriter3(os.Stdout)
	w.Write([]byte(str))
	fmt.Printf("w:%d\n", w)
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

	w, c := CountingWriter4(os.Stdout)

	var total int64 = 0

	for _, d := range data {
		bytes := []byte(d)
		w.Write(bytes)
		fmt.Println("*c:", *c)
		total += int64(len(bytes))

		if *c != total {
			t.Errorf("Count iis %d, want %d", *c, total)
		}
	}
}
