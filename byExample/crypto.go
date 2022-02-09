package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	sha := sha1.New()

	sha.Write([]byte(s))
	fmt.Println(sha.Size(), sha.BlockSize())

	shaSum := sha.Sum(nil)
	fmt.Printf("%+v\n", sha)
	fmt.Printf("%x\n", shaSum)
	fmt.Println(shaSum)
}
