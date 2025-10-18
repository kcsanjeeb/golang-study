package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("letters.txt")
	if err != nil {
		panic(err)
	}
	n, err := CountAlphabets(f)
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
