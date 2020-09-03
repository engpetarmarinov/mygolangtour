package rot13reader

import (
	"io"
	"strings"
	"testing"
)

func TestRot13Reader_Read(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Lbh penpxrq gur pbqr!", "You cracked the code!"},
	}
	for _, c := range cases {
		s := strings.NewReader(c.in)
		rot13Reader := Rot13Reader{s}
		buff := make([]byte, 1)
		var got string
		for {
			_, err := rot13Reader.Read(buff)
			if err == io.EOF {
				break
			}
			got = got + string(buff)
		}

		if got != c.want {
			t.Errorf("rot13Reader.Read(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
