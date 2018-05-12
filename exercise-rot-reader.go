package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	if e == nil {
		for i, v := range b {
			if v >= 'a' && v <= 'z' {
				b[i] = (v - 'a' + 13) % 26 + 'a'
			} else if v >= 'A' && v <= 'Z' {
				b[i] = (v - 'A' + 13) % 26 + 'A'
			}
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
