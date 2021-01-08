package main

import "io"

type Wrapper3 struct {
	w io.Writer
	c int64
}

func (wp *Wrapper3) Write(p []byte) (n int, err error) {
	n, err = wp.w.Write(p)
	wp.c += int64(n)
	return
}

func CountingWriter3(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper3
	wp.w = w
	return &wp, &wp.c
}
