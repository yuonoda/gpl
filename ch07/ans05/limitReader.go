package main

import "io"

type LimitedReader struct {
	r     io.Reader
	limit int64
	next  int
}

func (lr *LimitedReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	if int64(lr.next) >= lr.limit {
		return 0, io.EOF
	}

	nbytes := int(lr.limit - int64(lr.next))
	if nbytes > len(p) {
		nbytes = len(p)
	}
	n, err := lr.r.Read(p[:nbytes])
	if err != nil {
		return n, err
	}
	lr.next += nbytes
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n, 0}
}
