package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello World Hey")

	n, err := CountAlphabets(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Letters %d\n", n)
}

func CountAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)
	for {
		n, err := r.Read(buffer)
		for _, l := range buffer[:n] {
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}
	}
}
