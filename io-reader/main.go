package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := strings.NewReader("Hello, my friends")

	b := make([]byte, 8)
	for {

		n, err := s.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
