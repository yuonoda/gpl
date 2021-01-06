package main

import (
	"bufio"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
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
