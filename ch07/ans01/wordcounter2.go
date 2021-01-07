package main

import (
	"bufio"
)

type WordCounter2 int

func (c *WordCounter2) Write(p []byte) (n int, err error) {
	nBytes := len(p)
	for {
		advance, token, err := bufio.ScanWords(p, true)
		if err != nil {
			return 0, err
		}

		if token != nil {
			*c++
		}
		p = p[advance:]
		if len(p) == 0 {
			return nBytes, nil
		}

	}
}
