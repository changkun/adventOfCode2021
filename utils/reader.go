package utils

import (
	"bufio"
	"bytes"
	"os"
)

func Read(fname string) (*bufio.Scanner, error) {
	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(bytes.NewReader(b)), nil
}
