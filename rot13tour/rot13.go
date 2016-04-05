// Package rot13tour implements a tester for the rot13 excercise in the go tour
package rot13tour

import (
	"fmt"
	"io"
	"strings"
)

// TestReader takes a function that creates a rot13's Reader
// and executes a set of tests on Readers created with the function.
func TestReader(newRot13Reader func(io.Reader) io.Reader) error {
	p := make([]byte, 4096)
	for _, test := range tests {
		r13 := newRot13Reader(strings.NewReader(test.plain))
		n, err := r13.Read(p[:len(test.plain)])
		if n != len(test.plain) {
			return fmt.Errorf("unexpected number of bytes read (n) reading the stream: %q", test.plain)
		}
		if err != nil {
			return fmt.Errorf("unexpected error by the Rot13Reader reading the stream: %q", test.plain)
		}
		got := string(p[:n])
		if got != test.cipher {
			return fmt.Errorf("error encrypting %q\nwanted: %q\ngot:%q", test.plain, test.cipher, got)
		}
	}

	// after an error the n bytes read must have been encrypted
	var r shortReader
	r13 := newRot13Reader(r)
	n, err := r13.Read(p[:1])
	if err != io.EOF {
		return fmt.Errorf("base reader returned io.EOF but the Rot13Reader returned %v", err)
	}
	if n != 1 {
		return fmt.Errorf("base reader returned 1 byte wanted the Rot13Reader to have returned 1 byte but it returned %d instead", n)
	}
	if p[0] != 'n' {
		return fmt.Errorf("base reader returned 1 byte with value: 'a' wanted to have read 'n' but read %c instead", p[0])
	}
	return nil
}

type shortReader struct{}

func (s shortReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	p[0] = 'a'
	return 1, io.EOF
}

var tests = []struct {
	plain, cipher string
}{
	{
		plain:  "abcdefghijklmnopqrstuvwxyz",
		cipher: "nopqrstuvwxyzabcdefghijklm",
	},
	{
		"abcde@!#$%N",
		"nopqr@!#$%A",
	},
	{
		"",
		"",
	},
}
