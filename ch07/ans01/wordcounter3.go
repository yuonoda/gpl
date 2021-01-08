package main

import "bufio"

type WordCounter3 int

func (c *WordCounter3) Write(p []byte) (n int, err error) {
	nByte := len(p)
	for {
		advance, token, err := bufio.ScanWords(p, true)
		if err != nil {
			return 0, err
		}

		if token != nil {
			*c++
		}

		if len(p) == 0 {
			return nByte, nil
		}
		p = p[advance:]
	}

}
