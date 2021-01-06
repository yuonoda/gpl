package main

import "io"

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper
	wp.w = w
	return &wp, &(wp.c)
}

type Wrapper struct {
	c int64
	w io.Writer
}

func (wp *Wrapper) Write(p []byte) (n int, err error) {
	n, err = wp.w.Write(p)
	wp.c += int64(n)
	return
}
