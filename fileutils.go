package goutils

import (
	"bufio"
	"os"
	"strings"
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

func IsFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func RemoveTrailingSlash(path string) string {
	if len(path) > 1 && strings.HasSuffix(path, "/") {
		return path[:len(path)-1]
	}
	return path
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		panic("error: " + err.Error())
	}

	mode := stat.Mode()
	return mode.IsDir()
}

func IsRegular(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		panic("error: " + err.Error())
	}

	mode := stat.Mode()
	return mode.IsRegular()
}
