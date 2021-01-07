package main

import (
	"bufio"
	"bytes"
)

type LineCounter2 int

func (c *LineCounter2) Write(p []byte) (n int, err error) {
	nByte := len(p)
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		*c++
	}

	return nByte, err
}
