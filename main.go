package main

import (
	"crypto/sha256"
	"fmt"
  "flag"
	"strings"
  "os"
)

// Bytes with value 0 are not shown in string but do change the hash.
// 0 should not be used as value.

var start = flag.String("start", "", "The starting bytes for the miner to increment. (Default is 32 spaces)")

func main() {
  flag.Parse()
  bytes, err := getStartingBytes(*start)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

	increment(0, bytes)
}

func getStartingBytes(s string) ([]byte, error) {
  var out []byte
  if s == "" {
    out = make([]byte, 32)
    for i := range out {
      out[i] = 32
    }

    return out, nil
  }

  out = []byte(s)
  if len(out) > 64 {
    return nil, fmt.Errorf("start cannot be greater than 64 bytes")
  }

  return out, nil
}

func increment(index int, bytes []byte) {
  f := func() {
    increment(index+1, bytes)
    hash(bytes)
  }

	if index == len(bytes)-1 {
    f = func() {
      hash(bytes)
    }
	}


  incrementByte(&bytes[index], f)
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

func incrementByte(b *byte, f func()) {
  // printable ascii starts at 32 and ends at 127.
  // treat 128 as newline (ascii 10)
  for i := 32; i < 129; i++ {
    *b = byte(i)
    if *b == 128 {
      *b = 10
    }
    f()
  }

  // reset byte back to 32
  *b = byte(32)
}
