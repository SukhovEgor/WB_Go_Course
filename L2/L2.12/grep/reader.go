package grep

import (
	"bufio"
	"os"
)

func ReadFromStdin() []string {
	return readLines(os.Stdin)
}

func ReadFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return readLines(file), nil
}

func readLines(file *os.File) []string {
	var result []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}