package goutils

import (
	"bufio"
	"os"
)

// ReadLine read file line by line and calls handler with each line and index.
func ReadLine(filepath string, handler func(line string, index uint64)) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	var idx uint64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		handler(scanner.Text(), idx)
		idx++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
