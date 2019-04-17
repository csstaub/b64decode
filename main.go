package main

import (
	"encoding/base64"
	"io"
	"os"
)

type reader struct {
	in io.Reader
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.in.Read(p)

	for i := n; i > 0; i-- {
		switch p[i] {
		// Map - to +
		case 0x2D:
			p[i] = 0x2B
		// Map _ to /
		case 0x5F:
			p[i] = 0x2F
		// Strip =
		case 0x3D:
			n = i
			p = p[:i]
		default:
		}
	}

	return
}

func main() {
	in := reader{os.Stdin}
	io.Copy(os.Stdout, base64.NewDecoder(base64.RawStdEncoding, &in))
}
