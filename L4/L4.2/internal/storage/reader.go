package storage

import (
	"bufio"
	"os"
)

func ReadLines(filename string) (<-chan string, <-chan error) {

	lines := make(chan string)
	errCh := make(chan error, 1)

	go func() {

		defer close(lines)
		defer close(errCh)

		file, err := os.Open(filename)
		if err != nil {
			errCh <- err
			return
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			errCh <- err
		}
	}()

	return lines, errCh
}