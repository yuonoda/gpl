package main

import (
	"io"
)

var counter *int64

type wrapper struct {
	w io.Writer
	c int64
}

func (wp wrapper) Write(p []byte) (n int, err error) {
	n, err = wp.w.Write(p)
	wp.c += int64(n)
	return
}

func CountingWriter2(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper
	wp.w = w
	return &wp, &(wp.c)
}
