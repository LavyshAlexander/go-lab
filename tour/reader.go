package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	runeA := 'A'
	for i := 0; i < len(b); i++ {
		b[i] = byte(runeA)
	}

	return len(b), nil
}

func main() {
	runeA := 'A'

	fmt.Println(byte(runeA))
	fmt.Println(int(runeA))
}
