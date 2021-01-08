package main

import "bufio"

type LineCounter3 int

func (c *LineCounter3) Write(p []byte) (n int, err error) {
	nByte := len(p)
	for {
		advance, token, err := bufio.ScanLines(p, true)
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
