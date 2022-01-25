package main

import (
	"io"
	"os"
	"strings"
)

var rot13Cipher = map[rune]rune{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'A',
	'O': 'B',
	'P': 'C',
	'Q': 'D',
	'R': 'E',
	'S': 'F',
	'T': 'G',
	'U': 'H',
	'V': 'I',
	'W': 'J',
	'X': 'K',
	'Y': 'L',
	'Z': 'M',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
}

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	innerBuff := make([]byte, len(b))
	rr.r.Read(innerBuff)

	length := 0
	for i, r := range innerBuff {
		if val, exist := rot13Cipher[rune(r)]; exist {
			b[i] = byte(val)
		} else {
			b[i] = r
		}
		length++
	}

	if length <= len(b) {
		return length, io.EOF
	}

	return length, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
