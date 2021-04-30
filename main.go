package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	bytes := make([]byte, 3) // 64 is max length
	increment(0, bytes)
}

func increment(index int, bytes []byte) {
	if index == len(bytes)-1 {
		// printable ascii starts at 32. treat 31 as newline (ascii 10)
		for i := 31; i < 128; i++ {
			bytes[index] = byte(i)
			if i == 31 {
				bytes[index] = byte(10)
			}
			hash(bytes)
		}

		// reset byte back to 0
		bytes[index] = byte(0)
		return
	}

	for i := 31; i < 128; i++ {
		increment(index+1, bytes)
		bytes[index] = byte(i)
		if i == 31 {
			bytes[index] = byte(10)
		}
		hash(bytes)
	}

	// reset byte back to 0
	bytes[index] = byte(0)
	return
}

func hash(bytes []byte) {
	hash := sha256.New()
	hash.Write(bytes)
	hashed := fmt.Sprintf("%x", hash.Sum(nil))
	if strings.Contains(hashed, "8888") {
		fmt.Println(string(bytes), hashed)
	}
}

