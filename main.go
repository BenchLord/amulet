package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Bytes with value 0 are not shown in string but do change the hash.
// 0 should not be used as value.

func main() {
	bytes := make([]byte, 3) // 64 is max length
  for i := range bytes {
    bytes[i] = 32 // initialize all values as 32 (space character)
  }

	increment(0, bytes)
}

func increment(index int, bytes []byte) {
	if index == len(bytes)-1 {
    // TODO: break this out into seperate func instead of
    // copy/pasting same code below.

		// printable ascii starts at 32 and ends at 127.
    // treat 128 as newline (ascii 10)
		for i := 32; i < 129; i++ {
			bytes[index] = byte(i)
			if i == 128 {
				bytes[index] = byte(10)
			}
			hash(bytes)
		}

		// reset byte back to 32
		bytes[index] = byte(32)
		return
	}

	for i := 32; i < 129; i++ {
		increment(index+1, bytes)
		bytes[index] = byte(i)
		if i == 128 {
			bytes[index] = byte(10)
		}
		hash(bytes)
	}

	// reset byte back to 32
	bytes[index] = byte(32)
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

