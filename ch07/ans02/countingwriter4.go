package main

import (
	"fmt"
	"io"
)

type Wrapper4 struct {
	w io.Writer
	c int64
}

func (wp *Wrapper4) Write(p []byte) (n int, err error) {
	fmt.Println("Write")
	n, err = wp.w.Write(p)
	fmt.Println("wp.c:", wp.c)
	wp.c += int64(n)
	fmt.Println("wp.c:", wp.c)
	return
}

func CountingWriter4(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper4
	wp.w = w
	return &wp, &(wp.c)
}
