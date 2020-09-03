package rot13reader

import (
	"io"
)

type Rot13Reader struct {
	r io.Reader
}

func (rot *Rot13Reader) Read(buffer []byte) (n int, err error) {
	n, err = rot.r.Read(buffer)
	for i := 0; i < len(buffer); i++ {
		if buffer[i] >= 'A' && buffer[i] < 'Z' {
			buffer[i] = 65 + (((buffer[i] - 65) + 13) % 26)
		} else if buffer[i] >= 'a' && buffer[i] <= 'z' {
			buffer[i] = 97 + (((buffer[i] - 97) + 13) % 26)
		}
	}
	return
}
