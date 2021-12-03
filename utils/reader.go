package utils

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func Read(t *testing.T, fname string) *bufio.Scanner {
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}

	return bufio.NewScanner(bytes.NewReader(b))
}
